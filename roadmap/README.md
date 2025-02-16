## Priority
- [x] setup EC2 instance startup
- [x] setup S3 Ccache & log upload
- [x] first ok compilation start//process of chromium
- [x] specify specific chromium version [chromiumdash.appspot.com/releases](https://chromiumdash.appspot.com/releases?platform=Linux)
```bash
 df -h /home/ubuntu/.cache/ccache/
Filesystem      Size  Used Avail Use% Mounted on
/dev/root       6.8G  5.8G  987M  86% /
```
- [x] fix ccache location
- [x] first succesfull compilation
- [ ] test ccache speedup
  - [x] local
  - [ ] EC2
- [ ] setup publishing releases using fine-grained tokens
- [ ] ensure EC2s are automatically terminated after build or n hours
- [ ] start with modification of [os_crypt](https://source.chromium.org/search?q=(EncryptString%20OR%20DecryptString)%20AND%20file:os_crypt_%20-unittest%20-browsertest&ss=chromium%2Fchromium%2Fsrc)


## Optimization 
- [ ] Use [OICD](https://github.com/aws-actions/configure-aws-credentials?tab=readme-ov-file#oidc) for authenticating to AWS
- [ ] Use S3 for caching APT packages
- [ ] automatically choose instance and build timeout, based on whether Ccache exists in S3
- [ ] use [git-semantic-version](https://github.com/marketplace/actions/git-semantic-version) (`$VERSION.$SEMVER`) instead of [`$VERSION/$GITHUB_SHA`](https://github.com/chrxer/safe-chrx-proto/blob/b6df1b6855c1f2ca52625ff126c3ebc6c117ee84/entrypoint.sh#L94)
- [ ] pack based on the official [chromium](https://salsa.debian.org/chromium-team/chromium/) debian source (patch `.gn` files using [gn_ast.py](https://chromium.googlesource.com/chromium/src/+/refs/heads/main/build/gn_ast/gn_ast.py))
- [ ] automatic local configuration for [deb_startup](https://github.com/chrxer/safe-chrx-proto/blob/main/deb_startup.md)
