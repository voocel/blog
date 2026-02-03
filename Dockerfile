# Stage 1: Build
FROM docker.m.daocloud.io/golang:1.25-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git

WORKDIR /app

# Set Go proxy for China (use multiple mirrors for reliability)
ENV GOPROXY=https://goproxy.cn,https://mirrors.aliyun.com/goproxy/,https://goproxy.io,direct
ENV GO111MODULE=on
ENV GOSUMDB=off

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies with retry and timeout settings
RUN go env -w GOPROXY=${GOPROXY} && \
    go env -w GOSUMDB=off && \
    go mod download

# Copy source code
COPY . .

# Build the application with optimizations for smaller binary and less memory usage
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-w -s" -trimpath \
    -o /app/bin/blog ./cmd/blog

# Stage 2: Runtime
FROM docker.1ms.run/alpine:latest

# Install ca-certificates for HTTPS
RUN apk --no-cache add ca-certificates tzdata

# Set timezone
ENV TZ=Asia/Shanghai

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/bin/blog .

# Copy config files
COPY --from=builder /app/config ./config

# Create config.yaml from example.yaml
RUN cp config/example.yaml config/config.yaml

# Copy static files
COPY --from=builder /app/static ./static

# Create logs directory
RUN mkdir -p logs

# Expose port
EXPOSE 8080

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --quiet --tries=1 --spider http://localhost:8080/api/v1/health || exit 1

# Run the application
CMD ["./blog"]
