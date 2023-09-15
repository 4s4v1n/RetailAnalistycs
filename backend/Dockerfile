FROM golang:1.20-alpine AS builder
WORKDIR /app

COPY ./ ./
RUN go mod download
RUN go build -o apg6 ./cmd/apg6/main.go
EXPOSE 8080
CMD [ "/app/apg6" ]

FROM golang:latest

# Working directory
WORKDIR /app

# Copy everything at /app
COPY . /app

# Build the go app
RUN go mod download
RUN go build -o main ./cmd/apg6/main.go

# Expose port
EXPOSE 8080

# Define the command to run the app
CMD ["./main"]
