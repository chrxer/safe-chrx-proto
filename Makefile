.PHONY: deps patch build

deps:
	scripts/deps.sh

patch:
	scripts/patch.sh

build:
	scripts/build.sh

clean:
	scripts/clean.sh