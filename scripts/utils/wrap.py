#!scripts/.venv/bin/python3

from typing import Iterable, Union, Dict
from os import PathLike
from os.path import relpath
from pathlib import Path

import subprocess
import sys
import os
import shlex
import shutil
import datetime
from .initenv import WRK, PYTHON

class CommandFailed(Exception):
    def __init__(self, code:int, cmd:Iterable[str]):
        super().__init__(f"Command: {shlex.join(cmd)} failed with exit code {code}")

def fmtpath(path:PathLike, base=WRK, _relpath=False) -> str:
    # format path in a readable way (relative)
    try:
        if _relpath is True:
            return relpath(str(path), str(base))
        return "./"+str(Path(path).relative_to(base))
    except ValueError:
        return str(path)

def exists_in_PATH(target):
    # Check if target is a command
    if shutil.which(target):
        return True

    # Check if target is a directory in PATH
    paths = os.environ.get("PATH", "").split(os.pathsep)
    return any(os.path.abspath(target) == os.path.abspath(p) for p in paths)

# execute a system command
def exc(*cmd:Iterable[str],dbg:bool=True, _bytes:bool=False,timeout:float=None, cwd:PathLike=WRK, env:Dict[str, str]=os.environ, _pidx:int=0) -> Union[bytes, str]:
    stamp = datetime.datetime.now().strftime("%m-%d %H:%M:%S")

    hcmd = list(cmd).copy()
   
    for i in range(0, _pidx):
        hcmd[i] = fmtpath(hcmd[i],base=cwd, _relpath=False)

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

# execute a python file
def pyexc(*cmd:Iterable[str],dbg:bool=True, _bytes:bool=False,timeout:float=None, cwd:PathLike=WRK, _pidx:int=0, python:PathLike=PYTHON,env:Dict[str,str]=os.environ) -> Union[bytes, str]:
    return exc(str(python), *cmd, dbg=dbg, _bytes=_bytes, timeout=timeout, cwd=cwd, _pidx=_pidx+1, env=env)