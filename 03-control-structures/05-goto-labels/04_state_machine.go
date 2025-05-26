package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("=== GoCoffee State Machine Examples ===\n")
	
	rand.Seed(time.Now().UnixNano())

	// Example 1: Goto-based state machine (NOT RECOMMENDED)
	fmt.Println("Example 1: Coffee Machine with Goto")
	fmt.Println("-----------------------------------")
	coffeeMachineGoto()
	
	fmt.Println("\n" + strings.Repeat("=", 50) + "\n")
	
	// Example 2: Better state machine without goto
	fmt.Println("Example 2: Coffee Machine with Switch")
	fmt.Println("-------------------------------------")
	coffeeMachineSwitch()
	
	fmt.Println("\n" + strings.Repeat("=", 50) + "\n")
	
	// Example 3: Order processing state machine
	fmt.Println("Example 3: Order Processing States")
	fmt.Println("---------------------------------")
	orderProcessingComparison()
}

// Example 1: State machine using goto (hard to maintain)
func coffeeMachineGoto() {
	beans := 100
	water := 1000
	cups := 0
	
	fmt.Println("Starting coffee machine (goto version)...")
	fmt.Printf("Initial: %dg beans, %dml water\n\n", beans, water)
	
Idle:
	fmt.Println("State: IDLE - Waiting for button press")
	time.Sleep(500 * time.Millisecond)
	
	// Simulate button press
	if rand.Float32() < 0.7 {
		goto CheckResources
	}
	goto Idle
	
CheckResources:
	fmt.Println("State: CHECKING RESOURCES")
	
	if beans < 18 {
		goto OutOfBeans
	}
	if water < 200 {
		goto OutOfWater
	}
	goto Heating
	
Heating:
	fmt.Println("State: HEATING - Warming up...")
	time.Sleep(800 * time.Millisecond)
	goto Grinding
	
Grinding:
	fmt.Println("State: GRINDING - Grinding beans...")
	beans -= 18
	time.Sleep(600 * time.Millisecond)
	goto Brewing
	
Brewing:
	fmt.Println("State: BREWING - Making coffee...")
	water -= 200
	time.Sleep(1000 * time.Millisecond)
	goto Dispensing
	
Dispensing:
	fmt.Println("State: DISPENSING - Pouring coffee...")
	cups++
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("☕ Cup #%d ready! (%dg beans, %dml water left)\n", cups, beans, water)
	
	if cups < 3 {
		goto Idle
	}
	goto Shutdown
	
OutOfBeans:
	fmt.Println("State: ERROR - Out of beans!")
	goto Maintenance
	
OutOfWater:
	fmt.Println("State: ERROR - Out of water!")
	goto Maintenance
	
Maintenance:
	fmt.Println("State: MAINTENANCE - Refilling...")
	beans = 100
	water = 1000
	time.Sleep(1000 * time.Millisecond)
	goto Idle
	
Shutdown:
	fmt.Println("State: SHUTDOWN - Machine stopping...")
	fmt.Printf("Total cups made: %d\n", cups)
}

// Example 2: Better state machine using switch
func coffeeMachineSwitch() {
	type MachineState struct {
		state  string
		beans  int
		water  int
		cups   int
	}
	
	machine := MachineState{
		state: "idle",
		beans: 100,
		water: 1000,
		cups:  0,
	}
	
	fmt.Println("Starting coffee machine (switch version)...")
	fmt.Printf("Initial: %dg beans, %dml water\n\n", machine.beans, machine.water)
	
	for machine.state != "shutdown" {
		fmt.Printf("State: %s", strings.ToUpper(machine.state))
		
		switch machine.state {
		case "idle":
			fmt.Println(" - Waiting for button press")
			time.Sleep(500 * time.Millisecond)
			// Simulate button press
			if rand.Float32() < 0.7 {
				machine.state = "check_resources"
			}
			
		case "check_resources":
			fmt.Println(" - Checking resources")
			if machine.beans < 18 {
				machine.state = "error_beans"
			} else if machine.water < 200 {
				machine.state = "error_water"
			} else {
				machine.state = "heating"
			}
			
		case "heating":
			fmt.Println(" - Warming up...")
			time.Sleep(800 * time.Millisecond)
			machine.state = "grinding"
			
		case "grinding":
			fmt.Println(" - Grinding beans...")
			machine.beans -= 18
			time.Sleep(600 * time.Millisecond)
			machine.state = "brewing"
			
		case "brewing":
			fmt.Println(" - Making coffee...")
			machine.water -= 200
			time.Sleep(1000 * time.Millisecond)
			machine.state = "dispensing"
			
		case "dispensing":
			fmt.Println(" - Pouring coffee...")
			machine.cups++
			time.Sleep(500 * time.Millisecond)
			fmt.Printf("☕ Cup #%d ready! (%dg beans, %dml water left)\n", 
				machine.cups, machine.beans, machine.water)
			
			if machine.cups < 3 {
				machine.state = "idle"
			} else {
				machine.state = "shutdown"
			}
			
		case "error_beans":
			fmt.Println(" - Out of beans!")
			machine.state = "maintenance"
			
		case "error_water":
			fmt.Println(" - Out of water!")
			machine.state = "maintenance"
			
		case "maintenance":
			fmt.Println(" - Refilling...")
			machine.beans = 100
			machine.water = 1000
			time.Sleep(1000 * time.Millisecond)
			machine.state = "idle"
		}
	}
	
	fmt.Println("State: SHUTDOWN - Machine stopping...")
	fmt.Printf("Total cups made: %d\n", machine.cups)
}

// Example 3: Comparison of order processing approaches
func orderProcessingComparison() {
	fmt.Println("\n🔴 Goto-based order processing:")
	fmt.Println("--------------------------------")
	processOrderGoto("Latte", true)
	
	fmt.Println("\n\n🟢 Function-based order processing:")
	fmt.Println("-----------------------------------")
	processOrderFunctional("Latte", true)
}

// Goto version - harder to test and maintain
func processOrderGoto(drink string, rush bool) {
	fmt.Printf("Processing %s order (rush: %v)\n", drink, rush)
	
Received:
	fmt.Println("  → Order received")
	time.Sleep(200 * time.Millisecond)
	
	if rush {
		goto RushValidation
	}
	goto NormalValidation
	
RushValidation:
	fmt.Println("  → Rush order validation")
	if rand.Float32() < 0.8 {
		goto Preparation
	}
	goto Rejected
	
NormalValidation:
	fmt.Println("  → Normal validation")
	goto Preparation
	
Preparation:
	fmt.Println("  → Preparing drink")
	time.Sleep(500 * time.Millisecond)
	
	if rand.Float32() < 0.9 {
		goto QualityCheck
	}
	goto PreparationError
	
QualityCheck:
	fmt.Println("  → Quality check")
	time.Sleep(200 * time.Millisecond)
	goto Ready
	
Ready:
	fmt.Println("  → Order ready!")
	goto Complete
	
PreparationError:
	fmt.Println("  → Preparation failed")
	goto Retry
	
Rejected:
	fmt.Println("  → Order rejected")
	goto Failed
	
Retry:
	fmt.Println("  → Retrying...")
	goto Preparation
	
Failed:
	fmt.Println("  ❌ Order failed")
	return
	
Complete:
	fmt.Println("  ✅ Order complete")
}

// Functional version - cleaner and more testable
func processOrderFunctional(drink string, rush bool) {
	fmt.Printf("Processing %s order (rush: %v)\n", drink, rush)
	
	order := &Order{
		drink: drink,
		rush:  rush,
		state: "received",
	}
	
	// Process through states
	for order.state != "complete" && order.state != "failed" {
		order.state = processOrderState(order)
	}
	
	if order.state == "complete" {
		fmt.Println("  ✅ Order complete")
	} else {
		fmt.Println("  ❌ Order failed")
	}
}

type Order struct {
	drink string
	rush  bool
	state string
}

func processOrderState(order *Order) string {
	fmt.Printf("  → %s\n", strings.ToTitle(order.state))
	time.Sleep(200 * time.Millisecond)
	
	switch order.state {
	case "received":
		if order.rush {
			return "rush_validation"
		}
		return "normal_validation"
		
	case "rush_validation":
		if rand.Float32() < 0.8 {
			return "preparation"
		}
		return "rejected"
		
	case "normal_validation":
		return "preparation"
		
	case "preparation":
		time.Sleep(300 * time.Millisecond) // Extra time for preparation
		if rand.Float32() < 0.9 {
			return "quality_check"
		}
		return "retry"
		
	case "quality_check":
		return "ready"
		
	case "ready":
		return "complete"
		
	case "retry":
		fmt.Println("  → Retrying...")
		return "preparation"
		
	case "rejected":
		return "failed"
		
	default:
		return "failed"
	}
}

// Helper string functions
var strings = struct {
	Repeat  func(string, int) string
	ToUpper func(string) string
	ToTitle func(string) string
}{
	Repeat: func(s string, n int) string {
		result := ""
		for i := 0; i < n; i++ {
			result += s
		}
		return result
	},
	ToUpper: func(s string) string {
		result := ""
		for _, r := range s {
			if r >= 'a' && r <= 'z' {
				result += string(r - 32)
			} else {
				result += string(r)
			}
		}
		return result
	},
	ToTitle: func(s string) string {
		result := ""
		capitalizeNext := true
		for _, r := range s {
			if r == '_' || r == '-' {
				result += " "
				capitalizeNext = true
			} else if capitalizeNext && r >= 'a' && r <= 'z' {
				result += string(r - 32)
				capitalizeNext = false
			} else {
				result += string(r)
				capitalizeNext = false
			}
		}
		return result
	},
}