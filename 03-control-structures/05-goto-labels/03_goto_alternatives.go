package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== GoCoffee: Goto vs Better Alternatives ===\n")

	// Example 1: Loop control
	fmt.Println("Example 1: Loop Control")
	fmt.Println("-----------------------")
	
	fmt.Println("❌ Using goto (confusing):")
	loopWithGoto()
	
	fmt.Println("\n✅ Using break (clear):")
	loopWithBreak()

	fmt.Println("\n" + strings.Repeat("=", 50) + "\n")

	// Example 2: Nested loop escape
	fmt.Println("Example 2: Escaping Nested Loops")
	fmt.Println("--------------------------------")
	
	fmt.Println("❌ Using goto:")
	nestedLoopGoto()
	
	fmt.Println("\n✅ Using labeled break:")
	nestedLoopLabeledBreak()

	fmt.Println("\n" + strings.Repeat("=", 50) + "\n")

	// Example 3: State machine
	fmt.Println("Example 3: State Machine")
	fmt.Println("------------------------")
	
	fmt.Println("❌ Using goto (hard to follow):")
	stateMachineGoto()
	
	fmt.Println("\n✅ Using switch (structured):")
	stateMachineSwitch()

	fmt.Println("\n" + strings.Repeat("=", 50) + "\n")

	// Example 4: Error handling flow
	fmt.Println("Example 4: Error Handling Flow")
	fmt.Println("-----------------------------")
	
	fmt.Println("❌ Using goto:")
	errorFlowGoto()
	
	fmt.Println("\n✅ Using functions:")
	errorFlowFunctions()

	fmt.Println("\n" + strings.Repeat("=", 50) + "\n")

	// Example 5: Complex conditions
	fmt.Println("Example 5: Complex Conditional Logic")
	fmt.Println("-----------------------------------")
	
	fmt.Println("❌ Using goto (spaghetti):")
	complexConditionsGoto()
	
	fmt.Println("\n✅ Using structured if/else:")
	complexConditionsStructured()
}

// Example 1: Loop control alternatives

func loopWithGoto() {
	i := 0
Loop:
	if i >= 5 {
		goto End
	}
	fmt.Printf("  Brewing coffee #%d\n", i+1)
	i++
	goto Loop
End:
	fmt.Println("  Done brewing!")
}

func loopWithBreak() {
	for i := 0; i < 5; i++ {
		fmt.Printf("  Brewing coffee #%d\n", i+1)
	}
	fmt.Println("  Done brewing!")
}

// Example 2: Nested loop alternatives

func nestedLoopGoto() {
	fmt.Println("  Searching for empty table...")
	for floor := 1; floor <= 3; floor++ {
		for table := 1; table <= 4; table++ {
			if floor == 2 && table == 3 {
				fmt.Printf("  Found: Floor %d, Table %d\n", floor, table)
				goto Found
			}
		}
	}
Found:
	fmt.Println("  Search complete")
}

func nestedLoopLabeledBreak() {
	fmt.Println("  Searching for empty table...")
FloorLoop:
	for floor := 1; floor <= 3; floor++ {
		for table := 1; table <= 4; table++ {
			if floor == 2 && table == 3 {
				fmt.Printf("  Found: Floor %d, Table %d\n", floor, table)
				break FloorLoop
			}
		}
	}
	fmt.Println("  Search complete")
}

// Example 3: State machine alternatives

func stateMachineGoto() {
	orderCount := 0
	
Idle:
	fmt.Println("  State: Idle")
	time.Sleep(200 * time.Millisecond)
	goto TakeOrder
	
TakeOrder:
	orderCount++
	fmt.Printf("  State: Taking order #%d\n", orderCount)
	time.Sleep(200 * time.Millisecond)
	goto Prepare
	
Prepare:
	fmt.Println("  State: Preparing")
	time.Sleep(200 * time.Millisecond)
	goto Serve
	
Serve:
	fmt.Println("  State: Serving")
	time.Sleep(200 * time.Millisecond)
	if orderCount < 3 {
		goto Idle
	}
	goto Done
	
Done:
	fmt.Println("  State: Done")
}

func stateMachineSwitch() {
	state := "idle"
	orderCount := 0
	
	for state != "done" {
		fmt.Printf("  State: %s\n", state)
		time.Sleep(200 * time.Millisecond)
		
		switch state {
		case "idle":
			state = "take_order"
		case "take_order":
			orderCount++
			fmt.Printf("    (Order #%d)\n", orderCount)
			state = "prepare"
		case "prepare":
			state = "serve"
		case "serve":
			if orderCount < 3 {
				state = "idle"
			} else {
				state = "done"
			}
		}
	}
}

// Example 4: Error handling alternatives

func errorFlowGoto() {
	step := 1
	
	fmt.Printf("  Step %d: Check supplies\n", step)
	if !checkSupplies() {
		goto HandleError
	}
	
	step++
	fmt.Printf("  Step %d: Heat water\n", step)
	if !heatWater() {
		goto HandleError
	}
	
	step++
	fmt.Printf("  Step %d: Brew coffee\n", step)
	if !brewCoffee() {
		goto HandleError
	}
	
	fmt.Println("  ✅ Success!")
	return
	
HandleError:
	fmt.Printf("  ❌ Error at step %d\n", step)
}

func errorFlowFunctions() {
	steps := []struct {
		name string
		fn   func() bool
	}{
		{"Check supplies", checkSupplies},
		{"Heat water", heatWater},
		{"Brew coffee", brewCoffee},
	}
	
	for i, step := range steps {
		fmt.Printf("  Step %d: %s\n", i+1, step.name)
		if !step.fn() {
			fmt.Printf("  ❌ Error at step %d\n", i+1)
			return
		}
	}
	
	fmt.Println("  ✅ Success!")
}

// Example 5: Complex conditional alternatives

func complexConditionsGoto() {
	hasWater := true
	hasBeans := true
	machineReady := false
	
	if !hasWater {
		goto NoWater
	}
	if !hasBeans {
		goto NoBeans
	}
	if !machineReady {
		goto MachineNotReady
	}
	goto CanMakeCoffee
	
NoWater:
	fmt.Println("  ❌ No water available")
	goto CannotMakeCoffee
	
NoBeans:
	fmt.Println("  ❌ No beans available")
	goto CannotMakeCoffee
	
MachineNotReady:
	fmt.Println("  ❌ Machine not ready")
	goto CannotMakeCoffee
	
CanMakeCoffee:
	fmt.Println("  ✅ Can make coffee!")
	return
	
CannotMakeCoffee:
	fmt.Println("  Cannot make coffee")
}

func complexConditionsStructured() {
	hasWater := true
	hasBeans := true
	machineReady := false
	
	canMake := true
	reason := ""
	
	if !hasWater {
		canMake = false
		reason = "No water available"
	} else if !hasBeans {
		canMake = false
		reason = "No beans available"
	} else if !machineReady {
		canMake = false
		reason = "Machine not ready"
	}
	
	if canMake {
		fmt.Println("  ✅ Can make coffee!")
	} else {
		fmt.Printf("  ❌ %s\n", reason)
		fmt.Println("  Cannot make coffee")
	}
}

// Helper functions

func checkSupplies() bool {
	return true
}

func heatWater() bool {
	return true
}

func brewCoffee() bool {
	return false // Simulate failure
}

// strings package substitute
var strings = struct {
	Repeat func(string, int) string
}{
	Repeat: func(s string, n int) string {
		result := ""
		for i := 0; i < n; i++ {
			result += s
		}
		return result
	},
}