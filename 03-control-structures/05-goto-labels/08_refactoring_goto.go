package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("=== GoCoffee: Refactoring Away From Goto ===\n")
	
	fmt.Println("This example shows how to refactor goto-based code")
	fmt.Println("into cleaner, more maintainable alternatives.\n")
	
	rand.Seed(time.Now().UnixNano())
	
	// Example 1: Order processing
	fmt.Println("Example 1: Order Processing Refactoring")
	fmt.Println("--------------------------------------")
	
	fmt.Println("\n❌ BEFORE (with goto):")
	processOrderGoto("Latte", 5.50)
	
	fmt.Println("\n✅ AFTER (refactored):")
	processOrderRefactored("Latte", 5.50)
	
	fmt.Println("\n" + strings.Repeat("=", 50) + "\n")
	
	// Example 2: Resource management
	fmt.Println("Example 2: Resource Management Refactoring")
	fmt.Println("-----------------------------------------")
	
	fmt.Println("\n❌ BEFORE (with goto):")
	manageResourcesGoto()
	
	fmt.Println("\n✅ AFTER (with defer):")
	manageResourcesDefer()
	
	fmt.Println("\n" + strings.Repeat("=", 50) + "\n")
	
	// Example 3: Complex validation
	fmt.Println("Example 3: Validation Logic Refactoring")
	fmt.Println("--------------------------------------")
	
	order := Order{
		CustomerID: "CUST123",
		Items:      []string{"Latte", "Muffin"},
		Total:      8.50,
		PaymentMethod: "card",
	}
	
	fmt.Println("\n❌ BEFORE (with goto):")
	validateOrderGoto(order)
	
	fmt.Println("\n✅ AFTER (structured):")
	validateOrderStructured(order)
	
	fmt.Println("\n" + strings.Repeat("=", 50) + "\n")
	
	// Example 4: State machine refactoring
	fmt.Println("Example 4: State Machine Refactoring")
	fmt.Println("-----------------------------------")
	
	fmt.Println("\n❌ BEFORE (goto-based):")
	runCoffeeMachineGoto()
	
	fmt.Println("\n✅ AFTER (table-driven):")
	runCoffeeMachineRefactored()
}

// Example 1: Order processing refactoring

func processOrderGoto(item string, amount float64) {
	var discount float64
	var tax float64
	var total float64
	
	fmt.Printf("Processing order: %s ($%.2f)\n", item, amount)
	
	// Check amount
	if amount <= 0 {
		goto InvalidAmount
	}
	
	// Calculate discount
	if amount > 10 {
		discount = amount * 0.1
		goto CalculateTax
	}
	
	if amount > 5 {
		discount = amount * 0.05
		goto CalculateTax
	}
	
	discount = 0
	
CalculateTax:
	tax = (amount - discount) * 0.08
	goto CalculateTotal
	
CalculateTotal:
	total = amount - discount + tax
	goto PrintReceipt
	
PrintReceipt:
	fmt.Printf("  Subtotal: $%.2f\n", amount)
	fmt.Printf("  Discount: -$%.2f\n", discount)
	fmt.Printf("  Tax: +$%.2f\n", tax)
	fmt.Printf("  Total: $%.2f\n", total)
	goto Success
	
InvalidAmount:
	fmt.Println("  ❌ Invalid amount!")
	goto Failure
	
Success:
	fmt.Println("  ✅ Order processed!")
	return
	
Failure:
	fmt.Println("  ❌ Order failed!")
}

func processOrderRefactored(item string, amount float64) {
	fmt.Printf("Processing order: %s ($%.2f)\n", item, amount)
	
	// Validate amount
	if amount <= 0 {
		fmt.Println("  ❌ Invalid amount!")
		fmt.Println("  ❌ Order failed!")
		return
	}
	
	// Calculate discount
	discount := calculateDiscount(amount)
	
	// Calculate tax
	tax := (amount - discount) * 0.08
	
	// Calculate total
	total := amount - discount + tax
	
	// Print receipt
	fmt.Printf("  Subtotal: $%.2f\n", amount)
	fmt.Printf("  Discount: -$%.2f\n", discount)
	fmt.Printf("  Tax: +$%.2f\n", tax)
	fmt.Printf("  Total: $%.2f\n", total)
	fmt.Println("  ✅ Order processed!")
}

func calculateDiscount(amount float64) float64 {
	switch {
	case amount > 10:
		return amount * 0.1
	case amount > 5:
		return amount * 0.05
	default:
		return 0
	}
}

// Example 2: Resource management refactoring

type Resource struct {
	name string
}

func (r *Resource) Acquire() error {
	if rand.Float32() < 0.2 {
		return fmt.Errorf("failed to acquire %s", r.name)
	}
	fmt.Printf("  ✓ Acquired %s\n", r.name)
	return nil
}

func (r *Resource) Release() {
	fmt.Printf("  → Released %s\n", r.name)
}

func manageResourcesGoto() {
	var r1, r2, r3 *Resource
	var err error
	
	fmt.Println("Acquiring resources...")
	
	r1 = &Resource{name: "Database Connection"}
	err = r1.Acquire()
	if err != nil {
		goto Cleanup
	}
	
	r2 = &Resource{name: "File Handle"}
	err = r2.Acquire()
	if err != nil {
		goto Cleanup
	}
	
	r3 = &Resource{name: "Network Socket"}
	err = r3.Acquire()
	if err != nil {
		goto Cleanup
	}
	
	// Do work
	fmt.Println("  → Doing work...")
	goto Cleanup
	
Cleanup:
	fmt.Println("Cleanup:")
	if r3 != nil {
		r3.Release()
	}
	if r2 != nil {
		r2.Release()
	}
	if r1 != nil {
		r1.Release()
	}
	
	if err != nil {
		fmt.Printf("  ❌ Operation failed: %v\n", err)
	} else {
		fmt.Println("  ✅ Operation completed!")
	}
}

func manageResourcesDefer() {
	fmt.Println("Acquiring resources...")
	
	r1 := &Resource{name: "Database Connection"}
	if err := r1.Acquire(); err != nil {
		fmt.Printf("  ❌ Operation failed: %v\n", err)
		return
	}
	defer r1.Release()
	
	r2 := &Resource{name: "File Handle"}
	if err := r2.Acquire(); err != nil {
		fmt.Printf("  ❌ Operation failed: %v\n", err)
		return
	}
	defer r2.Release()
	
	r3 := &Resource{name: "Network Socket"}
	if err := r3.Acquire(); err != nil {
		fmt.Printf("  ❌ Operation failed: %v\n", err)
		return
	}
	defer r3.Release()
	
	// Do work
	fmt.Println("  → Doing work...")
	fmt.Println("  ✅ Operation completed!")
	
	fmt.Println("Cleanup (automatic with defer):")
}

// Example 3: Validation refactoring

type Order struct {
	CustomerID    string
	Items         []string
	Total         float64
	PaymentMethod string
}

func validateOrderGoto(order Order) {
	fmt.Println("Validating order...")
	
	// Check customer ID
	if order.CustomerID == "" {
		goto InvalidCustomer
	}
	
	// Check items
	if len(order.Items) == 0 {
		goto NoItems
	}
	
	// Check total
	if order.Total <= 0 {
		goto InvalidTotal
	}
	
	// Check payment method
	if order.PaymentMethod != "cash" && order.PaymentMethod != "card" {
		goto InvalidPayment
	}
	
	goto Valid
	
InvalidCustomer:
	fmt.Println("  ❌ Invalid customer ID")
	goto Invalid
	
NoItems:
	fmt.Println("  ❌ No items in order")
	goto Invalid
	
InvalidTotal:
	fmt.Println("  ❌ Invalid total amount")
	goto Invalid
	
InvalidPayment:
	fmt.Println("  ❌ Invalid payment method")
	goto Invalid
	
Valid:
	fmt.Println("  ✅ Order is valid!")
	return
	
Invalid:
	fmt.Println("  ❌ Order validation failed!")
}

func validateOrderStructured(order Order) {
	fmt.Println("Validating order...")
	
	// Create validation rules
	validations := []struct {
		name  string
		check func() bool
		error string
	}{
		{
			name:  "Customer ID",
			check: func() bool { return order.CustomerID != "" },
			error: "Invalid customer ID",
		},
		{
			name:  "Items",
			check: func() bool { return len(order.Items) > 0 },
			error: "No items in order",
		},
		{
			name:  "Total",
			check: func() bool { return order.Total > 0 },
			error: "Invalid total amount",
		},
		{
			name:  "Payment Method",
			check: func() bool {
				return order.PaymentMethod == "cash" || order.PaymentMethod == "card"
			},
			error: "Invalid payment method",
		},
	}
	
	// Run validations
	for _, v := range validations {
		if !v.check() {
			fmt.Printf("  ❌ %s\n", v.error)
			fmt.Println("  ❌ Order validation failed!")
			return
		}
	}
	
	fmt.Println("  ✅ Order is valid!")
}

// Example 4: State machine refactoring

func runCoffeeMachineGoto() {
	state := "idle"
	cycles := 0
	
	fmt.Println("Coffee machine (goto-based):")
	
Idle:
	cycles++
	fmt.Printf("  Cycle %d: Idle\n", cycles)
	time.Sleep(200 * time.Millisecond)
	state = "heating"
	goto Heating
	
Heating:
	fmt.Printf("  Cycle %d: Heating\n", cycles)
	time.Sleep(200 * time.Millisecond)
	state = "brewing"
	goto Brewing
	
Brewing:
	fmt.Printf("  Cycle %d: Brewing\n", cycles)
	time.Sleep(200 * time.Millisecond)
	state = "ready"
	goto Ready
	
Ready:
	fmt.Printf("  Cycle %d: Ready\n", cycles)
	if cycles < 2 {
		state = "idle"
		goto Idle
	}
	fmt.Println("  ✅ Complete!")
}

func runCoffeeMachineRefactored() {
	fmt.Println("Coffee machine (table-driven):")
	
	type State struct {
		name     string
		duration time.Duration
		next     string
	}
	
	// Define state transitions
	states := map[string]State{
		"idle":    {name: "Idle", duration: 200 * time.Millisecond, next: "heating"},
		"heating": {name: "Heating", duration: 200 * time.Millisecond, next: "brewing"},
		"brewing": {name: "Brewing", duration: 200 * time.Millisecond, next: "ready"},
		"ready":   {name: "Ready", duration: 0, next: "idle"},
	}
	
	currentState := "idle"
	cycles := 0
	
	for cycles < 2 {
		cycles++
		state := states[currentState]
		
		fmt.Printf("  Cycle %d: %s\n", cycles, state.name)
		if state.duration > 0 {
			time.Sleep(state.duration)
		}
		
		if currentState == "ready" {
			if cycles < 2 {
				currentState = "idle"
			} else {
				break
			}
		} else {
			currentState = state.next
		}
	}
	
	fmt.Println("  ✅ Complete!")
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