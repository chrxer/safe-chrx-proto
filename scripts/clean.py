#!scripts/.venv/bin/python3

from utils.fs import rmtree
from utils import OUT, SRC
from utils import git_

def clean(_target=None):
    git_.reset(SRC)
    git_.sub_update()
    if _target is not None:
        rmtree(OUT.joinpath(_target).resolve())

if __name__ == "__main__":
    clean()