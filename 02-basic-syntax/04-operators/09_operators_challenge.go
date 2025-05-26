package main

import (
    "fmt"
    "math"
)

// Challenge: Implement a complete coffee shop pricing system
// Requirements:
// 1. Calculate subtotal for multiple items
// 2. Apply time-based discounts
// 3. Apply customer status discounts
// 4. Calculate tax
// 5. Calculate tip options
// 6. Handle split bills

func main() {
    fmt.Println("=== GoCoffee Operators Challenge ===\n")
    fmt.Println("TODO: Implement a complete pricing system")
    fmt.Println("\nRequirements:")
    fmt.Println("- Support multiple items with quantities")
    fmt.Println("- Time-based pricing (happy hour, etc)")
    fmt.Println("- Member and bulk discounts")
    fmt.Println("- Tax calculation")
    fmt.Println("- Tip suggestions (15%, 18%, 20%)")
    fmt.Println("- Split bill evenly or by items")
    
    // Your implementation here...
    
    // Starter structure:
    type MenuItem struct {
        name  string
        price float64
    }
    
    type OrderItem struct {
        item     MenuItem
        quantity int
        mods     []string // size, milk type, etc
    }
    
    type Order struct {
        items      []OrderItem
        customerID string
        isMember   bool
        timestamp  string
    }
    
    // Example usage:
    _ = Order{} // Remove this when implementing
    
    fmt.Println("\nImplement your solution above!")
}

// Hint: Consider these helper functions
func calculateSubtotal(items []OrderItem) float64 {
    // TODO: Implement
    return 0
}

func applyDiscounts(subtotal float64, isMember bool) float64 {
    // TODO: Implement time and status based discounts
    return subtotal
}

func calculateTax(amount float64, rate float64) float64 {
    // TODO: Implement
    return amount * rate
}

func suggestTips(amount float64) map[string]float64 {
    // TODO: Return tip amounts for 15%, 18%, 20%
    return nil
}

func splitBillEvenly(total float64, ways int) float64 {
    // TODO: Implement fair splitting (round up cents)
    return math.Ceil(total*100/float64(ways)) / 100
}