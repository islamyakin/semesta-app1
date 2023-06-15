GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=web-app_semesta_service
LINTER=golangci-lint

.PHONY: all test build lint
all: test build

test:
      $(GOTEST) ./... -v

build:
      $(GOBUILD) -o $(BINARY_NAME) -v

lint:
      $(LINTER) run