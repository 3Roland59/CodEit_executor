# Stage 1: Pull required language images
FROM docker:24.0.7-cli AS puller

RUN docker pull python:3.12-alpine && \
    docker pull node:20-alpine && \
    docker pull openjdk:17-alpine && \
    docker pull golang:1.24-alpine && \
    docker pull gcc:13.2.1-alpine

# Stage 2: Build the Go binary
FROM golang:1.24.2-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o executor ./cmd/server

# Stage 3: Final runtime image
FROM alpine:3.19

RUN apk --no-cache add ca-certificates docker-cli curl

RUN mkdir -p /code && chmod 777 /code

WORKDIR /root/

COPY --from=builder /app/executor .
COPY --from=builder /app/configs ./configs

EXPOSE 8080

CMD ["./executor"]

