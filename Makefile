BIN_NAME := go-cli-example
BIN_PATH := bin/$(BIN_NAME)
BUILD_FLAGS := -installsuffix "static"

.PHONY: all build clean deps

all: build

build:
	@echo Compiling binary...
	@CGO_ENABLED=1 \
	GO111MODULE=on \
	GOARCH=amd64 \
	go build \
	$(BUILD_FLAGS) \
	-o $(BIN_PATH) \
	$$(find example/*.go)
	@echo Generated assets:
	@find bin/

clean:
	@echo Cleaning bin/ directory...
	@rm -rfv bin/

