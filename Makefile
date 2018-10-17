APP = junos_exporter

VERSION=$(shell \
        grep "version string =" main.go \
        |awk -F'=' '{print $$2}' \
        |sed -e "s/[^0-9.]//g" \
	|sed -e "s/ //g")

SHELL = /bin/bash
GO = go
DIR = $(shell pwd)

NO_COLOR=\033[0m
OK_COLOR=\033[32;01m
ERROR_COLOR=\033[31;01m
WARN_COLOR=\033[33;01m

MAKE_COLOR=\033[33;01m%-20s\033[0m

MAIN = github.com/czerwonk/junos_exporter
SRCS = $(shell git ls-files '*.go' | grep -v '^vendor/')
PKGS = $(shell go list ./... | grep -v '/vendor/')

.DEFAULT_GOAL := help

.PHONY: help
help:
	@echo -e "$(OK_COLOR)==== $(APP) [$(VERSION)] ====$(NO_COLOR)"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "$(MAKE_COLOR) : %s\n", $$1, $$2}'

.PHONY: init
init: ## Install requirements
	@echo -e "$(OK_COLOR)[$(APP)] Install requirements$(NO_COLOR)"
	@go get -u github.com/golang/dep/cmd/dep
	@go get -u golang.org/x/lint/golint

.PHONY: deps
deps: ## Update dependencies
	@echo -e "$(OK_COLOR)[$(APP)] Update dependencies$(NO_COLOR)"
	@dep ensure

.PHONY: build
build: ## Make binary
	@echo -e "$(OK_COLOR)[$(APP)] Build $(NO_COLOR)"
	@go build .

.PHONY: test
test: ## Launch unit tests
	@echo -e "$(OK_COLOR)[$(APP)] Launch unit tests $(NO_COLOR)"
	@go test .

.PHONY: lint
lint: ## Launch golint
	@$(foreach file,$(SRCS),golint $(file) || exit;)

.PHONY: vet
vet: ## Launch go vet
	@$(foreach file,$(SRCS),$(GO) vet $(file) || exit;)

.PHONY: coverage
coverage: ## Launch code coverage
	@$(foreach pkg,$(PKGS),$(GO) test -cover $(pkg) || exit;)

# for goprojectile
.PHONY: gopath
gopath:
	@echo `pwd`:`pwd`/vendor
