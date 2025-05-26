#!/bin/bash

# Script to run all variadic function examples in order

echo "=== Running Variadic Functions Examples ==="
echo

# Run each example
for i in {1..9}; do
    file=$(printf "%02d_*.go" $i)
    if [ -f $file ]; then
        echo "----------------------------------------"
        echo "Running: $file"
        echo "----------------------------------------"
        go run $file
        echo
        echo "Press Enter to continue..."
        read
    fi
done

echo "=== All Variadic Functions Examples Completed ==="