# Use the official Golang image to create a build environment
FROM golang:1.23-alpine AS builder

# Install build dependencies (SQLite and C compiler)
RUN apk add --no-cache gcc musl-dev sqlite-dev

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application with CGO enabled
RUN CGO_ENABLED=1 GOOS=linux go build -o to-do-list-app-backend ./cmd

# Use a minimal Alpine image for the final stage
FROM alpine:latest

# Install SQLite in the final image
RUN apk add --no-cache sqlite-dev

# Set the working directory
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/to-do-list-app-backend .

# Expose the application port
EXPOSE 8080

# Command to run the application
CMD ["./to-do-list-app-backend"]