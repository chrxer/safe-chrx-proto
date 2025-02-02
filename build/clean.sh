#!/bin/bash

WRK=$(pwd)  # working directory
DEPOT="$WRK/depot_tools"
CHROMIUM="$WRK/chromium"

rm -rf "$CHROMIUM/src/out/*"

# revert git to latest

cd "$CHROMIUM/src"
set +e
git clean -d --force && git reset --hard --recurse-submodules
set -e
cd $WRK
