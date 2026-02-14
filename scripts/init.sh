#!/bin/bash
# Initialization script - Create necessary directories and config files

echo "Creating directory structure..."
mkdir -p logs static/uploads static/avatar config nginx/conf.d ssl nginx/certbot

# Create backend config.yaml (if not exists)
if [ ! -f "config/config.yaml" ]; then
    echo "Creating config/config.yaml (from example.yaml)..."
    cp config/example.yaml config/config.yaml
    echo "✓ config.yaml created"
else
    echo "✓ config.yaml already exists"
fi

# Create frontend .env (if not exists)
if [ ! -f "web/.env" ]; then
    echo "Creating web/.env (from .env.example)..."
    cp web/.env.example web/.env
    echo "✓ web/.env created (configure VITE_API_KEY for AI features)"
else
    echo "✓ web/.env already exists"
fi

echo ""
echo "✓ Initialization complete"
echo ""
echo "Next steps:"
echo "  1. Configure backend: vim config/config.yaml"
echo "  2. Configure compose vars: cp .env.example .env && vim .env"
echo "  3. Configure frontend vars: vim web/.env"
echo "  4. Start services: docker compose up -d"
