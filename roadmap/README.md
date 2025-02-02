## Priority
- [x] setup EC2 instance startup
- [x] setup S3 Ccache & log upload
- [x] first ok compillation start//process of chromium
- [ ] specify specific chromium version
```bash
 df -h /home/ubuntu/.cache/ccache/
Filesystem      Size  Used Avail Use% Mounted on
/dev/root       6.8G  5.8G  987M  86% /
```
- [ ] fix ccache location (use disc, not EBS)
- [ ] first succesfull compillation
- [ ] setup?publishing releases using fine-grained tokens
- [ ] ensure EC2s are automatically terminated after n hours

## Optimization 
- [ ] Use [OICD](https://github.com/aws-actions/configure-aws-credentials?tab=readme-ov-file#oidc) for authenticating to AWS
- [ ] Use S3 for caching APT packages
- [ ] automatically choose instance and build timeout, based on whether Ccache exists in S3
