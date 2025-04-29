#!scripts/.venv/bin/python3

from utils.fs import rmtree
from utils import OUT, SRC
from ptcx import patch as ptcxpatch

# resets chromium/src
def clean(_target=None):
    ptcxpatch.reset(SRC)
    if _target is not None:
        rmtree(OUT.joinpath(_target).resolve())

if __name__ == "__main__":
    clean()