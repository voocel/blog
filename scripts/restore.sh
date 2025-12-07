#!/bin/bash

# Data restore script
set -e

GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m'

echo "======================================="
echo "   Data Restore Tool"
echo "======================================="
echo ""

# Input backup file
read -p "Please enter backup file name: " backup_file

if [ ! -f "$backup_file" ]; then
    echo -e "${RED}Error: File $backup_file not found${NC}"
    exit 1
fi

echo -e "${GREEN}Starting data restore...${NC}"
echo ""

# Extract
BACKUP_DIR=$(basename "$backup_file" .tar.gz)
tar -xzf "$backup_file"

# 1. Restore config files
echo "1/6 Restoring config files..."
if [ -f "$BACKUP_DIR/.env" ]; then
    cp "$BACKUP_DIR/.env" .
    echo -e "${GREEN}✓ Config files restored${NC}"
else
    echo -e "${YELLOW}! No .env in backup, please create manually${NC}"
fi

# 2. Restore SSL certificates (if exists)
echo "2/6 Restoring SSL certificates..."
if [ -d "$BACKUP_DIR/ssl" ]; then
    cp -r "$BACKUP_DIR/ssl" .
    mkdir -p nginx/conf.d
    cp -r "$BACKUP_DIR/conf.d/"* nginx/conf.d/
    echo -e "${GREEN}✓ SSL certificates restored${NC}"
else
    echo "Skipping SSL restore"
fi

# 3. Restore GeoIP database (if exists)
echo "3/6 Restoring GeoIP database..."
if [ -f "$BACKUP_DIR/config/GeoLite2-City.mmdb" ]; then
    mkdir -p config
    cp "$BACKUP_DIR/config/GeoLite2-City.mmdb" config/
    echo -e "${GREEN}✓ GeoIP database restored${NC}"
else
    echo "Skipping GeoIP restore"
fi

# 4. Start database
echo "4/6 Starting database..."
docker-compose up -d postgres
echo "Waiting for database to start..."
sleep 15
echo -e "${GREEN}✓ Database started${NC}"

# 5. Restore database data
echo "5/6 Restoring database..."
docker exec -i blog-postgres psql -U postgres blog < "$BACKUP_DIR/database.sql"
echo -e "${GREEN}✓ Database data restored${NC}"

# 6. Restore static files and start all services
echo "6/6 Restoring static files and starting services..."
if [ -f "$BACKUP_DIR/static.tar.gz" ]; then
    tar -xzf "$BACKUP_DIR/static.tar.gz"
fi
docker-compose up -d
echo -e "${GREEN}✓ All services started${NC}"

# Cleanup
echo ""
read -p "Delete temporary files? [y/N]: " cleanup
if [[ $cleanup =~ ^[Yy]$ ]]; then
    rm -rf "$BACKUP_DIR"
    echo -e "${GREEN}✓ Temporary files cleaned${NC}"
fi

echo ""
echo "======================================"
echo -e "${GREEN}Restore complete!${NC}"
echo "======================================"
echo ""
echo "Service status:"
docker-compose ps
echo ""

# Display access URL
if [ -f "ssl/fullchain.pem" ]; then
    DOMAIN=$(grep -m1 "server_name" nginx/conf.d/default.conf 2>/dev/null | awk '{print $2}' | sed 's/;//' || echo "")
    if [ ! -z "$DOMAIN" ] && [ "$DOMAIN" != "_" ]; then
        echo "Access URL: https://${DOMAIN}"
        echo "Note: Please make sure domain DNS is resolved to this server"
    else
        echo "Access URL: http://localhost"
    fi
else
    echo "Access URL: http://localhost"
fi
echo ""
