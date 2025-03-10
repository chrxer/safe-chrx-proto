.PHONY: deps patch build

deps:
	scripts/deps.sh

patch:
	scripts/patch.py

build:
	scripts/build.py && scripts/pack.py

os_crypt:
	scripts/build_os_crypt.py

test:
	scripts/autotest.py os_crypt_unittest.cc

clean:
	scripts/clean.py
