BIN_DIR := $(GOPATH)/bin
VERSION := $(shell git describe --tags)
BUILD := $(shell git rev-parse --short HEAD)
LDFLAGS=-ldflags "-X=main.Version=$(VERSION) -X=main.Build=$(BUILD)"

install-deps:
	@echo "Getting dependencies"
	@go get -v -t ./...
test: install-deps
	go test -v -coverprofile=coverage.txt -covermode=atomic ./...
lint:
	@golint ./...
clean:
	@rm -rf build
	@go clean
build: install-deps
	@mkdir -p build
	@echo "Building"
	@go build -o ./linesink github.com/bahadrix/linesink
	@echo "Done"
