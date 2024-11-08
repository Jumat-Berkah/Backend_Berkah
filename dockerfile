# Use a Debian-based Go image for broader compatibility
FROM golang:1.20

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the application code
COPY . .

# Build the application with verbose output for debugging
RUN go build -x -o main .

# Expose the port that the app runs on
EXPOSE 8080

# Run the application
CMD ["./main"]
