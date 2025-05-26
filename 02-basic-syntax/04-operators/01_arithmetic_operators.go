package main

import (
    "fmt"
    "math"
)

func main() {
    fmt.Println("=== GoCoffee Arithmetic Operations ===\n")
    
    // Basic arithmetic
    lattePrice := 4.50
    muffinPrice := 2.50
    
    // Addition
    subtotal := lattePrice + muffinPrice
    fmt.Printf("Latte ($%.2f) + Muffin ($%.2f) = $%.2f\n", 
        lattePrice, muffinPrice, subtotal)
    
    // Multiplication
    quantity := 3
    totalMuffins := muffinPrice * float64(quantity)
    fmt.Printf("%d Muffins × $%.2f = $%.2f\n", 
        quantity, muffinPrice, totalMuffins)
    
    // Division
    totalBill := 45.00
    numPeople := 3
    perPerson := totalBill / float64(numPeople)
    fmt.Printf("$%.2f ÷ %d people = $%.2f each\n", 
        totalBill, numPeople, perPerson)
    
    // Modulo (remainder)
    totalCents := 1234
    dollars := totalCents / 100
    cents := totalCents % 100
    fmt.Printf("%d cents = $%d.%02d\n", totalCents, dollars, cents)
    
    // Order of operations
    fmt.Println("\nORDER OF OPERATIONS:")
    
    // Wrong way (no parentheses)
    wrong := 10.00 + 5.00 * 0.08  // Tax calculated on $5 only
    fmt.Printf("Wrong: $10 + $5 × 8%% tax = $%.2f\n", wrong)
    
    // Right way (with parentheses)
    right := (10.00 + 5.00) * 0.08  // Tax on total
    fmt.Printf("Right: ($10 + $5) × 8%% tax = $%.2f\n", right)
    
    // Integer division pitfalls
    fmt.Println("\nINTEGER DIVISION:")
    
    // Integer division truncates
    items := 10
    groups := 3
    perGroup := items / groups
    leftover := items % groups
    fmt.Printf("%d items ÷ %d groups = %d per group, %d leftover\n", 
        items, groups, perGroup, leftover)
    
    // Float division for accuracy
    accuratePerGroup := float64(items) / float64(groups)
    fmt.Printf("Accurate division: %.2f items per group\n", accuratePerGroup)
    
    // Increment and decrement
    fmt.Println("\nCOUNTER OPERATIONS:")
    
    orderNumber := 1000
    fmt.Printf("Current order: #%d\n", orderNumber)
    
    orderNumber++ // Increment
    fmt.Printf("Next order: #%d\n", orderNumber)
    
    orderNumber-- // Decrement
    fmt.Printf("Previous order: #%d\n", orderNumber)
    
    // Compound calculations
    fmt.Println("\nCOMPLEX ORDER CALCULATION:")
    
    // Customer order
    espressos := 2
    lattes := 1
    croissants := 3
    
    espressoPrice := 3.00
    lattePrice = 4.50
    croissantPrice := 3.50
    
    // Calculate subtotal
    orderSubtotal := float64(espressos)*espressoPrice + 
                    float64(lattes)*lattePrice + 
                    float64(croissants)*croissantPrice
    
    // Apply discount
    memberDiscount := 0.10
    discountAmount := orderSubtotal * memberDiscount
    afterDiscount := orderSubtotal - discountAmount
    
    // Add tax
    taxRate := 0.085
    tax := afterDiscount * taxRate
    finalTotal := afterDiscount + tax
    
    // Add tip
    tipPercent := 0.18
    tip := finalTotal * tipPercent
    withTip := finalTotal + tip
    
    fmt.Printf("\nOrder Summary:\n")
    fmt.Printf("Subtotal:      $%.2f\n", orderSubtotal)
    fmt.Printf("Member (10%%): -$%.2f\n", discountAmount)
    fmt.Printf("After discount: $%.2f\n", afterDiscount)
    fmt.Printf("Tax (8.5%%):    $%.2f\n", tax)
    fmt.Printf("Total:         $%.2f\n", finalTotal)
    fmt.Printf("Tip (18%%):     $%.2f\n", tip)
    fmt.Printf("Final:         $%.2f\n", withTip)
    
    // Math package functions
    fmt.Println("\nMATH FUNCTIONS:")
    
    // Rounding
    precise := 4.567
    fmt.Printf("Original: $%.3f\n", precise)
    fmt.Printf("Round: $%.2f\n", math.Round(precise*100)/100)
    fmt.Printf("Ceil:  $%.2f\n", math.Ceil(precise*100)/100)
    fmt.Printf("Floor: $%.2f\n", math.Floor(precise*100)/100)
    
    // Power operations
    dailyGrowth := 1.05 // 5% daily growth
    days := 7
    weeklyGrowth := math.Pow(dailyGrowth, float64(days))
    fmt.Printf("\n5%% daily growth for %d days = %.2f%% total\n", 
        days, (weeklyGrowth-1)*100)
}