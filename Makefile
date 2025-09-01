# Makefile for Task Tracker CLI
# Compiler to use
GO = go

# Binary name
BINARY = expense-tracker

# Source directory
SRC_DIR = cmd

# Build flags
GOFLAGS = -v

# Default target
.PHONY: all
all: build

# Build the binary
.PHONY: build
build:
	$(GO) build $(GOFLAGS) -o $(BINARY) ./$(SRC_DIR)/...

# Run the application
.PHONY: run
run: build
	./$(BINARY)

# Clean up generated files
.PHONY: clean
clean:
	rm -f $(BINARY)

# Install dependencies
.PHONY: deps
deps:
	$(GO) mod tidy
	$(GO) mod download

# Format code
.PHONY: fmt
fmt:
	$(GO) fmt ./...

# Run tests
.PHONY: test
test:
	$(GO) test ./... -v