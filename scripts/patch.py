#!scripts/.venv/bin/python3

from utils import SRC, git_, gclient_, PATCH, WRK
import shutil



def patch(reset=True):
    if reset:
        git_.reset(SRC)
    shutil.copytree(PATCH.joinpath("chromium"), WRK.joinpath("chromium"), dirs_exist_ok=True)
    gclient_.sync()


if __name__ == "__main__":
    patch()