package main

import "fmt"

func main() {
	fmt.Println("=== GoCoffee Goto Basics ===\n")
	
	fmt.Println("⚠️  WARNING: This example demonstrates goto usage.")
	fmt.Println("In real code, avoid goto whenever possible!\n")

	// Example 1: Basic goto
	fmt.Println("Example 1: Basic Goto")
	fmt.Println("--------------------")
	
	count := 0
	
Start:
	count++
	fmt.Printf("Count: %d\n", count)
	
	if count < 3 {
		goto Start // Jump back to Start label
	}
	
	fmt.Println("Done counting!\n")

	// Example 2: Forward jump
	fmt.Println("Example 2: Forward Jump")
	fmt.Println("----------------------")
	
	fmt.Println("Checking coffee machine...")
	
	machineReady := false
	
	if !machineReady {
		fmt.Println("Machine not ready!")
		goto SkipBrewing
	}
	
	fmt.Println("Brewing coffee...")
	fmt.Println("Coffee ready!")
	
SkipBrewing:
	fmt.Println("Continuing with other tasks...\n")

	// Example 3: Multiple labels
	fmt.Println("Example 3: Multiple Labels")
	fmt.Println("-------------------------")
	
	orderType := "espresso"
	
	fmt.Printf("Processing %s order...\n", orderType)
	
	switch orderType {
	case "espresso":
		goto MakeEspresso
	case "latte":
		goto MakeLatte
	default:
		goto UnknownOrder
	}
	
MakeEspresso:
	fmt.Println("→ Grinding 18g beans")
	fmt.Println("→ Extracting for 25 seconds")
	goto OrderComplete
	
MakeLatte:
	fmt.Println("→ Making espresso base")
	fmt.Println("→ Steaming milk")
	fmt.Println("→ Creating latte art")
	goto OrderComplete
	
UnknownOrder:
	fmt.Println("→ Unknown order type!")
	goto EndOrder
	
OrderComplete:
	fmt.Println("→ Order ready for pickup!")
	
EndOrder:
	fmt.Println("Order processing finished\n")

	// Example 4: Goto with conditions
	fmt.Println("Example 4: Conditional Jumps")
	fmt.Println("---------------------------")
	
	ingredients := map[string]int{
		"beans": 100,
		"milk":  500,
		"sugar": 200,
	}
	
	fmt.Println("Checking ingredients...")
	
	if ingredients["beans"] < 20 {
		goto LowBeans
	}
	
	if ingredients["milk"] < 100 {
		goto LowMilk
	}
	
	fmt.Println("✓ All ingredients sufficient!")
	goto IngredientsOK
	
LowBeans:
	fmt.Println("⚠️  Low on coffee beans!")
	goto OrderSupplies
	
LowMilk:
	fmt.Println("⚠️  Low on milk!")
	goto OrderSupplies
	
OrderSupplies:
	fmt.Println("📦 Need to order supplies")
	goto EndCheck
	
IngredientsOK:
	fmt.Println("✅ Ready to serve customers!")
	
EndCheck:
	fmt.Println("Ingredient check complete\n")

	// Example 5: Why goto can be confusing
	fmt.Println("Example 5: Spaghetti Code Warning")
	fmt.Println("--------------------------------")
	fmt.Println("The following shows why goto creates hard-to-follow code:\n")
	
	x := 1
	y := 1
	
First:
	fmt.Printf("x=%d, y=%d → ", x, y)
	if x < 3 {
		x++
		goto Second
	}
	goto Done
	
Second:
	fmt.Printf("In Second → ")
	if y < 3 {
		y++
		goto Third
	}
	goto First
	
Third:
	fmt.Printf("In Third → ")
	if x == y {
		goto First
	}
	goto Second
	
Done:
	fmt.Printf("\nFinal: x=%d, y=%d\n", x, y)
	fmt.Println("\n🤯 Confused? That's why we avoid goto!")

	// Show the restriction
	fmt.Println("\n\nGo's Safety Restrictions:")
	fmt.Println("------------------------")
	fmt.Println("Go prevents some dangerous goto uses:")
	fmt.Println("1. Cannot jump over variable declarations")
	fmt.Println("2. Cannot jump into a block from outside")
	fmt.Println("3. Must be within the same function")
	fmt.Println("\nThese restrictions prevent many goto abuses.")
}