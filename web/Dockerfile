# Stage 1: Build
FROM docker.1ms.run/node:20-alpine AS builder

WORKDIR /app

# Copy package files
COPY package*.json ./

# Inject build-time env for Vite (Gemini API key)
ARG VITE_API_KEY
ENV VITE_API_KEY=${VITE_API_KEY}
ARG VITE_DEFAULT_HOMEPAGE=brand
ENV VITE_DEFAULT_HOMEPAGE=${VITE_DEFAULT_HOMEPAGE}

# Install dependencies
RUN npm ci --legacy-peer-deps

# Copy source code
COPY . .

# Build the application
RUN npm run build

# Stage 2: Production with Nginx
FROM docker.1ms.run/nginx:alpine

# Copy built files from builder
COPY --from=builder /app/dist /usr/share/nginx/html

# Copy nginx configuration
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Expose port
EXPOSE 80

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --quiet --tries=1 --spider http://localhost/health || exit 1

# Run nginx
CMD ["nginx", "-g", "daemon off;"]
