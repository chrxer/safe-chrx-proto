export GIT_COMMIT=$(aws ec2 describe-tags \
  --filters "Name=resource-id,Values=$INSTANCE_ID" "Name=key,Values=GITHUB_SHA" \
  --query "Tags[0].Value" \
  --output text)

  echo $GIT_COMMIT
  echo uname -a