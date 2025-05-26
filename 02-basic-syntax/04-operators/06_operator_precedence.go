package main

import "fmt"

func main() {
    fmt.Println("=== GoCoffee Operator Precedence ===\n")
    
    // Common precedence mistakes
    fmt.Println("PRECEDENCE MISTAKES:")
    
    // Multiplication before addition
    price1 := 10.00
    price2 := 5.00
    taxRate := 0.08
    
    // Wrong: tax only on second item
    wrong := price1 + price2 * taxRate
    fmt.Printf("Wrong: $%.2f + $%.2f × %.0f%% = $%.2f\n", 
        price1, price2, taxRate*100, wrong)
    
    // Right: tax on total
    right := (price1 + price2) * taxRate
    fmt.Printf("Right: ($%.2f + $%.2f) × %.0f%% = $%.2f\n", 
        price1, price2, taxRate*100, right)
    
    // Comparison vs arithmetic
    fmt.Println("\nCOMPARISON PRECEDENCE:")
    
    minOrder := 10.00
    discount := 2.00
    total := 15.00
    
    // This works as expected (> has lower precedence than -)
    qualifies := total - discount > minOrder
    fmt.Printf("$%.2f - $%.2f > $%.2f = %v\n", 
        total, discount, minOrder, qualifies)
    
    // Logical operator precedence
    fmt.Println("\nLOGICAL PRECEDENCE:")
    
    isMember := true
    orderAmount := 25.00
    itemCount := 3
    
    // AND has higher precedence than OR
    // This is interpreted as: isMember || (orderAmount > 30 && itemCount > 5)
    eligible1 := isMember || orderAmount > 30 && itemCount > 5
    fmt.Printf("member || amount > 30 && items > 5 = %v\n", eligible1)
    
    // Use parentheses for clarity
    eligible2 := (isMember || orderAmount > 30) && itemCount > 5
    fmt.Printf("(member || amount > 30) && items > 5 = %v\n", eligible2)
    
    // Complex calculation
    fmt.Println("\nCOMPLEX CALCULATION:")
    
    basePrice := 20.00
    memberDiscount := 0.10
    bulkDiscount := 0.05
    quantity := 3
    tax := 0.085
    
    // Step by step
    step1 := basePrice * float64(quantity)
    fmt.Printf("Step 1: $%.2f × %d = $%.2f\n", basePrice, quantity, step1)
    
    step2 := step1 * (1 - memberDiscount)
    fmt.Printf("Step 2: $%.2f × (1 - %.0f%%) = $%.2f\n", 
        step1, memberDiscount*100, step2)
    
    step3 := step2 * (1 - bulkDiscount)
    fmt.Printf("Step 3: $%.2f × (1 - %.0f%%) = $%.2f\n", 
        step2, bulkDiscount*100, step3)
    
    step4 := step3 * (1 + tax)
    fmt.Printf("Step 4: $%.2f × (1 + %.1f%%) = $%.2f\n", 
        step3, tax*100, step4)
    
    // All in one (hard to read!)
    allInOne := basePrice * float64(quantity) * (1 - memberDiscount) * 
                (1 - bulkDiscount) * (1 + tax)
    fmt.Printf("\nAll in one: $%.2f\n", allInOne)
    
    // Better with intermediate variables
    subtotal := basePrice * float64(quantity)
    afterMember := subtotal * (1 - memberDiscount)
    afterBulk := afterMember * (1 - bulkDiscount)
    finalTotal := afterBulk * (1 + tax)
    
    fmt.Printf("With variables: $%.2f\n", finalTotal)
    
    // Precedence table
    fmt.Println("\nOPERATOR PRECEDENCE (highest to lowest):")
    fmt.Println("1. () [] . ->")
    fmt.Println("2. ! ++ -- + - * & (unary)")
    fmt.Println("3. * / %")
    fmt.Println("4. + -")
    fmt.Println("5. << >>")
    fmt.Println("6. < <= > >=")
    fmt.Println("7. == !=")
    fmt.Println("8. &")
    fmt.Println("9. ^")
    fmt.Println("10. |")
    fmt.Println("11. &&")
    fmt.Println("12. ||")
    fmt.Println("13. = += -= *= /= %= etc.")
}