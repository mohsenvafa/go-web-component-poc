.PHONY: help setup run-all clean-all stop-all stop-port

# Show help
help:
	@echo "Go Web Component POC - Available targets:"
	@echo ""
	@echo "Setup:"
	@echo "  setup      - Setup all projects (install dependencies)"
	@echo ""
	@echo "Running:"
	@echo "  run-all    - Run all services (requires 3 terminals)"
	@echo ""
	@echo "Individual services:"
	@echo "  run-webcomponent  - Run PatientProfileGoWebComponent (port 8091)"
	@echo "  run-go-portal     - Run PocPortalGoSurfaceApp (port 8092)"
	@echo "  run-react-portal  - Run PocPortalReactApp (port 8093)"
	@echo ""
	@echo "Stopping services:"
	@echo "  stop-all   - Stop all POC services (ports 8091, 8092, 8093)"
	@echo "  stop-port  - Stop service on specific port (usage: make stop-port PORT=8091)"
	@echo ""
	@echo "Cleanup:"
	@echo "  clean-all  - Clean all build artifacts"
	@echo ""
	@echo "Help:"
	@echo "  help       - Show this help"

# Setup all projects
setup:
	@echo "Setting up PatientProfileGoWebComponent..."
	cd PatientProfileGoWebComponent && make deps
	@echo "Setting up PocPortalGoSurfaceApp..."
	cd PocPortalGoSurfaceApp && make deps
	@echo "Setting up PocPortalReactApp..."
	cd PocPortalReactApp && make install
	@echo "Setup complete!"

# Run all services (for demonstration - requires multiple terminals)
run-all:
	@echo "To run all services, open 3 terminals and run:"
	@echo "Terminal 1: make run-webcomponent"
	@echo "Terminal 2: make run-go-portal"
	@echo "Terminal 3: make run-react-portal"
	@echo ""
	@echo "Then visit:"
	@echo "- http://localhost:8091 (Web Component Demo)"
	@echo "- http://localhost:8092 (Go Portal)"
	@echo "- http://localhost:8093 (React Portal)"

# Run individual services
run-webcomponent:
	@echo "Starting PatientProfileGoWebComponent on port 8091..."
	cd PatientProfileGoWebComponent && make run

run-go-portal:
	@echo "Starting PocPortalGoSurfaceApp on port 8092..."
	cd PocPortalGoSurfaceApp && make run

run-react-portal:
	@echo "Starting PocPortalReactApp on port 8093..."
	cd PocPortalReactApp && make start

# Stop all POC services
stop-all:
	@echo "Stopping all POC services..."
	./stop-services.sh

# Stop service on specific port
stop-port:
	@if [ -z "$(PORT)" ]; then \
		echo "Usage: make stop-port PORT=<port_number>"; \
		echo "Example: make stop-port PORT=8091"; \
		exit 1; \
	fi
	@echo "Stopping service on port $(PORT)..."
	./kill-port.sh $(PORT)

# Clean all build artifacts
clean-all:
	@echo "Cleaning PatientProfileGoWebComponent..."
	cd PatientProfileGoWebComponent && make clean
	@echo "Cleaning PocPortalGoSurfaceApp..."
	cd PocPortalGoSurfaceApp && make clean
	@echo "Cleaning PocPortalReactApp..."
	cd PocPortalReactApp && make clean
	@echo "Cleanup complete!"
