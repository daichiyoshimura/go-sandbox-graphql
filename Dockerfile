# Stage 1: Build the Go application
FROM golang:1.22.6 as builder

WORKDIR /app

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the application (optional, depending on your needs)
RUN go build -o myapp ./cmd/server

# Stage 2: Lint and format the code
FROM golang:1.22.6 as linter

WORKDIR /app

# Copy the source code
COPY --from=builder /app /app

# Install linters and formatters
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
RUN go install golang.org/x/tools/cmd/goimports@latest

# Run golangci-lint
RUN golangci-lint run --config .golangci.yml

# Format the code
RUN goimports -w .

# Check for unformatted files
RUN git diff --exit-code