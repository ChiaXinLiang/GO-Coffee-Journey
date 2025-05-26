#!/bin/bash

# Script to run Go constants examples
# Usage: ./run_examples.sh

echo "=== Go Constants Examples ==="
echo "Select an example to run:"
echo "1. Basic Constants"
echo "2. Coffee Constants & Pricing"
echo "3. Iota Constants"
echo "4. Typed vs Untyped Constants"
echo "5. Business Constants"
echo "6. Constant Patterns"
echo "7. Constants Performance"
echo "8. Constants Challenge"
echo "9. Constants Summary"
echo "0. Exit"

read -p "Enter your choice (0-9): " choice

case $choice in
    1) go run 01_basic_constants.go ;;
    2) go run 02_coffee_constants.go ;;
    3) go run 03_iota_constants.go ;;
    4) go run 04_typed_constants.go ;;
    5) go run 05_business_constants.go ;;
    6) go run 06_constant_patterns.go ;;
    7) go run 07_constants_performance.go ;;
    8) go run 08_constants_challenge.go ;;
    9) go run 09_constants_summary.go ;;
    0) echo "Goodbye!" ;;
    *) echo "Invalid choice. Please run the script again." ;;
esac