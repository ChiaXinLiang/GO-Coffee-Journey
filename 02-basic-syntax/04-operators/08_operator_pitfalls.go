package main

import (
    "fmt"
    "math"
)

func main() {
    fmt.Println("=== GoCoffee Operator Pitfalls ===\n")
    
    // Pitfall 1: Integer division
    fmt.Println("PITFALL 1: Integer Division")
    
    totalBill := 100
    people := 3
    
    // Wrong: integer division truncates
    wrongSplit := totalBill / people
    fmt.Printf("Wrong: $%d ÷ %d = $%d per person\n", totalBill, people, wrongSplit)
    fmt.Printf("Total collected: $%d (Lost $%d!)\n", wrongSplit*people, totalBill-wrongSplit*people)
    
    // Right: convert to float
    rightSplit := float64(totalBill) / float64(people)
    fmt.Printf("Right: $%d ÷ %d = $%.2f per person\n", totalBill, people, rightSplit)
    
    // Pitfall 2: Floating point comparison
    fmt.Println("\nPITFALL 2: Float Comparison")
    
    price := 0.1 + 0.2
    expected := 0.3
    
    fmt.Printf("0.1 + 0.2 = %.17f\n", price)
    fmt.Printf("Expected:   %.17f\n", expected)
    fmt.Printf("Are they equal? %v\n", price == expected)
    
    // Solution: use epsilon comparison
    epsilon := 0.0001
    equal := math.Abs(price-expected) < epsilon
    fmt.Printf("Within epsilon? %v\n", equal)
    
    // Pitfall 3: Operator precedence
    fmt.Println("\nPITFALL 3: Operator Precedence")
    
    items := 5
    discount := 0.1
    price1 := 10.0
    
    // Wrong: multiplication before addition
    wrong := price1 + price1 * float64(items) * discount
    fmt.Printf("Wrong: %.2f + %.2f × %d × %.1f = %.2f\n", 
        price1, price1, items, discount, wrong)
    
    // Right: use parentheses
    right := (price1 + price1 * float64(items)) * discount
    fmt.Printf("Right: (%.2f + %.2f × %d) × %.1f = %.2f\n", 
        price1, price1, items, discount, right)
    
    // Pitfall 4: Assignment vs comparison
    fmt.Println("\nPITFALL 4: Assignment in Conditions")
    
    isOpen := true
    
    // This is valid Go (unlike C/C++)
    if isOpen {
        fmt.Println("Store is open")
    }
    
    // Go prevents accidental assignment
    // if isOpen = false { } // This won't compile!
    
    // Pitfall 5: Negative modulo
    fmt.Println("\nPITFALL 5: Negative Modulo")
    
    balance := -17
    adjustment := 5
    
    remainder := balance % adjustment
    fmt.Printf("%d %% %d = %d\n", balance, adjustment, remainder)
    
    // For always-positive result:
    positiveRem := ((balance % adjustment) + adjustment) % adjustment
    fmt.Printf("Positive remainder: %d\n", positiveRem)
    
    // Pitfall 6: Type mixing
    fmt.Println("\nPITFALL 6: Type Mixing")
    
    var count int32 = 10
    var multiplier int64 = 100
    
    // Wrong: can't mix types
    // result := count * multiplier // Error!
    
    // Right: explicit conversion
    result := int64(count) * multiplier
    fmt.Printf("%d × %d = %d\n", count, multiplier, result)
    
    // Pitfall 7: Short-circuit evaluation
    fmt.Println("\nPITFALL 7: Short-Circuit Side Effects")
    
    processed := 0
    
    // This function has side effects
    processOrder := func() bool {
        processed++
        return true
    }
    
    // Second function might not be called!
    if false && processOrder() {
        fmt.Println("This won't print")
    }
    fmt.Printf("Orders processed: %d (not 1!)\n", processed)
    
    // Better: evaluate separately if side effects matter
    shouldProcess := processOrder()
    if false && shouldProcess {
        fmt.Println("This won't print")
    }
    fmt.Printf("Orders processed: %d\n", processed)
}