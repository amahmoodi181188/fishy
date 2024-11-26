FROM golang:1.19
WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker layer caching
COPY go.mod .             

# Download dependencies
RUN go mod download  

# Copy the application code
COPY . .        

# Build the application
RUN go build -o app1 .

# Set the startup command
CMD ["./app1"]
