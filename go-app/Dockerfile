# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY src/go.mod src/go.sum ./

# Download dependencies
RUN go mod download

# Copy source code and Makefile
COPY src/ .
COPY Makefile .

# Build the application using Makefile
RUN make build

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/bin/go-app .

# Expose port
EXPOSE 8080

# Run the binary
CMD ["./go-app"] 