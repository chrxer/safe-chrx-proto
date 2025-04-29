#!scripts/.venv/bin/python3

from utils import SRC, gclient_, PATCH
from diff import diff
from ptcx import patch as ptcxpatch
from clean import clean

def patch():
    clean()
    gclient_.sync()
    ptcxpatch.path(srcroot=SRC, patchroot=PATCH.joinpath("chromium/src"))
    diff()


if __name__ == "__main__":
    patch()