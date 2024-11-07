# Use a lightweight, official Go image as the base
FROM golang:1.20-alpine AS builder

# Set the working directory inside the image
WORKDIR /app

# Copy the Go module and its dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Set environment variables for the build stage
ENV DB_HOST=db \
    DB_PORT=5432 \
    DB_USER=postgres \
    DB_PASSWORD=postgres \
    DB_NAME=auth_db \
    JWT_SECRET=mysecret \
    REDIS_HOST=redis \
    REDIS_PORT=6379

# Build the Go application
WORKDIR /app/cmd/auth-service
RUN go build -o main .

# Use a smaller, production-ready image for the final image
FROM alpine:latest

# Copy the built binary from the builder stage
COPY --from=builder /app/cmd/auth-service/main /app/main

# Set the working directory
WORKDIR /app

# Expose the port your application listens on (e.g., 8080)
EXPOSE 8080

# Command to run the application when the container starts
CMD ["./main"]