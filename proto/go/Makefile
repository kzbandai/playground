all: fmt build

.PHONY: edit
edit:
	vim src/example/example.go

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: build
build:
	go build ./build/gen/main ./...
