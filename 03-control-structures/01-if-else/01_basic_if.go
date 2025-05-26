package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println("=== GoCoffee Basic If Statements ===\n")
    
    // Simple if statement
    customerAge := 25
    if customerAge >= 18 {
        fmt.Println("✓ Customer can order coffee")
    }
    
    // If with else
    itemsInStock := 5
    orderQuantity := 3
    
    if itemsInStock >= orderQuantity {
        fmt.Printf("✓ Order confirmed: %d items\n", orderQuantity)
        itemsInStock -= orderQuantity
        fmt.Printf("  Remaining stock: %d\n", itemsInStock)
    } else {
        fmt.Printf("✗ Insufficient stock. Available: %d, Requested: %d\n", 
            itemsInStock, orderQuantity)
    }
    
    // If-else if-else chain
    currentHour := time.Now().Hour()
    
    fmt.Printf("\nCurrent time: %d:00\n", currentHour)
    
    if currentHour < 6 {
        fmt.Println("🌙 Sorry, we're closed (too early)")
    } else if currentHour < 12 {
        fmt.Println("☀️ Good morning! Breakfast menu available")
    } else if currentHour < 17 {
        fmt.Println("🌤️ Good afternoon! Full menu available")
    } else if currentHour < 22 {
        fmt.Println("🌆 Good evening! Dinner menu available")
    } else {
        fmt.Println("🌙 Sorry, we're closed (too late)")
    }
    
    // Multiple conditions
    isMember := true
    orderTotal := 25.50
    
    if isMember && orderTotal > 20 {
        discount := orderTotal * 0.10
        fmt.Printf("\n💳 Member discount applied: $%.2f\n", discount)
        fmt.Printf("Final total: $%.2f\n", orderTotal-discount)
    }
    
    // Nested if statements
    dayOfWeek := time.Now().Weekday()
    
    if dayOfWeek == time.Saturday || dayOfWeek == time.Sunday {
        fmt.Println("\n🎉 It's the weekend!")
        if currentHour >= 10 && currentHour <= 14 {
            fmt.Println("🍳 Brunch special available!")
        }
    } else {
        fmt.Println("\n💼 Weekday")
        if currentHour >= 7 && currentHour <= 9 {
            fmt.Println("☕ Morning rush hour - extra staff needed")
        }
    }
    
    // If with initialization
    if coffeeTemp := 165; coffeeTemp > 160 {
        fmt.Printf("\n🔥 Coffee temperature: %d°F (Perfect!)\n", coffeeTemp)
    } else {
        fmt.Printf("\n❄️ Coffee temperature: %d°F (Too cold!)\n", coffeeTemp)
    }
    // coffeeTemp is not accessible here - scoped to if statement
    
    // Practical example: Order validation
    fmt.Println("\n=== Order Validation Example ===")
    
    order := struct {
        item     string
        size     string
        quantity int
        payment  float64
    }{
        item:     "Latte",
        size:     "Large",
        quantity: 2,
        payment:  15.00,
    }
    
    // Calculate price
    basePrice := 4.50
    sizeMultiplier := 1.0
    
    if order.size == "Small" {
        sizeMultiplier = 0.8
    } else if order.size == "Medium" {
        sizeMultiplier = 1.0
    } else if order.size == "Large" {
        sizeMultiplier = 1.3
    }
    
    totalCost := basePrice * sizeMultiplier * float64(order.quantity)
    
    fmt.Printf("Order: %d %s %s(s)\n", order.quantity, order.size, order.item)
    fmt.Printf("Total cost: $%.2f\n", totalCost)
    fmt.Printf("Payment: $%.2f\n", order.payment)
    
    if order.payment >= totalCost {
        change := order.payment - totalCost
        fmt.Printf("✓ Payment accepted. Change: $%.2f\n", change)
    } else {
        shortfall := totalCost - order.payment
        fmt.Printf("✗ Insufficient payment. Need $%.2f more\n", shortfall)
    }
}