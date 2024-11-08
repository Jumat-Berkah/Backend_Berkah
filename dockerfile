# Use a Go 1.22 Alpine image as a base to match the go.mod version
FROM golang:1.22-alpine

# Install essential build tools
RUN apk update && apk add --no-cache git gcc musl-dev

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the application
RUN go build -o main .

# Expose the port that the app runs on
EXPOSE 8080

# Run the application
CMD ["./main"]
