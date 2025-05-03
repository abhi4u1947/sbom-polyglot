# Use a minimal base image
FROM alpine:3.21

WORKDIR /app

# Copy the pre-built binary and SBOM file, and make the binary executable
COPY ./bin/go-app ./bin/go-app.cdx.json . 
RUN chmod +x go-app

# Expose port
EXPOSE 8080

# Run the binary
CMD ["./go-app"]