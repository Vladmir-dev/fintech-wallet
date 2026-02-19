# Stage 1: Build the Go binary
FROM golang:1.23-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum first (for caching dependencies)
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the binary (replace 'fintech-wallet' with your binary name if different)
RUN go build -o fintech-wallet ./cmd/api/main.go

# Stage 2: Create a lightweight runtime image
FROM alpine:latest

# Set working directory
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/fintech-wallet .

# Expose the port your app runs on (8080 from your main.go)
EXPOSE 8080

# Run the binary
CMD ["./fintech-wallet"]