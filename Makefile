SOURCE_MAKE=. ./.make/make.sh
SHELL := /bin/bash

build:
	@${SOURCE_MAKE} && build

prepare:
	@${SOURCE_MAKE} && prepare

build_metactl:
	@${SOURCE_MAKE} && build_metactl

build_metamate:
	@${SOURCE_MAKE} && build_metamate

chore:
	@${SOURCE_MAKE} && chore

release:
	@${SOURCE_MAKE} && release

test_release:
	@${SOURCE_MAKE} && test_release

generate:
	@${SOURCE_MAKE} && generate

deploy:
	@${SOURCE_MAKE} && deploy

test:
	@${SOURCE_MAKE} && test

x_build_and_serve:
	@${SOURCE_MAKE} && x_build_and_serve

x_serve:
	@${SOURCE_MAKE} && x_serve