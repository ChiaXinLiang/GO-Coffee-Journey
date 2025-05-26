#!/bin/bash

# GoCoffee fmt Package Examples Runner
# This script runs all the fmt package examples in order

echo "=== Running GoCoffee fmt Package Examples ==="
echo

# Array of example files
examples=(
    "01_basic_printing.go"
    "02_format_verbs.go"
    "03_receipt_formatting.go"
    "04_menu_display.go"
    "05_table_formatting.go"
    "06_input_formatting.go"
    "07_color_output.go"
    "08_real_world_examples.go"
    "09_formatting_challenge.go"
)

# Run each example
for example in "${examples[@]}"; do
    echo "================================================"
    echo "Running $example"
    echo "================================================"
    
    go run "$example"
    
    echo
    echo "Press Enter to continue to the next example..."
    read
    clear
done

echo "All examples completed!"
echo "Great job learning about the fmt package!"