.PHONY: run test test-coverage dev docker-build docker-run clean

run:
	go run ./cmd/echo-server/main.go

test:
	go test -v ./tests/... ./internal/...

test-coverage:
	go test -v -coverprofile=coverage.out ./tests/... ./internal/...
	go tool cover -html=coverage.out -o coverage.html

dev:
	air -c .air.toml

docker-build:
	docker build -t echo-server .

docker-run:
	docker run -p 80:80 echo-server

clean:
	rm -rf tmp/ coverage.out coverage.html
