# Makefile for mpris-tui

# Binary name
BINARY_NAME=mpris-tui

# Default target
all: build

# Build the binary
build:
	@echo "Building $(BINARY_NAME)..."
	@go build -o $(BINARY_NAME) main.go

# Install the binary
install: build
	@echo "Installing $(BINARY_NAME) to /usr/local/bin..."
	@sudo mv $(BINARY_NAME) /usr/local/bin/$(BINARY_NAME)

# Clean the build artifacts
clean:
	@echo "Cleaning..."
	@rm -f $(BINARY_NAME)

# Uninstall the binary
uninstall:
	@echo "Uninstalling $(BINARY_NAME) from /usr/local/bin..."
	@sudo rm -f /usr/local/bin/$(BINARY_NAME)

.PHONY: all build install clean uninstall
