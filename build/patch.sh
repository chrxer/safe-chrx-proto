#!/bin/bash

WRK=$(pwd)  # working directory
DEPOT="$WRK/depot_tools"
CHROMIUM="$WRK/chromium"
PATCHED="$WRK/chromium_patched"
PATCH="$WRK/patch"

if [ ! -d $PATCHED ]; then
    cp -rs $CHROMIUM $PATCHED
fi

echo "Copying all files in patch/chromium to chromium_patched"
rm $PATCHED/.gclient
cp -H $PATCH/chromium/.gclient $PATCHED/.gclient