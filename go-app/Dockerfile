# Use a minimal base image
FROM alpine:3.21

WORKDIR /app

# Add OCI metadata as labels
ARG IMAGE_AUTHOR
ARG IMAGE_SOURCE
ARG IMAGE_LICENSE
LABEL org.opencontainers.image.authors="${IMAGE_AUTHOR}" \
      org.opencontainers.image.source="${IMAGE_SOURCE}" \
      org.opencontainers.image.licenses="${IMAGE_LICENSE}"

# Copy the pre-built binary and SBOM file, and make the binary executable
COPY ./bin/go-app ./bin/go-app.cdx.json . 
RUN chmod +x go-app

# Expose port
EXPOSE 8080

# Run the binary
CMD ["./go-app"]