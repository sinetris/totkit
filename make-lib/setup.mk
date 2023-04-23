SHELL := bash
.ONESHELL:
.SHELLFLAGS := -eu -o pipefail -c
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules

export GO111MODULE=on

RUNNING_IN_ENV?=local
XC_OS?=
XC_ARCH?=
# 📝 Setup based on "RUNNING_IN_ENV" environment variable
export RUNNING_IN_ENV?=
ifeq ($(RUNNING_IN_ENV), local) # setup to run in a development system
  UNAME_S := $(shell uname -s)
  ifeq ($(UNAME_S),Linux)
    XC_OS += "linux"
  endif
  ifeq ($(UNAME_S),Darwin)
    XC_OS += "darwin"
  endif
  UNAME_P := $(shell uname -p)
  ifeq ($(UNAME_P),x86_64)
    XC_ARCH += "amd64"
  endif
  ifeq ($(UNAME_P),i386)
    ifeq ($(UNAME_S),Darwin)
      # macOS report an i386 architecture, but can run 64bit software
      XC_ARCH += "amd64"
    else
      XC_ARCH += "386"
    endif
  endif
  ifneq ($(filter arm64%,$(UNAME_P)),)
    XC_ARCH += "arm64"
  endif
  XC_PARALLEL ?= "-1"
  XC_OSARCH ?=
else ifeq ($(RUNNING_IN_ENV), ci) # setup for running in CI
  XC_OS += "linux darwin freebsd windows"
  XC_ARCH += "amd64 arm64 386"
  XC_OSARCH ?= "!windows/arm64 !darwin/386"
  XC_PARALLEL ?= "2"
else # setup for anything else
  # setup for any other valid 'RUNNING_IN_ENV' in 'running_enabled_list'
endif
