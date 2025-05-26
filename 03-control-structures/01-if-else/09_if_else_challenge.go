package main

import (
    "fmt"
    "time"
)

// Challenge: Build a comprehensive order validation system
// Requirements:
// 1. Validate customer eligibility
// 2. Check item availability
// 3. Apply dynamic pricing
// 4. Handle special requests
// 5. Process loyalty rewards

type Customer struct {
    ID            string
    Name          string
    Age           int
    IsMember      bool
    MemberTier    string // "bronze", "silver", "gold"
    LoyaltyPoints int
    IsBanned      bool
    LastVisit     time.Time
}

type OrderItem struct {
    Name         string
    Category     string
    BasePrice    float64
    CustomOptions []string
}

func main() {
    fmt.Println("=== GoCoffee If-Else Challenge ===\n")
    fmt.Println("TODO: Implement a comprehensive validation system")
    fmt.Println("\nRequirements:")
    fmt.Println("1. Check customer age for certain items")
    fmt.Println("2. Validate opening hours and day-specific menus")
    fmt.Println("3. Apply tier-based discounts")
    fmt.Println("4. Handle combo deals")
    fmt.Println("5. Process loyalty point redemption")
    
    // Test data
    customer := Customer{
        ID:            "C123",
        Name:          "Test Customer",
        Age:           25,
        IsMember:      true,
        MemberTier:    "silver",
        LoyaltyPoints: 150,
        IsBanned:      false,
        LastVisit:     time.Now().AddDate(0, 0, -5),
    }
    
    items := []OrderItem{
        {
            Name:      "Espresso Martini",
            Category:  "special",
            BasePrice: 8.50,
        },
        {
            Name:      "Breakfast Combo",
            Category:  "combo",
            BasePrice: 12.00,
        },
    }
    
    // Your implementation here...
    fmt.Printf("\nCustomer: %+v\n", customer)
    fmt.Printf("Items: %+v\n", items)
    
    fmt.Println("\nImplement the validation logic above!")
}

// Hints:
// - Check age for alcohol-based coffee drinks
// - Breakfast combos only available before 11 AM
// - Gold members get 20% off, Silver 10%, Bronze 5%
// - Banned customers cannot order
// - Loyalty points: 100 points = $10 off
// - Weekend surcharge: 5% on Saturdays and Sundays
// - Happy hour: 3-5 PM weekdays, 15% off