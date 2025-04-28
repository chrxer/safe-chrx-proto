#!scripts/.venv/bin/python3
import shutil
from pathlib import Path
import os

from utils import BUILD, RELEASE

def pack_build():
    # https://salsa.debian.org/chromium-team/chromium/-/blob/master/BUILDian/chromium.install
    if (not RELEASE.is_dir()) or len(os.listdir(RELEASE))==0:
        raise ValueError("Chrome hasn't been compiled yet")
    
    BUILD.mkdir(parents=True, exist_ok=True)
    
    for pattern in ["chrome","chrome_*.pak", "resources.pak", "icudtl.dat", "chrome_crashpad_handler", "*snapshot*.bin", "lib*.so", "lib*.so.1", "locales/*.pak", "chrxCryptServer", "chrxCryptServer.exe"]:
        for file in RELEASE.glob(pattern):
            target_file = BUILD.joinpath(file.relative_to(RELEASE))
            if file.is_file() or file.is_symlink():
                target_file.parent.mkdir(exist_ok=True, parents=True)
                shutil.copy(file, target_file)
            elif file.is_dir():
                shutil.copytree(file,target_file, dirs_exist_ok=True)
    
    # BUILD/etc/apikeys etc/chromium.d
    # BUILD/etc/default-flags etc/chromium.d
    # BUILD/etc/extensions etc/chromium.d
    # BUILD/etc/master_preferences etc/chromium

if __name__ == "__main__":
    pack_build()