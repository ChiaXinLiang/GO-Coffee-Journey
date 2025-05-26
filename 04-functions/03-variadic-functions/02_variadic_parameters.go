package main

import "fmt"

// Inside a variadic function, the parameter is a slice
func analyzeOrders(orderIDs ...string) {
	// orderIDs is of type []string
	fmt.Printf("Type of variadic parameter: %T\n", orderIDs)
	fmt.Printf("Number of orders: %d\n", len(orderIDs))
	fmt.Printf("Capacity of slice: %d\n", cap(orderIDs))
	
	// Can use all slice operations
	if len(orderIDs) > 0 {
		fmt.Printf("First order: %s\n", orderIDs[0])
		fmt.Printf("Last order: %s\n", orderIDs[len(orderIDs)-1])
	}
	
	// Can modify the slice (but changes don't affect caller)
	for i := range orderIDs {
		orderIDs[i] = "PROCESSED-" + orderIDs[i]
	}
	
	fmt.Println("Processed orders:", orderIDs)
}

// Demonstrating that variadic creates a new slice
func modifyPrices(prices ...float64) {
	fmt.Println("Original prices in function:", prices)
	
	// Add tax to all prices
	for i := range prices {
		prices[i] = prices[i] * 1.08 // 8% tax
	}
	
	fmt.Println("Modified prices in function:", prices)
}

// Can use slice operations
func findHighestPrice(prices ...float64) (float64, bool) {
	if len(prices) == 0 {
		return 0, false
	}
	
	highest := prices[0]
	for _, price := range prices[1:] { // Skip first element
		if price > highest {
			highest = price
		}
	}
	
	return highest, true
}

func main() {
	fmt.Println("=== Working with Variadic Parameters ===")
	fmt.Println()

	// The variadic parameter is a regular slice inside the function
	fmt.Println("Order analysis:")
	analyzeOrders("ORD001", "ORD002", "ORD003")
	fmt.Println()

	// Original values are not affected by modifications
	coffeePrices := []float64{3.50, 4.25, 5.00}
	fmt.Println("Original prices before call:", coffeePrices)
	
	// Note: we're passing individual values, not the slice
	modifyPrices(3.50, 4.25, 5.00)
	fmt.Println("Original prices after call:", coffeePrices)
	fmt.Println()

	// Using slice operations
	highest, found := findHighestPrice(3.50, 4.25, 5.00, 3.75)
	if found {
		fmt.Printf("Highest price: $%.2f\n", highest)
	}
	
	// Empty call
	highest2, found2 := findHighestPrice()
	fmt.Printf("Highest from empty: $%.2f, found: %v\n", highest2, found2)
}

// Key insights:
// - Variadic parameters are slices inside the function
// - A new slice is created for each call
// - Modifications don't affect the original arguments
// - Can use all slice operations (len, cap, range, indexing, etc.)