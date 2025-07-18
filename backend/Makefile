# NestMate Backend Makefile

.PHONY: build run test clean deps lint

# Build the application
build:
	go build -o bin/server cmd/server/main.go

# Run the application
run:
	go run cmd/server/main.go

# Run tests
test:
	go test -v ./...

# Run tests with coverage
test-coverage:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# Clean build artifacts
clean:
	rm -rf bin/
	rm -f coverage.out coverage.html

# Download dependencies
deps:
	go mod download
	go mod tidy

# Run linter
lint:
	golangci-lint run

# Run the application in development mode with hot reload
dev:
	air

# Docker commands
docker-build:
	docker build -t nestmate-backend .

docker-run:
	docker run -p 8080:8080 nestmate-backend

# Database commands
db-migrate:
	@echo "Database migrations will be implemented in later tasks"

db-seed:
	@echo "Database seeding will be implemented in later tasks"