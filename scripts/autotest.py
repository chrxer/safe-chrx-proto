#!/usr/bin/python3

import sys

from build import gn
from utils import OUT, SRC, ccache_
from utils.wrap import pyexc

def test(target:str):
    outdir=OUT.joinpath("Debug")
    OUTR=str(outdir.relative_to(SRC))

    gn(outdir=OUTR, debug=True)
    ccache_.sv()
    ccache_.z()
    pyexc(str(SRC.joinpath("tools/autotest.py")), "-C" ,OUTR, "os_crypt_unittest.cc", cwd=SRC)

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: scripts/autotest.py os_crypt_unittest.cc", file=sys.stderr)
        sys.exit(1)
    
    target = sys.argv[1]
    
    test(target=target)