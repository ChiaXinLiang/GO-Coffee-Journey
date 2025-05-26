package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    fmt.Println("=== GoCoffee Input Formatting ===\n")
    
    reader := bufio.NewReader(os.Stdin)
    
    // Get customer name
    fmt.Print("Enter customer name: ")
    name, _ := reader.ReadString('\n')
    name = strings.TrimSpace(name)
    
    // Get coffee choice with validation
    fmt.Println("\nCoffee Menu:")
    fmt.Println("1. Espresso   ($3.00)")
    fmt.Println("2. Latte      ($4.50)")
    fmt.Println("3. Cappuccino ($4.00)")
    fmt.Println("4. Americano  ($3.50)")
    
    var choice int
    var coffee string
    var price float64
    
    for {
        fmt.Print("\nEnter choice (1-4): ")
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)
        
        var err error
        choice, err = strconv.Atoi(input)
        
        if err != nil || choice < 1 || choice > 4 {
            fmt.Printf("❌ Invalid choice: %s. Please enter 1-4.\n", input)
            continue
        }
        
        break
    }
    
    // Set coffee and price based on choice
    switch choice {
    case 1:
        coffee, price = "Espresso", 3.00
    case 2:
        coffee, price = "Latte", 4.50
    case 3:
        coffee, price = "Cappuccino", 4.00
    case 4:
        coffee, price = "Americano", 3.50
    }
    
    // Get quantity
    var quantity int
    for {
        fmt.Print("Enter quantity: ")
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)
        
        var err error
        quantity, err = strconv.Atoi(input)
        
        if err != nil || quantity < 1 {
            fmt.Printf("❌ Invalid quantity: %s. Please enter a positive number.\n", input)
            continue
        }
        
        break
    }
    
    // Confirm order
    total := price * float64(quantity)
    
    fmt.Println("\n" + strings.Repeat("=", 40))
    fmt.Println("ORDER CONFIRMATION")
    fmt.Println(strings.Repeat("=", 40))
    fmt.Printf("Customer: %s\n", name)
    fmt.Printf("Item:     %d × %s @ $%.2f\n", quantity, coffee, price)
    fmt.Printf("Total:    $%.2f\n", total)
    fmt.Println(strings.Repeat("=", 40))
    
    // Get confirmation
    fmt.Print("\nConfirm order? (y/n): ")
    confirm, _ := reader.ReadString('\n')
    confirm = strings.ToLower(strings.TrimSpace(confirm))
    
    if confirm == "y" || confirm == "yes" {
        fmt.Println("\n✅ Order placed successfully!")
        fmt.Printf("Thank you, %s! Your %s will be ready soon.\n", name, coffee)
    } else {
        fmt.Println("\n❌ Order cancelled.")
    }
    
    // Formatted input example
    fmt.Println("\n" + strings.Repeat("-", 40))
    fmt.Println("FORMATTED INPUT EXAMPLES:")
    
    // Phone number formatting
    fmt.Print("\nEnter phone (10 digits): ")
    phone, _ := reader.ReadString('\n')
    phone = strings.TrimSpace(phone)
    
    // Remove non-digits
    digitsOnly := ""
    for _, ch := range phone {
        if ch >= '0' && ch <= '9' {
            digitsOnly += string(ch)
        }
    }
    
    if len(digitsOnly) == 10 {
        formatted := fmt.Sprintf("(%s) %s-%s",
            digitsOnly[:3], digitsOnly[3:6], digitsOnly[6:])
        fmt.Printf("Formatted: %s\n", formatted)
    } else {
        fmt.Println("❌ Invalid phone number")
    }
    
    // Credit card formatting (demo only - last 4 digits)
    fmt.Print("\nEnter last 4 digits of card: ")
    card, _ := reader.ReadString('\n')
    card = strings.TrimSpace(card)
    
    if len(card) == 4 {
        fmt.Printf("Card: •••• •••• •••• %s\n", card)
    }
}