package main

import "fmt"

func main() {
	fmt.Println("=== GoCoffee Type Safety Demo ===\n")

	// Go prevents type mixing errors
	fmt.Println("TYPE SAFETY EXAMPLES:")

	// Correct: Same types
	espressoShots := 2
	extraShots := 1
	totalShots := espressoShots + extraShots
	fmt.Printf("Total shots: %d\n", totalShots)

	// Would cause compile error:
	// cups := "5"
	// total := cups + 3  // Error: mismatched types string and int

	// Must explicitly convert
	cupsString := "5"
	cupsInt := 5 // Simulating conversion
	totalCups := cupsInt + 3
	fmt.Printf("String '%s' converted to int: %d, total: %d\n",
		cupsString, cupsInt, totalCups)

	// Function type safety
	fmt.Println("\nFUNCTION TYPE SAFETY:")

	// Functions enforce parameter types
	makeCoffee("Latte", 2, true)
	// makeCoffee(2, "Latte", true)  // Error: wrong parameter order

	// Return type safety
	price := calculatePrice(4.50, 3)
	fmt.Printf("Total price: $%.2f\n", price)

	// Custom type safety
	type CoffeeSize string
	type Temperature int

	const (
		Small  CoffeeSize = "small"
		Medium CoffeeSize = "medium"
		Large  CoffeeSize = "large"
	)

	// Using custom types
	orderSize := Large
	brewTemp := Temperature(195)

	fmt.Printf("\nCustom types: Size=%s, Temp=%d°F\n", orderSize, brewTemp)

	// Type aliases for clarity
	type Cents int
	type Quantity int

	itemPrice := Cents(450)
	itemCount := Quantity(2)

	// Still need conversion between custom types
	total := int(itemPrice) * int(itemCount)
	fmt.Printf("Price: %d cents × Quantity: %d = %d cents\n",
		itemPrice, itemCount, total)
}

func makeCoffee(coffeeType string, shots int, addMilk bool) {
	fmt.Printf("Making %s with %d shots", coffeeType, shots)
	if addMilk {
		fmt.Print(" and milk")
	}
	fmt.Println()
}

func calculatePrice(basePrice float64, quantity int) float64 {
	return basePrice * float64(quantity)
}