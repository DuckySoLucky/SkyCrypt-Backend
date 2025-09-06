# Build stage
FROM golang:1.24-alpine AS builder

# Install git for dependency management and ca-certificates for HTTPS requests
RUN apk add --no-cache git ca-certificates

# Set working directory
WORKDIR /app

# Copy go mod files first for better layer caching
COPY go.mod go.sum ./

# Download dependencies (cached layer if go.mod/go.sum unchanged)
RUN go mod download

# Copy source code
COPY . .

# Build the application with optimizations 
RUN CGO_ENABLED=0 GOOS=linux go build \
    -a -installsuffix cgo \
    -buildvcs=false \
    -ldflags="-w -s" \ 
    -o main .

# Final stage - use distroless for security and size
FROM gcr.io/distroless/static-debian12:nonroot

# Copy the binary from builder stage
COPY --from=builder /app/main /main

# Copy assets and other necessary files
COPY --from=builder /app/assets /assets
COPY --from=builder /app/NotEnoughUpdates-REPO /NotEnoughUpdates-REPO

# Expose port
EXPOSE 8080

# Command to run the application
ENTRYPOINT ["/main"]
