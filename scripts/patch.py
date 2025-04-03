#!scripts/.venv/bin/python3

from utils import SRC, gclient_, PATCH, WRK
from diff import diff
from ptcx import patch as ptcxpatch
import shutil
from pathlib import Path
from clean import clean

def _logpath(path, names):
    for name in names:
        _path = Path(path).joinpath(name).absolute()
        if not _path.is_dir():
            try:
                _path = _path.relative_to(PATCH)
            except ValueError:
                _path = _path.relative_to(WRK)
                print(f"\033[92m[cp] {_path}\033[0m")
            else:
                print(f"\033[92m[patch cp] {_path}\033[0m")
    return []   # nothing will be ignored

def cpr(src, dst):
    shutil.copytree(src, dst, dirs_exist_ok=True, ignore=_logpath)

def patch():
    clean()
    gclient_.sync()
    ptcxpatch.path(srcroot=SRC, patchroot=PATCH.joinpath("chromium/src"))
    diff()


if __name__ == "__main__":
    patch()