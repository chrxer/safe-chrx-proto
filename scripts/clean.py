#!/usr/bin/python3

from utils.fs import rmtree
from utils import OUT, SRC
from utils import git
import sys

def clean(_target):
    rmtree(OUT.joinpath(_target).resolve())

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: scripts/clean.py <target>", file=sys.stderr)
        sys.exit(1)
    
    _target = sys.argv[1]
    clean(_target)
    git.reset(SRC)