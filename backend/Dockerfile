# Dockerfile for Go backend
FROM golang:alpine AS build

# Set working directory inside the container
WORKDIR /app

# Copy Go module files
COPY backend/go.mod backend/go.sum ./

# Download dependencies
RUN go mod tidy

# Copy the rest of the Go code
COPY backend/ .

# Build the Go application
RUN go build -o server .

# Use a minimal Alpine image to run the application
FROM alpine:latest

# Set working directory inside the container
WORKDIR /app

# Copy the built binary from the build stage
COPY --from=build /app/server /app/

# Copy the images directory into the container
COPY backend/images/ /app/images/

# Expose the port that your backend will use
EXPOSE 8080

# Start the Go server
CMD ["./server"]
