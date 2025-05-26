package main

import "fmt"

func main() {
	// Variables without initial values get "zero values"
	var (
		coffeeBags   int     // 0
		price        float64 // 0.0
		customerName string  // ""
		isOpen       bool    // false
	)

	fmt.Println("=== Zero Values Demo ===")
	fmt.Printf("Coffee bags: %d (zero value for int)\n", coffeeBags)
	fmt.Printf("Price: $%.2f (zero value for float64)\n", price)
	fmt.Printf("Customer name: '%s' (zero value for string)\n", customerName)
	fmt.Printf("Is open: %v (zero value for bool)\n", isOpen)

	// This prevents accidental null/nil errors!
	fmt.Println("\nThis is why Go is safer than many languages!")
}