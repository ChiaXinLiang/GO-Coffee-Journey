package main

import "fmt"

// calculateTotal accepts any number of prices and returns the sum
func calculateTotal(prices ...float64) float64 {
	var total float64
	
	// Inside the function, prices is a slice of float64
	fmt.Printf("Received %d items\n", len(prices))
	
	for i, price := range prices {
		fmt.Printf("  Item %d: $%.2f\n", i+1, price)
		total += price
	}
	
	return total
}

// greetCustomers can greet any number of customers
func greetCustomers(greeting string, names ...string) {
	for _, name := range names {
		fmt.Printf("%s, %s!\n", greeting, name)
	}
}

func main() {
	fmt.Println("=== Basic Variadic Functions at GoCoffee ===")
	fmt.Println()

	// Calling with different numbers of arguments
	fmt.Println("Single item order:")
	total1 := calculateTotal(3.50)
	fmt.Printf("Total: $%.2f\n\n", total1)

	fmt.Println("Two item order:")
	total2 := calculateTotal(3.50, 4.25)
	fmt.Printf("Total: $%.2f\n\n", total2)

	fmt.Println("Group order:")
	total3 := calculateTotal(3.50, 4.25, 5.00, 3.75, 4.50)
	fmt.Printf("Total: $%.2f\n\n", total3)

	// Can even call with zero arguments
	fmt.Println("Empty order:")
	total4 := calculateTotal()
	fmt.Printf("Total: $%.2f\n\n", total4)

	// Using variadic function with mixed parameters
	fmt.Println("Morning greetings:")
	greetCustomers("Good morning", "Sarah", "John", "Emma")
	fmt.Println()

	fmt.Println("Afternoon greeting:")
	greetCustomers("Good afternoon", "Marcus")
}

// Output shows:
// - Variadic parameters become slices inside the function
// - Can be called with any number of arguments (including zero)
// - The ... syntax indicates a variadic parameter
// - Regular parameters can come before variadic ones