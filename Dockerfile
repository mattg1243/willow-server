FROM golang:1.22-bullseye

WORKDIR /app
# Copy go mod and sum files
COPY go.mod go.sum ./
# Install deps
RUN go mod download
# Copy all other files 
COPY . .
# Run all tests
RUN go test ./...
# Build the app
RUN go build -o main 
# Expose the port
EXPOSE 8080
# Run the server
CMD ["./main"]