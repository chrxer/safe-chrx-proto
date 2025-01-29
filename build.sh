#!/bin/bash

set +e

build () {
  echo "Building chromium.."
}

# ensure curl is installed
if ! command -v curl &> /dev/null; then
  sudo apt install -y curl
fi

EC2ID=$(curl -s -m 5 http://169.254.169.254/latest/meta-data/instance-id)
if [ $EC2ID ]; then
  echo "Running on AWS, attempting to fetch commit";
  exec > >(tee /tmp/build.log) 2>&1
  sudo apt-get update

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

  git init
  git remote add origin https://github.com/$GIT_REPO
  git fetch origin $GITHUB_SHA
  git checkout $GITHUB_SHA
fi

echo "Running on"
echo $(uname -a)
if [ ! -f build.sh ]; then
  echo "Git repo not propperly initialized"
else
  build
fi

if [ "$EC2ID" ]; then
  sudo shutdown -h now
fi
