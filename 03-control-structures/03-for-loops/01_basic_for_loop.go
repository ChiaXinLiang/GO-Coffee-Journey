package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== GoCoffee Morning Brewing Routine ===\n")

	// Traditional for loop
	fmt.Println("Starting morning coffee preparation...")
	for i := 1; i <= 5; i++ {
		fmt.Printf("☕ Brewing coffee batch #%d\n", i)
		time.Sleep(500 * time.Millisecond) // Simulate brewing time
	}

	fmt.Println("\n--- Preparing individual cups ---")
	
	// For loop with different increment
	for cup := 0; cup < 10; cup += 2 {
		fmt.Printf("Preparing cups %d and %d\n", cup+1, cup+2)
	}

	fmt.Println("\n--- Quality check countdown ---")
	
	// Countdown loop
	for count := 5; count > 0; count-- {
		fmt.Printf("Quality check in %d...\n", count)
		time.Sleep(300 * time.Millisecond)
	}
	fmt.Println("✓ Quality check complete!")

	// Loop with multiple variables
	fmt.Println("\n--- Tracking beans and water ---")
	for beans, water := 100, 1000; beans >= 20 && water >= 200; beans -= 20, water -= 200 {
		fmt.Printf("Resources: %dg beans, %dml water - Can make %d more cups\n", 
			beans, water, beans/20)
	}

	fmt.Println("\n✓ Morning brewing routine complete!")
}