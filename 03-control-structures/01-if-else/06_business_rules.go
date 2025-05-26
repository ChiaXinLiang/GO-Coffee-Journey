package main

import (
    "fmt"
    "time"
)

type Customer struct {
    ID           string
    Name         string
    Type         string // "regular", "member", "employee", "vip"
    JoinDate     time.Time
    TotalSpent   float64
    OrderCount   int
    LastOrderDate time.Time
}

type Promotion struct {
    Code        string
    Description string
    Discount    float64
    MinOrder    float64
    ValidUntil  time.Time
    UsageLimit  int
    UsageCount  int
}

func main() {
    fmt.Println("=== GoCoffee Business Rules Engine ===\n")
    
    // Sample customers
    customers := []Customer{
        {
            ID:           "C001",
            Name:         "Alice Johnson",
            Type:         "member",
            JoinDate:     time.Now().AddDate(0, -6, 0),
            TotalSpent:   450.00,
            OrderCount:   35,
            LastOrderDate: time.Now().AddDate(0, 0, -2),
        },
        {
            ID:           "C002",
            Name:         "Bob Smith",
            Type:         "regular",
            TotalSpent:   85.00,
            OrderCount:   8,
            LastOrderDate: time.Now().AddDate(0, 0, -30),
        },
        {
            ID:           "C003",
            Name:         "Carol Davis",
            Type:         "employee",
            JoinDate:     time.Now().AddDate(-1, 0, 0),
            TotalSpent:   1200.00,
            OrderCount:   150,
            LastOrderDate: time.Now(),
        },
        {
            ID:           "C004",
            Name:         "David Wilson",
            Type:         "vip",
            JoinDate:     time.Now().AddDate(-2, 0, 0),
            TotalSpent:   3500.00,
            OrderCount:   200,
            LastOrderDate: time.Now().AddDate(0, 0, -1),
        },
    }
    
    // Sample promotions
    promotions := []Promotion{
        {
            Code:        "WELCOME10",
            Description: "10% off for new customers",
            Discount:    0.10,
            MinOrder:    10.00,
            ValidUntil:  time.Now().AddDate(0, 1, 0),
            UsageLimit:  1,
        },
        {
            Code:        "LOYALTY20",
            Description: "20% off for loyal customers",
            Discount:    0.20,
            MinOrder:    20.00,
            ValidUntil:  time.Now().AddDate(0, 0, 7),
            UsageLimit:  5,
        },
    }
    
    // Test orders for each customer
    orderAmount := 25.00
    
    for _, customer := range customers {
        fmt.Printf("=== Processing order for %s ===\n", customer.Name)
        processBusinessRules(customer, orderAmount, promotions)
        fmt.Println()
    }
}

func processBusinessRules(customer Customer, orderAmount float64, promotions []Promotion) {
    fmt.Printf("Customer Type: %s\n", customer.Type)
    fmt.Printf("Order Amount: $%.2f\n", orderAmount)
    
    totalDiscount := 0.0
    appliedDiscounts := []string{}
    
    // Rule 1: Base discount by customer type
    baseDiscount := 0.0
    
    if customer.Type == "member" {
        baseDiscount = 0.05 // 5%
        appliedDiscounts = append(appliedDiscounts, "Member discount: 5%")
    } else if customer.Type == "employee" {
        baseDiscount = 0.25 // 25%
        appliedDiscounts = append(appliedDiscounts, "Employee discount: 25%")
    } else if customer.Type == "vip" {
        baseDiscount = 0.15 // 15%
        appliedDiscounts = append(appliedDiscounts, "VIP discount: 15%")
    }
    
    totalDiscount += baseDiscount
    
    // Rule 2: Loyalty bonus
    if customer.OrderCount >= 50 {
        loyaltyBonus := 0.05
        totalDiscount += loyaltyBonus
        appliedDiscounts = append(appliedDiscounts, "Loyalty bonus (50+ orders): 5%")
    } else if customer.OrderCount >= 25 {
        loyaltyBonus := 0.03
        totalDiscount += loyaltyBonus
        appliedDiscounts = append(appliedDiscounts, "Loyalty bonus (25+ orders): 3%")
    }
    
    // Rule 3: Reactivation discount
    daysSinceLastOrder := int(time.Since(customer.LastOrderDate).Hours() / 24)
    if daysSinceLastOrder > 30 {
        reactivationDiscount := 0.10
        totalDiscount += reactivationDiscount
        appliedDiscounts = append(appliedDiscounts, 
            fmt.Sprintf("Welcome back! (%d days): 10%%", daysSinceLastOrder))
    }
    
    // Rule 4: Birthday discount (simulated)
    today := time.Now()
    if today.Month() == time.January && customer.Name == "Alice Johnson" {
        birthdayDiscount := 0.20
        totalDiscount += birthdayDiscount
        appliedDiscounts = append(appliedDiscounts, "Birthday special: 20%")
    }
    
    // Rule 5: Time-based discounts
    hour := time.Now().Hour()
    if hour >= 14 && hour <= 16 {
        happyHourDiscount := 0.15
        totalDiscount += happyHourDiscount
        appliedDiscounts = append(appliedDiscounts, "Happy hour (2-4 PM): 15%")
    } else if hour >= 6 && hour <= 8 {
        earlyBirdDiscount := 0.10
        totalDiscount += earlyBirdDiscount
        appliedDiscounts = append(appliedDiscounts, "Early bird (6-8 AM): 10%")
    }
    
    // Rule 6: Check eligible promotions
    for _, promo := range promotions {
        if isEligibleForPromotion(customer, orderAmount, promo) {
            // Don't stack with other discounts if total > 50%
            if totalDiscount + promo.Discount <= 0.50 {
                totalDiscount += promo.Discount
                appliedDiscounts = append(appliedDiscounts, 
                    fmt.Sprintf("Promo %s: %.0f%%", promo.Code, promo.Discount*100))
            }
        }
    }
    
    // Rule 7: Maximum discount cap
    maxDiscount := 0.50 // 50% max
    if totalDiscount > maxDiscount {
        totalDiscount = maxDiscount
        appliedDiscounts = append(appliedDiscounts, 
            fmt.Sprintf("(Capped at %.0f%%)", maxDiscount*100))
    }
    
    // Calculate final amount
    discountAmount := orderAmount * totalDiscount
    finalAmount := orderAmount - discountAmount
    
    // Display results
    if len(appliedDiscounts) > 0 {
        fmt.Println("\nDiscounts applied:")
        for _, discount := range appliedDiscounts {
            fmt.Printf("  • %s\n", discount)
        }
        fmt.Printf("\nTotal discount: %.0f%% ($%.2f)\n", 
            totalDiscount*100, discountAmount)
    } else {
        fmt.Println("\nNo discounts applicable")
    }
    
    fmt.Printf("Final amount: $%.2f\n", finalAmount)
    
    // Additional perks based on customer type
    fmt.Println("\nAdditional perks:")
    
    if customer.Type == "vip" {
        fmt.Println("  • Free size upgrade")
        fmt.Println("  • Priority queue")
        fmt.Println("  • Free WiFi premium")
    } else if customer.Type == "employee" {
        fmt.Println("  • Free refills")
        fmt.Println("  • Break room access")
    } else if customer.Type == "member" && customer.OrderCount >= 10 {
        fmt.Println("  • Every 10th drink free")
        fmt.Println("  • Early access to new items")
    }
    
    // Suggestions based on history
    if customer.TotalSpent > 1000 {
        fmt.Println("\n💎 Eligible for VIP upgrade!")
    } else if customer.Type == "regular" && customer.OrderCount >= 5 {
        fmt.Println("\n💳 Consider joining our membership program!")
    }
}

func isEligibleForPromotion(customer Customer, orderAmount float64, promo Promotion) bool {
    // Check if promotion is still valid
    if time.Now().After(promo.ValidUntil) {
        return false
    }
    
    // Check minimum order amount
    if orderAmount < promo.MinOrder {
        return false
    }
    
    // Check usage limit
    if promo.UsageCount >= promo.UsageLimit {
        return false
    }
    
    // Special eligibility rules
    switch promo.Code {
    case "WELCOME10":
        // Only for customers with less than 3 orders
        if customer.OrderCount >= 3 {
            return false
        }
    case "LOYALTY20":
        // Only for customers with 25+ orders
        if customer.OrderCount < 25 {
            return false
        }
    }
    
    return true
}