# Stage 1: Build the Go application
FROM golang:1.23rc2 as builder

WORKDIR /app

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Install linters and formatters
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
RUN go install golang.org/x/tools/cmd/goimports@latest

# Build the application (optional, depending on your needs)
RUN go build -o myapp ./...

# Stage 2: Lint and format the code
FROM golang:1.23rc2 as linter

WORKDIR /app

# Copy the source code
COPY --from=builder /app /app

# Run golangci-lint
RUN golangci-lint run --config .golangci.yml

# Format the code
RUN goimports -w .

# Check for unformatted files
RUN git diff --exit-code