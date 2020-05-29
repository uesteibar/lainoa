.PHONY: help build test

help: ## Show this help message
	@echo "\nOptions:\n"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: ## Build binary
	go build -o bin/ ./...

install: ## Install the binary
	go install ./...

repl:
	go run ./cmd/lainoa/main.go repl

test: ## Run the tests
	go test ./... -count=1
