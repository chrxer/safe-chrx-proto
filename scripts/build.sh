#!/bin/bash

set -e
# running inside chrxer repo directory root

WRK=$(realpath $(dirname $(dirname "$0")))
DEPOT="$WRK/depot_tools"
CHROMIUM="$WRK/chromium"
VERSION=$(cat "$WRK/chromium.version")

export PATH="$DEPOT:$PATH"

echo "Configuring ccache"
ccache --max-size=30G
export CCACHE_CPP2=yes
export CCACHE_SLOPPINESS=time_macros

# Build Chromium
cd "$CHROMIUM/src"

if [ ! -d "out/Release" ]; then
    echo "gn gen Chromium release.."
    gn gen out/Release --args="is_debug=false is_official_build=true symbol_level=0 blink_symbol_level=0 v8_symbol_level=0 enable_nacl=false cc_wrapper=\"ccache\""
fi
echo "autoninja .."_chrome_branded=true s
autoninja -C out/Release chrome
echo "Chromium build complete"

cd "$WRK"
set +e
