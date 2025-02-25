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


- [ ] ensure EC2s are automatically terminated after build or n hours

- [ ] start with modification of [os_crypt](https://source.chromium.org/search?q=(EncryptString%20OR%20DecryptString)%20AND%20file:os_crypt_%20-unittest%20-browsertest&ss=chromium%2Fchromium%2Fsrc)


## Optimization 
- [ ] setup publishing releases using fine-grained tokens
- [ ] move to [hetzner.com/cloud](https://www.hetzner.com/cloud/) using [setup-hcloud](https://github.com/hetznercloud/setup-hcloud) github action ([cli docs](https://github.com/hetznercloud/cli))
- [ ] Use [OICD](https://github.com/aws-actions/configure-aws-credentials?tab=readme-ov-file#oidc) for authenticating to ~~AWS~~ Hetzner
- [ ] Use S3 for caching APT packages
- [ ] automatically choose instance and build timeout, based on whether Ccache exists in ~~S3~~ [storage-box](https://www.hetzner.com/storage/storage-box/)
- [ ] use [git-semantic-version](https://github.com/marketplace/actions/git-semantic-version) (`$VERSION.$SEMVER`) instead of [`$VERSION/$GITHUB_SHA`](https://github.com/chrxer/safe-chrx-proto/blob/b6df1b6855c1f2ca52625ff126c3ebc6c117ee84/entrypoint.sh#L94)
- [ ] pack based on the official [chromium](https://salsa.debian.org/chromium-team/chromium/) debian source (patch `.gn` files using [gn_ast.py](https://chromium.googlesource.com/chromium/src/+/refs/heads/main/build/gn_ast/gn_ast.py))
- [ ] automatic local configuration for [deb_startup](https://github.com/chrxer/safe-chrx-proto/blob/main/deb_startup.md)
    - [ ] embedd default prefst around [initial_preferences.c](https://source.chromium.org/chromium/chromium/src/+/main:chrome/installer/util/initial_preferences.cc;drc=9be37efad6ba9af197f8cc22921f63a229a3a840;l=188) automatically


## Local compillation
8GB RAM, 11th Gen Intel(R) Core(TM) i5-1135G7, external 1TB SSD over USB-c

dmesg message
```
Feb 25 02:17:23.238758 debian kernel: oom-kill:constraint=CONSTRAINT_NONE,nodemask=(null),cpuset=/,mems_allowed=0,global_oom,task_memcg=/user.slice/user-1000.slice/user@1000.service/app.slice/app-org.gnome.Term>
Feb 25 02:17:23.238768 debian kernel: Out of memory: Killed process 81816 (clang++) total-vm:794600kB, anon-rss:669776kB, file-rss:1788kB, shmem-rss:0kB, UID:1000 pgtables:1544kB oom_score_adj:200
```

<details>

<summary>Task states at this point</summary>

```
Feb 25 02:17:23.237466 debian kernel: Tasks state (memory values in pages):
Feb 25 02:17:23.237480 debian kernel: [  pid  ]   uid  tgid total_vm      rss pgtables_bytes swapents oom_score_adj name
Feb 25 02:17:23.237489 debian kernel: [   1079]     0  1079    16624     1703   114688        0          -250 systemd-journal
Feb 25 02:17:23.237496 debian kernel: [   1107]     0  1107     7286     1675    81920        0         -1000 systemd-udevd
Feb 25 02:17:23.237503 debian kernel: [   2487]     0  2487    60182     1652   106496        0             0 accounts-daemon
Feb 25 02:17:23.237512 debian kernel: [   2489]   107  2489     1768      555    45056        0             0 avahi-daemon
Feb 25 02:17:23.237519 debian kernel: [   2491]     0  2491     3240      708    61440        0             0 bluetoothd
Feb 25 02:17:23.237528 debian kernel: [   2493]     0  2493     2461      531    61440        0             0 cron
Feb 25 02:17:23.237538 debian kernel: [   2495]   103  2495     2663      906    61440        0          -900 dbus-daemon
Feb 25 02:17:23.237545 debian kernel: [   2497]     0  2497    57685     1364    81920        0             0 iio-sensor-prox
Feb 25 02:17:23.237554 debian kernel: [   2501]     0  2501    57669     2620    81920        0             0 low-memory-moni
Feb 25 02:17:23.237561 debian kernel: [   2502]   997  2502    77554     2709   102400        0             0 polkitd
Feb 25 02:17:23.237570 debian kernel: [   2503]     0  2503     2829      464    57344        0             0 smartd
Feb 25 02:17:23.237578 debian kernel: [   2505]     0  2505    59150     1435    86016        0             0 switcheroo-cont
Feb 25 02:17:23.237587 debian kernel: [   2507]   107  2507     1722       78    45056        0             0 avahi-daemon
Feb 25 02:17:23.237595 debian kernel: [   2508]     0  2508    12403     1187    94208        0             0 systemd-logind
Feb 25 02:17:23.237603 debian kernel: [   2509]     0  2509     4157      986    73728        0             0 systemd-machine
Feb 25 02:17:23.237611 debian kernel: [   2510]     0  2510    98948     2851   135168        0             0 udisksd
Feb 25 02:17:23.237621 debian kernel: [   2547]     0  2547    65601     2869   143360        0             0 NetworkManager
Feb 25 02:17:23.237629 debian kernel: [   2577]     0  2577     4352     1835    73728        0             0 wpa_supplicant
Feb 25 02:17:23.237637 debian kernel: [   2578]     0  2578    60897     1684   102400        0             0 ModemManager
Feb 25 02:17:23.237646 debian kernel: [   2638]     0  2638    60871     1661   106496        0             0 gdm3
Feb 25 02:17:23.237657 debian kernel: [   2679]     0  2679    42036     1661    98304        0             0 gdm-session-wor
Feb 25 02:17:23.237666 debian kernel: [   2713]  1000  2713     5040     1366    77824        0           100 systemd
Feb 25 02:17:23.237674 debian kernel: [   2714]  1000  2714    42403     1046    94208        0           100 (sd-pam)
Feb 25 02:17:23.237683 debian kernel: [   2859]  1000  2859    16966     2493   114688        0           200 pipewire
Feb 25 02:17:23.237691 debian kernel: [   2860]  1000  2860   156411    16169   348160        0           200 tracker-extract
Feb 25 02:17:23.237699 debian kernel: [   2861]  1000  2861    64912     2302   131072        0           200 wireplumber
Feb 25 02:17:23.237708 debian kernel: [   2862]  1000  2862    10016     2886    98304        0           200 pipewire-pulse
Feb 25 02:17:23.237718 debian kernel: [   2864]  1000  2864    97682     2061   118784        0           200 gnome-keyring-d
Feb 25 02:17:23.237725 debian kernel: [   2866] 65534  2866     2801      133    65536        0             0 dnsmasq
Feb 25 02:17:23.237742 debian kernel: [   2867]     0  2867     2801      131    65536        0             0 dnsmasq
Feb 25 02:17:23.237750 debian kernel: [   2888]  1000  2888     2287      717    57344        0           200 dbus-daemon
Feb 25 02:17:23.237759 debian kernel: [   2910]   112  2910    22059       66    57344        0             0 rtkit-daemon
Feb 25 02:17:23.237767 debian kernel: [   2913]  1000  2913    40629     1328    81920        0             0 gdm-wayland-ses
Feb 25 02:17:23.237775 debian kernel: [   2918]  1000  2918    75307     2243   151552        0             0 gnome-session-b
Feb 25 02:17:23.237781 debian kernel: [   2927]  1000  2927    60187     1953    98304        0           200 gvfsd
Feb 25 02:17:23.237789 debian kernel: [   2951]  1000  2951    95093     1879   102400        0           200 gvfsd-fuse
Feb 25 02:17:23.237797 debian kernel: [   3009]  1000  3009    22095      942    69632        0           200 gcr-ssh-agent
Feb 25 02:17:23.237806 debian kernel: [   3010]  1000  3010    23016      955    77824        0           200 gnome-session-c
Feb 25 02:17:23.237816 debian kernel: [   3011]  1000  3011     1955     1118    53248        0           200 ssh-agent
Feb 25 02:17:23.237825 debian kernel: [   3019]  1000  3019   149279     2484   180224        0           200 gnome-session-b
Feb 25 02:17:23.237834 debian kernel: [   3038]  1000  3038    77810     2018   106496        0           200 at-spi-bus-laun
Feb 25 02:17:23.237840 debian kernel: [   3042]  1000  3042  1354428    42514  2248704        0           100 gnome-shell
Feb 25 02:17:23.237848 debian kernel: [   3048]  1000  3048     1976      513    57344        0           200 dbus-daemon
Feb 25 02:17:23.237855 debian kernel: [   3060]  1000  3060    59985     1559   102400        0           200 xdg-permission-
Feb 25 02:17:23.237863 debian kernel: [   3183]   117  3183     2676     2672    57344        0             0 ntpd
Feb 25 02:17:23.237871 debian kernel: [   3215]  1000  3215   145308     2426   196608        0           200 gnome-shell-cal
Feb 25 02:17:23.237880 debian kernel: [   3218]     0  3218    59315     1546    90112        0             0 upowerd
Feb 25 02:17:23.237887 debian kernel: [   3226]  1000  3226   303979     4463   389120        0           200 evolution-sourc
Feb 25 02:17:23.237895 debian kernel: [   3240]   111  3240   238510     4404   266240        0             0 geoclue
Feb 25 02:17:23.237904 debian kernel: [   3244]  1000  3244    88768     1880   139264        0           200 gvfs-udisks2-vo
Feb 25 02:17:23.237912 debian kernel: [   3252]  1000  3252    59180     1223    90112        0           200 gvfs-mtp-volume
Feb 25 02:17:23.237920 debian kernel: [   3256]  1000  3256    59452     1225    90112        0           200 gvfs-gphoto2-vo
Feb 25 02:17:23.237927 debian kernel: [   3271]  1000  3271    59189     1351    86016        0           200 gvfs-goa-volume
Feb 25 02:17:23.237935 debian kernel: [   3279]  1000  3279   150365     3098   307200        0           200 goa-daemon
Feb 25 02:17:23.237943 debian kernel: [   3526]  1000  3526    78952     1346   106496        0           200 goa-identity-se
Feb 25 02:17:23.237952 debian kernel: [   3532]  1000  3532    78915     1830   106496        0           200 gvfs-afc-volume
Feb 25 02:17:23.237960 debian kernel: [   3534]    33  3534     1061      528    49152        0             0 lighttpd
Feb 25 02:17:23.237967 debian kernel: [   3551]   101  3551     7090     3012    98304        0             0 exim4
Feb 25 02:17:23.237975 debian kernel: [   3554]  1000  3554   224949     2577   241664        0           200 evolution-calen
Feb 25 02:17:23.237983 debian kernel: [   3562]  1000  3562   696928     3408   253952        0           200 gjs
Feb 25 02:17:23.237991 debian kernel: [   3564]  1000  3564    41080     1887    90112        0           200 at-spi2-registr
Feb 25 02:17:23.237998 debian kernel: [   3565]   113  3565    61444     2323   102400        0             0 colord
Feb 25 02:17:23.238005 debian kernel: [   3582]  1000  3582      644      202    45056        0           200 sh
Feb 25 02:17:23.238013 debian kernel: [   3583]  1000  3583    77652     1905    98304        0           200 gsd-a11y-settin
Feb 25 02:17:23.238020 debian kernel: [   3586]  1000  3586    85777     3295   159744        0           200 gsd-color
Feb 25 02:17:23.238029 debian kernel: [   3587]  1000  3587    89625     1592   139264        0           200 gsd-datetime
Feb 25 02:17:23.238039 debian kernel: [   3588]  1000  3588    79192     2936   114688        0           200 ibus-daemon
Feb 25 02:17:23.238046 debian kernel: [   3590]  1000  3590    96667     1510   114688        0           200 gsd-housekeepin
Feb 25 02:17:23.238055 debian kernel: [   3591]  1000  3591    85489     2439   159744        0           200 gsd-keyboard
Feb 25 02:17:23.238063 debian kernel: [   3592]  1000  3592   130992     2543   212992        0           200 gsd-media-keys
Feb 25 02:17:23.238075 debian kernel: [   3593]  1000  3593   130697     3579   200704        0           200 gsd-power
Feb 25 02:17:23.238083 debian kernel: [   3596]  1000  3596   249870     5229   462848        0           200 evolution-alarm
Feb 25 02:17:23.238092 debian kernel: [   3599]  1000  3599    62611     1511   118784        0           200 gsd-print-notif
Feb 25 02:17:23.238099 debian kernel: [   3602]  1000  3602   114499     1873   110592        0           200 gsd-rfkill
Feb 25 02:17:23.238108 debian kernel: [   3603]  1000  3603    59103     1412    90112        0           200 gsd-screensaver
Feb 25 02:17:23.238116 debian kernel: [   3604]  1000  3604   116629     2016   131072        0           200 gsd-sharing
Feb 25 02:17:23.238124 debian kernel: [   3606]  1000  3606    78806     2032   106496        0           200 gsd-smartcard
Feb 25 02:17:23.238131 debian kernel: [   3607]  1000  3607    80656     2008   131072        0           200 gsd-sound
Feb 25 02:17:23.238141 debian kernel: [   3613]  1000  3613    96353     1406   110592        0           200 gsd-usb-protect
Feb 25 02:17:23.238150 debian kernel: [   3614]  1000  3614    85640     2605   163840        0           200 gsd-wacom
Feb 25 02:17:23.238157 debian kernel: [   3633]  1000  3633    57946     1490    77824        0           200 gsd-disk-utilit
Feb 25 02:17:23.238167 debian kernel: [   3643]  1000  3643   287654    11570   401408        0           200 gnome-software
Feb 25 02:17:23.238173 debian kernel: [   3734]  1000  3734   217492    12216   552960        0           100 Xwayland
Feb 25 02:17:23.238180 debian kernel: [   3747]  1000  3747   698991     3087   245760        0           200 gjs
Feb 25 02:17:23.238189 debian kernel: [   3754]  1000  3754    59364     1403    90112        0           200 ibus-dconf
Feb 25 02:17:23.238197 debian kernel: [   3765]  1000  3765    86968     4357   184320        0           200 ibus-extension-
Feb 25 02:17:23.238205 debian kernel: [   3781]  1000  3781    59349     1237    86016        0           200 ibus-portal
Feb 25 02:17:23.238214 debian kernel: [   3814]  1000  3814   168305     2609   241664        0           200 evolution-addre
Feb 25 02:17:23.238221 debian kernel: [   3827]  1000  3827    86286     1503   159744        0           200 gsd-printer
Feb 25 02:17:23.238230 debian kernel: [   3908]  1000  3908   221772     1760   192512        0           200 xdg-desktop-por
Feb 25 02:17:23.238238 debian kernel: [   3909]  1000  3909    40898     1244    77824        0           200 ibus-engine-sim
Feb 25 02:17:23.238247 debian kernel: [   3926]  1000  3926   187577     1429   155648        0           200 xdg-document-po
Feb 25 02:17:23.238253 debian kernel: [   3931]  1000  3931      620      236    45056        0           200 fusermount3
Feb 25 02:17:23.238259 debian kernel: [   3939]  1000  3939   168182     4640   229376        0           200 xdg-desktop-por
Feb 25 02:17:23.238266 debian kernel: [   4018]  1000  4018   268928     6016   532480        0           200 gsd-xsettings
Feb 25 02:17:23.238273 debian kernel: [   4028]  1000  4028    39204     1931    69632        0           200 dconf-service
Feb 25 02:17:23.238281 debian kernel: [   4091]  1000  4091    87133     2776   176128        0           200 xdg-desktop-por
Feb 25 02:17:23.238289 debian kernel: [   4115]  1000  4115    48364     2395   139264        0           200 ibus-x11
Feb 25 02:17:23.238297 debian kernel: [   4117]  1000  4117   157275     3617   294912        0           200 tracker-miner-f
Feb 25 02:17:23.238305 debian kernel: [   4165]  1000  4165    40811     1402    77824        0           200 gvfsd-metadata
Feb 25 02:17:23.238323 debian kernel: [   5690]  1000  5690   167336    13989   360448        0           200 gnome-terminal-
Feb 25 02:17:23.238330 debian kernel: [   5716]  1000  5716     3598     1601    65536        0           200 bash
Feb 25 02:17:23.238339 debian kernel: [   5940]  1000  5940     3598     1558    61440        0           200 bash
Feb 25 02:17:23.238346 debian kernel: [   6161]  1000  6161     3598     1598    61440        0           200 bash
Feb 25 02:17:23.238354 debian kernel: [   6448]  1000  6448     3182     1234    61440        0           200 htop
Feb 25 02:17:23.238363 debian kernel: [  23585]  1000 23585     2542      490    53248        0           200 test_target.sh
Feb 25 02:17:23.238370 debian kernel: [  23735]  1000 23735     6501     2795    81920        0           200 python3
Feb 25 02:17:23.238376 debian kernel: [  23884]  1000 23884     2542      477    57344        0           200 bash
Feb 25 02:17:23.238384 debian kernel: [  23888]  1000 23888     8704     4653   110592        0           200 python3
Feb 25 02:17:23.238393 debian kernel: [  23938]  1000 23938   102635    99933   864256        0           200 ninja
Feb 25 02:17:23.238402 debian kernel: [  26662]     0 26662     7567      996    90112        0             0 cupsd
Feb 25 02:17:23.238410 debian kernel: [  26670]     7 26670     4090      368    73728        0             0 dbus
Feb 25 02:17:23.238419 debian kernel: [  26671]     0 26671    44056     1737   114688        0             0 cups-browsed
Feb 25 02:17:23.238429 debian kernel: [  81766]  1000 81766      646      199    45056        0           200 sh
Feb 25 02:17:23.238437 debian kernel: [  81767]  1000 81767     2079      659    65536        0           200 ccache
Feb 25 02:17:23.238445 debian kernel: [  81771]  1000 81771      646      207    45056        0           200 sh
Feb 25 02:17:23.238454 debian kernel: [  81772]  1000 81772     2102      672    57344        0           200 ccache
Feb 25 02:17:23.238462 debian kernel: [  81774]  1000 81774      646      211    45056        0           200 sh
Feb 25 02:17:23.238471 debian kernel: [  81775]  1000 81775     2126      695    53248        0           200 ccache
Feb 25 02:17:23.238479 debian kernel: [  81777]  1000 81777      646      204    49152        0           200 sh
Feb 25 02:17:23.238487 debian kernel: [  81778]  1000 81778     2102      654    61440        0           200 ccache
Feb 25 02:17:23.238496 debian kernel: [  81816]  1000 81816   198650   167891  1581056        0           200 clang++
Feb 25 02:17:23.238505 debian kernel: [  81933]  1000 81933   184846   153952  1490944        0           200 clang++
Feb 25 02:17:23.238512 debian kernel: [  81943]  1000 81943   190150   159250  1531904        0           200 clang++
Feb 25 02:17:23.238521 debian kernel: [  81953]  1000 81953   197687   166964  1593344        0           200 clang++
Feb 25 02:17:23.238530 debian kernel: [  82390]  1000 82390      646      205    45056        0           200 sh
Feb 25 02:17:23.238539 debian kernel: [  82392]  1000 82392     2102      677    49152        0           200 ccache
Feb 25 02:17:23.238546 debian kernel: [  82395]  1000 82395      646      205    40960        0           200 sh
Feb 25 02:17:23.238552 debian kernel: [  82396]  1000 82396     2098      691    53248        0           200 ccache
Feb 25 02:17:23.238560 debian kernel: [  82397]  1000 82397      646      202    45056        0           200 sh
Feb 25 02:17:23.238568 debian kernel: [  82398]  1000 82398     2093      663    61440        0           200 ccache
Feb 25 02:17:23.238576 debian kernel: [  82405]  1000 82405      646      206    40960        0           200 sh
Feb 25 02:17:23.238583 debian kernel: [  82406]  1000 82406     2102      674    49152        0           200 ccache
Feb 25 02:17:23.238591 debian kernel: [  82408]  1000 82408      646      208    45056        0           200 sh
Feb 25 02:17:23.238599 debian kernel: [  82409]  1000 82409     2098      679    53248        0           200 ccache
Feb 25 02:17:23.238608 debian kernel: [  82411]  1000 82411   160273   128376  1273856        0           200 clang++
Feb 25 02:17:23.238616 debian kernel: [  82412]  1000 82412   159324   127046  1269760        0           200 clang++
Feb 25 02:17:23.238624 debian kernel: [  82413]  1000 82413   147646   115498  1167360        0           200 clang++
Feb 25 02:17:23.238632 debian kernel: [  82414]  1000 82414   151299   121945  1220608        0           200 clang++
Feb 25 02:17:23.238632 debian kernel: [  82414]  1000 82414   151299   121945  1220608        0           200 clang++
Feb 25 02:17:23.238640 debian kernel: [  82415]  1000 82415   138106   110582  1114112        0           200 clang++
Feb 25 02:17:23.238649 debian kernel: [  82416]  1000 82416      646      208    45056        0           200 sh
Feb 25 02:17:23.238658 debian kernel: [  82417]  1000 82417     2058      660    57344        0           200 ccache
Feb 25 02:17:23.238666 debian kernel: [  82419]  1000 82419    45211    20040   356352        0           200 clang++
Feb 25 02:17:23.238672 debian kernel: [  82430]     0 82430     2947      634    61440        0             0 cron
Feb 25 02:17:23.238681 debian kernel: [  82453]     0 82453     2947      574    61440        0             0 cron
Feb 25 02:17:23.238687 debian kernel: [  82455]     0 82455     2947      575    61440        0             0 cron
Feb 25 02:17:23.238696 debian kernel: [  82469]     0 82469     2974      438    61440        0             0 cron
Feb 25 02:17:23.238705 debian kernel: [  82472]   109 82472     2149       67    49152        0             0 fwupdmgr
Feb 25 02:17:23.238714 debian kernel: [  82491]     0 82491     2576      446    57344        0             0 cron
Feb 25 02:17:23.238720 debian kernel: [  82523]     0 82523     2568      405    57344        0             0 cron
Feb 25 02:17:23.238727 debian kernel: [  82548]     0 82548     2547      407    57344        0             0 cron
Feb 25 02:17:23.238737 debian kernel: [  82561]     0 82561    42613     1780   102400        0             0 (boltd)
Feb 25 02:17:23.238743 debian kernel: [  82562]     0 82562     2574      384    57344        0             0 cron
Feb 25 02:17:23.238750 debian kernel: [  82569]     0 82569   220102      687   135168        0          -900 snapd
```
 
</details>

```bash
sudo fallocate -l 10G -x /swapfile
```
