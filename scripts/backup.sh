#!/bin/bash

# Data backup script
set -e

GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m'

echo "======================================="
echo "   Data Backup Tool"
echo "======================================="
echo ""

# Check if service is running
if ! docker ps | grep -q blog-postgres; then
    echo -e "${RED}Error: Database service not running${NC}"
    exit 1
fi

BACKUP_DIR="backup_$(date +%Y%m%d_%H%M%S)"
mkdir -p "$BACKUP_DIR"

echo -e "${GREEN}Starting backup...${NC}"
echo ""

# 1. Backup database
echo "1/5 Backing up database..."
docker exec blog-postgres pg_dump -U postgres blog > "$BACKUP_DIR/database.sql"
echo -e "${GREEN}✓ Database backup complete${NC}"

# 2. Backup static files
echo "2/5 Backing up static files..."
if [ -d "static" ]; then
    tar -czf "$BACKUP_DIR/static.tar.gz" static/
    echo -e "${GREEN}✓ Static files backup complete${NC}"
else
    echo "Skipping static files backup (directory not found)"
fi

# 3. Backup config files
echo "3/5 Backing up config files..."
cp .env "$BACKUP_DIR/.env" 2>/dev/null || echo "Skipping .env backup"
echo -e "${GREEN}✓ Config files backup complete${NC}"

# 4. Backup SSL certificates (if exists)
echo "4/5 Backing up SSL certificates..."
if [ -f "ssl/fullchain.pem" ]; then
    cp -r ssl "$BACKUP_DIR/"
    cp -r nginx/conf.d "$BACKUP_DIR/"
    echo -e "${GREEN}✓ SSL certificates backup complete${NC}"
else
    echo "Skipping SSL backup (HTTPS not configured)"
fi

# 5. Backup GeoIP database (if exists)
echo "5/5 Backing up GeoIP database..."
if [ -f "config/GeoLite2-City.mmdb" ]; then
    mkdir -p "$BACKUP_DIR/config"
    cp config/GeoLite2-City.mmdb "$BACKUP_DIR/config/"
    echo -e "${GREEN}✓ GeoIP database backup complete${NC}"
else
    echo "Skipping GeoIP backup"
fi

# Package
echo ""
echo "Packaging..."
tar -czf "${BACKUP_DIR}.tar.gz" "$BACKUP_DIR"
rm -rf "$BACKUP_DIR"

echo ""
echo "======================================"
echo -e "${GREEN}Backup complete!${NC}"
echo "======================================"
echo ""
echo "Backup file: ${BACKUP_DIR}.tar.gz"
echo "Size: $(du -h ${BACKUP_DIR}.tar.gz | cut -f1)"
echo ""
echo "To migrate to new server:"
echo "  1. Upload backup file to new server"
echo "  2. git clone <repo> blog && cd blog"
echo "  3. Run ./scripts/restore.sh"
echo ""
