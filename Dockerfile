# Oficial Golang Image
FROM golang:1.21.3-alpine3.18

# Working directory
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Download and install any required dependencies
RUN go mod download

# DON'T DO IT, BE BETTER THAN ME.
ENV MONGOURI=mongodb+srv://davidfelipehernandez1:rvVAz2hgfAFTSVt1@dbm.h04vkul.mongodb.net/

# Build the Go app
RUN go build -o main .

# Expose port 8080
EXPOSE 8080

# Run the container
CMD ["/app/main"]
