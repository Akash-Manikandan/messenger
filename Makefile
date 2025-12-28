.PHONY: help proto clean build run test deps fmt lint setup rebuild air update-validate

# Default target
help:
	@echo "Available targets:"
	@echo "  proto           - Generate Go code from proto files"
	@echo "  clean           - Clean generated proto files"
	@echo "  build           - Build the application"
	@echo "  run             - Run the application"
	@echo "  test            - Run tests"
	@echo "  deps            - Install Go dependencies"
	@echo "  fmt             - Format code"
	@echo "  lint            - Run linter"
	@echo "  setup           - Setup development environment"
	@echo "  rebuild         - Regenerate proto files and build the application"
	@echo "  air             - Start air for live reload"
	@echo "  update-validate - Update validate proto dependency"

# Generate proto files directly in respective modules
proto:
	@echo "Generating proto files..."
	@cd $(CURDIR) && buf generate
	@echo "Moving proto files to respective modules..."
	@for proto in $(CURDIR)/proto/*.proto; do \
		module=$$(basename $$proto .proto); \
		if [ -d "$(CURDIR)/app-backend/internal/modules/$$module" ]; then \
			mv -f $(CURDIR)/proto/$$module.pb.go $(CURDIR)/app-backend/internal/modules/$$module/ 2>/dev/null || true; \
			mv -f $(CURDIR)/proto/$${module}_grpc.pb.go $(CURDIR)/app-backend/internal/modules/$$module/ 2>/dev/null || true; \
			mv -f $(CURDIR)/proto/$$module.pb.validate.go $(CURDIR)/app-backend/internal/modules/$$module/ 2>/dev/null || true; \
		fi; \
	done
	@echo "Proto generation complete!"

# Clean generated proto files
clean:
	@echo "Cleaning generated proto files..."
	@rm -f $(CURDIR)/proto/*.pb.go
	@rm -f $(CURDIR)/proto/*.pb.validate.go
	@rm -f $(CURDIR)/proto/*_grpc.pb.go
	@find $(CURDIR)/app-backend/internal/modules -type f \( -name "*.pb.go" -o -name "*.pb.validate.go" -o -name "*_grpc.pb.go" \) -delete
	@rm -rf $(CURDIR)/app-backend/bin
	@rm -rf $(CURDIR)/app-backend/tmp
	@echo "Clean complete!"

# Build the application
build:
	@echo "Building application..."
	@cd $(CURDIR)/app-backend && go build -o bin/server cmd/server/main.go
	@echo "Build complete! Binary: app-backend/bin/server"

# Run the application
run:
	@echo "Starting application..."
	@cd $(CURDIR)/app-backend && go run cmd/server/main.go

# Run tests
test:
	@echo "Running tests..."
	@cd $(CURDIR)/app-backend && go test ./...

# Install dependencies
deps:
	@echo "Installing Go dependencies..."
	@cd $(CURDIR)/app-backend && go mod download
	@cd $(CURDIR)/app-backend && go mod tidy
	@echo "Dependencies installed!"

# Format code
fmt:
	@echo "Formatting code..."
	@cd $(CURDIR)/app-backend && go fmt ./...
	@echo "Format complete!"

# Run linter
lint:
	@echo "Running linter..."
	@cd $(CURDIR)/app-backend && golangci-lint run ./...

# Development setup
setup: deps
	@echo "Setting up development environment..."
	@go install github.com/envoyproxy/protoc-gen-validate@latest
	@echo "Setup complete!"

# Regenerate proto and build
rebuild: clean proto build
	@echo "Rebuild complete!"

air:
	@echo "Starting air for live reload..."
	@cd $(CURDIR)/app-backend && air

# Update validate proto dependency
update-validate:
	@echo "Updating validate proto dependency..."
	@cd $(CURDIR) && buf export buf.build/envoyproxy/protoc-gen-validate --output proto
	@echo "Validate proto updated!"
