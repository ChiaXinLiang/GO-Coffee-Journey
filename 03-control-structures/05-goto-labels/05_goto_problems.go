package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== GoCoffee: Why Goto Is Dangerous ===\n")
	
	fmt.Println("This example demonstrates common problems with goto")
	fmt.Println("and why it makes code difficult to understand and maintain.\n")
	
	// Example 1: Spaghetti code
	fmt.Println("Example 1: Spaghetti Code")
	fmt.Println("------------------------")
	fmt.Println("Try to follow the flow of this code:\n")
	
	demonstrateSpaghetti()
	
	fmt.Println("\n" + strings.Repeat("=", 50) + "\n")
	
	// Example 2: Skipping initialization
	fmt.Println("Example 2: Variable Initialization Issues")
	fmt.Println("----------------------------------------")
	
	demonstrateInitializationIssues()
	
	fmt.Println("\n" + strings.Repeat("=", 50) + "\n")
	
	// Example 3: Unclear flow
	fmt.Println("Example 3: Unclear Control Flow")
	fmt.Println("-------------------------------")
	
	demonstrateUnclearFlow()
	
	fmt.Println("\n" + strings.Repeat("=", 50) + "\n")
	
	// Example 4: Maintenance nightmare
	fmt.Println("Example 4: Maintenance Nightmare")
	fmt.Println("-------------------------------")
	
	demonstrateMaintenanceIssues()
	
	fmt.Println("\n" + strings.Repeat("=", 50) + "\n")
	
	// Example 5: Go's restrictions
	fmt.Println("Example 5: Go's Safety Restrictions")
	fmt.Println("----------------------------------")
	
	demonstrateGoRestrictions()
}

// Example 1: Classic spaghetti code
func demonstrateSpaghetti() {
	x := 0
	y := 0
	z := 0
	
	fmt.Println("Starting values: x=0, y=0, z=0")
	fmt.Println("Following the execution...")
	
Start:
	x++
	fmt.Printf("  Start: x=%d\n", x)
	if x > 3 {
		goto End
	}
	goto Middle
	
Bottom:
	z++
	fmt.Printf("  Bottom: z=%d\n", z)
	if z < 2 {
		goto Top
	}
	goto Start
	
Middle:
	y++
	fmt.Printf("  Middle: y=%d\n", y)
	if y > 2 {
		goto Bottom
	}
	goto Top
	
Top:
	fmt.Printf("  Top: checking x=%d, y=%d\n", x, y)
	if x == y {
		goto Middle
	}
	goto Bottom
	
End:
	fmt.Printf("\nFinal values: x=%d, y=%d, z=%d\n", x, y, z)
	fmt.Println("\n🤯 Can you trace the execution path? This is spaghetti code!")
}

// Example 2: Variable initialization problems
func demonstrateInitializationIssues() {
	fmt.Println("Attempting to show initialization issues...")
	
	// This would be problematic in other languages, but Go prevents it
	fmt.Println("\nIn Go, you CANNOT do this:")
	fmt.Println(`
	goto Skip
	x := 5  // Error: goto jumps over declaration
Skip:
	fmt.Println(x)
	`)
	
	fmt.Println("\nGo compiler prevents this dangerous pattern!")
	
	// Show what could happen without this protection
	fmt.Println("\nWithout Go's protection, you might have:")
	
	var coffeeCount int
	var initialized bool
	
	if false { // This simulates what could happen
		goto UseVariable
	}
	
	// Initialization that might be skipped
	coffeeCount = 10
	initialized = true
	
UseVariable:
	if !initialized {
		fmt.Printf("  ⚠️  Using uninitialized variable: coffeeCount=%d\n", coffeeCount)
		fmt.Println("  This could lead to bugs!")
	} else {
		fmt.Printf("  ✓ Variable properly initialized: coffeeCount=%d\n", coffeeCount)
	}
}

// Example 3: Unclear control flow
func demonstrateUnclearFlow() {
	orders := []string{"Latte", "Espresso", "Mocha", "Cappuccino"}
	processed := 0
	failed := 0
	i := 0
	
	fmt.Println("Processing orders with goto (confusing flow):\n")
	
CheckOrder:
	if i >= len(orders) {
		goto Summary
	}
	
	fmt.Printf("Order %d: %s", i+1, orders[i])
	
	// Random validation
	if orders[i] == "Mocha" {
		goto FailOrder
	}
	
	// Random processing
	if i == 1 {
		goto SpecialProcess
	}
	
	goto NormalProcess
	
SpecialProcess:
	fmt.Print(" → Special processing")
	processed++
	goto NextOrder
	
NormalProcess:
	fmt.Print(" → Normal processing")
	processed++
	goto NextOrder
	
FailOrder:
	fmt.Print(" → FAILED")
	failed++
	goto NextOrder
	
NextOrder:
	fmt.Println()
	i++
	goto CheckOrder
	
Summary:
	fmt.Printf("\nProcessed: %d, Failed: %d\n", processed, failed)
	fmt.Println("\n😵 The flow jumps all over the place!")
}

// Example 4: Maintenance nightmare
func demonstrateMaintenanceIssues() {
	fmt.Println("Imagine maintaining this payment processing code:\n")
	
	// Simulate complex business logic with goto
	amount := 50.0
	customerType := "regular"
	hasLoyaltyCard := true
	paymentMethod := "card"
	
	var discount float64
	var fee float64
	var total float64
	
	fmt.Printf("Processing payment: $%.2f\n", amount)
	
	// Complex goto-based logic (DON'T DO THIS!)
	if customerType == "vip" {
		goto VIPDiscount
	}
	
	if hasLoyaltyCard {
		goto LoyaltyDiscount
	}
	
	goto NoDiscount
	
VIPDiscount:
	discount = amount * 0.20
	fmt.Println("  VIP discount: 20%")
	goto CheckPayment
	
LoyaltyDiscount:
	discount = amount * 0.10
	fmt.Println("  Loyalty discount: 10%")
	if amount > 100 {
		goto BonusDiscount
	}
	goto CheckPayment
	
BonusDiscount:
	discount += amount * 0.05
	fmt.Println("  Bonus discount: +5%")
	goto CheckPayment
	
NoDiscount:
	discount = 0
	fmt.Println("  No discount")
	
CheckPayment:
	if paymentMethod == "cash" {
		goto CashPayment
	}
	goto CardPayment
	
CashPayment:
	fee = 0
	fmt.Println("  Cash payment: no fee")
	goto Calculate
	
CardPayment:
	fee = 2.50
	fmt.Println("  Card payment: $2.50 fee")
	goto Calculate
	
Calculate:
	total = amount - discount + fee
	goto DisplayTotal
	
DisplayTotal:
	fmt.Printf("\nAmount: $%.2f\n", amount)
	fmt.Printf("Discount: -$%.2f\n", discount)
	fmt.Printf("Fee: +$%.2f\n", fee)
	fmt.Printf("Total: $%.2f\n", total)
	
	fmt.Println("\n🔧 Now imagine adding a new discount type or payment method!")
	fmt.Println("The goto statements make it very hard to modify safely.")
}

// Example 5: Go's restrictions prevent some issues
func demonstrateGoRestrictions() {
	fmt.Println("Go prevents many dangerous goto patterns:\n")
	
	// Restriction 1: Cannot jump over declarations
	fmt.Println("1. Cannot jump over variable declarations:")
	fmt.Println(`
// This is ILLEGAL in Go:
goto Label
x := 5  // Error: goto jumps over declaration
Label:
	fmt.Println(x)
`)
	
	// Restriction 2: Cannot jump into blocks
	fmt.Println("\n2. Cannot jump into a block from outside:")
	fmt.Println(`
// This is ILLEGAL in Go:
if condition {
Label:  // Label inside block
    fmt.Println("Inside")
}
goto Label  // Error: cannot jump into block
`)
	
	// Restriction 3: Function scope
	fmt.Println("\n3. Labels are function-scoped:")
	fmt.Println(`
// Cannot goto a label in another function
func foo() {
Label:
    // ...
}

func bar() {
    goto Label  // Error: label not defined
}
`)
	
	// Show what IS allowed (but still not recommended)
	fmt.Println("\n4. What Go DOES allow (but you shouldn't use):")
	
	condition := true
	
	if condition {
		goto Allowed
	}
	
	fmt.Println("  This won't print")
	
Allowed:
	fmt.Println("  ✓ Jumping forward is allowed")
	
	// Demonstrate scope issues
	fmt.Println("\n5. Scope confusion example:")
	
	for i := 0; i < 3; i++ {
		if i == 1 {
			goto LoopEnd
		}
		fmt.Printf("  Loop iteration: %d\n", i)
	}
	
LoopEnd:
	fmt.Println("  Jumped out of loop (but 'i' is now out of scope)")
	
	fmt.Println("\n📚 Even with restrictions, goto still makes code hard to understand!")
}

// Helper functions
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