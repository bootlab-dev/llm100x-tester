.PHONY: build test clean run install uninstall

# Build tester
build:
	go build -o llm100x-tester .

# Install to $GOPATH/bin (global access)
install:
	go install .

# Uninstall
uninstall:
	rm -f $(shell go env GOPATH)/bin/llm100x-tester

# Run tests
test:
	go test -v ./...

# Clean build artifacts
clean:
	rm -f llm100x-tester
	go clean

# Run tester (requires working directory)
run:
	go run . $(ARGS)

# Install dependencies
deps:
	go mod download
	go mod tidy

# Format code
fmt:
	go fmt ./...

# Lint code
lint:
	golangci-lint run ./...
