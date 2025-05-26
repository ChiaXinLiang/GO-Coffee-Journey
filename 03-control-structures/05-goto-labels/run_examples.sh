#!/bin/bash

echo "Running Go Goto and Labels Examples"
echo "==================================="
echo
echo "⚠️  WARNING: These examples demonstrate goto usage."
echo "In real code, avoid goto whenever possible!"
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
echo
echo "Remember: Understanding goto helps you appreciate"
echo "why structured programming is better!"