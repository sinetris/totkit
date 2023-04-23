this_makefile_path := $(realpath $(dir $(realpath $(lastword $(MAKEFILE_LIST)))))

ifeq ($(filter oneshell,$(.FEATURES)),)
$(error Your version of make is too old, please update make)
endif

-include ${this_makefile_path}/make-lib/setup.mk
-include ${this_makefile_path}/make-lib/helpers.mk

BUILDDIR?="${this_makefile_path}/build"
SRC=$(shell find . -name "*.go")

running_enabled_list := ci local
dependencies := git go golangci-lint gox richgo
local_dependencies := pre-commit direnv
ci_dependencies :=

define project_help_header
Development tasks for the TechOps Toolkit project
  Used environment variables: ${blue_text}${bold_text}RUNNING_IN_ENV, USE_COLORS${reset_text}
    ${bold_text}RUNNING_IN_ENV${reset_text}: specify where make is running
      valid values: one of '${bold_text}${running_enabled_list}${reset_text}'
      current value: ${bold_text}${RUNNING_IN_ENV}${reset_text}
    ${bold_text}USE_COLORS${reset_text}: use colors in output
      valid values: one of '${bold_text}true false${reset_text}'
      current value: ${bold_text}${USE_COLORS}${reset_text}
endef
export project_help_header

# >> helper targets
_pre-check-if-allowed:
	@if ! [[ " ${running_enabled_list} " =~ " ${RUNNING_IN_ENV} " ]]; then \
		echo "The environment variable RUNNING_IN_ENV is '${RUNNING_IN_ENV}' and should be one of '${running_enabled_list}'" \
			&& echo "Check the README.md for a list of requirements." && exit 1; \
	fi;
	@$(info ${blue_text}- Running in '${bold_text}$(RUNNING_IN_ENV)${reset_text}${blue_text}' environment ...${reset_text})
.PHONY: _pre-check-if-allowed

_pre-check-dependencies-local: export dependencies += $(local_dependencies)
_pre-check-dependencies-%:
	@$(info ${blue_text}- Checking dependencies '${bold_text}$(dependencies)${reset_text}${blue_text}' ...${reset_text})
	@for dependency in $(dependencies); do \
		if ! [ -x "$$(command -v "$${dependency}")" ]; then \
			echo "${bold_text}Dependency not found: '${red_text}$${dependency}${reset_text}'" \
				&& echo "Check the README.md for a list of requirements." \
				&& exit 1; \
		fi;
	done
# << helper targets

check: _pre-check-if-allowed _pre-check-dependencies-$(RUNNING_IN_ENV) ## Check prerequisites
.PHONY: check

format: check tidy lint-fix  ## Format code (will change files that need formatting)
	@$(info ${blue_text}- Formatting code ...${reset_text})
	@go fmt ./...
.PHONY: format

tests: check install_project_deps lint-check vet  ## Test code quality
	@$(info ${blue_text}- Starting tests ...${reset_text})
	@richgo test ./...
.PHONY: tests

vet:  ## Run go vet against code.
	@$(info ${blue_text}- Running vetting tools ...${reset_text})
	@go vet ./...
.PHONY: vet

lint-check:  ## Check linter
	@$(info ${blue_text}- Running lint check tools ...${reset_text})
	@golangci-lint run --verbose
.PHONY: lint-check

lint-fix:  ## Fix linter found issues
	@$(info ${blue_text}- Running lint fix tools ...${reset_text})
	@golangci-lint run --fix --verbose
.PHONY: lint-fix

tidy:  ## Run go mod tidy on every mod file in the repo
	@$(info ${blue_text}- Running tidy ...${reset_text})
	@go mod tidy
.PHONY: tidy

build: check install_project_deps ## Build the project executable
	@$(info ${blue_text}- Build ...${reset_text})
	@gox \
		-os=$(XC_OS) \
		-arch=$(XC_ARCH) \
		-osarch=$(XC_OSARCH) \
		-parallel=$(XC_PARALLEL) \
		-output=$(BUILDDIR)/{{.Dir}}_{{.OS}}_{{.Arch}} \
		;
.PHONY: build

install_project_deps:
	@$(info ${blue_text}- Downloading dependencies ...${reset_text})
	@go get -v ./...
.PHONY: install_project_deps

clean:  ## Remove generated artifacts
	@$(info ${blue_text}- Clean generated artifacts ...${reset_text})
	@rm -rf $(BUILDDIR)
.PHONY: clean

.DEFAULT_GOAL := help
help: Makefile  ## This help
	$(call print_help,project_help_header)
.PHONY: help
