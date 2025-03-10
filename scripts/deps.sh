#!/bin/bash

# requires sudo!

set -e

# Loop through arguments
for arg in "$@"; do
    if [[ "$arg" == "--force" || "$arg" == "-f" ]]; then
        echo "--force provided"
        FORCE=1
    fi
done


WRK=$(realpath $(dirname $(dirname "$0")))
CHROMIUM="$WRK/chromium"

DEPOT="$WRK/depot_tools"
export PATH="$DEPOT:$PATH"

stmp() {
    date +"%m-%d %T"
}

printf "\033[94m[EXC %s]\033[0m sudo %s\n" "$(stmp)" "apt install python3"
sudo apt install python3

USER=$(awk -F: '$3 >= 1000 && $3 < 60000 {print $1; exit}' /etc/passwd)



nsu() {
    printf "\033[94m[EXC %s]\033[0m %s\n" "$(stmp)" "$*"
    sudo -u "$USER" env "PATH=$PATH" "$@"
}

asu() {
    printf "\033[94m[EXC %s]\033[0m sudo %s\n" "$(stmp)" "$*"
    sudo env "PATH=$PATH" "$@"
}

nsu python3 -m venv scripts/.venv
nsu scripts/.venv/bin/python -m pip install -r scripts/requirements.txt

VERSION=$(cat "$WRK/chromium.version")
COMMIT=$(sudo -u "$USER" env "PATH=$PATH" "$WRK/scripts/utils/git_.py")
echo $COMMIT

gsync() {
    nsu gclient sync --force --nohooks --no-history --shallow --jobs 8 --revision src@$COMMIT
}



if [ ! -d "$WRK/depot_tools" ]; then
    nsu git clone https://chromium.googlesource.com/chromium/tools/depot_tools.git
fi


if [[ -d "$CHROMIUM/src" || $FORCE ]]; then
    nsu mkdir -p "$CHROMIUM"
    cd "$CHROMIUM"
    nsu git config protocol.version 2
    nsu cp "$WRK/patch/chromium/.gclient" "$CHROMIUM"
    nsu sed -i 's|\$revision|$COMMIT|g' "$CHROMIUM/.gclient"
    gsync
    cd "$CHROMIUM/src"
    gsync
fi

# W: An error occurred during the signature verification
# subprocess.CalledProcessError: Command '['sudo', 'apt-get', 'update']' returned non-zero exit status 100.
set +e
nsu sed -i '/subprocess\.check_call\s*(\s*\["sudo",\s*"apt-get",\s*"update"\s*\]\s*)/d' "$CHROMIUM/src/build/install-build-deps.py"
asu "$CHROMIUM/src/build/install-build-deps.sh"
cd $WRK

set +e
