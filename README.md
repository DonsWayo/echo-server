# Echo Server

A feature-rich echo server built with Go and Fiber v3, inspired by [Ealenn/Echo-Server](https://github.com/Ealenn/Echo-Server).

## Features

- Supports GET, POST, PUT, PATCH, DELETE methods
- Customizable responses via headers or query parameters
- Custom HTTP status codes
- Custom response bodies
- Custom response headers
- Custom response latency
- Environment variable access
- File/folder exploration
- Health check endpoint

## Requirements

- Go 1.24 or higher (for development)
- Docker (for containerized deployment)

## Configuration

The server can be configured using environment variables:

| Environment Variable | Default | Description |
|----------------------|---------|-------------|
| PORT | 80 | Port to listen on |
| LOGS__IGNORE__PING | false | Ignore ping requests in logs |
| ENABLE__HOST | true | Enable host information in response |
| ENABLE__HTTP | true | Enable HTTP information in response |
| ENABLE__REQUEST | true | Enable request information in response |
| ENABLE__COOKIES | true | Enable cookies information in response |
| ENABLE__HEADER | true | Enable header information in response |
| ENABLE__ENVIRONMENT | true | Enable environment variables in response |
| ENABLE__FILE | true | Enable file/folder exploration |
| CONTROLS__TIMES__MIN | 0 | Minimum time delay in milliseconds |
| CONTROLS__TIMES__MAX | 60000 | Maximum time delay in milliseconds |


## Installation

### Using Go

```bash
# Clone the repository
git clone https://github.com/a-safe-digital/echo-server.git
cd echo-server

# Install dependencies
go mod download

# Run the server
go run ./cmd/echo-server/main.go

# Or use the Makefile
make run
```

### Using Docker

```bash
# Build the Docker image
docker build -t echo-server .

# Run the container
docker run -p 8080:80 echo-server

# Or use the Makefile
make docker-build
make docker-run
```

## API Usage

### Custom HTTP Status Code

```bash
# Using headers
curl -I --header 'X-ECHO-CODE: 404' localhost:8080

# Using query parameters
curl -I localhost:8080/?echo_code=404

# Multiple status codes (random selection)
curl -I localhost:8080/?echo_code=200-400-500
```

### Custom Response Body

```bash
# Using headers
curl --header 'X-ECHO-BODY: Hello World' localhost:8080

# Using query parameters
curl localhost:8080/?echo_body=Hello%20World
```

### Custom Response Headers

```bash
# Using headers
curl --header 'X-ECHO-HEADER: Custom-Header:Value' localhost:8080

# Using query parameters
curl "localhost:8080/?echo_header=Custom-Header:Value"
```

### Custom Response Latency

```bash
# Add a 5-second delay
curl --header 'X-ECHO-TIME: 5000' localhost:8080
curl "localhost:8080/?echo_time=5000"
```

### File/Folder Exploration

```bash
# List directory contents
curl --header 'X-ECHO-FILE: /' localhost:8080
curl "localhost:8080/?echo_file=/"
```

### Environment Variable Access

```bash
# Get an environment variable
curl --header 'X-ECHO-ENV-BODY: HOSTNAME' localhost:8080
curl "localhost:8080/?echo_env_body=HOSTNAME"
```

### Combine Custom Actions

```bash
# Combine body and status code
curl --header 'X-ECHO-CODE: 401' --header 'X-ECHO-BODY: Unauthorized' localhost:8080
curl "localhost:8080/?echo_body=Unauthorized&echo_code=401"
```

### Health Check

```bash
curl localhost:8080/health
```
