package main

import "fmt"

func main() {
	// Method 1: var with type
	var espressoShots int = 2

	// Method 2: var with type inference
	var milkFoam = true

	// Method 3: short declaration (inside functions only)
	sugarPackets := 3

	// Method 4: multiple variables
	var (
		customerName = "Marcus"
		orderNumber  = 1001
		isPriority   = false
	)

	// Method 5: multiple variables same line
	temp, size := 165, "large"

	fmt.Println("=== Order Details ===")
	fmt.Printf("Customer: %s (Order #%d)\n", customerName, orderNumber)
	fmt.Printf("Espresso shots: %d\n", espressoShots)
	fmt.Printf("Milk foam: %v\n", milkFoam)
	fmt.Printf("Sugar packets: %d\n", sugarPackets)
	fmt.Printf("Temperature: %d°F, Size: %s\n", temp, size)
	fmt.Printf("Priority order: %v\n", isPriority)
}