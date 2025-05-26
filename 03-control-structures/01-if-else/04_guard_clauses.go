package main

import (
    "fmt"
    "strings"
    "time"
)

type MenuItem struct {
    Name      string
    Category  string
    Price     float64
    Available bool
    PrepTime  int // minutes
}

type OrderRequest struct {
    CustomerName string
    Items        []string
    Quantity     map[string]int
    IsDelivery   bool
    Distance     float64 // km
    PaymentType  string
}

var menu = map[string]MenuItem{
    "espresso":    {"Espresso", "coffee", 3.00, true, 2},
    "latte":       {"Latte", "coffee", 4.50, true, 3},
    "cappuccino":  {"Cappuccino", "coffee", 4.00, true, 3},
    "americano":   {"Americano", "coffee", 3.50, true, 2},
    "mocha":       {"Mocha", "coffee", 5.00, false, 4},
    "croissant":   {"Croissant", "food", 3.50, true, 1},
    "sandwich":    {"Sandwich", "food", 6.50, true, 5},
    "muffin":      {"Muffin", "food", 2.50, true, 1},
}

func main() {
    fmt.Println("=== GoCoffee Guard Clauses Pattern ===\n")
    
    // Test various order scenarios
    testOrders := []OrderRequest{
        {
            CustomerName: "Alice",
            Items:        []string{"latte", "croissant"},
            Quantity:     map[string]int{"latte": 1, "croissant": 1},
            IsDelivery:   false,
            PaymentType:  "card",
        },
        {
            CustomerName: "Bob",
            Items:        []string{"mocha", "sandwich"},
            Quantity:     map[string]int{"mocha": 2, "sandwich": 1},
            IsDelivery:   true,
            Distance:     3.5,
            PaymentType:  "cash",
        },
        {
            CustomerName: "",
            Items:        []string{"espresso"},
            Quantity:     map[string]int{"espresso": 1},
            IsDelivery:   false,
            PaymentType:  "card",
        },
        {
            CustomerName: "Charlie",
            Items:        []string{"latte", "unknown_item"},
            Quantity:     map[string]int{"latte": 1, "unknown_item": 1},
            IsDelivery:   true,
            Distance:     15.0,
            PaymentType:  "bitcoin",
        },
    }
    
    for i, order := range testOrders {
        fmt.Printf("=== Order %d ===\n", i+1)
        result, err := validateOrder(order)
        
        if err != nil {
            fmt.Printf("❌ Order rejected: %s\n\n", err)
            continue
        }
        
        fmt.Printf("✅ Order accepted!\n")
        fmt.Printf("Total: $%.2f\n", result.total)
        fmt.Printf("Prep time: %d minutes\n\n", result.prepTime)
    }
}

type OrderResult struct {
    total    float64
    prepTime int
}

// validateOrder uses guard clauses for early returns
func validateOrder(order OrderRequest) (OrderResult, error) {
    // Guard: Check customer name
    if strings.TrimSpace(order.CustomerName) == "" {
        return OrderResult{}, fmt.Errorf("customer name is required")
    }
    
    // Guard: Check if we're open
    hour := time.Now().Hour()
    if hour < 6 || hour >= 22 {
        return OrderResult{}, fmt.Errorf("sorry, we're closed (hours: 6 AM - 10 PM)")
    }
    
    // Guard: Check items exist
    if len(order.Items) == 0 {
        return OrderResult{}, fmt.Errorf("no items in order")
    }
    
    // Guard: Validate payment type
    validPayments := map[string]bool{
        "cash": true,
        "card": true,
        "app":  true,
    }
    
    if !validPayments[order.PaymentType] {
        return OrderResult{}, fmt.Errorf("invalid payment type: %s", order.PaymentType)
    }
    
    // Guard: Check delivery constraints
    if order.IsDelivery {
        if order.Distance > 10 {
            return OrderResult{}, fmt.Errorf("delivery distance too far: %.1f km (max: 10 km)", 
                order.Distance)
        }
        
        if order.PaymentType == "cash" {
            return OrderResult{}, fmt.Errorf("cash payment not accepted for delivery")
        }
    }
    
    // Process order items
    total := 0.0
    maxPrepTime := 0
    
    for _, itemName := range order.Items {
        item, exists := menu[strings.ToLower(itemName)]
        
        // Guard: Check if item exists
        if !exists {
            return OrderResult{}, fmt.Errorf("item not found: %s", itemName)
        }
        
        // Guard: Check if item is available
        if !item.Available {
            return OrderResult{}, fmt.Errorf("item not available: %s", item.Name)
        }
        
        // Get quantity
        qty := order.Quantity[itemName]
        if qty <= 0 {
            qty = 1
        }
        
        // Guard: Check quantity limits
        if qty > 10 {
            return OrderResult{}, fmt.Errorf("quantity too high for %s: %d (max: 10)", 
                item.Name, qty)
        }
        
        // Calculate totals
        total += item.Price * float64(qty)
        
        if item.PrepTime > maxPrepTime {
            maxPrepTime = item.PrepTime
        }
    }
    
    // Guard: Minimum order for delivery
    if order.IsDelivery && total < 15.00 {
        return OrderResult{}, fmt.Errorf("minimum order for delivery is $15.00 (current: $%.2f)", 
            total)
    }
    
    // Add delivery fee
    if order.IsDelivery {
        deliveryFee := 3.00
        if order.Distance > 5 {
            deliveryFee = 5.00
        }
        total += deliveryFee
    }
    
    // All validations passed!
    return OrderResult{
        total:    total,
        prepTime: maxPrepTime,
    }, nil
}