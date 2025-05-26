package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("=== GoCoffee Infinite Loop Examples ===\n")
	
	rand.Seed(time.Now().UnixNano())

	// Example 1: Coffee shop service loop
	fmt.Println("☕ Coffee Shop Service System")
	fmt.Println("-----------------------------")
	fmt.Println("(Simulating 10 customer interactions)\n")
	
	customerCount := 0
	
	for {
		customerCount++
		
		// Simulate customer arrival
		fmt.Printf("🔔 Customer #%d arrived\n", customerCount)
		
		// Random order
		orders := []string{"Latte", "Espresso", "Cappuccino", "Americano"}
		order := orders[rand.Intn(len(orders))]
		
		fmt.Printf("   Order: %s\n", order)
		fmt.Println("   Processing...")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("   ✓ Order complete!\n")
		
		// Exit condition
		if customerCount >= 10 {
			fmt.Println("Demo complete - 10 customers served!")
			break
		}
		
		// Random delay between customers
		time.Sleep(time.Duration(rand.Intn(500)+300) * time.Millisecond)
	}

	// Example 2: Event monitoring loop
	fmt.Println("\n\n🔍 Coffee Machine Monitoring")
	fmt.Println("----------------------------")
	
	machineTemp := 90.0 // Starting temperature
	optimalTemp := 93.0
	iterations := 0
	
	for {
		iterations++
		
		// Simulate temperature changes
		change := (rand.Float64() - 0.3) * 2 // -0.6 to +1.4
		machineTemp += change
		
		fmt.Printf("Check #%d: Temperature: %.1f°C", iterations, machineTemp)
		
		if machineTemp < 88.0 {
			fmt.Println(" ⚠️  TOO COLD! Heating...")
			machineTemp += 3.0
		} else if machineTemp > 96.0 {
			fmt.Println(" ⚠️  TOO HOT! Cooling...")
			machineTemp -= 3.0
		} else if machineTemp >= 92.0 && machineTemp <= 94.0 {
			fmt.Println(" ✓ Perfect temperature!")
			if iterations >= 10 {
				fmt.Println("\nMachine stabilized at optimal temperature!")
				break
			}
		} else {
			fmt.Println(" ~ Acceptable")
		}
		
		time.Sleep(400 * time.Millisecond)
		
		// Safety break
		if iterations > 20 {
			fmt.Println("\nMaximum monitoring cycles reached!")
			break
		}
	}

	// Example 3: User input simulation (menu system)
	fmt.Println("\n\n📱 Interactive Coffee Shop Menu")
	fmt.Println("-------------------------------")
	
	// Simulate user choices
	userChoices := []string{"1", "2", "3", "1", "4", "5"}
	choiceIndex := 0
	
	for {
		fmt.Println("\n--- MAIN MENU ---")
		fmt.Println("1. Order Coffee")
		fmt.Println("2. View Menu")
		fmt.Println("3. Check Balance")
		fmt.Println("4. Leave Review")
		fmt.Println("5. Exit")
		
		// Simulate user input
		if choiceIndex >= len(userChoices) {
			fmt.Println("\nDemo ended - no more simulated inputs")
			break
		}
		
		choice := userChoices[choiceIndex]
		choiceIndex++
		
		fmt.Printf("\nYour choice: %s\n", choice)
		time.Sleep(500 * time.Millisecond)
		
		switch choice {
		case "1":
			fmt.Println("→ Opening order menu...")
		case "2":
			fmt.Println("→ Displaying full menu...")
		case "3":
			fmt.Println("→ Your balance: $25.50")
		case "4":
			fmt.Println("→ Thank you for your feedback!")
		case "5":
			fmt.Println("→ Thank you for visiting GoCoffee!")
			break // This only breaks the switch, not the loop
		default:
			fmt.Println("→ Invalid choice, please try again")
			continue
		}
		
		// Check if user chose to exit
		if choice == "5" {
			break // This breaks the for loop
		}
		
		time.Sleep(300 * time.Millisecond)
	}

	// Example 4: Retry mechanism
	fmt.Println("\n\n🔄 Connection Retry System")
	fmt.Println("--------------------------")
	
	maxRetries := 5
	retryCount := 0
	
	for {
		retryCount++
		fmt.Printf("Attempt #%d: Connecting to payment system...", retryCount)
		
		// Simulate connection attempt (70% success rate)
		connected := rand.Float32() < 0.7
		
		time.Sleep(500 * time.Millisecond)
		
		if connected {
			fmt.Println(" ✓ Connected!")
			break
		}
		
		fmt.Println(" ✗ Failed")
		
		if retryCount >= maxRetries {
			fmt.Println("\n❌ Could not establish connection after 5 attempts")
			break
		}
		
		fmt.Printf("   Retrying in %d seconds...\n", 3-retryCount/2)
		time.Sleep(time.Duration(3-retryCount/2) * time.Second)
	}

	// Example 5: State machine
	fmt.Println("\n\n🎮 Coffee Making State Machine")
	fmt.Println("------------------------------")
	
	state := "IDLE"
	stateCount := 0
	
	for {
		stateCount++
		fmt.Printf("State %d: %s", stateCount, state)
		
		switch state {
		case "IDLE":
			fmt.Println(" - Waiting for order...")
			state = "GRINDING"
			
		case "GRINDING":
			fmt.Println(" - Grinding beans...")
			state = "BREWING"
			
		case "BREWING":
			fmt.Println(" - Brewing coffee...")
			state = "POURING"
			
		case "POURING":
			fmt.Println(" - Pouring into cup...")
			state = "READY"
			
		case "READY":
			fmt.Println(" - ☕ Coffee ready!")
			state = "COMPLETE"
			
		case "COMPLETE":
			fmt.Println(" - Process complete!")
			break
		}
		
		if state == "COMPLETE" {
			break
		}
		
		time.Sleep(600 * time.Millisecond)
		
		// Safety check
		if stateCount > 10 {
			fmt.Println("\n⚠️  State machine error - too many states!")
			break
		}
	}

	fmt.Println("\n\n✓ All infinite loop examples completed!")
}