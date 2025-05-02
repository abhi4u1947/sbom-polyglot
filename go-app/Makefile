.PHONY: build run clean test deps mod-download docker-build

# Detect OS
ifeq ($(OS),Windows_NT)
    DETECTED_OS := Windows
    PATH_SEP := \\
    RM := cmd.exe /c "if exist $(BINARY_DIR) rmdir /s /q $(BINARY_DIR)"
    MKDIR := mkdir
else
    DETECTED_OS := Linux
    PATH_SEP := /
    RM := rm -rf $(BINARY_DIR)
    MKDIR := mkdir -p
endif

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
BINARY_NAME=go-app
BINARY_DIR=bin

# Default target
all: clean mod-download build

# Build targets
build:
	@echo "Building application on $(DETECTED_OS)..."
	@$(MKDIR) $(BINARY_DIR)
	cd src && $(GOBUILD) -o $(BINARY_DIR)$(PATH_SEP)$(BINARY_NAME)

# Run targets
run:
	@echo "Running application..."
	cd src && $(GOCMD) run main.go

# Clean targets
clean:
	@echo "Cleaning..."
	@$(RM)
	cd src && $(GOCLEAN)

# Test targets
test:
	@echo "Running tests..."
	cd src && $(GOTEST) -v ./...

# Dependency targets
deps:
	@echo "Installing dependencies..."
	cd src && $(GOGET) -v ./...

mod-download:
	@echo "Downloading Go modules..."
	cd src && $(GOMOD) download
