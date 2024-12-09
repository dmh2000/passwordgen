# Binary name
BINARY=passwordgen

# Default OS is linux
OS ?= linux

# Build directory
BUILD_DIR=build

# Go build command
GOBUILD=go build

# Version information
VERSION ?= 1.0.0
BUILD_TIME=$(shell date +%FT%T%z)

# Build flags
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.BuildTime=${BUILD_TIME}"

# Default target
.DEFAULT_GOAL := build

# Create build directory
$(BUILD_DIR):
	mkdir -p $(BUILD_DIR)

# Clean build directory
.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)

# Build for current platform (defaults to linux)
.PHONY: build
build: $(BUILD_DIR)
	GOOS=$(OS) $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY) main.go

# Build for Linux
.PHONY: linux
linux: OS=linux
linux: build

# Build for macOS
.PHONY: macos
macos: OS=darwin
macos: build

# Build for Windows
.PHONY: windows
windows: OS=windows
windows: BINARY=passwordgen.exe
windows: build

# Build for all platforms
.PHONY: all
all: clean
	GOOS=linux $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY)-linux main.go
	GOOS=darwin $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY)-macos main.go
	GOOS=windows $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY)-windows.exe main.go

# Help target
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  build   - Build for current platform (default: linux)"
	@echo "  linux   - Build for Linux"
	@echo "  macos   - Build for macOS"
	@echo "  windows - Build for Windows"
	@echo "  all     - Build for all platforms"
	@echo "  clean   - Remove build directory"
	@echo "  help    - Show this help message"
