package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== GoCoffee Break Statement Basics ===\n")

	// Example 1: Simple break on condition
	fmt.Println("☕ Processing Morning Orders")
	fmt.Println("---------------------------")
	
	coffeeBeans := 200 // grams
	orderCount := 0
	
	for order := 1; order <= 20; order++ {
		beansNeeded := 18 // grams per cup
		
		fmt.Printf("Order #%d: Checking inventory... ", order)
		
		if coffeeBeans < beansNeeded {
			fmt.Println("\n❌ BREAK: Insufficient coffee beans!")
			fmt.Printf("   (Need %dg, only have %dg)\n", beansNeeded, coffeeBeans)
			break
		}
		
		coffeeBeans -= beansNeeded
		orderCount++
		fmt.Printf("✓ Prepared (%dg beans remaining)\n", coffeeBeans)
		
		time.Sleep(300 * time.Millisecond)
	}
	
	fmt.Printf("\nTotal orders completed: %d\n", orderCount)

	// Example 2: Break with user action simulation
	fmt.Println("\n\n🔔 Customer Service (Press 'Q' simulation)")
	fmt.Println("------------------------------------------")
	
	// Simulate user pressing 'Q' after 5 customers
	customerCount := 0
	maxBeforeQuit := 5
	
	for {
		customerCount++
		fmt.Printf("Serving customer #%d... ", customerCount)
		
		// Simulate service time
		time.Sleep(500 * time.Millisecond)
		fmt.Println("✓")
		
		// Simulate checking for 'Q' press
		if customerCount >= maxBeforeQuit {
			fmt.Println("\n[Manager pressed 'Q' - Emergency meeting!]")
			break
		}
	}
	
	fmt.Printf("Served %d customers before break\n", customerCount)

	// Example 3: Break in switch statement (common mistake)
	fmt.Println("\n\n⚠️  Break in Switch vs Loop")
	fmt.Println("---------------------------")
	
	for i := 1; i <= 5; i++ {
		fmt.Printf("Iteration %d: ", i)
		
		switch i {
		case 3:
			fmt.Print("Special case! ")
			break // This breaks the switch, NOT the loop!
		default:
			fmt.Print("Normal case. ")
		}
		
		fmt.Println("Loop continues...")
	}
	
	fmt.Println("\nNotice: 'break' in switch didn't exit the loop!")

	// Example 4: Breaking from infinite loop
	fmt.Println("\n\n♾️  Coffee Machine Status Monitor")
	fmt.Println("--------------------------------")
	
	temperature := 85.0
	targetTemp := 93.0
	heating := true
	
	for {
		if heating {
			temperature += 1.5
			fmt.Printf("🔥 Heating: %.1f°C", temperature)
		}
		
		if temperature >= targetTemp {
			fmt.Println(" → ✓ Target reached!")
			break
		}
		
		fmt.Println(" → Still heating...")
		time.Sleep(400 * time.Millisecond)
	}
	
	fmt.Printf("Final temperature: %.1f°C\n", temperature)

	// Example 5: Multiple break conditions
	fmt.Println("\n\n🏪 Shop Operations Monitor")
	fmt.Println("-------------------------")
	
	hour := 6 // Start at 6 AM
	sales := 0.0
	targetSales := 1000.0
	
	for {
		// Simulate hourly sales
		hourlySales := float64((hour * 20) + 50)
		sales += hourlySales
		
		fmt.Printf("%02d:00 - Sales: $%.2f (Total: $%.2f)", hour, hourlySales, sales)
		
		// Check multiple break conditions
		if hour >= 22 {
			fmt.Println(" → 🌙 Closing time!")
			break
		}
		
		if sales >= targetSales {
			fmt.Println(" → 🎯 Sales target reached!")
			break
		}
		
		fmt.Println()
		hour++
		time.Sleep(300 * time.Millisecond)
	}
	
	fmt.Printf("\nFinal: %02d:00, Total sales: $%.2f\n", hour, sales)
}