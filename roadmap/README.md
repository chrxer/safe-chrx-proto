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
  - [x] EC2
    - [x] fix Ccache no hits 
- [x] fix compillation error (maybe related to. [use_custom_libcxx=false](https://github.com/chrxer/safe-chrx-proto/blob/b7d8b4ddf8c3c6e3dd099d61267cae7d9cf5cfb4/scripts/build.sh#L34) [crbugs#41455655#comment313](https://issues.chromium.org/issues/41455655#comment313) [diff](https://chromium.googlesource.com/chromium/src/+/2e14a3ac178ee87aa9154e5a15dcd986af1b6059%5E%21/#F0)):

  <details>

  ```
  3037/51480] CXX obj/third_party/dawn/src/dawn/common/common/StringViewUtils.o
  FAILED: obj/third_party/dawn/src/dawn/common/common/StringViewUtils.o 
  env CCACHE_SLOPPINESS=time_macros CCACHE_NOHASHDIR=1 CCACHE_LOGFILE=/tmp/ccache_log.log ccache ../../third_party/llvm-build/Release+Asserts/bin/clang++ -MMD -MF obj/third_party/dawn/src/dawn/common/common/StringViewUtils.o.d -DUSE_UDEV -DUSE_AURA=1 -DUSE_GLIB=1 -DUSE_OZONE=1 -DOFFICIAL_BUILD -D__STDC_CONSTANT_MACROS -D__STDC_FORMAT_MACROS -D_FORTIFY_SOURCE=2 -D_FILE_OFFSET_BITS=64 -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -DNO_UNWIND_TABLES -D_GNU_SOURCE -DCR_CLANG_REVISION=\"llvmorg-20-init-6794-g3dbd929e-1\" -D_LIBCPP_HARDENING_MODE=_LIBCPP_HARDENING_MODE_NONE -D_GLIBCXX_ASSERTIONS=1 -DCR_SYSROOT_KEY=20230611T210420Z-2 -DNDEBUG -DNVALGRIND -DDYNAMIC_ANNOTATIONS_ENABLED=0 -DDAWN_ABORT_ON_ASSERT -DDAWN_ENABLE_BACKEND_NULL -DDAWN_ENABLE_BACKEND_OPENGL -DDAWN_ENABLE_BACKEND_DESKTOP_GL -DDAWN_ENABLE_BACKEND_OPENGLES -DDAWN_ENABLE_BACKEND_VULKAN -DDAWN_USE_X11 -DVK_USE_PLATFORM_XCB_KHR -DVK_USE_PLATFORM_WAYLAND_KHR -I../.. -Igen -Igen/third_party/dawn/src -I../../third_party/dawn/src -Igen/third_party/dawn/include -I../../third_party/dawn/include -I../../third_party/abseil-cpp -I../../base/allocator/partition_allocator/src -Igen/base/allocator/partition_allocator/src -I../../third_party/dawn -I../../third_party/vulkan-headers/src/include -I../../third_party/wayland/src/src -I../../third_party/wayland/include/src -Wall -Wextra -Wimplicit-fallthrough -Wextra-semi -Wunreachable-code-aggressive -Wthread-safety -Wno-missing-field-initializers -Wno-unused-parameter -Wno-psabi -Wloop-analysis -Wno-unneeded-internal-declaration -Wno-cast-function-type -Wno-thread-safety-reference-return -Wshadow -fno-delete-null-pointer-checks -fno-ident -fno-strict-aliasing -fstack-protector -fno-unwind-tables -fno-asynchronous-unwind-tables -fPIC -pthread -fcolor-diagnostics -fmerge-all-constants -fno-sized-deallocation -fcrash-diagnostics-dir=../../tools/clang/crashreports -mllvm -instcombine-lower-dbg-declare=0 -mllvm -split-threshold-for-reg-with-hint=0 -ffp-contract=off -flto=thin -fsplit-lto-unit -mllvm -inlinehint-threshold=360 -fwhole-program-vtables -m64 -msse3 -ffile-compilation-dir=. -no-canonical-prefixes -ftrivial-auto-var-init=pattern -O2 -fdata-sections -ffunction-sections -fno-unique-section-names -fno-math-errno -fno-omit-frame-pointer -g0 -fvisibility=hidden -Wheader-hygiene -Wstring-conversion -Wtautological-overlap-compare -Wno-redundant-parens -Wno-invalid-offsetof -Wenum-compare-conditional -Wno-c++11-narrowing-const-reference -Wno-missing-template-arg-list-after-template-kw -Wno-dangling-assignment-gsl -std=c++20 -Wno-trigraphs -gsimple-template-names -fno-exceptions -fno-rtti --sysroot=../../build/linux/debian_bullseye_amd64-sysroot -fvisibility-inlines-hidden -c ../../third_party/dawn/src/dawn/common/StringViewUtils.cpp -o obj/third_party/dawn/src/dawn/common/common/StringViewUtils.o
  ../../third_party/dawn/src/dawn/common/StringViewUtils.cpp:51:21: error: no member named 'strlen' in namespace 'std'
    51 |     return {s, std::strlen(s)};
        |                ~~~~~^
  ```
  fix: commented `use_custom_libcxx=false`

  </details>

- [ ] setup publishing releases using fine-grained tokens
- [ ] ensure EC2s are automatically terminated after build or n hours
- [ ] start with modification of [os_crypt](https://source.chromium.org/search?q=(EncryptString%20OR%20DecryptString)%20AND%20file:os_crypt_%20-unittest%20-browsertest&ss=chromium%2Fchromium%2Fsrc)


## Optimization 
- [ ] move to [hetzner.com/cloud](https://www.hetzner.com/cloud/) using [setup-hcloud](https://github.com/hetznercloud/setup-hcloud) github action ([cli docs](https://github.com/hetznercloud/cli))
- [ ] Use [OICD](https://github.com/aws-actions/configure-aws-credentials?tab=readme-ov-file#oidc) for authenticating to ~~AWS~~ Hetzner
- [ ] Use S3 for caching APT packages
- [ ] automatically choose instance and build timeout, based on whether Ccache exists in ~~S3~~ [storage-box](https://www.hetzner.com/storage/storage-box/)
- [ ] use [git-semantic-version](https://github.com/marketplace/actions/git-semantic-version) (`$VERSION.$SEMVER`) instead of [`$VERSION/$GITHUB_SHA`](https://github.com/chrxer/safe-chrx-proto/blob/b6df1b6855c1f2ca52625ff126c3ebc6c117ee84/entrypoint.sh#L94)
- [ ] pack based on the official [chromium](https://salsa.debian.org/chromium-team/chromium/) debian source (patch `.gn` files using [gn_ast.py](https://chromium.googlesource.com/chromium/src/+/refs/heads/main/build/gn_ast/gn_ast.py))
- [ ] automatic local configuration for [deb_startup](https://github.com/chrxer/safe-chrx-proto/blob/main/deb_startup.md)
