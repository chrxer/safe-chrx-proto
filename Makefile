.PHONY: deps patch build

deps:
	scripts/deps.sh

patch:
	scripts/patch.sh

build:
	scripts/build.sh && scripts/pack.sh

clean:
	scripts/clean.sh