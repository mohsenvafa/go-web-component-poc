.PHONY: build run clean generate

# Generate templ files
generate:
	templ generate

# Build the application
build: generate
	go build -o portal-server .

# Run the application
run: generate
	go run .

# Clean build artifacts
clean:
	rm -f portal-server
	rm -rf components/*.go

# Install dependencies
deps:
	go mod tidy
	go mod download

# Run with hot reload (requires air)
dev:
	air

# Help
help:
	@echo "Available targets:"
	@echo "  generate  - Generate templ files"
	@echo "  build     - Build the application"
	@echo "  run       - Run the application"
	@echo "  clean     - Clean build artifacts"
	@echo "  deps      - Install dependencies"
	@echo "  dev       - Run with hot reload (requires air)"
	@echo "  help      - Show this help"
