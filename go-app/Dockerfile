# Use a minimal base image
FROM alpine:3.21

WORKDIR /app

# Copy the pre-built binary from the CI pipeline
COPY ./go-app/bin/go-app .

# Expose port
EXPOSE 8080

# Run the binary
CMD ["./go-app"]