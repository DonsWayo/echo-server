.PHONY: run test

# Run the application
run:
	go run ./cmd/echo-server/main.go

# Run tests
test:
	go test -v ./...

# Run with hot reload (requires air: https://github.com/cosmtrek/air)
dev:
	air -c .air.toml

# Build Docker image
docker-build:
	docker build -t echo-server .

# Run Docker container
docker-run:
	docker run -p 80:80 echo-server
