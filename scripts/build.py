#!/usr/bin/python3

from utils import ccache_

def build(target:str="chrome",debug:bool=False):
    ccache_.set_max_g(30)
    # TODO: complete building script..

if __name__ == "__main__":
    build()