# Variables
APP_NAME := aoc24
SRC := $(shell find . -type f -name '*.go')

GO := go
GOBUILD := $(GO) build
GOTEST := $(GO) test
GOCLEAN := $(GO) clean
GOMOD := $(GO) mod
GOLINT := golint
GOFMT := gofmt

BUILD_DIR := ./bin

.PHONY: all
all: build

.PHONY: build
build: $(SRC)
	@echo "Building $(APP_NAME)..."
	$(GOBUILD) -o $(BUILD_DIR)/$(APP_NAME)

.PHONY: test
test:
	@echo "Running tests..."
	$(GOTEST) ./...

.PHONY: lint
lint:
	@echo "Linting code..."
	$(GOLINT) ./...

.PHONY: fmt
fmt:
	@echo "Formatting code..."
	$(GOFMT) -w $(SRC)

.PHONY: run
run: build
	@echo "Running $(APP_NAME)..."
	$(BUILD_DIR)/$(APP_NAME)

.PHONY: clean
clean:
	@echo "Cleaning up..."
	$(GOCLEAN)
	rm -rf $(BUILD_DIR)

.PHONY: mod
mod:
	@echo "Tidying Go modules..."
	$(GOMOD) tidy
