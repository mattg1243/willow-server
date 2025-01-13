FROM golang:1.22-bullseye

WORKDIR /app
# Copy all files 
COPY . .
# Install deps
RUN go mod download
# Build the app
RUN go build -o main 

EXPOSE 8080

CMD ["./main"]