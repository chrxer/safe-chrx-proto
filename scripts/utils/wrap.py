#!/usr/bin/python3

from typing import Iterable, Union
from os import PathLike

import subprocess
import sys
import shlex
import datetime

from .initenv import WRK

class CommandFailed(Exception):
    def __init__(self, code:int, cmd:Iterable[str]):
        super().__init__(f"Command: {shlex.join(cmd)} failed with exit code {code}")

def exc(*cmd:Iterable[str],dbg:bool=True, bytes:bool=False,timeout:float=None, cwd:PathLike=WRK) -> Union[bytes, str]:
    stamp = datetime.datetime.now().strftime("%m-%d %H:%M:%S")
    if dbg:
        print(f"[EXC {stamp}] {shlex.join(cmd)}")
        stdout = sys.stdout
    else:
        stdout = subprocess.PIPE
    proc = subprocess.Popen(cmd, stdout=stdout, stdin=sys.stdin, stderr=sys.stderr, cwd=cwd)
    proc.wait(timeout=timeout)
    if proc.returncode != 0:
        raise CommandFailed(proc.returncode, cmd)
    if proc.stdout:
        if bytes:
            return proc.stdout.read()
        return proc.stdout.read().decode("utf-8")