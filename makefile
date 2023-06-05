.PHONY: build test install-deps lint staticcheck gosec vet update-minor upgrade-major check-outdated prune-unused vendor-dependencies

# Set the binary name and desired output directory
BINARY_NAME=mycli
OUTPUT_DIR=bin

# Build the app
build:
	go build -o $(OUTPUT_DIR)/$(BINARY_NAME)

# Run tests
test:
	go test -v ./...

# check
check-all: lint staticcheck gosec vet

lint:
	@echo "Running GoLint..."
	@golint ./...

staticcheck:
	@echo "Running Staticcheck..."
	@staticcheck ./...

gosec:
	@echo "Running GoSec..."
	@gosec ./...

vet:
	@echo "Running Go Vet..."
	@go vet ./...



## paving packageing

# Install dependencies
install-deps:
	go mod download

# delete unused
prune-unused:
	go mod tidy


# vendor dependencies for airgapped building
vendor-dependencies:
	go mod vendor


### update lifecycle

install-libs:
	go get -u golang.org/x/lint/golint
	go get -u honnef.co/go/tools/cmd/staticcheck
	curl -sfL https://raw.githubusercontent.com/securego/gosec/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v2.15.0

# check for outdated
check-outdated:
	go list -u -m -json all

# To update all your dependencies to their latest minor or patch versions
update-minor:
	go get -u ./...

# ! maybe breaking ! To update all your dependencies to their latest major versions
upgrade-major:
	go get -u all
