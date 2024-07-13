# Use the official Golang image to build the binary
FROM golang:1.21 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app for Linux architecture
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/api

# Use a minimal Docker image to reduce the size of the final image
FROM alpine:latest

# Install any necessary runtime dependencies (if any)
RUN apk add --no-cache ca-certificates

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Ensure the binary is executable
RUN chmod +x ./main

# Expose port 4001 to the outside world
EXPOSE 4001

# Command to run the executable
CMD ["./main"]
