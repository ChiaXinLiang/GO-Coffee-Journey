package main

import "fmt"

func main() {
    fmt.Println("=== GoCoffee Basic Printing ===\n")
    
    // Print vs Println vs Printf
    fmt.Print("Welcome ")
    fmt.Print("to ")
    fmt.Print("GoCoffee")
    fmt.Print("\n\n") // Need manual newline
    
    // Println adds newline automatically
    fmt.Println("Welcome")
    fmt.Println("to")
    fmt.Println("GoCoffee")
    fmt.Println() // Empty line
    
    // Printf allows formatting
    storeName := "GoCoffee Downtown"
    year := 2024
    rating := 4.8
    
    fmt.Printf("Welcome to %s!\n", storeName)
    fmt.Printf("Established: %d\n", year)
    fmt.Printf("Rating: %.1f stars\n", rating)
    
    // Multiple values
    fmt.Printf("\n%s (est. %d) - %.1f★\n", storeName, year, rating)
    
    // Common mistakes
    fmt.Println("\nCOMMON MISTAKES:")
    
    // Forgetting newline in Printf
    fmt.Printf("This needs a newline")
    fmt.Printf("This will be on same line\n")
    
    // Using wrong verb
    price := 4.50
    fmt.Printf("\nPrice: %d\n", price) // Wrong! %d for integers
    fmt.Printf("Price: %.2f\n", price) // Correct! %f for floats
    
    // Building strings
    fmt.Println("\nBUILDING STRINGS:")
    
    // Using + operator (inefficient for many strings)
    greeting1 := "Hello, " + "Marcus" + "!"
    fmt.Println(greeting1)
    
    // Using fmt.Sprintf (better)
    customerName := "Sarah"
    greeting2 := fmt.Sprintf("Hello, %s!", customerName)
    fmt.Println(greeting2)
    
    // Sprint variations
    coffee := "Latte"
    size := "Large"
    
    // Sprint concatenates with spaces
    order1 := fmt.Sprint(coffee, size)
    fmt.Printf("Sprint: '%s'\n", order1)
    
    // Sprintln adds spaces and newline
    order2 := fmt.Sprintln(coffee, size)
    fmt.Printf("Sprintln: '%s'", order2) // Already has newline
    
    // Sprintf with format
    order3 := fmt.Sprintf("%s (%s)", coffee, size)
    fmt.Printf("Sprintf: '%s'\n", order3)
}