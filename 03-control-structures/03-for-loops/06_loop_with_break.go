package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("=== GoCoffee Break Statement Examples ===\n")
	
	rand.Seed(time.Now().UnixNano())

	// Example 1: Breaking when stock runs out
	fmt.Println("☕ Morning Coffee Service")
	fmt.Println("------------------------")
	
	coffeeBeans := 100 // grams
	cupsServed := 0
	
	for i := 1; i <= 20; i++ {
		if coffeeBeans < 18 {
			fmt.Println("\n❌ OUT OF COFFEE BEANS! Cannot serve more customers.")
			break
		}
		
		coffeeBeans -= 18
		cupsServed++
		fmt.Printf("Served cup #%d (Beans left: %dg)\n", i, coffeeBeans)
	}
	
	fmt.Printf("\nTotal cups served: %d\n", cupsServed)

	// Example 2: Breaking on error condition
	fmt.Println("\n\n💳 Payment Processing")
	fmt.Println("--------------------")
	
	for attempt := 1; attempt <= 5; attempt++ {
		fmt.Printf("Processing payment attempt #%d...", attempt)
		
		// Simulate payment processing (80% success rate)
		success := rand.Float32() < 0.8
		
		if success {
			fmt.Println(" ✓ Success!")
			break
		} else {
			fmt.Println(" ✗ Failed")
			if attempt == 5 {
				fmt.Println("❌ Maximum attempts reached. Transaction failed.")
			}
		}
		
		time.Sleep(500 * time.Millisecond)
	}

	// Example 3: Search and break
	fmt.Println("\n\n🔍 Finding Customer Order")
	fmt.Println("-------------------------")
	
	orders := []struct {
		id       int
		customer string
		drink    string
	}{
		{101, "Alice", "Latte"},
		{102, "Bob", "Espresso"},
		{103, "Charlie", "Cappuccino"},
		{104, "Diana", "Americano"},
		{105, "Eve", "Mocha"},
	}
	
	targetID := 103
	found := false
	
	for _, order := range orders {
		fmt.Printf("Checking order #%d...", order.id)
		
		if order.id == targetID {
			fmt.Printf(" ✓ Found!\n")
			fmt.Printf("Customer: %s, Drink: %s\n", order.customer, order.drink)
			found = true
			break
		}
		fmt.Println(" not this one")
	}
	
	if !found {
		fmt.Println("Order not found!")
	}

	// Example 4: Breaking from nested loops with labels
	fmt.Println("\n\n📊 Finding Available Table")
	fmt.Println("--------------------------")
	
	// 3 floors, 5 tables each
	occupied := [3][5]bool{
		{true, true, false, true, true},
		{true, true, true, true, true},
		{false, true, true, false, true},
	}
	
	foundTable := false
	
FloorLoop:
	for floor := 0; floor < 3; floor++ {
		fmt.Printf("Checking floor %d:\n", floor+1)
		
		for table := 0; table < 5; table++ {
			if !occupied[floor][table] {
				fmt.Printf("  ✓ Table %d is available!\n", table+1)
				foundTable = true
				break FloorLoop // Break from both loops
			}
			fmt.Printf("  Table %d: occupied\n", table+1)
		}
	}
	
	if !foundTable {
		fmt.Println("❌ No tables available!")
	}

	// Example 5: Break with user input simulation
	fmt.Println("\n\n🛒 Order Taking System")
	fmt.Println("---------------------")
	
	orderItems := []string{}
	maxItems := 5
	
	// Simulate user adding items
	availableItems := []string{"Latte", "Cookie", "Muffin", "Water", "DONE"}
	
	for {
		// Simulate random item selection
		selectedIndex := rand.Intn(len(availableItems))
		selected := availableItems[selectedIndex]
		
		if selected == "DONE" {
			fmt.Println("Customer finished ordering")
			break
		}
		
		orderItems = append(orderItems, selected)
		fmt.Printf("Added: %s\n", selected)
		
		if len(orderItems) >= maxItems {
			fmt.Println("Maximum items reached!")
			break
		}
		
		time.Sleep(300 * time.Millisecond)
	}
	
	fmt.Printf("\nFinal order: %v\n", orderItems)

	// Example 6: Breaking infinite loop
	fmt.Println("\n\n⏰ Coffee Shop Hours")
	fmt.Println("-------------------")
	
	hour := 6 // Start at 6 AM
	customersServed := 0
	
	for {
		fmt.Printf("%02d:00 - ", hour)
		
		// Random customers per hour
		customers := rand.Intn(10) + 5
		customersServed += customers
		fmt.Printf("Served %d customers\n", customers)
		
		hour++
		
		// Break at closing time
		if hour >= 22 { // 10 PM
			fmt.Println("\n🌙 Closing time!")
			break
		}
		
		// Break if we've served enough customers
		if customersServed >= 100 {
			fmt.Printf("\n🎉 Served %d customers today! Great job!\n", customersServed)
			break
		}
		
		time.Sleep(200 * time.Millisecond)
	}
	
	fmt.Printf("\nShop closed at %02d:00 after serving %d customers\n", hour, customersServed)
}