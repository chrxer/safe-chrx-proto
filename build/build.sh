#!/bin/bash

set -e
# running inside chrxer repo directory root
WRK=$(pwd)  # working directory
DEPOT="$WRK/depot_tools"
CHROMIUM="$WRK/chromium"
PATCHED="$WRK/chromium_patched"

export PATH="$DEPOT:$PATH"

if [ ! -d $PATCHED ]; then
    "$WRK/build/patch.sh"
fi


echo "Configuring ccache"
ccache --max-size=30G
export CCACHE_CPP2=yes
export CCACHE_SLOPPINESS=time_macros

# Build Chromium
cd "$PATCHED/src"

echo "Building Chromium release now"
gn gen out/Release --args='is_debug=false is_official_build=true symbol_level=0 cc_wrapper="ccache"'
autoninja -C out/Release chrome

cd "$WRK"
set +e
