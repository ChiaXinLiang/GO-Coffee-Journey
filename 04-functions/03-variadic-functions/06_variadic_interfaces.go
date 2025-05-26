package main

import (
	"fmt"
	"reflect"
)

// Using interface{} (or 'any' in Go 1.18+) for mixed types
func printAny(values ...interface{}) {
	for i, v := range values {
		fmt.Printf("%d: %v (type: %T)\n", i, v, v)
	}
}

// Type-safe logging function
func logEvent(event string, details ...interface{}) {
	fmt.Printf("[EVENT] %s", event)
	
	if len(details) > 0 {
		fmt.Print(" - Details: ")
		for i, detail := range details {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Printf("%v", detail)
		}
	}
	fmt.Println()
}

// Building formatted strings (like fmt.Sprintf)
func formatMessage(format string, args ...interface{}) string {
	// Simple implementation for demonstration
	result := format
	
	for i, arg := range args {
		placeholder := fmt.Sprintf("{%d}", i)
		value := fmt.Sprintf("%v", arg)
		result = strings.ReplaceAll(result, placeholder, value)
	}
	
	return result
}

// Type checking in variadic interface functions
func sumNumbers(values ...interface{}) (float64, error) {
	var sum float64
	
	for i, v := range values {
		switch num := v.(type) {
		case int:
			sum += float64(num)
		case float64:
			sum += num
		case float32:
			sum += float64(num)
		default:
			return 0, fmt.Errorf("argument %d is not a number: %v (type %T)", i, v, v)
		}
	}
	
	return sum, nil
}

// More complex example: building a coffee order
func buildOrder(customerName string, items ...interface{}) {
	fmt.Printf("Building order for %s:\n", customerName)
	
	for _, item := range items {
		switch v := item.(type) {
		case string:
			fmt.Printf("  - Item: %s\n", v)
		case float64:
			fmt.Printf("  - Price adjustment: $%.2f\n", v)
		case int:
			fmt.Printf("  - Quantity: %d\n", v)
		case map[string]interface{}:
			fmt.Printf("  - Options: %v\n", v)
		default:
			fmt.Printf("  - Unknown item type: %v\n", v)
		}
	}
}

// Using reflection for more advanced type handling
func debugPrint(label string, values ...interface{}) {
	fmt.Printf("=== %s ===\n", label)
	
	for i, v := range values {
		rv := reflect.ValueOf(v)
		rt := reflect.TypeOf(v)
		
		fmt.Printf("[%d] Type: %v, Kind: %v, Value: %v\n", 
			i, rt, rv.Kind(), v)
	}
	fmt.Println()
}

func main() {
	fmt.Println("=== Variadic Functions with Interfaces ===")
	fmt.Println()

	// Mixed types in one call
	printAny("Hello", 42, 3.14, true, []string{"a", "b"})
	fmt.Println()

	// Logging with mixed types
	logEvent("OrderPlaced", "OrderID", "ORD-123", "Amount", 25.50)
	logEvent("SystemStartup") // No additional details
	fmt.Println()

	// Custom formatting
	msg := formatMessage("Customer {0} ordered {1} items for ${2}", 
		"Marcus", 3, 15.75)
	fmt.Println("Formatted:", msg)
	fmt.Println()

	// Type-safe operations
	if total, err := sumNumbers(10, 20.5, 30); err == nil {
		fmt.Printf("Sum of numbers: %.2f\n", total)
	}
	
	if _, err := sumNumbers(10, "20", 30); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Println()

	// Complex order building
	buildOrder("Emma",
		"Latte",
		2,                                    // quantity
		5.50,                                // price
		map[string]interface{}{              // options
			"size": "large",
			"milk": "oat",
		},
		"Croissant",
		1,
		3.75,
	)
	fmt.Println()

	// Debug printing with reflection
	debugPrint("Coffee Shop Data",
		"GoCoffee",
		42,
		3.14159,
		[]string{"Espresso", "Latte"},
		map[string]int{"tables": 10, "chairs": 40},
	)
	
	// Type aliases for clarity
	type OrderItem interface{}
	
	processItems := func(items ...OrderItem) {
		fmt.Printf("Processing %d items\n", len(items))
		// Process items...
	}
	
	processItems("Coffee", 2.50, map[string]string{"size": "medium"})
}

// Important considerations:
// 1. Type safety is lost with interface{}
// 2. Runtime type checking/assertion may be needed
// 3. Performance overhead from boxing/unboxing
// 4. Consider if generics (Go 1.18+) might be better
// 5. Document expected types clearly