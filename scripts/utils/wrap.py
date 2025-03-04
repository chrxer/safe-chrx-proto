#!/usr/bin/python3

from typing import Iterable, Union, Dict
from os import PathLike
from pathlib import Path

import subprocess
import sys
import os
import shlex
import datetime
from .initenv import WRK, PYTHON

class CommandFailed(Exception):
    def __init__(self, code:int, cmd:Iterable[str]):
        super().__init__(f"Command: {shlex.join(cmd)} failed with exit code {code}")

def fmtpath(path:PathLike) -> str:
    try:
        return "./"+str(Path(path).relative_to(WRK))
    except ValueError:
        return str(path)

def exc(*cmd:Iterable[str],dbg:bool=True, _bytes:bool=False,timeout:float=None, cwd:PathLike=WRK, env:Dict[str, str]=os.environ, _pidx:int=0) -> Union[bytes, str]:
    stamp = datetime.datetime.now().strftime("%m-%d %H:%M:%S")

    hcmd = list(cmd).copy()
    hcmd[_pidx] = fmtpath(hcmd[_pidx])

    if dbg:
        print(f"\033[94m[EXC {stamp}]\033[0m {shlex.join(hcmd)}")
        stdout = sys.stdout
    else:
        stdout = subprocess.PIPE
    proc = subprocess.Popen(cmd, stdout=stdout, stdin=sys.stdin, stderr=sys.stderr, cwd=cwd, env=env)
    proc.wait(timeout=timeout)
    if proc.returncode != 0:
        
        raise CommandFailed(proc.returncode, hcmd)
    if proc.stdout:
        if _bytes:
            return proc.stdout.read()
        return proc.stdout.read().decode("utf-8")
    
def pyexc(*cmd:Iterable[str],dbg:bool=True, _bytes:bool=False,timeout:float=None, cwd:PathLike=WRK, _pidx:int=0) -> Union[bytes, str]:
    return exc(PYTHON, *cmd, dbg=dbg, _bytes=_bytes, timeout=timeout, cwd=cwd, _pidx=_pidx+1)