#!/bin/bash

set -e

WRK=$(pwd)  # working directory
CHROMIUM="$WRK/chromium"
PATCH="$WRK/patch"
export PATH="$WRK/depot_tools:$PATH"

# revert git to latest

cd "$CHROMIUM/src"
set +e
git clean -d --force && git reset --hard --recurse-submodules
set -e

echo "Patching chromium.."
cp -a "$PATCH/chromium/." $CHROMIUM/
echo "patched chromium"
gclient sync

cd $WRK

set +e