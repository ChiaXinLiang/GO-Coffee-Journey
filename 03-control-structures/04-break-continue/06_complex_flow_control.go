package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("=== GoCoffee Complex Flow Control ===\n")
	
	rand.Seed(time.Now().UnixNano())

	// Example 1: Order processing pipeline
	fmt.Println("☕ Order Processing Pipeline")
	fmt.Println("---------------------------")
	
	type Order struct {
		id         string
		items      []string
		priority   string
		paid       bool
		cupSize    string
		allergies  []string
	}
	
	orders := []Order{
		{"ORD001", []string{"Latte", "Cookie"}, "normal", true, "large", nil},
		{"ORD002", []string{"Espresso"}, "VIP", false, "small", nil},
		{"ORD003", []string{"Cappuccino"}, "normal", true, "medium", []string{"lactose"}},
		{"ORD004", []string{"Mocha", "Muffin"}, "VIP", true, "large", nil},
		{"ORD005", []string{"Decaf Latte"}, "normal", true, "medium", nil},
		{"ORD006", []string{"Americano"}, "rush", true, "invalid", nil},
	}
	
	// Process orders with complex rules
	for _, order := range orders {
		fmt.Printf("\nProcessing %s (%s priority):\n", order.id, order.priority)
		
		// Step 1: Payment verification
		if !order.paid {
			fmt.Println("  ❌ Payment pending - skipping order")
			continue
		}
		fmt.Println("  ✓ Payment verified")
		
		// Step 2: Priority handling
		if order.priority == "VIP" {
			fmt.Println("  ⭐ VIP order - expedited processing")
		} else if order.priority == "rush" {
			if rand.Float32() > 0.7 { // 30% chance to handle rush
				fmt.Println("  ⚡ Rush order accepted")
			} else {
				fmt.Println("  ❌ Cannot handle rush order now")
				continue
			}
		}
		
		// Step 3: Validate cup size
		validSizes := map[string]bool{"small": true, "medium": true, "large": true}
		if !validSizes[order.cupSize] {
			fmt.Printf("  ❌ Invalid cup size: %s - canceling order\n", order.cupSize)
			break // Critical error - stop processing
		}
		
		// Step 4: Check each item
		itemsProcessed := 0
		for _, item := range order.items {
			fmt.Printf("  Preparing %s... ", item)
			
			// Skip unavailable items
			if item == "Decaf Latte" {
				fmt.Println("unavailable")
				continue
			}
			
			// Check allergies
			if len(order.allergies) > 0 {
				if item == "Latte" || item == "Cappuccino" {
					fmt.Printf("contains milk (customer has allergies: %v) - skipped\n", order.allergies)
					continue
				}
			}
			
			itemsProcessed++
			fmt.Println("✓")
			time.Sleep(300 * time.Millisecond)
		}
		
		if itemsProcessed == 0 {
			fmt.Println("  ❌ No items could be prepared")
			continue
		}
		
		fmt.Printf("  ✅ Order complete (%d/%d items)\n", itemsProcessed, len(order.items))
	}

	// Example 2: Multi-stage quality control
	fmt.Println("\n\n🔍 Coffee Bean Quality Control")
	fmt.Println("------------------------------")
	
	type Bean struct {
		batch      string
		origin     string
		roastDate  int // days ago
		moisture   float64
		defects    int
		aroma      int // 1-10
	}
	
	beans := []Bean{
		{"B001", "Colombia", 2, 11.5, 3, 8},
		{"B002", "Ethiopia", 15, 12.5, 8, 9},  // Too old
		{"B003", "Brazil", 5, 13.5, 2, 7},     // Too moist
		{"B004", "Kenya", 3, 11.0, 12, 9},     // Too many defects
		{"B005", "Peru", 4, 11.2, 1, 6},       // Low aroma
		{"B006", "Jamaica", 1, 10.8, 0, 10},   // Perfect
	}
	
	approvedBatches := []string{}
	
	for _, bean := range beans {
		fmt.Printf("\nBatch %s from %s:\n", bean.batch, bean.origin)
		passedTests := 0
		totalTests := 4
		
		// Test 1: Freshness
		fmt.Printf("  Freshness test (<%d days): ", 10)
		if bean.roastDate > 10 {
			fmt.Printf("❌ Failed (%d days old)\n", bean.roastDate)
			if bean.roastDate > 14 {
				fmt.Println("  ⛔ Critical: Too old - rejecting batch")
				continue
			}
		} else {
			fmt.Printf("✓ Passed (%d days)\n", bean.roastDate)
			passedTests++
		}
		
		// Test 2: Moisture
		fmt.Printf("  Moisture test (10-12%%): ")
		if bean.moisture < 10 || bean.moisture > 12 {
			fmt.Printf("❌ Failed (%.1f%%)\n", bean.moisture)
			if bean.moisture > 13 {
				fmt.Println("  ⛔ Critical: Risk of mold - rejecting batch")
				continue
			}
		} else {
			fmt.Printf("✓ Passed (%.1f%%)\n", bean.moisture)
			passedTests++
		}
		
		// Test 3: Defects
		fmt.Printf("  Defect count (<5): ")
		if bean.defects >= 5 {
			fmt.Printf("❌ Failed (%d defects)\n", bean.defects)
			if bean.defects > 10 {
				fmt.Println("  ⛔ Critical: Too many defects - rejecting batch")
				continue
			}
		} else {
			fmt.Printf("✓ Passed (%d defects)\n", bean.defects)
			passedTests++
		}
		
		// Test 4: Aroma
		fmt.Printf("  Aroma score (>7): ")
		if bean.aroma <= 7 {
			fmt.Printf("⚠️  Warning (%d/10)\n", bean.aroma)
			// Not critical - don't continue
		} else {
			fmt.Printf("✓ Excellent (%d/10)\n", bean.aroma)
			passedTests++
		}
		
		// Final decision
		passingGrade := float64(passedTests) / float64(totalTests) * 100
		fmt.Printf("  Score: %.0f%% (%d/%d tests passed) - ", passingGrade, passedTests, totalTests)
		
		if passingGrade >= 75 {
			fmt.Println("✅ APPROVED")
			approvedBatches = append(approvedBatches, bean.batch)
		} else {
			fmt.Println("❌ REJECTED")
		}
	}
	
	fmt.Printf("\nApproved batches: %v\n", approvedBatches)

	// Example 3: State machine with complex transitions
	fmt.Println("\n\n🎮 Coffee Machine State Controller")
	fmt.Println("---------------------------------")
	
	type MachineState struct {
		state       string
		waterLevel  int
		beanLevel   int
		temperature float64
		errors      []string
	}
	
	machine := MachineState{
		state:       "idle",
		waterLevel:  80,
		beanLevel:   60,
		temperature: 20.0,
		errors:      []string{},
	}
	
	// Simulate machine operation
	for cycle := 1; cycle <= 10; cycle++ {
		fmt.Printf("\nCycle %d - State: %s\n", cycle, machine.state)
		
		// Check for critical errors first
		if machine.waterLevel < 10 {
			machine.errors = append(machine.errors, "Low water")
			machine.state = "error"
			fmt.Println("  ⛔ Critical: Water too low!")
			break
		}
		
		if machine.beanLevel < 5 {
			machine.errors = append(machine.errors, "No beans")
			machine.state = "error"
			fmt.Println("  ⛔ Critical: Out of beans!")
			break
		}
		
		// State transitions
		switch machine.state {
		case "idle":
			fmt.Println("  Waiting for order...")
			if rand.Float32() < 0.6 { // 60% chance of order
				machine.state = "heating"
				fmt.Println("  📥 Order received!")
			}
			
		case "heating":
			machine.temperature += 15
			fmt.Printf("  🔥 Heating: %.1f°C\n", machine.temperature)
			
			if machine.temperature >= 90 {
				machine.state = "ready"
				fmt.Println("  ✓ Temperature reached!")
			}
			
		case "ready":
			if machine.waterLevel < 20 {
				fmt.Println("  ⚠️  Low water warning")
				machine.state = "maintenance"
				continue
			}
			
			machine.state = "brewing"
			fmt.Println("  ☕ Starting brew...")
			
		case "brewing":
			machine.waterLevel -= 10
			machine.beanLevel -= 5
			fmt.Printf("  Brewing... (Water: %d%%, Beans: %d%%)\n", 
				machine.waterLevel, machine.beanLevel)
			
			machine.state = "dispensing"
			
		case "dispensing":
			fmt.Println("  🥤 Dispensing coffee...")
			machine.temperature = 85 // Cool down slightly
			machine.state = "idle"
			fmt.Println("  ✓ Coffee ready!")
			
		case "maintenance":
			fmt.Println("  🔧 Maintenance mode")
			machine.waterLevel = 100
			machine.temperature = 20
			machine.state = "idle"
			fmt.Println("  ✓ Maintenance complete")
		}
		
		time.Sleep(500 * time.Millisecond)
	}
	
	if machine.state == "error" {
		fmt.Printf("\n❌ Machine stopped with errors: %v\n", machine.errors)
	} else {
		fmt.Println("\n✓ Simulation complete!")
	}
}