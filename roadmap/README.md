## Priority
- [x] setup EC2 instance startup
- [x] setup S3 Ccache & log upload
- [x] first ok compilation start//process of chromium
- [x] specify specific chromium version [chromiumdash.appspot.com/releases](https://chromiumdash.appspot.com/releases?platform=Linux)
- [x] fix ccache location
- [x] first succesfull compilation
- [x] setup Ccache
- [ ] move from bash to python
- [ ] start with patching of [os_crypt](https://source.chromium.org/search?q=(EncryptString%20OR%20DecryptString)%20AND%20file:os_crypt_%20-unittest%20-browsertest&ss=chromium%2Fchromium%2Fsrc) using clang [python bindings](https://source.chromium.org/chromium/chromium/src/+/main:third_party/angle/third_party/llvm/src/clang/bindings/python/)  or [py-tree-sitter](https://github.com/tree-sitter/py-tree-sitter) ( or maybe [mergiraf](https://mergiraf.org/) ? )

## Optimization
- [ ] Remote build
  - [ ] ensure EC2s are automatically terminated after build or n hours
  - [ ] setup publishing releases using fine-grained tokens
  - [ ] move to [hetzner.com/cloud](https://www.hetzner.com/cloud/) using [setup-hcloud](https://github.com/hetznercloud/setup-hcloud) github action ([cli docs](https://github.com/hetznercloud/cli))
  - [ ] Use [OICD](https://github.com/aws-actions/configure-aws-credentials?tab=readme-ov-file#oidc) for authenticating to ~~AWS~~ Hetzner
  - [ ] Use S3 for caching APT packages
  - [ ] automatically choose instance and build timeout, based on whether Ccache exists in ~~S3~~ [storage-box](https://www.hetzner.com/storage/storage-box/)
  - [ ] use [git-semantic-version](https://github.com/marketplace/actions/git-semantic-version) (`$VERSION.$SEMVER`) instead of [`$VERSION/$GITHUB_SHA`](https://github.com/chrxer/safe-chrx-proto/blob/b6df1b6855c1f2ca52625ff126c3ebc6c117ee84/entrypoint.sh#L94)
- [ ] match with [deb package](https://salsa.debian.org/chromium-team/chromium/-/tree/master/debian)
    - [ ] embedd 
    [master prefs](https://salsa.debian.org/chromium-team/chromium/-/blob/master/debian/etc/master_preferences) around [initial_preferences.c](https://source.chromium.org/chromium/chromium/src/+/main:chrome/installer/util/initial_preferences.cc;drc=9be37efad6ba9af197f8cc22921f63a229a3a840;l=188) automatically
    - [ ] compare//match other modifications (patch `.gn` files using [gn_ast.py](https://chromium.googlesource.com/chromium/src/+/refs/heads/main/build/gn_ast/gn_ast.py))
