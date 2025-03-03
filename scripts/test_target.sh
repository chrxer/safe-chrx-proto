#!/bin/bash

set -e
# running inside chrxer repo directory root

WRK=$(realpath $(dirname $(dirname "$0")))
DEPOT="$WRK/depot_tools"
CHROMIUM="$WRK/chromium"
VERSION=$(cat "$WRK/chromium.version")

export PATH="$DEPOT:$PATH"

cd $CHROMIUM/src

# pseudo-parsed args
# supported: 
#   New lines (become " ")
#   Empty lines & lines starting with # (are deleted)
#   regular arguments ('"' is automatically escaped to '\"')
ARGS=$(cat <<EOF
# https://salsa.debian.org/chromium-team/chromium/-/blob/master/debian/rules
clang_use_chrome_plugins=false

# disabled features
#is_debug=false
use_libjpeg_turbo=true

# ./../third_party/dawn/src/dawn/common/StringViewUtils.cpp:51:21: error: no member named 'strlen' in namespace 'std'
# use_custom_libcxx=false

use_unofficial_version_number=false
safe_browsing_use_unrar=false
enable_vr=false
enable_nacl=false
#build_dawn_tests=false
enable_reading_list=false
#enable_iterator_debugging=false
enable_hangout_services_extension=false
angle_has_histograms=false
#angle_build_tests=false
#build_angle_perftests=false
treat_warnings_as_errors=false
use_qt=false
is_cfi=false
chrome_pgo_phase=0

# enabled features
use_gio=true
#is_official_build=true
#symbol_level=0
#blink_symbol_level=0
#v8_symbol_level=0
use_pulseaudio=true
link_pulseaudio=true
rtc_use_pipewire=true
icu_use_data_file=true
enable_widevine=true
v8_enable_backtrace=true
#use_system_zlib=true
#use_system_lcms2=true
#use_system_libjpeg=true
#use_system_libpng=true
#use_system_libtiff=false
#use_system_freetype=true
#use_system_harfbuzz=true
#use_system_libopenjpeg2=true
proprietary_codecs=true
ffmpeg_branding="Chrome"
disable_fieldtrial_testing_config=true

# https://salsa.debian.org/chromium-team/chromium/-/blob/master/debian/etc/apikeys?ref_type=heads
google_api_key="AIzaSyCkfPOPZXDKNn8hhgu3JrA62wIgC93d44k"
google_default_client_id="811574891467.apps.googleusercontent.com"
google_default_client_secret="kdloedMFGdGla2P1zacGjAQh"

cc_wrapper="env CCACHE_SLOPPINESS=time_macros CCACHE_NOHASHDIR=1 CCACHE_LOGFILE=/tmp/ccache_log.log ccache"
is_component_build=true
EOF
)

ARGS=$(echo "$ARGS" | sed '/^\s*#/d' | sed '/^\s*$/d' | paste -sd " ")

# components/BUILD.gn
# components/os_crypt/sync/os_crypt_unittest.cc

gn gen out/Test --root-target=//components/os_crypt/sync --args="$ARGS"
# gn ls out/Test | grep os_crypt
# gn refs out/Tests --testonly=true --type=executable --all components/os_crypt/sync/os_crypt_unittest.cc

echo "Testing now: `date "+%Y-%m-%d %H:%M:%S"`"

# autoninja -C out/Test components/os_crypt/sync:unit_tests

# /usr/lib/xorg/Xorg.wrap: Only console users are allowed to run the X server
# Solution: set
# echo "allowed_users=anybody" | sudo tee -a /etc/X11/Xwrapper.config

tools/autotest.py -C out/Test os_crypt_unittest.cc
#out/Test/installer_util_unittests

echo "Chromium test complete: `date "+%Y-%m-%d %H:%M:%S"`"

cd $WRK