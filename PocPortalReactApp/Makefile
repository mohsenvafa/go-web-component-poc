.PHONY: install build start clean

# Install dependencies
install:
	npm install

# Build the application
build: install
	npm run build

# Start the development server
start: install
	npm start

# Clean build artifacts
clean:
	rm -rf build
	rm -rf node_modules

# Help
help:
	@echo "Available targets:"
	@echo "  install   - Install dependencies"
	@echo "  build     - Build the application"
	@echo "  start     - Start the development server"
	@echo "  clean     - Clean build artifacts"
	@echo "  help      - Show this help"
