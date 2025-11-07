# Build stage
FROM golang:1.23-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o smr-api ./cmd/main.go

# Final stage
FROM golang:1.23-alpine

WORKDIR /app/bin

# Copy the binary from builder
COPY --from=builder /app/smr-api .

# Expose port
EXPOSE 8080

# Command to run
CMD ["./smr-api"]
