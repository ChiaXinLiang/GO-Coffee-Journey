package main

import (
	"fmt"
	"strings"
)

// Variadic parameter must be last
func createOrder(customerName string, items ...string) string {
	order := fmt.Sprintf("Order for %s:\n", customerName)
	
	if len(items) == 0 {
		order += "  No items ordered\n"
	} else {
		for i, item := range items {
			order += fmt.Sprintf("  %d. %s\n", i+1, item)
		}
	}
	
	return order
}

// Multiple regular parameters before variadic
func generateReceipt(customerName string, paymentMethod string, items ...float64) {
	fmt.Printf("=== Receipt for %s ===\n", customerName)
	fmt.Printf("Payment Method: %s\n", paymentMethod)
	fmt.Println("Items:")
	
	var total float64
	for i, price := range items {
		fmt.Printf("  Item %d: $%.2f\n", i+1, price)
		total += price
	}
	
	fmt.Printf("Total: $%.2f\n", total)
	fmt.Println("==================")
}

// Complex example with struct
type OrderOption struct {
	Name  string
	Price float64
}

func processComplexOrder(customerID string, basePrice float64, options ...OrderOption) {
	fmt.Printf("Processing order for customer: %s\n", customerID)
	fmt.Printf("Base price: $%.2f\n", basePrice)
	
	total := basePrice
	
	if len(options) > 0 {
		fmt.Println("Add-ons:")
		for _, opt := range options {
			fmt.Printf("  + %s: $%.2f\n", opt.Name, opt.Price)
			total += opt.Price
		}
	}
	
	fmt.Printf("Final total: $%.2f\n", total)
}

// Error: Variadic parameter must be last
// func invalid(items ...string, customerName string) {} // This won't compile!

// Error: Only one variadic parameter allowed
// func alsoInvalid(drinks ...string, foods ...string) {} // This won't compile!

func main() {
	fmt.Println("=== Mixing Regular and Variadic Parameters ===")
	fmt.Println()

	// Regular parameter followed by variadic
	order1 := createOrder("Marcus", "Cappuccino", "Croissant", "Muffin")
	fmt.Print(order1)
	fmt.Println()

	// Can omit variadic arguments
	order2 := createOrder("Emma")
	fmt.Print(order2)
	fmt.Println()

	// Multiple regular parameters
	generateReceipt("Sarah", "Credit Card", 3.50, 2.75, 4.25)
	fmt.Println()

	// With struct types
	fmt.Println("Complex order:")
	processComplexOrder("CUST-123", 5.00,
		OrderOption{"Extra Shot", 0.50},
		OrderOption{"Oat Milk", 0.75},
		OrderOption{"Caramel Syrup", 0.50},
	)
	fmt.Println()

	// Building the variadic arguments dynamically
	extras := []OrderOption{
		{"Whipped Cream", 0.50},
		{"Chocolate Drizzle", 0.75},
	}
	
	// Note: we'll see how to pass slices to variadic functions in the next example
	fmt.Println("Dynamic order:")
	processComplexOrder("CUST-456", 4.50, extras[0], extras[1])
}

// Rules for mixing parameters:
// 1. Variadic parameter MUST be the last parameter
// 2. Only ONE variadic parameter per function
// 3. Regular parameters get their arguments first
// 4. Remaining arguments go to the variadic parameter