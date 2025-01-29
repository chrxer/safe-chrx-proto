## Setup build instance


### Setup
Fork this repository

Create a IAM policy in the [AWS console](https://console.aws.amazon.com/iamv2/home#/policies) with the following minimum permissions
```json
{s
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
				"ec2:DescribeInstanceStatus"
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
				"s3:ListMultipartUploadParts"
			],
			"Resource": [
				"arn:aws:s3:::s3-chrxer/*",
				"arn:aws:s3:::s3-chrxer"
			]
		}
	]
}
```
Create an EC2 [IAM role](https://console.aws.amazon.com/iamv2/home#/roles) named `chrxer` with the same policy.

Create an [IAM user](https://console.aws.amazon.com/iamv2/home#/users) named with the same policy. \
Add the repository secrets for `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY`

