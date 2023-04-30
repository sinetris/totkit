SHELL := bash
.ONESHELL:
.SHELLFLAGS := -eu -o pipefail -c
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules

export GO111MODULE=on

RUNNING_IN_ENV?=local
ifeq ($(RUNNING_IN_ENV), ci) # setup for running in CI
  XC_OS ?= "linux darwin freebsd windows"
  XC_ARCH ?= "amd64 arm64 386"
  XC_OSARCH ?= "!windows/arm64 !darwin/386"
  XC_PARALLEL ?= "2"
endif
