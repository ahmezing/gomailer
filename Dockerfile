# Stage 1: Build the application
FROM golang:1.24-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o gomailer ./cmd/gomailer




# Stage 2: Create a minimal image
FROM gcr.io/distroless/static:nonroot

# Set working directory
WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/gomailer /app/gomailer

# Copy templates directory
COPY --from=builder /app/templates /app/templates

# Set the command to run the executable
ENTRYPOINT ["/app/gomailer"]