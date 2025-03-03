#!/usr/bin/python3

from os import PathLike
from .initenv import IS_LINUX, SRC, VERSION
from .wrap import exc
from .fetch import fetch
import json


def fcount(cwd:PathLike=SRC) -> int:
    if not IS_LINUX:
        raise NotImplemented("Only implemented for linux yet")
    count = exc("bash","-c","git ls-files | wc -l", dbg=False, cwd=cwd)
    return int(count)

def reset(cwd=SRC):
    exc("git", "clean", "-d", "--force", cwd=cwd)
    exc("git","reset", "--hard", "--recurse-submodules", cwd=cwd)

def get_commit_from_tag(tag:str=VERSION):
    url=f"https://chromium.googlesource.com/chromium/src.git/+/{tag}?format=JSON"

    data = fetch(url)
    data = json.loads((data[5:]).decode("utf-8"))

    # Remove Gitiles security prefix and extract commit hash
    return data["commit"]