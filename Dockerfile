# Use Go 1.23 bookworm as base image
FROM golang:1.23-bookworm AS base

# Development stage
# =============================================================================
# Create a development stage based on the "base" image
FROM base AS development

# Change the working directory to /app
WORKDIR /app

# Install the air CLI for auto-reloading
RUN go install github.com/air-verse/air@latest

# Copy the go.mod and go.sum files to the /app directory
COPY go.mod go.sum ./

# Install dependencies
RUN go mod download

# Start air for live reloading
CMD ["air"]

# Builder stage
# =============================================================================
# Create a builder stage based on the "base" image
FROM base AS builder

# Move to working directory /build
WORKDIR /build

# Copy the go.mod and go.sum files to the /build directory
COPY go.mod go.sum ./

# Install dependencies
RUN go mod download

# Copy the entire source code into the container
COPY . .

# Build the application
RUN CGO_ENABLED=0 go build -o tutorme

# Production stage
# =============================================================================
# Use a minimal image for production
FROM debian:bookworm-slim AS production

# Move to working directory /prod
WORKDIR /prod

# Install certificates (explicitly ensuring they are up-to-date)
RUN \
  apt-get update && \
  apt-get install -y ca-certificates && \
  apt-get clean

# Copy binary from builder stage
COPY --from=builder /build/tutorme ./

LABEL org.opencontainers.image.source https://github.com/tutorme-verse/tutorme-backend

# Start the application
CMD ["/prod/tutorme"]
