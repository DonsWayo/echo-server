FROM golang:1.24-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o echo-server ./cmd/echo-server

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app

COPY --from=builder /app/echo-server .

USER appuser

EXPOSE 80

ENV PORT=80 \
    ENABLE__HOST=true \
    ENABLE__HTTP=true \
    ENABLE__REQUEST=true \
    ENABLE__COOKIES=true \
    ENABLE__HEADER=true \
    ENABLE__ENVIRONMENT=true \
    ENABLE__FILE=true

# Run the application
CMD ["./echo-server"]
