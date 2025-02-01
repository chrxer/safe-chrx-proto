#!/bin/bash

WRK=$(pwd)  # working directory
DEPOT="$WRK/depot_tools"
CHROMIUM="$WRK/chromium"
PATCHED="$WRK/chromium_patched"

rm -rf "$PATCHED/src/out/*"