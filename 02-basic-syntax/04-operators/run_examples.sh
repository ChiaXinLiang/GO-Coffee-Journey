#!/bin/bash

# Run all Go examples for the operators chapter

echo "🔧 Go Operators Examples"
echo "========================"
echo

examples=(
    "01_arithmetic_operators.go"
    "02_comparison_operators.go"
    "03_logical_operators.go"
    "04_assignment_operators.go"
    "05_bitwise_operators.go"
    "06_operator_precedence.go"
    "07_practical_examples.go"
    "08_operator_pitfalls.go"
    "09_operators_challenge.go"
)

for example in "${examples[@]}"; do
    if [ -f "$example" ]; then
        echo "▶️  Running $example"
        echo "----------------------------------------"
        go run "$example"
        echo
        echo "Press Enter to continue..."
        read
        clear
    else
        echo "⚠️  File not found: $example"
    fi
done

echo "✅ All examples completed!"