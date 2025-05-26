package main

import (
	"fmt"
	"time"
)

func main() {
	// Customer information
	customerName := "Sarah"
	membershipLevel := "Gold"

	// Order details
	coffeeType := "Cappuccino"
	size := "Large"
	quantity := 2

	// Pricing
	basePrice := 4.50
	sizeMultiplier := 1.5  // Large = 1.5x price
	memberDiscount := 0.10 // Gold members get 10% off

	// Calculate total
	subtotal := basePrice * sizeMultiplier * float64(quantity)
	discount := subtotal * memberDiscount
	total := subtotal - discount

	// Order metadata
	orderID := 1001
	orderTime := time.Now()

	// Print receipt
	fmt.Println("=================================")
	fmt.Println("        GoCoffee Receipt         ")
	fmt.Println("=================================")
	fmt.Printf("Order #%d\n", orderID)
	fmt.Printf("Time: %s\n", orderTime.Format("2006-01-02 15:04:05"))
	fmt.Println("---------------------------------")
	fmt.Printf("Customer: %s (%s Member)\n", customerName, membershipLevel)
	fmt.Printf("Item: %s %s x%d\n", size, coffeeType, quantity)
	fmt.Println("---------------------------------")
	fmt.Printf("Subtotal:  $%.2f\n", subtotal)
	fmt.Printf("Discount:  -$%.2f\n", discount)
	fmt.Printf("Total:     $%.2f\n", total)
	fmt.Println("=================================")
	fmt.Println("    Thank you for your order!    ")
	fmt.Println("=================================")
}