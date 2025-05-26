package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("=== GoCoffee Customer Queue System ===\n")
	
	rand.Seed(time.Now().UnixNano())

	// While-style loop for customer queue
	customersInQueue := rand.Intn(8) + 3 // 3-10 customers
	fmt.Printf("Starting with %d customers in queue\n\n", customersInQueue)

	for customersInQueue > 0 {
		fmt.Printf("👥 Queue length: %d\n", customersInQueue)
		fmt.Println("Serving next customer...")
		
		// Simulate order processing
		orderTime := rand.Intn(3) + 1
		fmt.Printf("   Processing order (%d seconds)...\n", orderTime)
		time.Sleep(time.Duration(orderTime) * time.Second)
		
		customersInQueue--
		
		// Random chance of new customer joining
		if rand.Float32() < 0.3 && customersInQueue < 10 {
			customersInQueue++
			fmt.Println("   🔔 New customer joined the queue!")
		}
		
		fmt.Println("   ✓ Order complete!\n")
	}

	fmt.Println("--- No more customers in queue! ---")

	// Another while-style pattern
	fmt.Println("\n=== Inventory Check ===")
	
	coffeeBeans := 1000.0 // grams
	cupsPerDay := 0
	
	for coffeeBeans >= 18.0 { // Need at least 18g for one espresso
		coffeeBeans -= 18.0
		cupsPerDay++
	}
	
	fmt.Printf("With 1kg of beans, we can make %d cups\n", cupsPerDay)
	fmt.Printf("Leftover beans: %.1fg\n", coffeeBeans)

	// While loop with complex condition
	fmt.Println("\n=== Shop Hours ===")
	
	hour := 6 // Start at 6 AM
	isOpen := true
	
	for isOpen {
		fmt.Printf("🕐 %d:00 - Shop is OPEN\n", hour)
		
		// Check if we should close
		if hour >= 18 { // Close at 6 PM
			isOpen = false
		} else {
			hour++
			time.Sleep(500 * time.Millisecond)
		}
	}
	
	fmt.Println("\n🌙 Shop is now CLOSED. See you tomorrow!")
}