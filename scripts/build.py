#!scripts/.venv/bin/python3

from utils import ccache_, DEPOT_TOOLS, OUT, SRC, mkargs
from utils.wrap import pyexc
from pack import pack_build

import os
import sys
import datetime
from pathlib import Path

def gn(outdir:str,target:str="chrome",debug:bool=False):
    # find root-target
    root_target=None
    targs = target.split(":")
    if len(targs) > 1:
        root_target=targs[0]

    ccache_.set_max_g(15)
    poutdir = SRC.joinpath(outdir)
    args = mkargs.make(debug=debug)
    pargs= ' '.join(args)
    gnargs = [str(DEPOT_TOOLS.joinpath("gn.py")), "gen", outdir, f"--args={pargs}"]
    if root_target is not None:
        gnargs.append(f"--root-target={root_target}")
    pyexc(*gnargs, cwd=SRC)


def build(target:str="chrome",debug:bool=False):
    if debug:
        tag="Debug"
    else:
        tag="Release"
    OUTD=OUT.joinpath(tag)
    OUTR=str(OUTD.relative_to(SRC))
    
    gn(outdir=OUTR, target=target, debug=debug)
    
    ccache_.show()
    ccache_.z()
    pyexc(str(DEPOT_TOOLS.joinpath("autoninja.py")), "-C", OUTR, target, cwd=SRC)

    stamp = datetime.datetime.now().strftime("%m-%d %H:%M:%S")
    print(f"\033[94m[MOD {stamp}]\033[0m Building finished")

    if target=="chrome" and (not debug):
        pack_build()

if __name__ == "__main__":
    target="chrome"
    if len(sys.argv) == 2:
        target = sys.argv[1]
    
    build(target=target)