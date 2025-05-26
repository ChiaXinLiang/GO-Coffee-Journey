package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println("=== GoCoffee Logical Operations ===\n")
    
    // Store status
    currentHour := time.Now().Hour()
    dayOfWeek := time.Now().Weekday()
    
    isWeekday := dayOfWeek >= time.Monday && dayOfWeek <= time.Friday
    isWeekend := dayOfWeek == time.Saturday || dayOfWeek == time.Sunday
    
    // Store hours logic
    weekdayOpen := 6
    weekdayClose := 22
    weekendOpen := 7
    weekendClose := 23
    
    var isOpen bool
    if isWeekday {
        isOpen = currentHour >= weekdayOpen && currentHour < weekdayClose
    } else {
        isOpen = currentHour >= weekendOpen && currentHour < weekendClose
    }
    
    fmt.Printf("Current time: %d:00 on %s\n", currentHour, dayOfWeek)
    fmt.Printf("Is weekday: %v\n", isWeekday)
    fmt.Printf("Is weekend: %v\n", isWeekend)
    fmt.Printf("Store is open: %v\n", isOpen)
    
    // Complex order eligibility
    fmt.Println("\nFREE DELIVERY ELIGIBILITY:")
    
    type Customer struct {
        name      string
        isPremium bool
        orderTotal float64
        distance  float64 // miles
    }
    
    customers := []Customer{
        {"Alice", true, 25.00, 2.5},
        {"Bob", false, 35.00, 1.5},
        {"Charlie", false, 15.00, 3.0},
        {"Diana", true, 10.00, 4.0},
    }
    
    const (
        freeDeliveryMinimum = 30.00
        maxDeliveryDistance = 5.0
        premiumMinimum      = 15.00
    )
    
    for _, customer := range customers {
        // Free delivery if:
        // - Premium AND order >= $15 OR
        // - Non-premium AND order >= $30
        // AND always: distance <= 5 miles
        
        qualifiesForFree := (customer.isPremium && customer.orderTotal >= premiumMinimum) ||
                           (!customer.isPremium && customer.orderTotal >= freeDeliveryMinimum)
        
        withinRange := customer.distance <= maxDeliveryDistance
        
        getsFreeDelivery := qualifiesForFree && withinRange
        
        fmt.Printf("\n%s: Premium=%v, $%.2f order, %.1f miles\n", 
            customer.name, customer.isPremium, customer.orderTotal, customer.distance)
        fmt.Printf("  Qualifies by amount: %v\n", qualifiesForFree)
        fmt.Printf("  Within range: %v\n", withinRange)
        fmt.Printf("  → Free delivery: %v\n", getsFreeDelivery)
    }
    
    // Short-circuit evaluation
    fmt.Println("\nSHORT-CIRCUIT EVALUATION:")
    
    // This demonstrates how Go stops evaluating when result is known
    inventory := map[string]int{
        "coffee": 50,
        "milk": 20,
        "sugar": 100,
    }
    
    // Check if we can make a latte
    canMakeLatte := checkInventory("coffee", 1, inventory) &&
                    checkInventory("milk", 8, inventory) &&
                    checkInventory("sugar", 2, inventory)
    
    fmt.Printf("Can make latte: %v\n", canMakeLatte)
    
    // Negation operator
    fmt.Println("\nNEGATION EXAMPLES:")
    
    isClosingSoon := currentHour >= 21
    stayOpen := !isClosingSoon
    
    fmt.Printf("Closing soon (after 9 PM): %v\n", isClosingSoon)
    fmt.Printf("Stay open: %v\n", stayOpen)
    
    // Complex business rules
    fmt.Println("\nHAPPY HOUR LOGIC:")
    
    isHappyHour := (isWeekday && currentHour >= 15 && currentHour < 17) ||
                   (isWeekend && currentHour >= 14 && currentHour < 16)
    
    fmt.Printf("Happy hour active: %v\n", isHappyHour)
    
    // Conditional assignments
    fmt.Println("\nCONDITIONAL PRICING:")
    
    basePrice := 4.50
    var finalPrice float64
    
    if isHappyHour && isOpen {
        finalPrice = basePrice * 0.8 // 20% off
        fmt.Printf("Happy hour price: $%.2f (20%% off)\n", finalPrice)
    } else if isWeekend && isOpen {
        finalPrice = basePrice * 1.1 // 10% weekend surcharge
        fmt.Printf("Weekend price: $%.2f (10%% surcharge)\n", finalPrice)
    } else if isOpen {
        finalPrice = basePrice
        fmt.Printf("Regular price: $%.2f\n", finalPrice)
    } else {
        finalPrice = 0
        fmt.Println("Store is closed")
    }
}

func checkInventory(item string, needed int, inventory map[string]int) bool {
    available, exists := inventory[item]
    result := exists && available >= needed
    
    if result {
        fmt.Printf("  ✓ %s: need %d, have %d\n", item, needed, available)
    } else {
        fmt.Printf("  ✗ %s: need %d, have %d\n", item, needed, available)
    }
    
    return result
}