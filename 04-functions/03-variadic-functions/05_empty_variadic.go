package main

import (
	"fmt"
	"strings"
)

// Variadic functions can be called with zero arguments
func logMessage(prefix string, messages ...string) {
	if len(messages) == 0 {
		fmt.Printf("[%s] No messages to log\n", prefix)
		return
	}
	
	combined := strings.Join(messages, " | ")
	fmt.Printf("[%s] %s\n", prefix, combined)
}

// Default behavior for empty variadic
func serveCoffee(customerName string, additions ...string) {
	fmt.Printf("Serving coffee to %s", customerName)
	
	if len(additions) == 0 {
		fmt.Println(" (black coffee)")
	} else {
		fmt.Printf(" with: %s\n", strings.Join(additions, ", "))
	}
}

// Safe aggregation with empty check
func findMin(numbers ...int) (int, bool) {
	if len(numbers) == 0 {
		return 0, false // No minimum for empty input
	}
	
	min := numbers[0]
	for _, n := range numbers[1:] {
		if n < min {
			min = n
		}
	}
	
	return min, true
}

// Building flexible APIs
func createCoffeeMenu(shopName string, items ...string) []string {
	menu := []string{
		fmt.Sprintf("=== %s Menu ===", shopName),
		"",
		"Standard Offerings:",
		"- House Coffee",
		"- Espresso",
	}
	
	if len(items) > 0 {
		menu = append(menu, "", "Special Items:")
		for _, item := range items {
			menu = append(menu, fmt.Sprintf("- %s", item))
		}
	}
	
	menu = append(menu, "", "================")
	return menu
}

func main() {
	fmt.Println("=== Handling Empty Variadic Calls ===")
	fmt.Println()

	// Calling with zero variadic arguments
	logMessage("INFO", "Server started", "Port 8080")
	logMessage("WARNING") // No additional messages
	fmt.Println()

	// Default behavior example
	serveCoffee("Marcus", "milk", "sugar")
	serveCoffee("Emma") // Just black coffee
	fmt.Println()

	// Safe operations with empty input
	if min, ok := findMin(5, 2, 8, 1, 9); ok {
		fmt.Printf("Minimum value: %d\n", min)
	}
	
	if min, ok := findMin(); ok {
		fmt.Printf("Minimum value: %d\n", min)
	} else {
		fmt.Println("No values provided to find minimum")
	}
	fmt.Println()

	// Building with defaults
	basicMenu := createCoffeeMenu("GoCoffee Express")
	for _, line := range basicMenu {
		fmt.Println(line)
	}
	fmt.Println()

	specialMenu := createCoffeeMenu("GoCoffee Premium",
		"Caramel Macchiato",
		"Vanilla Latte",
		"Mocha Frappuccino",
	)
	for _, line := range specialMenu {
		fmt.Println(line)
	}
	fmt.Println()

	// Practical pattern: optional configuration
	type Option struct {
		Name  string
		Value interface{}
	}
	
	configure := func(base string, opts ...Option) {
		fmt.Printf("Base configuration: %s\n", base)
		if len(opts) == 0 {
			fmt.Println("Using all defaults")
		} else {
			fmt.Println("Custom options:")
			for _, opt := range opts {
				fmt.Printf("  %s = %v\n", opt.Name, opt.Value)
			}
		}
	}
	
	configure("production")
	configure("development",
		Option{"debug", true},
		Option{"port", 3000},
	)
}

// Best practices for empty variadic:
// 1. Always check len() before accessing elements
// 2. Provide sensible defaults for empty input
// 3. Document the behavior for zero arguments
// 4. Consider if a regular slice parameter might be clearer
// 5. Return appropriate zero values or use ok pattern