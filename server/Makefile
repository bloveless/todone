.PHONY: all
all: build
FORCE: ;

SHELL  := env BOOKMARK_ENV=$(BOOKMARK_ENV) $(SHELL)
BOOKMARK_ENV ?= dev

BIN_DIR = $(PWD)/bin

.PHONY: build

clean:
	rm -rf bin/*

dependencies:
	go mod download

build: dependencies build-server

build-server:
	go build -o $(BIN_DIR)/todone-server cmd/server/main.go

debug: dependencies
	reflex -c reflex.conf

linux-binaries:
	CGO_ENABLED=0 GOOS=linux go build -a -o $(BIN_DIR)/todone-server cmd/server/main.go

ci: dependencies test

test:
	go test ./...
