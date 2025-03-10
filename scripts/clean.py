#!scripts/.venv/bin/python3

from utils.fs import rmtree
from utils import OUT, SRC
from utils import git_
import sys

def clean(_target):
    rmtree(OUT.joinpath(_target).resolve())

if __name__ == "__main__":
    if len(sys.argv) == 2:
        target = sys.argv[1]
        clean(target)
    git_.reset(SRC)