# Use the official Golang image as a base
FROM golang:1.21-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the port your Echo server listens on
EXPOSE 8080

# Command to run the executable
CMD ["/app/main"]
