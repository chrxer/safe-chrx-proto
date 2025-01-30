#!/bin/bash

set +e

build () {
  echo "Building chromium.."
  echo "Installing deps.. https://chromium.googlesource.com/chromium/src/+/main/docs/linux/build_instructions.md#Install"
  git clone https://chromium.googlesource.com/chromium/tools/depot_tools.git
  export PATH="$(pwd)/depot_tools:$PATH"
  mkdir ~/chromium && cd ~/chromium
  fetch --no-history --nohooks chromium
  cd src
  ./build/install-build-deps.sh
  gclient runhooks

  echo "Compiling chromium release now"
  gn gen out/Release --args="is_debug=false is_official_build=true symbol_level=0"
}

TOKEN=`curl -X PUT "http://169.254.169.254/latest/api/token" -H "X-aws-ec2-metadata-token-ttl-seconds: 300"`

if [ $TOKEN ]; then
  EC2ID=$(curl -H "X-aws-ec2-metadata-token: $TOKEN" -s -m 5 http://169.254.169.254/latest/meta-data/instance-id)
  exec > >(tee /tmp/build.log) 2>&1
fi

if [ $EC2ID ]; then
  echo "Running on AWS, attempting to fetch commit";
  
  REGION=$(curl -H "X-aws-ec2-metadata-token: $TOKEN" -s -m 5 http://169.254.169.254/latest/meta-data/placement/region)
  sudo snap install aws-cli --classic
  aws configure set default.region $REGION

  GITHUB_SHA=$(aws ec2 describe-tags \
    --filters "Name=resource-id,Values=$EC2ID" "Name=key,Values=GIT_SHA" \
    --query "Tags[0].Value" \
    --output text)
    
  GIT_REPO=$(aws ec2 describe-tags \
    --filters "Name=resource-id,Values=$EC2ID" "Name=key,Values=GIT_REPO" \
    --query "Tags[0].Value" \
    --output text)

  # Debug output
  echo XXXXXXXXXXXXXXX
  echo $GIT_REPO
  echo $GITHUB_SHA
  echo XXXXXXXXXXXXXXX

  # fetch correct git commit
  mkdir -p $GIT_REPO
  cd $GIT_REPO

  git init && git remote add origin $GIT_REPO && git fetch origin $GITHUB_SHA && git checkout $GITHUB_SHA
fi

echo "Running on"
echo $(uname -a)
if [ ! -f build.sh ]; then
  echo "Git repo not propperly initialized"
else
  build
fi

if [ "$EC2ID" ]; then
  echo "sudo shutdown -h now"
fi
