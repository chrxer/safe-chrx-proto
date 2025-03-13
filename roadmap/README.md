## Priority
- [x] setup EC2 instance startup
- [x] setup S3 Ccache & log upload
- [x] first ok compilation start//process of chromium
- [x] specify specific chromium version [chromiumdash.appspot.com/releases](https://chromiumdash.appspot.com/releases?platform=Linux)
- [x] fix ccache location
- [x] first succesfull compilation
- [x] setup Ccache
- [x] move from bash to python
  - [ ] also for deps.sh
- [x] pass `os_crypt_unittest.cc` with external encryption server
- [ ] implement patching of [os_crypt](https://source.chromium.org/search?q=(EncryptString%20OR%20DecryptString)%20AND%20file:os_crypt_%20-unittest%20-browsertest&ss=chromium%2Fchromium%2Fsrc) over [py-tree-sitter](https://github.com/tree-sitter/py-tree-sitter) ( or maybe [mergiraf](https://mergiraf.org/) ? )

## Optimization
- [ ] Implement remote build over [hetzner.com/cloud](https://www.hetzner.com/cloud/) using [setup-hcloud](https://github.com/hetznercloud/setup-hcloud) github action ([cli docs](https://github.com/hetznercloud/cli)
  - [ ] use [git-semantic-version](https://github.com/marketplace/actions/git-semantic-version) (`$VERSION.$SEMVER`) for `libchrx.so`
- [ ] match with [deb package](https://salsa.debian.org/chromium-team/chromium/-/tree/master/debian)
    - [ ] embedd 
    [master prefs](https://salsa.debian.org/chromium-team/chromium/-/blob/master/debian/etc/master_preferences) around [initial_preferences.c](https://source.chromium.org/chromium/chromium/src/+/main:chrome/installer/util/initial_preferences.cc;drc=9be37efad6ba9af197f8cc22921f63a229a3a840;l=188) automatically
    - [ ] compare//match other modifications (patch `.gn` files using `tree-sitter`)
