## Setup build instance
Building requires running on a debian-based distro (ubuntu recommended)

### Setup
Fork this repository

Create a IAM policy in the [AWS console](https://console.aws.amazon.com/iamv2/home#/policies) with the following minimum permissions
```json
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "EC2RunChrxer",
            "Effect": "Allow",
            "Action": [
                "ec2:DescribeTags",
                "ec2:CreateTags",
                "ec2:RunInstances",
                "ssm:GetParameters",
                "ec2:StopInstances",
                "ec2:TerminateInstances",
                "ec2:DescribeInstances",
                "ec2:DescribeInstanceStatus",
                "ec2:DescribeSecurityGroups"
            ],
            "Resource": "*"
        },
        {
            "Sid": "PassRoleChrxer",
            "Effect": "Allow",
            "Action": "iam:PassRole",
            "Resource": "arn:aws:iam::*:role/chrxer"
        },
        {
            "Sid": "AllowS3Chrxer",
            "Effect": "Allow",
            "Action": [
                "s3:PutObject",
                "s3:GetObject",
                "s3:AbortMultipartUpload",
                "s3:ListBucket",
                "s3:DeleteObject",
                "s3:GetObjectVersion",
                "s3:ListMultipartUploadParts",
                "s3:CreateBucket",
                "s3:DeleteBucket"
            ],
            "Resource": [
                "arn:aws:s3:::amzn-s3-chrxer-bucket-v*/*",
                "arn:aws:s3:::amzn-s3-chrxer-bucket-v*"
            ]
        }
    ]
}
```
Create an EC2 [IAM role](https://console.aws.amazon.com/iamv2/home#/roles) named `chrxer` with the same policy.

Create an [IAM user](https://console.aws.amazon.com/iamv2/home#/users) named with the same policy. \
Add the repository secrets for `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY`

Go to [EC2 security group](https://console.aws.amazon.com/ec2/home#SecurityGroups:) and create a new security group named `chrxer` for the (in the [workflow](../.github/workflows/build.yml)) specified region. \
Add an inbound rule with `Type:ssh` and select `com.amazonaws.<your-region>.ec2-instance-connect` for `Source` 

### Debugging
You can use [ec2-instance-connect](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-connect-methods.html) for connecting to the instance.

Enter build directory
```bash
cd /data
```

You may uses
```bash
tail -f -n +1 /tmp/build.log
```
for retrieving the output of [build.sh](../build.sh)

Free space on SSD
```bash
df -h /data
```

Size stats for chromium source

### Retrieve log from S3
```bash
aws s3 cp --quiet s3://amzn-s3-chrxer-bucket-v1/build.log /dev/stdout
```

### S3 bucket naming
The bucket name must be in the form `amzn-s3-chrxer-bucket-v*`.
