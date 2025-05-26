package main

import "fmt"

func main() {
    fmt.Println("=== GoCoffee Assignment Operators ===\n")
    
    // Basic assignment
    total := 0.0
    fmt.Printf("Starting total: $%.2f\n", total)
    
    // Addition assignment
    fmt.Println("\nADDING ITEMS TO ORDER:")
    
    lattePrice := 4.50
    total += lattePrice
    fmt.Printf("Added latte ($%.2f): Total = $%.2f\n", lattePrice, total)
    
    muffinPrice := 3.00
    total += muffinPrice
    fmt.Printf("Added muffin ($%.2f): Total = $%.2f\n", muffinPrice, total)
    
    // Subtraction assignment
    discount := 1.00
    total -= discount
    fmt.Printf("Applied discount (-$%.2f): Total = $%.2f\n", discount, total)
    
    // Multiplication assignment
    fmt.Println("\nAPPLYING MULTIPLIERS:")
    
    // Apply tax (8.5%)
    taxRate := 1.085
    total *= taxRate
    fmt.Printf("After tax (×%.3f): Total = $%.2f\n", taxRate, total)
    
    // Division assignment
    fmt.Println("\nSPLITTING THE BILL:")
    
    people := 2
    perPerson := total
    perPerson /= float64(people)
    fmt.Printf("Split %d ways: $%.2f per person\n", people, perPerson)
    
    // Integer operations
    fmt.Println("\nINVENTORY MANAGEMENT:")
    
    coffeeBags := 100
    fmt.Printf("Starting inventory: %d bags\n", coffeeBags)
    
    // Daily usage
    morningUse := 15
    coffeeBags -= morningUse
    fmt.Printf("After morning (-15): %d bags\n", coffeeBags)
    
    afternoonUse := 20
    coffeeBags -= afternoonUse
    fmt.Printf("After afternoon (-20): %d bags\n", coffeeBags)
    
    // Delivery arrived
    delivery := 50
    coffeeBags += delivery
    fmt.Printf("After delivery (+50): %d bags\n", coffeeBags)
    
    // Modulo assignment
    fmt.Println("\nPOINTS CALCULATION:")
    
    loyaltyPoints := 1247
    fmt.Printf("Total points: %d\n", loyaltyPoints)
    
    pointsPerReward := 100
    rewards := loyaltyPoints / pointsPerReward
    loyaltyPoints %= pointsPerReward // Remaining points
    
    fmt.Printf("Rewards earned: %d\n", rewards)
    fmt.Printf("Points remaining: %d\n", loyaltyPoints)
    
    // Multiple assignments
    fmt.Println("\nMULTIPLE ASSIGNMENTS:")
    
    // Swap values
    hot, cold := "Latte", "Iced Coffee"
    fmt.Printf("Before swap: hot=%s, cold=%s\n", hot, cold)
    
    hot, cold = cold, hot
    fmt.Printf("After swap: hot=%s, cold=%s\n", hot, cold)
    
    // Parallel assignment
    name, price, available := "Mocha", 5.00, true
    fmt.Printf("\nProduct: %s, Price: $%.2f, Available: %v\n", 
        name, price, available)
    
    // Practical example: Running totals
    fmt.Println("\nDAILY SALES TRACKER:")
    
    type Sale struct {
        time   string
        amount float64
    }
    
    sales := []Sale{
        {"9:00 AM", 45.50},
        {"10:30 AM", 23.75},
        {"12:00 PM", 67.25},
        {"2:15 PM", 34.00},
        {"4:30 PM", 89.50},
    }
    
    dailyTotal := 0.0
    transactionCount := 0
    
    fmt.Println("Time      Amount    Running Total")
    fmt.Println("------------------------------------")
    
    for _, sale := range sales {
        dailyTotal += sale.amount
        transactionCount++
        
        fmt.Printf("%-9s $%6.2f   $%7.2f\n", 
            sale.time, sale.amount, dailyTotal)
    }
    
    averageSale := dailyTotal / float64(transactionCount)
    fmt.Printf("\nTotal sales: $%.2f\n", dailyTotal)
    fmt.Printf("Transactions: %d\n", transactionCount)
    fmt.Printf("Average sale: $%.2f\n", averageSale)
}