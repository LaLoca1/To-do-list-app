# Use an official Golang image to build the Go app with CGO support
FROM golang:1.23 AS builder 

# Set the Current Working Directory inside the container 
WORKDIR /app 

# Install dependencies required for CGO (SQLite and build tools)
RUN apt-get update && apt-get install -y \
    gcc \
    make \
    libc6-dev \
    libsqlite3-dev

# Copy go.mod and go.sum files for dependency resolution 
COPY go.mod go.sum ./ 

# Download all dependencies
RUN go mod tidy

# Copy entire backend source code into the container 
COPY . . 

# Build the Go app with CGO enabled
WORKDIR /app/cmd 
RUN CGO_ENABLED=1 go build -o to-do-list-app-backend 

# Start a new stage for the final image 
FROM alpine:latest

# Install necessary dependencies for the final image (SQLite)
RUN apk add --no-cache sqlite-dev

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the binary from the build stage
COPY --from=builder /app/cmd/to-do-list-app-backend .

# Expose port
EXPOSE 8080

# # Command to run the executable
# CMD ["./to-do-list-app-backend"]
