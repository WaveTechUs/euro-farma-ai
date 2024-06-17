# Simple Makefile for a Go project

# Build the application
all: build

build:
	@echo "Building..."
		
	@go build -o main.exe cmd/api/main.go

# Run the application + Swager Init
run:
	@swag init -o .\cmd\api\swagger\ -g .\internal\server\routes.go
	@go run cmd/api/main.go

# Create DB container
docker-run:
	@if docker compose up 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose up; \
	fi

# Shutdown DB container
docker-down:
	@if docker compose down 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose down; \
	fi

# Test the application
test:
	@echo "Testing..."
	@go test ./tests -v

migrate:
	@migrate -path cmd/migrate/ -database "mysql://wave:wave123@tcp(localhost:3306)/eurofarma?query" -verbose up

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main

# Swagger init
swag:
	@swag init -o .\cmd\api\swagger\ -g .\internal\server\routes.go

# Live Reload
watch:
	@if command -v air > /dev/null; then \
	    air; \
	    echo "Watching...";\
	else \
	    read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
	    if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
	        go install github.com/cosmtrek/air@latest; \
	        air; \
	        echo "Watching...";\
	    else \
	        echo "You chose not to install air. Exiting..."; \
	        exit 1; \
	    fi; \
	fi

.PHONY: all build run test clean
