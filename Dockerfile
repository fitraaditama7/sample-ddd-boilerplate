# Stage 1: Build the application
FROM golang:1.19 AS builder
WORKDIR /app
# Copy the source code
COPY . .
# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o http-server ./cmd/main.go

# Stage 2: Run the application
FROM alpine:latest
WORKDIR /app
# Copy the binary from the builder stage
COPY --from=builder /app/http-server .

# Expose port 8080
EXPOSE 8080
# Run the binary
CMD ["./http-server"]
