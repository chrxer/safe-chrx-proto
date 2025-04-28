#!scripts/.venv/bin/python3

import pathlib
import sys
import os

IS_LINUX = sys.platform in ["linux" , "linux2"]


WRK=pathlib.Path(__file__).parent.parent.parent.resolve()

PYTHON=sys.executable
_pypath="bin/python3"
if not IS_LINUX:
    _pypath="Scripts/python.exe"
GOOGLEPYTHON=WRK.joinpath("scripts/.googlevenv/").joinpath(_pypath)

SRC=WRK.joinpath("chromium/src")

DEPOT_TOOLS=WRK.joinpath("depot_tools")

VENVBIN=WRK.joinpath("scripts/.googlevenv/bin")
if not IS_LINUX:
    VENVBIN=WRK.joinpath("scripts/.googlevenv/Scripts")

sys.path.insert(0, DEPOT_TOOLS)
os.environ["PATH"] = str(DEPOT_TOOLS) + os.pathsep + os.environ["PATH"]
GOOGLEENV=os.environ.copy()
GOOGLEENV["PATH"] = str(VENVBIN) + os.pathsep + os.environ["PATH"]


OUT=SRC.joinpath("out")
RELEASE=OUT.joinpath("Release")

BUILD=WRK.joinpath("build")
PATCH=WRK.joinpath("patch")



with open(WRK.joinpath("chromium.version"), "rb") as f:
    VERSION=f.read().decode("utf-8")

del sys, pathlib, os