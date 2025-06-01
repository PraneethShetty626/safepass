.PHONY: fmt lint

# Format the code
fmt:
	go fmt ./...
	goimports -w .
	golines --max-len=100 -w .

# Run linter
lint:
	golangci-lint run

# Format + lint together
check: fmt lint

# Build the project
build:
	go build -o safepass ./

# Test the project
test:
	go test ./...

# Clean build artifacts
clean:
	rm -f safepass
