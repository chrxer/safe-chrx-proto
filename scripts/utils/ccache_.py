#!scripts/.venv/bin/python3
# wrapper for ccache

from .wrap import exc
from .initenv import GOOGLEENV

def set_max_g(size:int=30):
    exc("ccache", f"--max-size={size}G", env=GOOGLEENV)

def show():
    exc("ccache", "-s", env=GOOGLEENV)

def z():
    exc("ccache", "-z",env=GOOGLEENV)

def clear():
    res = input("Are you sure you want to clear all Ccache ?\n:")
    if res in ["y", "yes"]:
        exc("ccache", "-C", env=GOOGLEENV)
    else:
        raise ValueError("Clearing ccache aborted based on user-input")