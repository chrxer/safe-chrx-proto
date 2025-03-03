from .git import IS_LINUX
from .wrap import exc
from .initenv import SRC, WRK
from pathlib import Path

import shlex
from os import PathLike
import re


def fsize(path:PathLike=SRC) -> int:
    # size of path in bytes
    if not IS_LINUX:
        raise NotImplemented("Only implemented for linux yet")
    path = shlex.quote(str(Path(path).absolute()))
    size = exc("bash","-c",f"du -s -b {path} | awk '{{print $1}}'", dbg=False)
    return int(size)

def rmtree(path:PathLike):
    path = Path(path).absolute()
    if IS_LINUX:
        exc("rm", "-rf", str(path.relative_to(WRK)), cwd=WRK)
    else:
        import shutil
        shutil.rmtree(path, ignore_errors=True)

def fmt(num, base=1000, suffix=""):
    """Convert a number into a human-readable format using SI (decimal) or IEC (binary) units."""
    units = ("", "K", "M", "B", "T", "P", "E", "Z", "Y") if base == 1000 else ("", "Ki", "Mi", "Gi", "Ti", "Pi", "Ei", "Zi", "Yi")
    for unit in units:
        if abs(num) < base:
            return f"{num:3.1f}{unit}{suffix}"
        num /= base
    return f"{num:.1f}{units[-1]}{suffix}"
