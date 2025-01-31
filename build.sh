#!/bin/bash

# running inside chrxer repo directory root
WRK=$(pwd)  # working directory
echo "Building Chromium..."
echo "Downloading dependencies... https://chromium.googlesource.com/chromium/src/+/main/docs/linux/build_instructions.md#Install"

git clone https://chromium.googlesource.com/chromium/tools/depot_tools.git
export PATH="$WRK/depot_tools:$PATH"

CHROMIUM="$WRK/chromium"
mkdir -p "$CHROMIUM" && cd "$CHROMIUM"
fetch --no-history --nohooks chromium

# Configure ccache
ccache --max-size=30G
export CCACHE_CPP2=yes
export CCACHE_SLOPPINESS=time_macros

# Build Chromium
cd "$CHROMIUM/src"
./build/install-build-deps.sh
gclient runhooks

echo "Compiling Chromium release now"
gn gen out/Release --args='is_debug=false is_official_build=true symbol_level=0 cc_wrapper="ccache"'
ninja -C out/Release chrome

cd "$WRK"
