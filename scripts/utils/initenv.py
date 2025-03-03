#!/usr/bin/python3

import pathlib
import sys

WRK=pathlib.Path(__file__).parent.parent.parent.resolve()
SRC=WRK.joinpath("chromium/src")
OUT=SRC.joinpath("out")

IS_LINUX = sys.platform in ["linux" , "linux2"]

del sys, pathlib