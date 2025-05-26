#!/bin/bash

# Script to run Go examples individually
# Usage: ./run_examples.sh

echo "=== Go Variables Examples ==="
echo "Select an example to run:"
echo "1. Basic Variables"
echo "2. Declaration Styles"
echo "3. Zero Values"
echo "4. Variable Scope"
echo "5. Coffee Order System"
echo "6. Naming Conventions"
echo "7. Interactive Exercise"
echo "8. Common Mistakes"
echo "9. Inventory System"
echo "0. Exit"

read -p "Enter your choice (0-9): " choice

case $choice in
    1) go run 01_basic_variables.go ;;
    2) go run 02_declaration_styles.go ;;
    3) go run 03_zero_values.go ;;
    4) go run 04_scope.go ;;
    5) go run 05_coffee_order.go ;;
    6) go run 06_naming_conventions.go ;;
    7) go run 07_exercise.go ;;
    8) go run 08_common_mistakes.go ;;
    9) go run 09_inventory_system.go ;;
    0) echo "Goodbye!" ;;
    *) echo "Invalid choice. Please run the script again." ;;
esac