.PHONY: deps patch build

deps:
	scripts/deps.sh

patch:
	scripts/patch.py

build:
	scripts/build.sh && scripts/pack.py

clean:
	scripts/clean.py Release