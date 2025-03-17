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
- [ ] move from `third_party/crashpad/crashpad/third_party/cpp-httplib/` to [SimpleURLLoader](https://github.com/brave/brave-browser/wiki/Simple-Guide-to-SimpleURLLoader) (Guide)

## Optimization
- [ ] Implement remote build over [hetzner.com/cloud](https://www.hetzner.com/cloud/) using [setup-hcloud](https://github.com/hetznercloud/setup-hcloud) github action ([cli docs](https://github.com/hetznercloud/cli)
  - [ ] use [git-semantic-version](https://github.com/marketplace/actions/git-semantic-version) (`$VERSION.$SEMVER`) for `libchrx.so`
- [ ] match with [deb package](https://salsa.debian.org/chromium-team/chromium/-/tree/master/debian)
    - [ ] embedd 
    [master prefs](https://salsa.debian.org/chromium-team/chromium/-/blob/master/debian/etc/master_preferences) around [initial_preferences.c](https://source.chromium.org/chromium/chromium/src/+/main:chrome/installer/util/initial_preferences.cc;drc=9be37efad6ba9af197f8cc22921f63a229a3a840;l=188) automatically
    - [ ] compare//match other modifications (patch `.gn` files using `tree-sitter`)
- [ ] fix failing [`OSCryptTest.Metrics`](https://source.chromium.org/chromium/chromium/src/+/7e3d5c978c6d3a6eda25692cfac7f893a2b20dd0:components/os_crypt/sync/os_crypt_unittest.cc;l=156-201)

<details>

```
Running test: /home/user/VSCProjects/chrxer/safe-chrx-proto/chromium/src/out/Debug/bin/run_components_unittests '--gtest_filter=OSCryptTest.*:OSCryptTestWin.*:OSCryptConcurrencyTest.*:*/OSCryptTest.*/*:*/OSCryptTestWin.*/*:*/OSCryptConcurrencyTest.*/*:*/OSCryptTest/*.*:*/OSCryptTestWin/*.*:*/OSCryptConcurrencyTest/*.*:OSCryptTest.*/*:OSCryptTestWin.*/*:OSCryptConcurrencyTest.*/*:OSCryptTest/*.*:OSCryptTestWin/*.*:OSCryptConcurrencyTest/*.*' --fast-local-dev

X.Org X Server 1.21.1.7
X Protocol Version 11, Revision 0
Current Operating System: Linux debian 6.1.0-22-amd64 #1 SMP PREEMPT_DYNAMIC Debian 6.1.94-1 (2024-06-21) x86_64
Kernel command line: BOOT_IMAGE=/live/vmlinuz-6.1.0-22-amd64 boot=live persistence components quiet splash findiso=
xorg-server 2:21.1.7-3+deb12u8 (https://www.debian.org/support) 
Current version of pixman: 0.42.2
        Before reporting problems, check http://wiki.x.org
        to make sure that you have the latest version.
Markers: (--) probed, (**) from config file, (==) default setting,
        (++) from command line, (!!) notice, (II) informational,
        (WW) warning, (EE) error, (NI) not implemented, (??) unknown.
(==) Log file: "/home/user/.local/share/xorg/Xorg.119.log", Time: Fri Mar 14 15:56:46 2025
(++) Using config file: "/tmp/xorg-5f4f0099dbb1476088b085ce7d269b62.config"
(==) Using system config directory "/usr/share/X11/xorg.conf.d"
(EE) wacom: Wacom HID 526B Pen: Error opening /dev/input/event5 (Permission denied)
(EE) wacom: Wacom HID 526B Finger: Error opening /dev/input/event6 (Permission denied)
Openbox-Message: Unable to find a valid menu file "/var/lib/openbox/debian-menu.xml"
Additional test environment:
    CHROME_HEADLESS=1
    LANG=en_US.UTF-8
Command: out/Debug/components_unittests --test-launcher-bot-mode --gtest_filter=OSCryptTest.*:OSCryptTestWin.*:OSCryptConcurrencyTest.*:*/OSCryptTest.*/*:*/OSCryptTestWin.*/*:*/OSCryptConcurrencyTest.*/*:*/OSCryptTest/*.*:*/OSCryptTestWin/*.*:*/OSCryptConcurrencyTest/*.*:OSCryptTest.*/*:OSCryptTestWin.*/*:OSCryptConcurrencyTest.*/*:OSCryptTest/*.*:OSCryptTestWin/*.*:OSCryptConcurrencyTest/*.* --fast-local-dev

IMPORTANT DEBUGGING NOTE: batches of tests are run inside their
own process. For debugging a test inside a debugger, use the
--gtest_filter=<your_test_name> flag along with
--single-process-tests.
Using sharding settings from environment. This is shard 0/1
Using 8 parallel jobs.
[0314/155649.551906:INFO:test_launcher.cc(1213)] Starting [OSCryptTest.String16EncryptionDecryption, OSCryptTest.EncryptionDecryption, OSCryptTest.CypherTextDiffers, OSCryptTest.DecryptError, OSCryptTest.Metrics, OSCryptConcurrencyTest.ConcurrentInitialization]
[1/6] OSCryptTest.String16EncryptionDecryption (5 ms)
[2/6] OSCryptTest.EncryptionDecryption (2 ms)
[3/6] OSCryptTest.CypherTextDiffers (2 ms)
[4/6] OSCryptTest.DecryptError (1 ms)
[ RUN      ] OSCryptTest.Metrics
../../base/test/metrics/histogram_tester.cc:57: Failure
Expected equality of these values:
  0
  expected_bucket_count
    Which is: 1
Zero samples found for Histogram "OSCrypt.EncryptionPrefixVersion".
(expected at TestBody@components/os_crypt/sync/os_crypt_unittest.cc:175)
Stack trace:
#0 0x555f3526783b base::HistogramTester::ExpectUniqueSample() [../../base/test/metrics/histogram_tester.cc:57:5]
#1 0x555f2a3506fb base::HistogramTester::ExpectUniqueSample<>() [../../base/test/metrics/histogram_tester.h:65:5]
#2 0x555f2a34f064 (anonymous namespace)::OSCryptTest_Metrics_Test::TestBody() [../../components/os_crypt/sync/os_crypt_unittest.cc:175:16]
#3 0x555f2fd494f4 testing::internal::HandleSehExceptionsInMethodIfSupported<>() [../../third_party/googletest/src/googletest/src/gtest.cc:2653:10]
#4 0x555f2fd377b0 testing::internal::HandleExceptionsInMethodIfSupported<>() [../../third_party/googletest/src/googletest/src/gtest.cc:2708:12]

../../base/test/metrics/histogram_tester.cc:57: Failure
Expected equality of these values:
  0
  expected_bucket_count
    Which is: 1
Zero samples found for Histogram "OSCrypt.EncryptionPrefixVersion".
(expected at TestBody@components/os_crypt/sync/os_crypt_unittest.cc:196)
Stack trace:
#0 0x555f3526783b base::HistogramTester::ExpectUniqueSample() [../../base/test/metrics/histogram_tester.cc:57:5]
#1 0x555f2a3506fb base::HistogramTester::ExpectUniqueSample<>() [../../base/test/metrics/histogram_tester.h:65:5]
#2 0x555f2a34f2b8 (anonymous namespace)::OSCryptTest_Metrics_Test::TestBody() [../../components/os_crypt/sync/os_crypt_unittest.cc:196:16]
#3 0x555f2fd494f4 testing::internal::HandleSehExceptionsInMethodIfSupported<>() [../../third_party/googletest/src/googletest/src/gtest.cc:2653:10]
#4 0x555f2fd377b0 testing::internal::HandleExceptionsInMethodIfSupported<>() [../../third_party/googletest/src/googletest/src/gtest.cc:2708:12]

[  FAILED  ] OSCryptTest.Metrics (727 ms)
[5/6] OSCryptTest.Metrics (727 ms)
[6/6] OSCryptConcurrencyTest.ConcurrentInitialization (2 ms)
Retrying 1 test (retry #0)
[0314/155651.979653:INFO:test_launcher.cc(1213)] Starting [OSCryptTest.Metrics]
[ RUN      ] OSCryptTest.Metrics
../../base/test/metrics/histogram_tester.cc:57: Failure
Expected equality of these values:
  0
  expected_bucket_count
    Which is: 1
Zero samples found for Histogram "OSCrypt.EncryptionPrefixVersion".
(expected at TestBody@components/os_crypt/sync/os_crypt_unittest.cc:175)
Stack trace:
#0 0x5617d712583b base::HistogramTester::ExpectUniqueSample() [../../base/test/metrics/histogram_tester.cc:57:5]
#1 0x5617cc20e6fb base::HistogramTester::ExpectUniqueSample<>() [../../base/test/metrics/histogram_tester.h:65:5]
#2 0x5617cc20d064 (anonymous namespace)::OSCryptTest_Metrics_Test::TestBody() [../../components/os_crypt/sync/os_crypt_unittest.cc:175:16]
#3 0x5617d1c074f4 testing::internal::HandleSehExceptionsInMethodIfSupported<>() [../../third_party/googletest/src/googletest/src/gtest.cc:2653:10]
#4 0x5617d1bf57b0 testing::internal::HandleExceptionsInMethodIfSupported<>() [../../third_party/googletest/src/googletest/src/gtest.cc:2708:12]

../../base/test/metrics/histogram_tester.cc:57: Failure
Expected equality of these values:
  0
  expected_bucket_count
    Which is: 1
Zero samples found for Histogram "OSCrypt.EncryptionPrefixVersion".
(expected at TestBody@components/os_crypt/sync/os_crypt_unittest.cc:196)
Stack trace:
#0 0x5617d712583b base::HistogramTester::ExpectUniqueSample() [../../base/test/metrics/histogram_tester.cc:57:5]
#1 0x5617cc20e6fb base::HistogramTester::ExpectUniqueSample<>() [../../base/test/metrics/histogram_tester.h:65:5]
#2 0x5617cc20d2b8 (anonymous namespace)::OSCryptTest_Metrics_Test::TestBody() [../../components/os_crypt/sync/os_crypt_unittest.cc:196:16]
#3 0x5617d1c074f4 testing::internal::HandleSehExceptionsInMethodIfSupported<>() [../../third_party/googletest/src/googletest/src/gtest.cc:2653:10]
#4 0x5617d1bf57b0 testing::internal::HandleExceptionsInMethodIfSupported<>() [../../third_party/googletest/src/googletest/src/gtest.cc:2708:12]

[  FAILED  ] OSCryptTest.Metrics (707 ms)
[7/7] OSCryptTest.Metrics (707 ms)
1 test failed:
    OSCryptTest.Metrics (../../components/os_crypt/sync/os_crypt_unittest.cc:156)
Tests took 6 seconds.
(II) Server terminated successfully (0). Closing log file.
Traceback (most recent call last):
  File "/home/user/VSCProjects/chrxer/safe-chrx-proto/scripts/autotest.py", line 25, in <module>
    test(target=target)
  File "/home/user/VSCProjects/chrxer/safe-chrx-proto/scripts/autotest.py", line 16, in test
    pyexc(str(SRC.joinpath("tools/autotest.py")), "-C" ,OUTR, target, cwd=SRC)
  File "/home/user/VSCProjects/chrxer/safe-chrx-proto/scripts/utils/wrap.py", line 61, in pyexc
    return exc(PYTHON, *cmd, dbg=dbg, _bytes=_bytes, timeout=timeout, cwd=cwd, _pidx=_pidx+1)
           ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
  File "/home/user/VSCProjects/chrxer/safe-chrx-proto/scripts/utils/wrap.py", line 54, in exc
    raise CommandFailed(proc.returncode, hcmd)
utils.wrap.CommandFailed: Command: /home/user/VSCProjects/chrxer/safe-chrx-proto/scripts/.venv/bin/python3 /home/user/VSCProjects/chrxer/safe-chrx-proto/chromium/src/tools/autotest.py -C out/Debug os_crypt_unittest.cc failed with exit code 1
make: *** [Makefile:19: test] Error 1
```

</details>
