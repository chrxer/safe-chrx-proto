#!/usr/bin/python3
import datetime

from utils.git_ import fcount
from utils.fs import fsize, fmt
from utils import SRC, WRK, ccache_

def show_stats():
    size=0
    if SRC.exists():
        count = fcount(SRC)
        size = fsize(SRC)
    avg_size = size / count if count > 0 else 0

    stamp = datetime.datetime.now().strftime("%m-%d %H:%M:%S")
    print(f"\033[94m[MOD {stamp}]\033[0m Size of ./{SRC.relative_to(WRK)}")

    print(f"Size: {fmt(size, 1024, 'B')}")
    print(f"Files: {fmt(count).lower()}")
    print(f"Avg size: {fmt(avg_size, 1024, 'B')}")
    ccache_.sv()

if __name__ == "__main__":
    show_stats()
