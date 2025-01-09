# Use the official Golang image as the base image
FROM golang:1.21-alpine

# Install Git (required to clone the repository)
RUN apk add --no-cache git

# Set the working directory inside the container
WORKDIR /app

# Clone the Git repository
RUN git clone -b project/web-server https://github.com/your-username/go-server.git .

# Download and install Go dependencies
RUN go mod init go-server
RUN go mod tidy

# Build the Go application
RUN go build -o go-server .

# Expose the port your application will run on
EXPOSE 8080

# Command to run the application
CMD ["./go-server"]