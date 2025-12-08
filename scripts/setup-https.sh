#!/bin/bash

# HTTPS certificate setup script (using acme.sh)
set -e

GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m'

echo "======================================="
echo "   HTTPS Certificate Setup Tool"
echo "   Using acme.sh (Let's Encrypt)"
echo "======================================="
echo ""

# Check if service is running
if ! docker ps | grep -q blog-nginx; then
    echo -e "${RED}Error: Service not running. Please run: docker-compose up -d${NC}"
    exit 1
fi

# Input domain
echo -e "${YELLOW}Please enter your domain(s) (comma-separated, e.g., example.com,www.example.com):${NC}"
read -p "Domain(s): " DOMAIN_INPUT

# Normalize domain list
IFS=',' read -ra DOMAINS <<< "$DOMAIN_INPUT"
PRIMARY_DOMAIN=${DOMAINS[0]}
DOMAIN_ARGS=()
for d in "${DOMAINS[@]}"; do
    d_trimmed=$(echo "$d" | xargs)
    if [ -n "$d_trimmed" ]; then
        DOMAIN_ARGS+=("-d" "$d_trimmed")
    fi
done

if [ -z "$PRIMARY_DOMAIN" ]; then
    echo -e "${RED}Error: Domain cannot be empty${NC}"
    exit 1
fi

echo ""
echo -e "${YELLOW}Please make sure your domain is resolved to this server IP!${NC}"
echo -e "${YELLOW}DNS propagation may take a few minutes${NC}"
echo ""
read -p "Confirm domain is resolved? [y/N]: " confirm

if [[ ! $confirm =~ ^[Yy]$ ]]; then
    echo "Cancelled"
    exit 0
fi

echo ""
echo "======================================"
echo "Starting HTTPS setup..."
echo "======================================"

# Create necessary directories
mkdir -p nginx/certbot ssl

# 1. Update domain in Nginx config
echo "1/5 Updating Nginx configuration..."
SERVER_NAMES=$(printf " %s" "${DOMAINS[@]}")
SERVER_NAMES=${SERVER_NAMES:1}
sed -i.bak "s/server_name _;/server_name ${SERVER_NAMES};/g" nginx/conf.d/default.conf
rm -f nginx/conf.d/default.conf.bak
echo -e "${GREEN}✓ Nginx config updated${NC}"

# 2. Reload Nginx
echo "2/5 Reloading Nginx..."
docker exec blog-nginx nginx -s reload
echo -e "${GREEN}✓ Nginx reloaded${NC}"

# 3. Install acme.sh
echo "3/5 Installing acme.sh..."
if [ ! -d "$HOME/.acme.sh" ]; then
    curl https://get.acme.sh | sh -s email=admin@${PRIMARY_DOMAIN}
    echo -e "${GREEN}✓ acme.sh installed${NC}"
else
    echo -e "${GREEN}✓ acme.sh already installed${NC}"
fi

# 4. Issue certificate (using webroot mode)
echo "4/5 Issuing certificate (may take a few minutes)..."
if [ ${#DOMAIN_ARGS[@]} -eq 0 ]; then
    echo -e "${RED}Error: No valid domains provided${NC}"
    exit 1
fi

$HOME/.acme.sh/acme.sh --issue \
    "${DOMAIN_ARGS[@]}" \
    --webroot "$(pwd)/nginx/certbot" \
    --server letsencrypt \
    --keylength 2048 \
    --force

if [ $? -ne 0 ]; then
    echo -e "${RED}✗ Certificate issuance failed${NC}"
    echo ""
    echo "Possible reasons:"
    echo "1. Domain DNS not correctly resolved to this server"
    echo "2. Port 80 not open or occupied"
    echo "3. Firewall blocking connection"
    echo ""
    echo "You can re-run this script later to retry"
    exit 1
fi

echo -e "${GREEN}✓ Certificate issued successfully${NC}"

# 5. Install certificate
echo "5/5 Installing certificate..."
$HOME/.acme.sh/acme.sh --install-cert -d ${PRIMARY_DOMAIN} \
    --key-file       $(pwd)/ssl/key.pem  \
    --fullchain-file $(pwd)/ssl/fullchain.pem \
    --reloadcmd      "docker exec blog-nginx nginx -s reload"

echo -e "${GREEN}✓ Certificate installed${NC}"

# 6. Enable HTTPS config
echo "6/6 Enabling HTTPS..."
# Uncomment HTTPS configuration
sed -i.bak '/^# server {/,/^# }/s/^# //' nginx/conf.d/default.conf
rm -f nginx/conf.d/default.conf.bak

# Reload Nginx
docker exec blog-nginx nginx -s reload

# Setup auto-renewal (crontab)
echo "Setting up auto-renewal..."
(crontab -l 2>/dev/null | grep -v "acme.sh --cron"; echo "0 0 * * * $HOME/.acme.sh/acme.sh --cron --home $HOME/.acme.sh > /dev/null") | crontab -
echo -e "${GREEN}✓ Auto-renewal configured${NC}"

echo ""
echo "======================================"
echo -e "${GREEN}HTTPS setup complete!${NC}"
echo "======================================"
echo ""
echo "Certificate information:"
echo "  Domains: ${SERVER_NAMES}"
echo "  Certificate: ./ssl/fullchain.pem"
echo "  Private key: ./ssl/key.pem"
echo "  Validity: 90 days (auto-renewal)"
echo ""
echo "Access URLs:"
echo "  HTTP:  http://${PRIMARY_DOMAIN} (auto redirect to HTTPS)"
echo "  HTTPS: https://${PRIMARY_DOMAIN}"
echo ""
echo "Certificate will auto-renew, no manual operation needed"
echo ""
