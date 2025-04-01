.DEFAULT_GOAL := build

fmt:
	go fmt ./...
.PHONY: fmt

lint:
	golint ./...
.PHONY: lint

test:
	go test -v ./...
.PHONY: test

vet: fmt
	go vet ./...
.PHONY: vet

build: fmt lint test vet
	go build -o hello
.PHONY: build

run: build
	./hello
.PHONY: run
