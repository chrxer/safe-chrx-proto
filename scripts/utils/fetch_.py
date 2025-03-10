#!scripts/.venv/bin/python3
import requests

def fetch(url:str) -> bytes:
    resp = requests.get(url, allow_redirects=True)
    if resp.status_code != 200:
        raise ValueError(f"Expected status 200, but got {resp.status_code}")
    return resp.content