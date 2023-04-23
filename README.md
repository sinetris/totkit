# totkit - TechOps ToolKit

## Development dependencies

- [Go](https://go.dev/):
    open-source programming language supported by Google
- [golangci-lint](https://golangci-lint.run/):
    Go linters aggregator
- [gox](https://github.com/mitchellh/gox):
    dead simple, no frills Go cross compile tool
- [direnv](https://direnv.net/):
    load and unload environment variables depending on the current directory
- [pre-commit](https://pre-commit.com/):
    framework for managing and maintaining multi-language pre-commit hooks
- [richgo](https://github.com/kyoh86/richgo):
    enrich `go test` outputs with text decorations
- [GNU Make](https://www.gnu.org/software/make):
    dependency management and automated software build

### Setup on macOS

> **Note:** You need to have [Homebrew](https://brew.sh/)
> and [Xcode](https://developer.apple.com/xcode/) installed.

```sh
# check that Xcode is installed
xcode-select -p
# check that homebrew is installed properly
brew doctor
# install dependencies
brew update
brew install pre-commit direnv make
```

> Is better to install [Go](https://go.dev/) following
> the official [install instructions](https://go.dev/doc/manage-install),
> but you can also use `brew`.

```sh
# if you want to install GO using brew
brew install golang
# install 'richgo'
brew install kyoh86/tap/richgo
# install 'gox'
go install github.com/mitchellh/gox@latest
# check and setup 'pre-commit'
pre-commit --version
pre-commit install
pre-commit run --all-files
# Use 'direnv' to set/unset environment variables automatically
cp .envrc.example .envrc
# ... modify .envrc to your needs and then run
direnv allow
```

## Run locally

To get a complete list of commands, run

```sh
make help
```

For example, to run tests:

```sh
make tests
```

## To Do

- [ ] add [go-plugin](https://github.com/hashicorp/go-plugin)
