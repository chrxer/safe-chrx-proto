#!/usr/bin/python3

from utils import SRC, DEPOT_TOOLS, VERSION, git
from .wrap import exc

GCLIENT=DEPOT_TOOLS.joinpath("gclient.py")

def sync():
    commit = git.get_commit_from_tag(VERSION)
    exc("python3", str(GCLIENT),"sync","--no-history", "--shallow", "--jobs","8", "--revision", f"src@{commit}", cwd=SRC)