#!scripts/.venv/bin/python3

from utils import SRC, WRK, PATCH, GOOGLEENV
from utils.wrap import exc
from pathlib import Path
import shutil, os
import datetime
from collections import Counter

def diff():

    stamp = datetime.datetime.now().strftime("%m-%d %H:%M:%S")
    print(f"\033[94m[MOD {stamp}]\033[0m creating git diff (excluding added files)")

    diff = exc("git","diff", "--diff-filter=MD","--patience", "--submodule=diff", cwd=SRC, _bytes=True, dbg=False, env=GOOGLEENV)
    with open(WRK.joinpath("diff.patch"), "wb") as f:
        f.seek(0)
        f.write(diff)
        f.truncate()

    stamp = datetime.datetime.now().strftime("%m-%d %H:%M:%S")
    print(f"\033[94m[MOD {stamp}]\033[0m copying added files to patch/chromium/src")

    ls = exc("git", "ls-files", "--others", "--exclude-standard", "--exclude", "out", cwd=SRC, dbg=False, env=GOOGLEENV)
    untracked = ls.strip().split("\n") if ls else []

    ls = exc("git", "ls-files", "--modified", "--exclude-standard", "--exclude", "out", cwd=SRC, dbg=False, env=GOOGLEENV)
    modified = ls.strip().split("\n") if ls else []

    for file in (list((Counter(untracked) - Counter(modified)).elements())):
        dst_path:Path=PATCH.joinpath("chromium/src").joinpath(file)

        os.makedirs(os.path.dirname(dst_path), exist_ok=True)
        shutil.copy2(SRC.joinpath(file), dst_path)
        print(f"\033[92m[cp] {file}\033[0m")



if __name__ == "__main__":
    diff()