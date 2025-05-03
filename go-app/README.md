# Go Application

This is a simple Go web application that demonstrates SBOM generation capabilities.

## Project Structure

```text
go-app/
├── src
├──── handlers/         # HTTP handlers and routing
│     └── router.go
├──── main.go          # Application entry point
└──── go.mod           # Go module definition
```

## Getting Started

1. Install dependencies:

    ```bash
    go mod tidy
    ```

2. Run the application:

    ```bash
    go run main.go
    ```

The application will start on port 8080 and can be accessed at <http://localhost:8080>

## Generating SBOM

To generate the SBOM for this Go application, follow the instructions in the main README.md file under the Go section.
