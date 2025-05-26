#!/bin/bash

echo "Running Go Break and Continue Examples"
echo "====================================="
echo

# Loop through all .go files
for file in *.go; do
    if [ -f "$file" ]; then
        echo "Running $file:"
        echo "-------------------"
        go run "$file"
        echo
        echo "Press Enter to continue to next example..."
        read
        clear
    fi
done

echo "All examples completed!"