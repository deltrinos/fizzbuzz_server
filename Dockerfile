# Stage 1: Build the Go binary
FROM golang:latest AS builder

WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o fizzbuzz_server cmd/main.go

# Stage 2: Create the final Docker image
FROM gcr.io/distroless/base-debian10

WORKDIR /app

# Copy the Go binary from the builder stage
COPY --from=builder /app/fizzbuzz_server .

# Expose the port on which the server will listen
EXPOSE 8000

# Command to run the server
CMD ["./fizzbuzz_server"]
