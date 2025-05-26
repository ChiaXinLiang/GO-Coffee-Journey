package main

import (
    "fmt"
    "strings"
    "time"
)

type Customer struct {
    Name         string
    IsMember     bool
    MemberLevel  string // "Bronze", "Silver", "Gold"
    OrderHistory int    // Number of previous orders
    TotalSpent   float64
}

type OrderItem struct {
    Name     string
    Category string
    Price    float64
    Quantity int
    IsHot    bool
}

func main() {
    fmt.Println("=== GoCoffee Complex Decision Making ===\n")
    
    // Create a customer
    customer := Customer{
        Name:         "Marcus",
        IsMember:     true,
        MemberLevel:  "Silver",
        OrderHistory: 45,
        TotalSpent:   567.89,
    }
    
    // Create an order
    items := []OrderItem{
        {"Cappuccino", "Coffee", 4.50, 2, true},
        {"Croissant", "Food", 3.50, 1, false},
        {"Iced Latte", "Coffee", 5.00, 1, false},
    }
    
    // Process the order
    processOrder(customer, items)
    
    // Loyalty program evaluation
    evaluateLoyaltyStatus(customer)
    
    // Special offer eligibility
    checkSpecialOffers(customer, items)
}

func processOrder(customer Customer, items []OrderItem) {
    fmt.Printf("Processing order for: %s\n", customer.Name)
    fmt.Println(strings.Repeat("-", 40))
    
    subtotal := 0.0
    coffeeCount := 0
    
    // Calculate subtotal and count coffee items
    for _, item := range items {
        itemTotal := item.Price * float64(item.Quantity)
        subtotal += itemTotal
        
        fmt.Printf("%d × %s @ $%.2f = $%.2f\n", 
            item.Quantity, item.Name, item.Price, itemTotal)
        
        if item.Category == "Coffee" {
            coffeeCount += item.Quantity
        }
    }
    
    fmt.Printf("\nSubtotal: $%.2f\n", subtotal)
    
    // Apply discounts based on multiple conditions
    discount := 0.0
    discountReasons := []string{}
    
    // Member discount
    if customer.IsMember {
        memberDiscount := 0.0
        
        if customer.MemberLevel == "Bronze" {
            memberDiscount = 0.05
        } else if customer.MemberLevel == "Silver" {
            memberDiscount = 0.10
        } else if customer.MemberLevel == "Gold" {
            memberDiscount = 0.15
        }
        
        if memberDiscount > 0 {
            amount := subtotal * memberDiscount
            discount += amount
            discountReasons = append(discountReasons, 
                fmt.Sprintf("%s member %.0f%%: -$%.2f", 
                    customer.MemberLevel, memberDiscount*100, amount))
        }
    }
    
    // Volume discount
    if coffeeCount >= 5 {
        volumeDiscount := subtotal * 0.05
        discount += volumeDiscount
        discountReasons = append(discountReasons, 
            fmt.Sprintf("5+ coffees: -$%.2f", volumeDiscount))
    }
    
    // Time-based discount
    hour := time.Now().Hour()
    if hour >= 14 && hour <= 16 {
        happyHourDiscount := subtotal * 0.20
        discount += happyHourDiscount
        discountReasons = append(discountReasons, 
            fmt.Sprintf("Happy hour 20%%: -$%.2f", happyHourDiscount))
    }
    
    // Apply discounts
    if discount > 0 {
        fmt.Println("\nDiscounts applied:")
        for _, reason := range discountReasons {
            fmt.Printf("  • %s\n", reason)
        }
        fmt.Printf("Total discount: -$%.2f\n", discount)
    }
    
    // Calculate final total
    finalTotal := subtotal - discount
    tax := finalTotal * 0.085
    totalWithTax := finalTotal + tax
    
    fmt.Printf("\nAfter discount: $%.2f\n", finalTotal)
    fmt.Printf("Tax (8.5%%): $%.2f\n", tax)
    fmt.Printf("Total: $%.2f\n", totalWithTax)
    
    // Payment recommendation
    if totalWithTax > 50 {
        fmt.Println("\n💳 Recommend card payment for large amount")
    } else if totalWithTax < 10 {
        fmt.Println("\n💵 Cash payment OK for small amount")
    }
}

func evaluateLoyaltyStatus(customer Customer) {
    fmt.Println("\n=== Loyalty Status Evaluation ===")
    
    // Determine if eligible for upgrade
    shouldUpgrade := false
    nextLevel := customer.MemberLevel
    reason := ""
    
    if !customer.IsMember {
        if customer.OrderHistory >= 5 || customer.TotalSpent >= 50 {
            shouldUpgrade = true
            nextLevel = "Bronze"
            reason = "Eligible for membership"
        }
    } else if customer.MemberLevel == "Bronze" {
        if customer.OrderHistory >= 25 && customer.TotalSpent >= 250 {
            shouldUpgrade = true
            nextLevel = "Silver"
            reason = "Met Silver requirements"
        }
    } else if customer.MemberLevel == "Silver" {
        if customer.OrderHistory >= 50 && customer.TotalSpent >= 500 {
            shouldUpgrade = true
            nextLevel = "Gold"
            reason = "Met Gold requirements"
        }
    }
    
    // Display status
    fmt.Printf("Customer: %s\n", customer.Name)
    
    if customer.IsMember {
        fmt.Printf("Current level: %s\n", customer.MemberLevel)
    } else {
        fmt.Println("Current level: Non-member")
    }
    
    fmt.Printf("Orders: %d, Total spent: $%.2f\n", 
        customer.OrderHistory, customer.TotalSpent)
    
    if shouldUpgrade {
        fmt.Printf("🎉 %s! Upgrade to %s\n", reason, nextLevel)
    } else {
        // Show progress
        if customer.MemberLevel == "Bronze" {
            ordersNeeded := 25 - customer.OrderHistory
            spentNeeded := 250.0 - customer.TotalSpent
            
            if ordersNeeded > 0 || spentNeeded > 0 {
                fmt.Println("\nProgress to Silver:")
                if ordersNeeded > 0 {
                    fmt.Printf("  Need %d more orders\n", ordersNeeded)
                }
                if spentNeeded > 0 {
                    fmt.Printf("  Need $%.2f more in purchases\n", spentNeeded)
                }
            }
        }
    }
}

func checkSpecialOffers(customer Customer, items []OrderItem) {
    fmt.Println("\n=== Special Offers ===")
    
    hasHotDrink := false
    hasColdDrink := false
    hasFood := false
    
    // Check what categories are in the order
    for _, item := range items {
        if item.Category == "Coffee" {
            if item.IsHot {
                hasHotDrink = true
            } else {
                hasColdDrink = true
            }
        } else if item.Category == "Food" {
            hasFood = true
        }
    }
    
    offersAvailable := false
    
    // Combo offer
    if hasHotDrink && hasFood {
        fmt.Println("☕🥐 Breakfast Combo: Add any pastry for 50% off!")
        offersAvailable = true
    }
    
    // Temperature mix offer
    if hasHotDrink && hasColdDrink {
        fmt.Println("🔥❄️ Temperature Mix: Get 3rd drink free!")
        offersAvailable = true
    }
    
    // Member-only offers
    if customer.IsMember {
        dayOfWeek := time.Now().Weekday()
        
        if dayOfWeek == time.Monday {
            fmt.Println("🎉 Member Monday: Double points on all purchases!")
            offersAvailable = true
        }
        
        if customer.MemberLevel == "Gold" && dayOfWeek == time.Friday {
            fmt.Println("🏆 Gold Friday: Free size upgrade on any drink!")
            offersAvailable = true
        }
    }
    
    // Birthday offer (simulated)
    if customer.Name == "Marcus" { // Pretend it's their birthday
        fmt.Println("🎂 Birthday Special: Free dessert with any purchase!")
        offersAvailable = true
    }
    
    if !offersAvailable {
        fmt.Println("No special offers available today")
    }
}