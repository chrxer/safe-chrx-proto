#!/bin/bash

# requires sudo!

WRK=$(pwd)  # working directory
DEPOT="$WRK/depot_tools"
CHROMIUM="$WRK/chromium"

set -e
USER=$(awk -F: '$3 >= 1000 && $3 < 60000 {print $1; exit}' /etc/passwd)

if [ ! -d "$WRK/depot_tools" ]; then
    echo "Depot tools missing, installing.."

    sudo -u $USER env "PATH=$PATH" git clone https://chromium.googlesource.com/chromium/tools/depot_tools.git
fi
export PATH="$DEPOT:$PATH"

if [ ! -d $CHROMIUM ]; then
    echo "Chromium missing, fetching now..."
    
    sudo -u $USER env "PATH=$PATH" mkdir -p "$CHROMIUM"
    cd "$CHROMIUM"
    sudo -u $USER env "PATH=$PATH" fetch --no-history --nohooks chromium
    cd "$CHROMIUM/src"
else
    echo "chromium src present, syncing & cleaning.."
    cd "$CHROMIUM/src"
    
    git config http.postBuffer 524288000
    git config protocol.version 2
    sudo -u $USER env "PATH=$PATH" git clean -d --force && git reset --hard --recurse-submodules
    sudo -u $USER env "PATH=$PATH" git rebase-update --current
    sudo -u $USER env "PATH=$PATH" gclient sync --reset
fi
sudo "$CHROMIUM/src/build/install-build-deps.sh"
cd $WRK

set +e
