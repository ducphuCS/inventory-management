#!/bin/bash

# Test script for the renamed parameters in the inventory-service
BASE_URL="http://localhost:8080/v1"

echo "--- 1. Checking server health ---"
curl -s http://localhost:8080/health | grep "ok" && echo "Server is up!" || { echo "Server might not be running. Please start it first with 'go run main.go'."; exit 1; }

echo -e "\n--- 2. Adding a new item (Laptop) with renamed parameters ---"
curl -s -X POST $BASE_URL/items \
  -H "Content-Type: application/json" \
  -d '{
    "product_name": "Pro Laptop",
    "stock_count": 15,
    "item_details": "Next-gen workplace laptop"
  }'

echo -e "\n--- 3. Adding another item (Mouse) with renamed parameters ---"
curl -s -X POST $BASE_URL/items \
  -H "Content-Type: application/json" \
  -d '{
    "product_name": "Master Mouse",
    "stock_count": 25,
    "item_details": "Ergonomic bluetooth mouse"
  }'

echo -e "\n--- 4. Getting total count of items (new response key: total_items) ---"
curl -s $BASE_URL/items/count

echo -e "\n\n--- 5. Getting Python-based analysis for an item (ID: 1) ---"
curl -s $BASE_URL/items/1/analysis
echo -e "\n"
