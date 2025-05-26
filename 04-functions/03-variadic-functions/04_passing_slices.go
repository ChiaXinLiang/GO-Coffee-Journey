package main

import "fmt"

// Our variadic function for processing coffee orders
func processCoffeeOrders(orders ...string) {
	fmt.Printf("Processing %d orders:\n", len(orders))
	for i, order := range orders {
		fmt.Printf("  %d. %s\n", i+1, order)
	}
}

// Calculate total with variadic
func calculateTotal(prices ...float64) float64 {
	var total float64
	for _, price := range prices {
		total += price
	}
	return total
}

// Combine multiple coffee blends
func createBlend(name string, coffees ...string) {
	fmt.Printf("Creating '%s' blend with:\n", name)
	for _, coffee := range coffees {
		fmt.Printf("  - %s\n", coffee)
	}
}

func main() {
	fmt.Println("=== Passing Slices to Variadic Functions ===")
	fmt.Println()

	// Sometimes we have a slice that we want to pass to a variadic function
	morningOrders := []string{"Espresso", "Latte", "Cappuccino"}
	afternoonOrders := []string{"Iced Coffee", "Cold Brew"}

	// Wrong way - this passes the slice as a single argument
	fmt.Println("WRONG: Passing slice directly")
	// processCoffeeOrders(morningOrders) // This would cause a type error!
	
	// Right way - use ... to expand the slice
	fmt.Println("RIGHT: Using ... to expand slice")
	processCoffeeOrders(morningOrders...)
	fmt.Println()

	// Can combine individual values and expanded slices
	fmt.Println("Morning and afternoon orders combined:")
	processCoffeeOrders("Early Bird Special", "Regular Drip")
	processCoffeeOrders(morningOrders...)
	processCoffeeOrders(afternoonOrders...)
	fmt.Println()

	// Working with numeric slices
	groupPrices := []float64{3.50, 4.25, 5.00, 3.75}
	total := calculateTotal(groupPrices...)
	fmt.Printf("Group order total: $%.2f\n\n", total)

	// Combining multiple slices
	blend1 := []string{"Colombian", "Brazilian"}
	blend2 := []string{"Ethiopian", "Kenyan"}
	
	// Can't combine slices directly in the call
	// createBlend("World Tour", blend1..., blend2...) // This won't work!
	
	// Must combine slices first
	allBeans := append(blend1, blend2...)
	createBlend("World Tour", allBeans...)
	fmt.Println()

	// Empty slice expands to zero arguments
	emptyOrders := []string{}
	fmt.Println("Empty slice expansion:")
	processCoffeeOrders(emptyOrders...)
	fmt.Println()

	// Practical example: filtering and passing
	allOrders := []string{"Latte", "Espresso", "Tea", "Cappuccino", "Hot Chocolate"}
	
	// Filter only coffee orders
	var coffeeOnly []string
	for _, order := range allOrders {
		if order != "Tea" && order != "Hot Chocolate" {
			coffeeOnly = append(coffeeOnly, order)
		}
	}
	
	fmt.Println("Filtered coffee orders:")
	processCoffeeOrders(coffeeOnly...)
}

// Key points:
// - Use ... after a slice to expand it into individual arguments
// - The slice type must match the variadic parameter type
// - Can't mix individual values and slice expansion in the same position
// - Can't expand multiple slices in the same call
// - Empty slices expand to zero arguments