#!/bin/bash

# currently not used, can be ignored

set +e
LOGFILE="/tmp/build.log"
USER=$(awk -F: '$3 >= 1000 && $3 < 60000 {print $1; exit}' /etc/passwd)

# Fetch EC2 instance metadata
TOKEN=$(curl -X PUT -H "X-aws-ec2-metadata-token-ttl-seconds: 300" -s -m 3 "http://169.254.169.254/latest/api/token" )

if [ -n "$TOKEN" ]; then
  exec > >(tee "$LOGFILE") 2>&1
  EC2ID=$(curl -H "X-aws-ec2-metadata-token: $TOKEN" -s http://169.254.169.254/latest/meta-data/instance-id)
fi

if [ -n "$EC2ID" ]; then
  echo "Running on AWS, setting up environment..."

  REGION=$(curl -H "X-aws-ec2-metadata-token: $TOKEN" -s -m 5 http://169.254.169.254/latest/meta-data/placement/region)
  sudo snap install aws-cli --classic
  aws configure set default.region "$REGION"

  GIT_SHA=$(aws ec2 describe-tags --filters "Name=resource-id,Values=$EC2ID" "Name=key,Values=GIT_SHA" --query "Tags[0].Value" --output text)
  GIT_REPO=$(aws ec2 describe-tags --filters "Name=resource-id,Values=$EC2ID" "Name=key,Values=GIT_REPO" --query "Tags[0].Value" --output text)
  BUCKET_NAME=$(aws ec2 describe-tags --filters "Name=resource-id,Values=$EC2ID" "Name=key,Values=BUCKET" --query "Tags[0].Value" --output text)

  echo "Repository: $GIT_REPO"
  echo "Commit SHA: $GIT_SHA"

  # S3 bucket setup
  if ! aws s3api head-bucket --bucket "$BUCKET_NAME" 2>/dev/null; then
    aws s3api create-bucket --bucket "$BUCKET_NAME" --region "$REGION"
  fi
  save-log() {
    aws s3 cp "$LOGFILE" "s3://$BUCKET_NAME/build.log"
  }
  
  # install deps
  sudo apt-get update && sudo apt-get install -y python3 ccache make

  # Set up SSD and ccache only on EC2
  if lsblk | grep -q "nvme1n1"; then
    # use nvme for ccache if available
    sudo mkfs -t xfs /dev/nvme1n1 && sudo mkdir -p /data && sudo mount /dev/nvme1n1 /data && cd /data
  fi

  

  echo "Downloading repo to $CHROMIUM"
  git init && git remote add origin "$GIT_REPO" && git fetch origin "$GIT_SHA" && git checkout "$GIT_SHA"
  save-log
  sudo chown -R $USER:$USER $(pwd)

  if [ -d /data ]; then
    # setup data dir
    TMP="/data/tmp"
    CCACHE_DIR="/data/ccache"

    UHOME=$(su $USER -c 'echo $HOME')
    UCCACHE=$UHOME/.cache/ccache/

    sudo -u $USER env "PATH=$PATH" mkdir -p "$TMP" "$CCACHE_DIR" "$UCCACHE"
    mount --bind "$CCACHE_DIR" "$UCCACHE"

    if aws s3 ls "s3://$BUCKET_NAME/ccache.tar.gz"; then
        echo "Fetching ccache from S3..."
        sudo -u $USER env "PATH=$PATH" aws s3 cp "s3://$BUCKET_NAME/ccache.tar.gz" "$TMP/ccache.tar.gz" --quiet

        echo "Extracting ccache..."
        sudo -u $USER env "PATH=$PATH" tar -xzf "$TMP/ccache.tar.gz" -C "$CCACHE_DIR" --strip-components=2
        # rm -f "$TMP/ccache.tar.gz"
    fi
  fi
else
  sudo apt-get install -y python3 ccache
fi

echo "Running on $(uname -a)"

if [ ! -f entrypoint.sh ]; then
  echo "Git repo not properly initialized."
else
  sudo make deps && \
    sudo -u $USER env "PATH=$PATH" make patch && \
    sudo -u $USER env "PATH=$PATH" make build

  if [ -n "$EC2ID" ]; then

    echo "Uploading chrome to S3..."
    
    VERSION=$(cat "chromium.version")

    # create directories if not exist
    if ! aws s3 ls "s3://$BUCKET_NAME/releases/" > /dev/null 2>&1; then
      aws s3api put-object --bucket "$BUCKET_NAME" --key "releases/"
    fi
    if ! aws s3 ls "s3://$BUCKET_NAME/releases/$VERSION" > /dev/null 2>&1; then
      aws s3api put-object --bucket "$BUCKET_NAME" --key "releases/$VERSION"
    fi

    # push release
    echo "Uploading release.."
    tar -czf "$TMP/release.tar.gz" "build"
    aws s3 cp "$TMP/release.tar.gz" "s3://$BUCKET_NAME/releases/$VERSION/$GIT_SHA.release.tar.gz"
    
    save-log
    
    if [ -d /data ]; then
      echo "Uploading ccache to S3..."
      
      echo "Creating ccache archive..."
      tar -czf "$TMP/ccache.tar.gz" "$CCACHE_DIR"
      echo "Uploading ccache to S3..."
      aws s3 cp "$TMP/ccache.tar.gz" "s3://$BUCKET_NAME/ccache.tar.gz"
    fi
    
  fi
fi

# Upload ccache back to S3 after building
if [ -n "$EC2ID" ]; then
  save-log
  sudo shutdown -h now
fi
