.PHONY: help

help: ## This help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

title:
	@echo "Hookz Makefile"
	@echo "--------------"

build: ## Builds the application
	go get -u ./...
	@go mod tidy
	@go build

check: build ## Tests the pre-commit hooks if they exist
	./hookz reset --verbose --debug --verbose-output 
	. .git/hooks/pre-commit

test: ## Runs tests and coverage
	@go test -v -coverprofile=coverage.out ./... && go tool cover -func=coverage.out
	@go tool cover -html=coverage.out -o coverage.html

install: build ## Builds an executable local version of Hookz and puts in in /usr/local/bin
	@sudo chmod +x hookz
	@sudo mv hookz /usr/local/bin

all: title build test ## Makes all targets

