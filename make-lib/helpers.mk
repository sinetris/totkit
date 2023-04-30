# Helper functions

# named colors for text
ifeq ($(origin NO_COLOR),undefined)
  no_color_description := false (will try to print colors)
  red_text := $(shell tput setaf 1)
  green_text := $(shell tput setaf 2)
  yellow_text := $(shell tput setaf 3)
  blue_text := $(shell tput setaf 4)
  magenta_text := $(shell tput setaf 5)
  cyan_text := $(shell tput setaf 6)
  white_text := $(shell tput setaf 7)
  bold_text := $(shell tput bold)
  reset_text := $(shell tput sgr 0)
else
  no_color_description := true (colors are disabled)
  red_text :=
  green_text :=
  yellow_text :=
  blue_text :=
  magenta_text :=
  cyan_text :=
  white_text :=
  bold_text :=
  reset_text :=
endif

error_text := ${red_text}
ok_text := ${green_text}
warning_text := ${yellow_text}
info_text := ${blue_text}

allowed_styles := red_text green_text yellow_text blue_text \
 magenta_text cyan_text white_text bold_text \
 error_text ok_text warning_text info_text

# 📝 You can NOT use commas (,) in make function arguments or leading spaces in the first argument.
# More info in: https://www.gnu.org/software/make/manual/html_node/Syntax-of-Functions.html
# If you need to use a comma in a parameter, you can use variable substitution.
# For example, if you want to use as firtst argument the string
# "  Hello, people!" (starting with 2 spaces and with a comma in the middle)
# you could use:
# msg = echo '$1'
# $(call msg,$(space) Hello$(comma) people!)
comma:= ,
empty:=
space:= $(empty) $(empty)
.PHONY: comma empty space

# $(call print_help,<header_variable_name>)
#   You need to 'export' <header_variable_name>
#   Example:
#   define project_help_header
#   	This message will be shown at the top
#   	when you run 'make help'
#   endef
#   export project_help_header
#   $(call print_help,help_header)
define print_help
	@echo "$(yellow_text)$${$1}$(reset_text)"
	@echo "$(bold_text)$(yellow_text)Usage:$(reset_text)"
	@grep -E '(^[a-zA-Z_-]+:.*?##.*$$)|(^##)' Makefile \
	| awk 'BEGIN {FS = ":.*?## "}; {printf "$(bold_text)$(green_text)make %-30s$(reset_text) %s\n", $$1, $$2}' \
	| sed -e 's/\[32m##/[33m/'
endef
