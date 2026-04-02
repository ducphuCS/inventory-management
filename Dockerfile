# --- Build Stage 1: Compile the Go Microservice ---
FROM golang:1.25 AS builder

WORKDIR /app

# Install build dependencies
# RUN apk add --no-cache gcc musl-dev

# Copy dependency files first for caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire source code
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -o inventory-service main.go

# --- Build Stage 2: Final Runtime Environment ---
FROM python:3.12-alpine

WORKDIR /app

# 1. Install system dependencies (including uv)
RUN pip install --no-cache-dir uv

# 2. Copy the compiled Go binary from the builder stage
COPY --from=builder /app/inventory-service .

# 3. Copy our Python data analysis engine and configuration
COPY scripts/ ./scripts/
COPY pyproject.toml .

# 4. Initialize the uv virtual environment for the Python part
RUN uv sync --no-dev

# 5. Prepare the data directory for SQLite persistence
RUN mkdir -p /app/data
ENV DB_PATH=/app/data/inventory.db

# Expose port for the service
EXPOSE 8080

# Start the microservice
CMD ["./inventory-service"]
