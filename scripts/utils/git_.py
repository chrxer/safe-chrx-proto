#!scripts/.venv/bin/python3

from pathlib import Path
import sys
sys.path.insert(0,str(Path(__file__).parent.parent))

from os import PathLike
from utils.initenv import IS_LINUX, SRC, VERSION
from utils.wrap import exc
from utils.fetch_ import fetch

import json
import os
from typing import Iterable

def peek(_iter:Iterable):
    try:
        return next(_iter)
    except StopIteration:
        return None

def fcount(cwd:PathLike=SRC) -> int:
    if not IS_LINUX:
        raise NotImplemented("Only implemented for linux yet")
    count = int(exc("bash","-c","git ls-files | wc -l", dbg=False, cwd=cwd))

    pcwd = Path(cwd)
    if count == 0 and pcwd.is_dir():
        idir = pcwd.iterdir()
        if peek(idir) is not None:
            count = sum(len(files) for _, _, files in os.walk(cwd))
    return count

def reset(cwd=SRC):
    exc("git", "clean", "-d", "--force", cwd=cwd)
    exc("git","reset", "--hard", "--recurse-submodules", cwd=cwd)

def sub_update(cwd=SRC):
    exc("git", "submodule", "update","--recursive", "--remote")

def get_commit_from_tag(tag:str=VERSION):
    url=f"https://chromium.googlesource.com/chromium/src.git/+/{tag}?format=JSON"

    data = fetch(url)
    data = json.loads((data[5:]).decode("utf-8"))

    # Remove Gitiles security prefix and extract commit hash
    return data["commit"]

if __name__ == "__main__":
    t=VERSION
    if len(sys.argv) >= 2:
        t=sys.argv[1]
    print(get_commit_from_tag(tag=t))