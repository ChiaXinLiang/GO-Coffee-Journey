package main

import "fmt"

func main() {
	fmt.Println("=== GoCoffee Basic Switch ===\n")

	// Example 1: Basic switch with single values
	fmt.Println("Example 1: Coffee Size Selection")
	size := "medium"
	
	switch size {
	case "small":
		fmt.Println("Small: 8oz cup selected")
	case "medium":
		fmt.Println("Medium: 12oz cup selected")
	case "large":
		fmt.Println("Large: 16oz cup selected")
	default:
		fmt.Println("Unknown size")
	}

	// Example 2: Switch with multiple cases
	fmt.Println("\n\nExample 2: Day of Week Specials")
	day := "Tuesday"
	
	switch day {
	case "Monday", "Wednesday", "Friday":
		fmt.Println("☕ Coffee special: 20% off all lattes!")
	case "Tuesday", "Thursday":
		fmt.Println("🥐 Pastry special: Buy 2 get 1 free!")
	case "Saturday", "Sunday":
		fmt.Println("🎉 Weekend special: Free size upgrade!")
	default:
		fmt.Println("No special today")
	}

	// Example 3: Switch without expression
	fmt.Println("\n\nExample 3: Price Range")
	price := 4.50
	
	switch {
	case price < 3:
		fmt.Println("Budget friendly option")
	case price >= 3 && price < 5:
		fmt.Println("Regular price range")
	case price >= 5 && price < 8:
		fmt.Println("Premium selection")
	case price >= 8:
		fmt.Println("Luxury item")
	}

	// Example 4: Switch with initialization
	fmt.Println("\n\nExample 4: Customer Status")
	
	switch customerType := "VIP"; customerType {
	case "Regular":
		fmt.Println("Welcome back!")
	case "VIP":
		fmt.Println("Welcome, valued customer! Enjoy 10% off!")
	case "New":
		fmt.Println("Welcome! First drink is on us!")
	default:
		fmt.Println("Welcome to GoCoffee!")
	}

	// Example 5: Type switch preview
	fmt.Println("\n\nExample 5: Order Type Check")
	processOrder("Latte")
	processOrder(3)
	processOrder(true)
}

// Function demonstrating switch with interface{}
func processOrder(order interface{}) {
	switch v := order.(type) {
	case string:
		fmt.Printf("String order: %s\n", v)
	case int:
		fmt.Printf("Order number: %d\n", v)
	case bool:
		fmt.Printf("Rush order: %v\n", v)
	default:
		fmt.Println("Unknown order type")
	}
}