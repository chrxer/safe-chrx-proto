#!/usr/bin/python3


from urllib.request import urlopen
def fetch(url:str) -> bytes:
    resp = urlopen(url)
    return resp.read()