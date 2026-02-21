# ------------------------------
# Stage 1: Build the Go binary
# ------------------------------
FROM golang:1.26-alpine AS builder

# Set working directory inside container
WORKDIR /app

# Copy dependency files first (better layer caching)
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy full source code
COPY . .

# Build static Linux binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app

# ------------------------------
# Stage 2: Minimal runtime image
# ------------------------------
FROM alpine:latest

WORKDIR /root/

# Copy compiled binary from builder stage
COPY --from=builder /app/app .

# Application runs on port 8080
EXPOSE 8080

# Start application
CMD ["./app"]