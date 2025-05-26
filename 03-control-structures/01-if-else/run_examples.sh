#!/bin/bash

# GoCoffee If-Else Examples Runner Script

clear
echo "🚀 GoCoffee If-Else Examples"
echo "============================"
echo ""

# Function to run an example with a header
run_example() {
    local file=$1
    local title=$2
    
    echo "📘 $title"
    echo "-------------------------------------------"
    
    if [ -f "$file" ]; then
        go run "$file"
        echo ""
        echo "Press Enter to continue..."
        read -r
        clear
    else
        echo "❌ File not found: $file"
        echo ""
    fi
}

# Main menu
while true; do
    echo "🚀 GoCoffee If-Else Examples"
    echo "============================"
    echo ""
    echo "Select an example to run:"
    echo ""
    echo "1. Basic If Statements"
    echo "2. Complex Decision Making"
    echo "3. Order Flow Control"
    echo "4. Guard Clauses Pattern"
    echo "5. State Machine Example"
    echo "6. Business Rules Engine"
    echo "7. Menu Navigation System"
    echo "8. Build Tags (Debug Mode)"
    echo "9. If-Else Challenge"
    echo ""
    echo "0. Exit"
    echo ""
    echo -n "Enter your choice (0-9): "
    read -r choice
    
    clear
    
    case $choice in
        1)
            run_example "01_basic_if.go" "Basic If Statements"
            ;;
        2)
            run_example "02_complex_decisions.go" "Complex Decision Making"
            ;;
        3)
            run_example "03_order_flow.go" "Order Flow Control"
            ;;
        4)
            run_example "04_guard_clauses.go" "Guard Clauses Pattern"
            ;;
        5)
            run_example "05_state_machine.go" "State Machine Example"
            ;;
        6)
            run_example "06_business_rules.go" "Business Rules Engine"
            ;;
        7)
            run_example "07_menu_navigation.go" "Menu Navigation System"
            ;;
        8)
            run_example "08_build_tags.go" "Build Tags (Debug Mode)"
            ;;
        9)
            run_example "09_if_else_challenge.go" "If-Else Challenge"
            ;;
        0)
            echo "👋 Thank you for exploring GoCoffee If-Else examples!"
            exit 0
            ;;
        *)
            echo "❌ Invalid choice. Please select a number between 0 and 9."
            echo ""
            echo "Press Enter to continue..."
            read -r
            clear
            ;;
    esac
done