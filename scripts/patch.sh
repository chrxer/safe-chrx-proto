#!/bin/bash

set -e

WRK=$(realpath $(dirname $(dirname "$0")))
CHROMIUM="$WRK/chromium"
PATCH="$WRK/patch"
VERSION=$(cat "$WRK/chromium.version")
COMMIT=$("$WRK/scripts/get_commit.sh" $VERSION)

export PATH="$WRK/depot_tools:$PATH"

# revert git to latest

cd "$CHROMIUM/src"
set +e
git clean -d --force && git reset --hard --recurse-submodules
set -e

echo "Patching chromium.."
cp -a "$PATCH/chromium/." $CHROMIUM/
echo "patched chromium"
gclient sync --no-history --shallow --jobs 8 --revision src@$COMMIT

cd $WRK

set +e