# Warren Clough Website Makefile

.PHONY: help build run test clean deploy seed-data

# Default target
help:
	@echo "Available commands:"
	@echo "  build      - Build the Go application"
	@echo "  run        - Run the application locally"
	@echo "  test       - Run tests"
	@echo "  clean      - Clean build artifacts"
	@echo "  deploy     - Deploy to Firebase"
	@echo "  seed-data  - Seed Firestore with sample data"
	@echo "  dev        - Run in development mode with hot reload"

# Build the application
build:
	@echo "Building application..."
	go mod tidy
	go build -o bin/warrenclough main.go
	go build -o bin/functions functions.go

# Run the application locally
run:
	@echo "Starting local server..."
	go run main.go

# Run tests
test:
	@echo "Running tests..."
	go test ./...

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	rm -rf bin/
	go clean

# Deploy to Firebase
deploy:
	@echo "Deploying to Firebase..."
	go mod tidy
	go build -o functions functions.go
	firebase deploy

# Seed Firestore with sample data
seed-data:
	@echo "Seeding Firestore with sample data..."
	go run scripts/seed-data.go

# Development mode with hot reload (requires air)
dev:
	@echo "Starting development server with hot reload..."
	@if command -v air > /dev/null; then \
		air; \
	else \
		echo "Air not installed. Install with: go install github.com/cosmtrek/air@latest"; \
		echo "Falling back to regular run..."; \
		go run main.go; \
	fi

# Install development dependencies
install-dev:
	@echo "Installing development dependencies..."
	go install github.com/cosmtrek/air@latest

# Format code
fmt:
	@echo "Formatting code..."
	go fmt ./...

# Lint code
lint:
	@echo "Linting code..."
	@if command -v golangci-lint > /dev/null; then \
		golangci-lint run; \
	else \
		echo "golangci-lint not installed. Install with: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; \
	fi
