# Stage 1: Build the Go binary
FROM golang:1.21 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum before copying the rest for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o executor ./cmd/server

# Stage 2: Create a small runtime image
FROM alpine:latest

# Install ca-certificates for HTTPS support
RUN apk --no-cache add ca-certificates

# Set the working directory inside the final image
WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /app/executor .

# Copy configs if needed (like languages.json)
COPY --from=builder /app/configs ./configs

# Expose the port your app runs on
EXPOSE 8080

# Command to run the binary
CMD ["./executor"]

