FROM golang:1.21.5-alpine AS builder

RUN mkdir /app
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the binary
RUN go build -o main .

# Runner Stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

# Volume for SSL/TLS Certificates(optional for development)
VOLUME /certs

EXPOSE 3000

CMD ["/app/main"]
