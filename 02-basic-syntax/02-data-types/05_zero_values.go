package main

import "fmt"

func main() {
	fmt.Println("=== GoCoffee Zero Values ===\n")

	// Numeric zero values
	var (
		ordersToday int
		temperature float64
		stockLevel  uint32
		revenue     float32
	)

	fmt.Println("NUMERIC ZERO VALUES:")
	fmt.Printf("Orders today: %d (zero value for int)\n", ordersToday)
	fmt.Printf("Temperature: %.1f°C (zero value for float64)\n", temperature)
	fmt.Printf("Stock level: %d (zero value for uint32)\n", stockLevel)
	fmt.Printf("Revenue: $%.2f (zero value for float32)\n", revenue)

	// String zero value
	var (
		customerName string
		orderNotes   string
	)

	fmt.Println("\nSTRING ZERO VALUES:")
	fmt.Printf("Customer name: '%s' (empty string)\n", customerName)
	fmt.Printf("Order notes: '%s' (empty string)\n", orderNotes)

	// Boolean zero value
	var (
		isReady bool
		isPaid  bool
	)

	fmt.Println("\nBOOLEAN ZERO VALUES:")
	fmt.Printf("Is ready: %v (always false)\n", isReady)
	fmt.Printf("Is paid: %v (always false)\n", isPaid)

	// Proper initialization
	fmt.Println("\nPROPER INITIALIZATION:")

	// Explicit initialization
	coffeeStock := 100
	milkLiters := 50.5
	shopName := "GoCoffee Downtown"
	isOpen := true

	fmt.Printf("Coffee stock: %d bags\n", coffeeStock)
	fmt.Printf("Milk: %.1f liters\n", milkLiters)
	fmt.Printf("Shop: %s\n", shopName)
	fmt.Printf("Open: %v\n", isOpen)

	// Multiple variable initialization
	var (
		espressoPrice   = 3.00
		lattePrice      = 4.50
		cappuccinoPrice = 4.00
	)

	fmt.Printf("\nPrices: Espresso $%.2f, Latte $%.2f, Cappuccino $%.2f\n",
		espressoPrice, lattePrice, cappuccinoPrice)

	// Checking for zero values
	fmt.Println("\nCHECKING ZERO VALUES:")

	if customerName == "" {
		fmt.Println("⚠️  No customer name provided!")
	}

	if ordersToday == 0 {
		fmt.Println("📊 No orders yet today")
	}

	if !isPaid {
		fmt.Println("💳 Payment pending")
	}
}