#!/bin/bash

set -e
# running inside chrxer repo directory root

WRK=$(realpath $(dirname $(dirname "$0")))
CHROMIUM="$WRK/chromium"
RELEASE=$CHROMIUM/src/out/Release

BUILD=$WRK/build

mkdir -p $BUILD
# https://salsa.BUILDian.org/chromium-team/chromium/-/blob/master/BUILDian/chromium.install?ref_type=heads

cp $RELEASE/chrome $BUILD

cp -r $RELEASE/chrome_*.pak $BUILD
cp $RELEASE/resources.pak $BUILD
cp $RELEASE/icudtl.dat $BUILD
cp $RELEASE/chrome_crashpad_handler $BUILD
cp -r $RELEASE/*snapshot*.bin $BUILD
cp -r $RELEASE/lib*.so $BUILD
cp -r $RELEASE/lib*.so.1 $BUILD

mkdir -p $BUILD/locales
cp -r $RELEASE/locales/*.pak $BUILD/locales

# BUILDian/etc/apikeys etc/chromium.d
# BUILDian/etc/default-flags etc/chromium.d
# BUILDian/etc/extensions etc/chromium.d
# BUILDian/etc/master_preferences etc/chromium