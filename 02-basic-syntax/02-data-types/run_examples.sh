#!/bin/bash

# Script to run Go data types examples
# Usage: ./run_examples.sh

echo "=== Go Data Types Examples ==="
echo "Select an example to run:"
echo "1. Number Types"
echo "2. Strings and Runes"
echo "3. Boolean Logic"
echo "4. Type Conversions"
echo "5. Zero Values"
echo "6. Money Handling"
echo "7. Type Safety"
echo "8. Real World Types"
echo "9. Type Challenge"
echo "0. Exit"

read -p "Enter your choice (0-9): " choice

case $choice in
    1) go run 01_number_types.go ;;
    2) go run 02_strings_and_runes.go ;;
    3) go run 03_boolean_logic.go ;;
    4) go run 04_type_conversions.go ;;
    5) go run 05_zero_values.go ;;
    6) go run 06_money_handling.go ;;
    7) go run 07_type_safety.go ;;
    8) go run 08_real_world_types.go ;;
    9) go run 09_type_challenge.go ;;
    0) echo "Goodbye!" ;;
    *) echo "Invalid choice. Please run the script again." ;;
esac