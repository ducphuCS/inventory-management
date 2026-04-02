#!/bin/bash

# Configuration for local testing
GIN_BASE_URL="http://localhost:8080/v1"
WEBAPP_BASE_URL="http://localhost:8000"

echo "=== GIN INVENTORY SERVICE (PORT 8080) ==="
echo "--- 1. Checking Gin health ---"
curl -s http://localhost:8080/health | grep "ok" && echo "Gin is up!" || { echo "Gin might not be running on Port 8080."; exit 1; }

echo -e "\n--- 2. Adding a new item (Phone) ---"
curl -s -X POST $GIN_BASE_URL/items \
  -H "Content-Type: application/json" \
  -d '{
    "product_name": "Smart Phone",
    "stock_count": 50,
    "item_details": "Latest gen smartphone"
  }'

echo -e "\n\n--- 3. Listing all items in Gin ---"
curl -s $GIN_BASE_URL/items

echo -e "\n\n--- 4. Getting Python-based analysis for an item (ID: 1) ---"
curl -s $GIN_BASE_URL/items/1/analysis

echo -e "\n\n=== WEBAPP FASTAPI BACKEND (PORT 8000) ==="
echo "--- 1. Checking WebApp health ---"
curl -s http://localhost:8000/ | grep "bridged" && echo "WebApp is up!" || { echo "WebApp might not be running on Port 8000."; exit 1; }

echo -e "\n--- 2. Fetching items from WebApp (Bridged) ---"
curl -s $WEBAPP_BASE_URL/items/

echo -e "\n\n--- 3. Requesting analysis from WebApp (Proxy) ---"
curl -s $WEBAPP_BASE_URL/items/1/analyze

echo -e "\n"
