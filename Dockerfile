# Stage 1: Build
FROM golang AS builder

WORKDIR /app

COPY . .

# Build binary and name it grpc-server to avoid naming conflict
RUN CGO_ENABLED=0 go build -o grpc-server /app/grpc/cmd/main.go

# Stage 2: Final image
FROM alpine:latest

WORKDIR /app

# You may need ca-certificates if using HTTPS
RUN apk --no-cache add ca-certificates

# Copy binary only
COPY --from=builder /app/grpc-server /app/grpc-server

# Run the binary
CMD ["./grpc-server"]
