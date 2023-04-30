.PHONY: build test install-deps

# Set the binary name and desired output directory
BINARY_NAME=mycli
OUTPUT_DIR=bin

# Build the app
build:
	go build -o $(OUTPUT_DIR)/$(BINARY_NAME)

# Run tests
test:
	go test -v ./...

# Install dependencies
install-deps:
	go mod download