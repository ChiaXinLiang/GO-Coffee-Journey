package main

import "fmt"

func main() {
	fmt.Println("=== Common Variable Mistakes ===")

	// Mistake 1: Unused variables
	// var unused = "This will cause a compile error"
	// Solution: Use _ for variables you must declare but don't use
	data, _ := getData()
	fmt.Printf("Data: %s\n", data)

	// Mistake 2: Shadowing
	price := 4.50
	fmt.Printf("Original price: $%.2f\n", price)
	{
		price := 3.00 // This shadows the outer price
		fmt.Printf("Inner price: $%.2f\n", price)
	}
	fmt.Printf("Price after block: $%.2f\n", price) // Still 4.50

	// Mistake 3: Wrong type operations
	quantity := 5
	discount := 0.1
	// total := quantity * discount // Error: mismatched types
	total := float64(quantity) * discount // Correct: explicit conversion
	fmt.Printf("Discount amount: $%.2f\n", total)

	// Mistake 4: Modifying strings (strings are immutable)
	coffee := "Latte"
	// coffee[0] = 'M' // Error: cannot assign to coffee[0]
	// Solution: Create new string
	coffee = "Mocha"
	fmt.Printf("New coffee: %s\n", coffee)
}

func getData() (string, error) {
	return "Coffee data", nil
}