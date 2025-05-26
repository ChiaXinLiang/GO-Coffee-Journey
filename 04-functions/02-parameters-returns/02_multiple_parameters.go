package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	fmt.Println("=== GoCoffee Multiple Parameters ===\n")
	
	// Example 1: Different numbers of parameters
	fmt.Println("Example 1: Functions with Different Parameter Counts")
	fmt.Println("---------------------------------------------------")
	demonstrateParameterCounts()
	
	// Example 2: Parameter ordering
	fmt.Println("\n\nExample 2: Parameter Order Matters")
	fmt.Println("---------------------------------")
	demonstrateParameterOrder()
	
	// Example 3: Complex parameter combinations
	fmt.Println("\n\nExample 3: Complex Parameter Types")
	fmt.Println("---------------------------------")
	demonstrateComplexParameters()
	
	// Example 4: Struct parameters
	fmt.Println("\n\nExample 4: Using Structs for Many Parameters")
	fmt.Println("-------------------------------------------")
	demonstrateStructParameters()
}

// Example 1: Different parameter counts
func demonstrateParameterCounts() {
	// One parameter
	brewSimple("Espresso")
	
	// Two parameters
	brewWithSize("Latte", "Large")
	
	// Three parameters
	brewWithOptions("Cappuccino", "Medium", 2)
	
	// Many parameters
	brewComplete("Mocha", "Large", 1, true, false, "Almond")
}

func brewSimple(coffeeType string) {
	fmt.Printf("Brewing: %s\n", coffeeType)
}

func brewWithSize(coffeeType, size string) {
	fmt.Printf("Brewing: %s %s\n", size, coffeeType)
}

func brewWithOptions(coffeeType, size string, shots int) {
	fmt.Printf("Brewing: %s %s with %d shot(s)\n", size, coffeeType, shots)
}

func brewComplete(
	coffeeType string,
	size string,
	extraShots int,
	decaf bool,
	iced bool,
	milkType string,
) {
	fmt.Printf("\nComplete order:\n")
	fmt.Printf("  Drink: %s\n", coffeeType)
	fmt.Printf("  Size: %s\n", size)
	
	if extraShots > 0 {
		fmt.Printf("  Extra shots: %d\n", extraShots)
	}
	
	if decaf {
		fmt.Println("  ✓ Decaf")
	}
	
	if iced {
		fmt.Println("  ✓ Iced")
	}
	
	fmt.Printf("  Milk: %s\n", milkType)
}

// Example 2: Parameter order importance
func demonstrateParameterOrder() {
	// Correct order
	createOrder("Marcus", "Latte", 4.50)
	
	// Same function, different argument order = different result
	// This would be wrong: createOrder("Latte", "Marcus", 4.50)
	
	// Better design with clear parameter names
	createDetailedOrder(
		"ORD-123",     // orderID
		"Sarah",       // customerName
		"Cappuccino",  // item
		3.75,         // price
	)
}

func createOrder(customer, item string, price float64) {
	fmt.Printf("\nOrder for %s: %s ($%.2f)\n", customer, item, price)
}

func createDetailedOrder(orderID, customerName, item string, price float64) {
	fmt.Printf("\nDetailed Order:\n")
	fmt.Printf("  Order ID: %s\n", orderID)
	fmt.Printf("  Customer: %s\n", customerName)
	fmt.Printf("  Item: %s\n", item)
	fmt.Printf("  Price: $%.2f\n", price)
}

// Example 3: Complex parameter types
func demonstrateComplexParameters() {
	// Slice parameters
	items := []string{"Latte", "Croissant", "Muffin"}
	prices := []float64{4.50, 3.25, 3.00}
	calculateBill(items, prices, 0.08)
	
	// Map parameters
	inventory := map[string]int{
		"Coffee Beans": 50,
		"Milk":        30,
		"Sugar":       100,
	}
	checkInventory(inventory, 10)
	
	// Function parameters (callbacks)
	processOrderWithCallback(
		"Espresso",
		func(item string) {
			fmt.Printf("  Callback: Preparing %s\n", item)
		},
	)
}

func calculateBill(items []string, prices []float64, taxRate float64) {
	if len(items) != len(prices) {
		fmt.Println("\nError: Items and prices mismatch")
		return
	}
	
	fmt.Println("\n--- BILL ---")
	subtotal := 0.0
	
	for i, item := range items {
		fmt.Printf("%-15s $%.2f\n", item, prices[i])
		subtotal += prices[i]
	}
	
	tax := subtotal * taxRate
	total := subtotal + tax
	
	fmt.Println(strings.Repeat("-", 25))
	fmt.Printf("Subtotal:       $%.2f\n", subtotal)
	fmt.Printf("Tax (%.0f%%):      $%.2f\n", taxRate*100, tax)
	fmt.Printf("Total:          $%.2f\n", total)
}

func checkInventory(stock map[string]int, minLevel int) {
	fmt.Println("\n--- INVENTORY CHECK ---")
	
	for item, quantity := range stock {
		status := "✓ OK"
		if quantity < minLevel {
			status = "⚠️  LOW"
		}
		fmt.Printf("%-15s: %3d units %s\n", item, quantity, status)
	}
}

func processOrderWithCallback(item string, callback func(string)) {
	fmt.Printf("\nProcessing order with callback:\n")
	fmt.Printf("  Item: %s\n", item)
	callback(item)
	fmt.Println("  Order complete!")
}

// Example 4: Struct parameters for complex data
func demonstrateStructParameters() {
	// Define order struct
	type Order struct {
		ID         string
		Customer   string
		Items      []string
		Quantities []int
		Timestamp  time.Time
		Paid       bool
		Notes      string
	}
	
	// Create order
	order := Order{
		ID:         "ORD-456",
		Customer:   "Bob Wilson",
		Items:      []string{"Americano", "Bagel"},
		Quantities: []int{2, 1},
		Timestamp:  time.Now(),
		Paid:       true,
		Notes:      "Extra hot",
	}
	
	// Pass struct to function
	processFullOrder(order)
	
	// Alternative: Order configuration
	type CoffeeConfig struct {
		Type        string
		Size        string
		Temperature int
		Shots       int
		Decaf       bool
		MilkType    string
		Sweetener   string
		ExtraFoam   bool
	}
	
	config := CoffeeConfig{
		Type:        "Latte",
		Size:        "Large",
		Temperature: 140,
		Shots:       2,
		Decaf:       false,
		MilkType:    "Oat",
		Sweetener:   "Vanilla",
		ExtraFoam:   true,
	}
	
	makeCustomCoffee(config)
}

func processFullOrder(order Order) {
	fmt.Printf("\n=== ORDER %s ===\n", order.ID)
	fmt.Printf("Customer: %s\n", order.Customer)
	fmt.Printf("Time: %s\n", order.Timestamp.Format("15:04:05"))
	
	fmt.Println("\nItems:")
	for i, item := range order.Items {
		fmt.Printf("  %d × %s\n", order.Quantities[i], item)
	}
	
	if order.Notes != "" {
		fmt.Printf("\nNotes: %s\n", order.Notes)
	}
	
	if order.Paid {
		fmt.Println("\nStatus: PAID ✓")
	} else {
		fmt.Println("\nStatus: AWAITING PAYMENT")
	}
}

func makeCustomCoffee(cfg CoffeeConfig) {
	fmt.Printf("\n=== CUSTOM COFFEE ===\n")
	fmt.Printf("Type: %s %s\n", cfg.Size, cfg.Type)
	
	if cfg.Decaf {
		fmt.Println("✓ Decaf")
	}
	
	fmt.Printf("Shots: %d\n", cfg.Shots)
	fmt.Printf("Temperature: %d°F\n", cfg.Temperature)
	fmt.Printf("Milk: %s\n", cfg.MilkType)
	
	if cfg.Sweetener != "" {
		fmt.Printf("Sweetener: %s\n", cfg.Sweetener)
	}
	
	if cfg.ExtraFoam {
		fmt.Println("✓ Extra foam")
	}
}

// Best practices demonstration
func parameterBestPractices() {
	// Good: Clear parameter names and order
	func(customerID string, orderDate time.Time, items []string, total float64) {
		// Required params first, optional last
	}("CUST-123", time.Now(), []string{"Coffee"}, 4.50)
	
	// Good: Group related parameters
	type Address struct {
		Street, City, State, Zip string
	}
	
	func(name string, addr Address, phone string) {
		// Cleaner than many string parameters
	}("John", Address{}, "555-1234")
	
	// Good: Use constants for options
	const (
		SizeSmall  = "small"
		SizeMedium = "medium"
		SizeLarge  = "large"
	)
	
	func(drink string, size string) {
		// Size parameter is constrained
	}("Latte", SizeLarge)
}