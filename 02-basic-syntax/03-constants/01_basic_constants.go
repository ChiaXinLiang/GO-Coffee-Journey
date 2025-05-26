package main

import "fmt"

// Package-level constants
const (
	// Company information
	CompanyName = "GoCoffee"
	FoundedYear = 2020
	TaxRate     = 0.085 // 8.5%

	// Store hours
	OpeningHour = 6  // 6 AM
	ClosingHour = 22 // 10 PM
)

func main() {
	fmt.Println("=== GoCoffee Constants ===\n")

	// Single constant declaration
	const maxOrderItems = 10

	// Multiple constants
	const (
		minOrderAmount = 3.00
		maxOrderAmount = 500.00
		deliveryFee    = 2.50
	)

	// Typed constants
	const (
		storeName     string  = "GoCoffee Downtown"
		storeZipCode  int     = 98101
		isOpen24Hours bool    = false
		latitude      float64 = 47.6062
	)

	fmt.Printf("Welcome to %s!\n", CompanyName)
	fmt.Printf("Founded: %d\n", FoundedYear)
	fmt.Printf("Tax Rate: %.1f%%\n", TaxRate*100)
	fmt.Printf("Hours: %d:00 - %d:00\n", OpeningHour, ClosingHour)

	fmt.Printf("\nOrder Limits:\n")
	fmt.Printf("Min order: $%.2f\n", minOrderAmount)
	fmt.Printf("Max order: $%.2f\n", maxOrderAmount)
	fmt.Printf("Max items per order: %d\n", maxOrderItems)

	fmt.Printf("\nStore Info:\n")
	fmt.Printf("Location: %s (ZIP: %d)\n", storeName, storeZipCode)
	fmt.Printf("24 hours: %v\n", isOpen24Hours)

	// Constants are immutable
	// maxOrderItems = 20 // Error: cannot assign to maxOrderItems
}