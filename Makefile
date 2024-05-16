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
