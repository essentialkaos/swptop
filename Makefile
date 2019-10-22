################################################################################

# This Makefile generated by GoMakeGen 1.2.0 using next command:
# gomakegen .
#
# More info: https://kaos.sh/gomakegen

################################################################################

.DEFAULT_GOAL := help
.PHONY = fmt all clean git-config deps help

################################################################################

all: swptop ## Build all binaries

swptop: ## Build swptop binary
	go build swptop.go

install: ## Install all binaries
	cp swptop /usr/bin/swptop

uninstall: ## Uninstall all binaries
	rm -f /usr/bin/swptop

git-config: ## Configure git redirects for stable import path services
	git config --global http.https://pkg.re.followRedirects true

deps: git-config ## Download dependencies
	go get -d -v pkg.re/essentialkaos/ek.v11

fmt: ## Format source code with gofmt
	find . -name "*.go" -exec gofmt -s -w {} \;

clean: ## Remove generated files
	rm -f swptop

help: ## Show this info
	@echo -e '\n\033[1mSupported targets:\033[0m\n'
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
		| awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[33m%-12s\033[0m %s\n", $$1, $$2}'
	@echo -e ''
	@echo -e '\033[90mGenerated by GoMakeGen 1.2.0\033[0m\n'

################################################################################
