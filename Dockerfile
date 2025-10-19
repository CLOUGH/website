# Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copy source code
COPY . .

# Ensure dependencies are correct
RUN go mod tidy

# Download dependencies
RUN go mod download

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Run stage
FROM alpine:latest

# Install ca-certificates for HTTPS
RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy the binary and necessary files from the builder stage
COPY --from=builder /app/main .
COPY --from=builder /app/web ./web

# Expose port (Cloud Run will set PORT env var)
EXPOSE 8080
ENV PORT=8080

# Run the application
CMD ["./main"]
