#!scripts/.venv/bin/python3

from utils import SRC, DEPOT_TOOLS, VERSION, git_
from .wrap import pyexc

GCLIENT=DEPOT_TOOLS.joinpath("gclient.py")

def sync():
    commit = git_.get_commit_from_tag(VERSION)
    pyexc(str(GCLIENT),"sync","--no-history", "--shallow", "--jobs","8", "--revision", f"src@{commit}", cwd=SRC)