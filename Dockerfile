##### Stage 1: Build Stage
# Use a lightweight Go image with Alpine for the build environment
FROM golang:1.23.2-alpine3.20 AS builder

# Install necessary tools
RUN apk add --no-cache git

# Set the working directory inside the container
WORKDIR /pawnshop_backend

# Copy Go dependency files (go.mod, go.sum) for efficient layer caching
COPY go.mod go.sum ./

# Use Go module proxy to speed up dependency downloads
ENV GOPROXY=direct

# Copy the entire source code into the working directory
COPY . .

# Disable CGO for a fully static binary, enhancing cross-platform compatibility
ENV CGO_ENABLED=0

# Compile the Go application targeting Linux OS
RUN go build -o server cmd/server/main.go

##### Stage 2: Deployment Stage
# Use a minimal Alpine Linux image for the final application image
FROM alpine:latest

# Set the working directory for the application
WORKDIR /pawnshop_backend

# Copy the built server binary and necessary files from the build stage
COPY --from=builder /pawnshop_backend/server .
COPY --from=builder /pawnshop_backend/go.mod .
COPY --from=builder /pawnshop_backend/env ./env

# Set environment variables for time zone, Go Gin mode, and environment configuration
ENV TZ=Asia/Ho_Chi_Minh
ENV GIN_MODE=release 
ENV ENVIRONMENT=production

# Specify the entry point to start the server
ENTRYPOINT ./server

# Expose the application port (5000) for external access
EXPOSE 8000
