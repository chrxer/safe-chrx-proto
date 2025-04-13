.PHONY: deps patch build

deps:
	scripts/deps.bat

diff:
	scripts/diff.py

patch:
	scripts/patch.py

build:
	scripts/build.py && scripts/pack.py

test:
	scripts/autotest.py chrx/os_crypt_hook/os_crypt_hook_unittest.cc

clean:
	scripts/clean.py
