#!/bin/bash

set -e
exec > >(tee /tmp/build.log) 2>&1

sudo apt-get update

GITHUB_SHA=$(aws ec2 describe-tags \
  --filters "Name=resource-id,Values=$INSTANCE_ID" "Name=key,Values=GIT_SHA" \
  --query "Tags[0].Value" \
  --output text)
  
GIT_REPO=$(aws ec2 describe-tags \
  --filters "Name=resource-id,Values=$INSTANCE_ID" "Name=key,Values=GIT_REPO" \
  --query "Tags[0].Value" \
  --output text)

echo XXXXXXXXXXXXXXX
echo $GIT_REPO
echo $GIT_SHA
echo uname -a
echo XXXXXXXXXXXXXXX

mkdir -p $GIT_REPO
cd $GIT_REPO

git init
git remote add origin https://github.com/$GIT_REPO
git fetch origin $GIT_SHA
git checkout $GIT_SHA

cat LICENSE

echo XXXXXXXXXXX

sudo shutdown -h now
