# !/usr/bin/python3
from os import PathLike
from .initenv import IS_LINUX, SRC
from .wrap import exc


def fcount(cwd:PathLike=SRC) -> int:
    if not IS_LINUX:
        raise NotImplemented("Only implemented for linux yet")
    count = exc("bash","-c","git ls-files | wc -l", dbg=False, cwd=cwd)
    return int(count)

def reset(cwd=SRC):
    exc("git", "clean", "-d", "--force", cwd=cwd)
    exc("git","reset", "--hard", "--recurse-submodules", cwd=cwd)