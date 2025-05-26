package main

import "fmt"

func main() {
	// BAD: Unclear variable names
	var x = 4.50
	var y = 2
	var z = x * float64(y)

	// GOOD: Clear, descriptive names
	var coffeePrice = 4.50
	var quantity = 2
	var totalPrice = coffeePrice * float64(quantity)

	// Go naming conventions
	var (
		// Use camelCase for multi-word names
		customerFirstName = "Marcus"
		orderStatus       = "pending"

		// Acronyms should be all caps
		apiURL   = "https://api.gocoffee.com"
		httpPort = 8080

		// Boolean names should be clear
		isReady     = true
		hasDiscount = false
		canRefund   = true
	)

	fmt.Println("=== Naming Conventions Demo ===")
	fmt.Printf("Bad example: x=%.2f, y=%d, z=%.2f\n", x, y, z)
	fmt.Printf("Good example: price=%.2f, qty=%d, total=%.2f\n",
		coffeePrice, quantity, totalPrice)
	fmt.Println("\nOther examples:")
	fmt.Printf("Customer: %s\n", customerFirstName)
	fmt.Printf("Order status: %s\n", orderStatus)
	fmt.Printf("API URL: %s\n", apiURL)
	fmt.Printf("HTTP Port: %d\n", httpPort)
	fmt.Printf("Order ready: %v\n", isReady)
	fmt.Printf("Has discount: %v\n", hasDiscount)
	fmt.Printf("Can refund: %v\n", canRefund)
}