# Stage 1: Build
FROM docker.1ms.run/node:20-alpine AS builder

WORKDIR /app

# Copy package files
COPY package*.json ./

# Install dependencies
RUN npm ci --legacy-peer-deps

# Copy source code
COPY . .

# Build the application
RUN npm run build

# Stage 2: Production with Nginx
FROM docker.1ms.run/nginx:alpine

# Store build output in a source directory; on startup, sync to the volume mount point.
# This ensures the shared volume always has the latest build when the container restarts.
COPY --from=builder /app/dist /app/dist-source

# Copy nginx configuration
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Expose port
EXPOSE 80

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --quiet --tries=1 --spider http://localhost/health || exit 1

# Sync build output to volume, then start nginx
CMD cp -r /app/dist-source/* /usr/share/nginx/html/ && nginx -g "daemon off;"
