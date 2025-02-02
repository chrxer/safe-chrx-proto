#!/bin/bash

set -e
# running inside chrxer repo directory root
WRK=$(pwd)  # working directory
DEPOT="$WRK/depot_tools"
CHROMIUM="$WRK/chromium"

export PATH="$DEPOT:$PATH"

echo "Configuring ccache"
ccache --max-size=30G
export CCACHE_CPP2=yes
export CCACHE_SLOPPINESS=time_macros

# Build Chromium
cd "$CHROMIUM/src"

echo "gn gen Chromium release.."
gn gen out/Release --args='is_debug=false is_official_build=true symbol_level=0 cc_wrapper="ccache"'
echo "autoninja .."
autoninja -C out/Release chrome
echo "Chromium build complete"

cd "$WRK"
set +e
