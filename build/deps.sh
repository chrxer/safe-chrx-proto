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
    echo "chromium src present, syncing.."
    cd "$CHROMIUM/src"
    
    # sudo -u $USER env "PATH=$PATH" git rebase-update
    sudo -u $USER env "PATH=$PATH" gclient sync
fi
sudo "$CHROMIUM/src/build/install-build-deps.sh"
cd $WRK

set +e
