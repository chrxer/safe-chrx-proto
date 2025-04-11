#!scripts/.venv/bin/python3

import sys
import datetime
import shutil

from utils import ccache_, DEPOT_TOOLS, OUT, SRC, mkargs, GOOGLEPYTHON, GOOGLEENV, WRK
from utils.wrap import pyexc, exc
from pack import pack_build


def gn(outdir:str,target:str="chrome",debug:bool=False):
    # find root-target
    root_target=None
    targs = target.split(":")
    if len(targs) > 1:
        root_target=targs[0]

    ccache_.set_max_g(20)
    args = mkargs.make(debug=debug)
    pargs= ' '.join(args)
    gnargs = [str(DEPOT_TOOLS.joinpath("gn.py")), "gen", outdir, f"--args={pargs}"]
    if root_target is not None:
        gnargs.append(f"--root-target={root_target}")
    pyexc(*gnargs, cwd=SRC, python=GOOGLEPYTHON, env=GOOGLEENV)

def build_server(outdir:str):
    build_dir = WRK.joinpath("backend/server")
    out_dir = SRC.joinpath(outdir)
    exec_name="chrxCryptServer"
    exc("go", "build", exec_name, cwd=build_dir)
    shutil.copy2(build_dir.joinpath(exec_name), out_dir.joinpath(exec_name))



def build(target:str="chrome",debug:bool=False):
    if debug:
        tag="Debug"
    else:
        tag="Release"
    OUTD=OUT.joinpath(tag)
    OUTR=str(OUTD.relative_to(SRC))
    
    build_server(outdir=OUTR)
    gn(outdir=OUTR, target=target, debug=debug)
    
    ccache_.show()
    ccache_.z()
    pyexc(str(DEPOT_TOOLS.joinpath("autoninja.py")), "-C", OUTR, target, cwd=SRC, python=GOOGLEPYTHON, env=GOOGLEENV)

    stamp = datetime.datetime.now().strftime("%m-%d %H:%M:%S")
    print(f"\033[94m[MOD {stamp}]\033[0m Building finished")

    if target=="chrome" and (not debug):
        pack_build()

if __name__ == "__main__":
    target="chrome"
    if len(sys.argv) == 2:
        target = sys.argv[1]
    
    build(target=target)