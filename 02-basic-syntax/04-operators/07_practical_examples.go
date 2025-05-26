package main

import (
    "fmt"
    "math"
    "time"
)

// Order represents a coffee order
type Order struct {
    items    []OrderItem
    isMember bool
    tipPercent float64
}

type OrderItem struct {
    name     string
    price    float64
    quantity int
    size     string
}

func main() {
    fmt.Println("=== GoCoffee Practical Calculations ===\n")
    
    // Create an order
    order := Order{
        items: []OrderItem{
            {"Latte", 4.50, 2, "large"},
            {"Cappuccino", 4.00, 1, "medium"},
            {"Croissant", 3.50, 3, ""},
            {"Blueberry Muffin", 2.75, 2, ""},
        },
        isMember: true,
        tipPercent: 18,
    }
    
    // Calculate order
    calculateOrder(order)
    
    // Loyalty points calculation
    fmt.Println("\nLOYALTY POINTS CALCULATION:")
    calculateLoyaltyPoints(45.67)
    
    // Split bill calculation
    fmt.Println("\nSPLIT BILL CALCULATION:")
    splitBill(125.80, 4, 20)
    
    // Inventory calculations
    fmt.Println("\nINVENTORY CALCULATIONS:")
    calculateInventoryNeeds()
    
    // Time-based pricing
    fmt.Println("\nDYNAMIC PRICING:")
    calculateDynamicPrice("Latte", 4.50)
}

func calculateOrder(order Order) {
    fmt.Println("ORDER CALCULATION:")
    fmt.Println("==================")
    
    subtotal := 0.0
    
    // Calculate subtotal
    for _, item := range order.items {
        itemTotal := item.price * float64(item.quantity)
        subtotal += itemTotal
        
        fmt.Printf("%d × %s", item.quantity, item.name)
        if item.size != "" {
            fmt.Printf(" (%s)", item.size)
        }
        fmt.Printf(" @ $%.2f = $%.2f\n", item.price, itemTotal)
    }
    
    fmt.Printf("\nSubtotal: $%.2f\n", subtotal)
    
    // Apply discounts
    discount := 0.0
    if order.isMember {
        discount = subtotal * 0.10 // 10% member discount
        fmt.Printf("Member discount (10%%): -$%.2f\n", discount)
    }
    
    // Bulk discount (additional 5% if over $30)
    if subtotal > 30 {
        bulkDiscount := subtotal * 0.05
        discount += bulkDiscount
        fmt.Printf("Bulk discount (5%%): -$%.2f\n", bulkDiscount)
    }
    
    afterDiscount := subtotal - discount
    fmt.Printf("After discounts: $%.2f\n", afterDiscount)
    
    // Calculate tax
    tax := afterDiscount * 0.085
    fmt.Printf("Tax (8.5%%): $%.2f\n", tax)
    
    beforeTip := afterDiscount + tax
    fmt.Printf("Total before tip: $%.2f\n", beforeTip)
    
    // Calculate tip
    tip := beforeTip * (order.tipPercent / 100)
    fmt.Printf("Tip (%.0f%%): $%.2f\n", order.tipPercent, tip)
    
    finalTotal := beforeTip + tip
    fmt.Printf("\nFINAL TOTAL: $%.2f\n", finalTotal)
}

func calculateLoyaltyPoints(amount float64) {
    // 1 point per dollar spent
    basePoints := int(amount)
    
    // Bonus points
    bonusPoints := 0
    
    // Double points on weekends
    if time.Now().Weekday() == time.Saturday || 
       time.Now().Weekday() == time.Sunday {
        bonusPoints += basePoints
        fmt.Printf("Weekend bonus: +%d points\n", basePoints)
    }
    
    // Extra 50 points for orders over $50
    if amount > 50 {
        bonusPoints += 50
        fmt.Println("Large order bonus: +50 points")
    }
    
    totalPoints := basePoints + bonusPoints
    fmt.Printf("Order amount: $%.2f\n", amount)
    fmt.Printf("Base points: %d\n", basePoints)
    fmt.Printf("Bonus points: %d\n", bonusPoints)
    fmt.Printf("Total points earned: %d\n", totalPoints)
    
    // Check rewards
    pointsNeeded := 500
    if totalPoints >= pointsNeeded {
        rewards := totalPoints / pointsNeeded
        remaining := totalPoints % pointsNeeded
        fmt.Printf("Rewards earned: %d free drinks!\n", rewards)
        fmt.Printf("Points remaining: %d\n", remaining)
    } else {
        fmt.Printf("Points until next reward: %d\n", pointsNeeded-totalPoints)
    }
}

func splitBill(total float64, people int, tipPercent float64) {
    // Add tip to total
    tipAmount := total * (tipPercent / 100)
    grandTotal := total + tipAmount
    
    // Calculate per person
    perPerson := grandTotal / float64(people)
    
    // Round up to nearest cent
    perPersonRounded := math.Ceil(perPerson*100) / 100
    
    // Calculate actual total
    actualTotal := perPersonRounded * float64(people)
    overage := actualTotal - grandTotal
    
    fmt.Printf("Bill: $%.2f\n", total)
    fmt.Printf("Tip (%.0f%%): $%.2f\n", tipPercent, tipAmount)
    fmt.Printf("Grand total: $%.2f\n", grandTotal)
    fmt.Printf("Split %d ways: $%.2f per person\n", people, perPersonRounded)
    
    if overage > 0 {
        fmt.Printf("Total collected: $%.2f (%.2f extra)\n", actualTotal, overage)
    }
}

func calculateInventoryNeeds() {
    // Daily usage rates
    coffeeBagsPerDay := 5.5
    milkGallonsPerDay := 12.3
    cupsPerDay := 450
    
    // Current inventory
    currentCoffeeBags := 25.0
    currentMilkGallons := 45.0
    currentCups := 2500
    
    // Calculate days until out
    daysOfCoffee := currentCoffeeBags / coffeeBagsPerDay
    daysOfMilk := currentMilkGallons / milkGallonsPerDay
    daysOfCups := float64(currentCups) / float64(cupsPerDay)
    
    fmt.Println("Current inventory will last:")
    fmt.Printf("Coffee: %.1f days\n", daysOfCoffee)
    fmt.Printf("Milk: %.1f days\n", daysOfMilk)
    fmt.Printf("Cups: %.1f days\n", daysOfCups)
    
    // Calculate order quantities (order when 3 days left)
    reorderDays := 3.0
    deliveryDays := 2.0
    safetyStock := 1.5
    
    targetDays := reorderDays + deliveryDays + safetyStock
    
    if daysOfCoffee < targetDays {
        needed := (targetDays - daysOfCoffee) * coffeeBagsPerDay
        fmt.Printf("\n⚠️  Order %.0f coffee bags NOW!\n", math.Ceil(needed))
    }
    
    if daysOfMilk < targetDays {
        needed := (targetDays - daysOfMilk) * milkGallonsPerDay
        fmt.Printf("⚠️  Order %.0f gallons of milk NOW!\n", math.Ceil(needed))
    }
}

func calculateDynamicPrice(item string, basePrice float64) {
    hour := time.Now().Hour()
    dayOfWeek := time.Now().Weekday()
    
    price := basePrice
    reason := "regular price"
    
    // Happy hour: 3-5 PM weekdays
    if hour >= 15 && hour < 17 && dayOfWeek >= time.Monday && dayOfWeek <= time.Friday {
        price *= 0.8 // 20% off
        reason = "happy hour (-20%)"
    } else if hour >= 6 && hour < 9 {
        // Morning rush: 6-9 AM
        price *= 1.1 // 10% surge
        reason = "morning rush (+10%)"
    } else if dayOfWeek == time.Saturday || dayOfWeek == time.Sunday {
        // Weekend pricing
        if hour >= 10 && hour < 14 {
            price *= 1.15 // 15% brunch surge
            reason = "weekend brunch (+15%)"
        }
    }
    
    fmt.Printf("%s: $%.2f (%s)\n", item, price, reason)
    fmt.Printf("Base price: $%.2f\n", basePrice)
    
    if price != basePrice {
        difference := price - basePrice
        percent := (difference / basePrice) * 100
        fmt.Printf("Difference: $%.2f (%.0f%%)\n", difference, percent)
    }
}