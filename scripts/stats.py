#!/usr/bin/python3

from utils.git import fcount
from utils.fs import fsize, fmt
from utils import SRC, WRK

def show_stats():
    count = fcount(SRC)
    size = fsize(SRC)
    avg_size = size / count if count > 0 else 0

    print(f"stats for {SRC.relative_to(WRK)}")
    print(f"Files: {fmt(count)}")
    print(f"Size: {fmt(size, 1024)}")
    print(f"Average size: {fmt(avg_size, 1024)}")

if __name__ == "__main__":
    show_stats()
