#!scripts/.venv/bin/python3

from utils import SRC, git_, gclient_, PATCH, WRK
from utils.wrap import exc
import shutil
from pathlib import Path
from os.path import relpath
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

def patch(_clean=True):
    if _clean:
        clean()
        
    gclient_.sync()
    cpr(PATCH.joinpath("chromium"), WRK.joinpath("chromium"))
    # cpr(SRC.joinpasth("tools/vscode/"), SRC.joinpath(".vscode"))
    exc("git","apply", str(WRK.joinpath("os_crypt.patch")), cwd=SRC,_pidx=3)


if __name__ == "__main__":
    patch()