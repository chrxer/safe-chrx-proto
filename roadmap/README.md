- [x] setup EC2 instance startup
- [ ] setup S3 cache & logging
- [ ] setup?publishing releases using fine-grained tokens
- [ ] ensure EC2s are automatically terminated after n hours

## Optimization 
- [ ] Use [OICD](https://github.com/aws-actions/configure-aws-credentials?tab=readme-ov-file#oidc) for authenticating to AWS
- [ ] Use S3 for caching APT packages
- [ ] automatically choose instance and build timeout, based on whether Ccache exists in S3
