#!/bin/bash

# requires sudo!

WRK=$(realpath $(dirname $(dirname "$0")))
DEPOT="$WRK/depot_tools"
CHROMIUM="$WRK/chromium"
VERSION=$(cat "$WRK/chromium.version")
COMMIT=$("$WRK/scripts/get_commit.sh" $VERSION)

git_config() {
    sudo -u $USER env "PATH=$PATH" git config protocol.version 2
}

rerev() {
    sudo -u $USER env "PATH=$PATH" cp "$WRK/patch/chromium/.gclient" "$CHROMIUM"
    sed -i "s|\$revision|$COMMIT|g" "$CHROMIUM/.gclient"
}

gsync() {
    sudo -u $USER env "PATH=$PATH" gclient sync --force --nohooks --no-history --shallow --jobs 8 --revision src@$COMMIT
}

set -e
USER=$(awk -F: '$3 >= 1000 && $3 < 60000 {print $1; exit}' /etc/passwd)

if [ ! -d "$WRK/depot_tools" ]; then
    echo "Depot tools missing, installing.."

    sudo -u $USER env "PATH=$PATH" git clone https://chromium.googlesource.com/chromium/tools/depot_tools.git
fi
export PATH="$DEPOT:$PATH"

if [ ! -d "$CHROMIUM/src" ]; then
    sudo -u $USER env "PATH=$PATH" mkdir -p "$CHROMIUM"

    echo "Chromium missing, fetching now..."
    # ensure gclient is already configured
    
    
    cd "$CHROMIUM"
    git_config
    rerev
    echo "syncing gclient"
    gsync
    cd "$CHROMIUM/src"
    gsync

else
    echo "chromium src present, syncing & cleaning.."
    cd "$CHROMIUM/src"
    git_config
    rerev
    sudo -u $USER env "PATH=$PATH" git clean -d --force && git reset --hard --recurse-submodules
    gsync
fi

# W: An error occurred during the signature verification
# subprocess.CalledProcessError: Command '['sudo', 'apt-get', 'update']' returned non-zero exit status 100.
set +e
sed -i '/subprocess\.check_call\s*(\s*\["sudo",\s*"apt-get",\s*"update"\s*\]\s*)/d' "$CHROMIUM/src/build/install-build-deps.py"
set -e
sudo "$CHROMIUM/src/build/install-build-deps.sh"
cd $WRK

set +e
