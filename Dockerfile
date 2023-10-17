# Oficial Golang Image
FROM golang:latest

# Working directory
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Download and install any required dependencies
RUN go mod download

# Build the Go app
RUN go build -o main .

# Expose port 8080
EXPOSE 8080

# Run the container
CMD ["/app/main"]
