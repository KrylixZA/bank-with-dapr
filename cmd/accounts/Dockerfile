FROM golang:1.22.5

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project or the necessary directories, preserving the structure
COPY . .

# Build the Go application
RUN go build -o /dist cmd/accounts/main.go

# Expose the application port
EXPOSE 8080

# Options to run in debug or release mode
ARG CONFIG=debug
ENV GIN_MODE=${CONFIG}

# Run the compiled application
CMD ["/dist"]
