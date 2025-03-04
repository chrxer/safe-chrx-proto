#!/usr/bin/python3

from build import build as obuild

# /usr/lib/xorg/Xorg.wrap: Only console users are allowed to run the X server
# Solution: set
# echo "allowed_users=anybody" | sudo tee -a /etc/X11/Xwrapper.config

def build():
    obuild(target="components/os_crypt/sync:unit_tests",debug=False)

if __name__ == "__main__":
    build()