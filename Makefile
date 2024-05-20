#!/bin/zsh

install:
	@echo "Started installing process."
	@go mod tidy
	@go install
	@echo "Finished installed."

lint:
	@echo "Started linting process."
	@golangci-lint run
	@echo "Finished linting."

test:
	@echo "Started testing process."
	@go test -v ./...
	@echo "Finished testing."

run:
	@echo "Started running process."
	@go run ./main.go
	@echo "Finished running."