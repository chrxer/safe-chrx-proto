#!scripts/.venv/bin/python3

from utils.fs import rmtree
from utils import OUT, SRC
from utils import git_
import sys

def clean(_target):
    rmtree(OUT.joinpath(_target).resolve())

if __name__ == "__main__":
    git_.reset(SRC)
    git_.sub_update()