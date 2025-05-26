package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("=== GoCoffee Parameter Basics ===\n")
	
	// Example 1: Pass by value
	fmt.Println("Example 1: Pass by Value")
	fmt.Println("------------------------")
	demonstratePassByValue()
	
	// Example 2: Different parameter types
	fmt.Println("\n\nExample 2: Different Parameter Types")
	fmt.Println("-----------------------------------")
	demonstrateParameterTypes()
	
	// Example 3: Parameter evaluation
	fmt.Println("\n\nExample 3: Parameter Evaluation")
	fmt.Println("-------------------------------")
	demonstrateParameterEvaluation()
	
	// Example 4: Empty parameters
	fmt.Println("\n\nExample 4: Functions with No Parameters")
	fmt.Println("--------------------------------------")
	demonstrateNoParameters()
}

// Example 1: Pass by value demonstration
func demonstratePassByValue() {
	// Original values
	coffeeType := "Latte"
	price := 4.50
	quantity := 2
	
	fmt.Printf("Before function call:\n")
	fmt.Printf("  Coffee: %s\n", coffeeType)
	fmt.Printf("  Price: $%.2f\n", price)
	fmt.Printf("  Quantity: %d\n", quantity)
	
	// Call function - parameters are copied
	processOrderByValue(coffeeType, price, quantity)
	
	// Original values unchanged
	fmt.Printf("\nAfter function call:\n")
	fmt.Printf("  Coffee: %s (unchanged)\n", coffeeType)
	fmt.Printf("  Price: $%.2f (unchanged)\n", price)
	fmt.Printf("  Quantity: %d (unchanged)\n", quantity)
}

func processOrderByValue(coffee string, price float64, qty int) {
	fmt.Printf("\nInside function:\n")
	fmt.Printf("  Received: %s, $%.2f, qty: %d\n", coffee, price, qty)
	
	// Modify parameters (only affects local copies)
	coffee = "Modified " + coffee
	price = price * 1.1  // Add 10%
	qty = qty * 2
	
	fmt.Printf("  Modified locally: %s, $%.2f, qty: %d\n", coffee, price, qty)
}

// Example 2: Different parameter types
func demonstrateParameterTypes() {
	// String parameters
	greetCustomer("Marcus", "morning")
	
	// Numeric parameters
	calculateTotal(3, 4.50, 0.08)
	
	// Boolean parameters
	prepareSpecialOrder("Mocha", true, false)
	
	// Mixed types
	generateReceipt("ORD-001", 3, 15.75, true)
}

func greetCustomer(name string, timeOfDay string) {
	greeting := "Good " + timeOfDay
	fmt.Printf("\n%s, %s! Welcome to GoCoffee!\n", greeting, name)
}

func calculateTotal(quantity int, unitPrice float64, taxRate float64) {
	subtotal := float64(quantity) * unitPrice
	tax := subtotal * taxRate
	total := subtotal + tax
	
	fmt.Printf("\nPrice calculation:\n")
	fmt.Printf("  %d items × $%.2f = $%.2f\n", quantity, unitPrice, subtotal)
	fmt.Printf("  Tax (%.0f%%): $%.2f\n", taxRate*100, tax)
	fmt.Printf("  Total: $%.2f\n", total)
}

func prepareSpecialOrder(drink string, extraShot bool, decaf bool) {
	fmt.Printf("\nSpecial order: %s\n", drink)
	
	if extraShot {
		fmt.Println("  → Adding extra shot")
	}
	
	if decaf {
		fmt.Println("  → Making it decaf")
	}
	
	if !extraShot && !decaf {
		fmt.Println("  → Standard preparation")
	}
}

func generateReceipt(orderID string, items int, total float64, paid bool) {
	fmt.Printf("\n--- RECEIPT ---\n")
	fmt.Printf("Order: %s\n", orderID)
	fmt.Printf("Items: %d\n", items)
	fmt.Printf("Total: $%.2f\n", total)
	
	if paid {
		fmt.Println("Status: PAID")
	} else {
		fmt.Println("Status: PENDING")
	}
	fmt.Println("---------------")
}

// Example 3: Parameter evaluation
func demonstrateParameterEvaluation() {
	counter := 0
	
	// Parameters are evaluated before function call
	fmt.Println("Calling function with increment expressions:")
	
	// All increments happen before function is called
	testEvaluation(
		increment(&counter, "first"),
		increment(&counter, "second"),
		increment(&counter, "third"),
	)
	
	fmt.Printf("Final counter value: %d\n", counter)
}

func increment(counter *int, label string) int {
	*counter++
	fmt.Printf("  Evaluating %s parameter (counter now: %d)\n", label, *counter)
	return *counter
}

func testEvaluation(a, b, c int) {
	fmt.Printf("\nInside function - received values: a=%d, b=%d, c=%d\n", a, b, c)
}

// Example 4: Functions with no parameters
func demonstrateNoParameters() {
	// Functions can have empty parameter lists
	showDailySpecial()
	
	fmt.Printf("\nTime check: %s\n", getCurrentTime())
	
	if isShopOpen() {
		fmt.Println("✓ Shop is open!")
	} else {
		fmt.Println("✗ Shop is closed!")
	}
}

func showDailySpecial() {
	fmt.Println("\n Today's Special:")
	fmt.Println("☕ Caramel Macchiato - 20% off!")
}

func getCurrentTime() string {
	// In real code, would use time.Now()
	return "10:30 AM"
}

func isShopOpen() bool {
	// In real code, would check actual time
	return true
}

// Additional examples showing parameter patterns
func parameterPatterns() {
	// Multiple parameters of same type (verbose)
	func(firstName string, lastName string, middleName string) {
		// ...
	}("John", "Doe", "A")
	
	// Multiple parameters of same type (concise)
	func(firstName, lastName, middleName string) {
		// ...
	}("John", "Doe", "A")
	
	// Mixed types require separate declarations
	func(name string, age int, active bool) {
		// ...
	}("Alice", 25, true)
}

// Demonstrating descriptive parameter names
func processPayment(
	customerID string,    // Who is paying
	amount float64,      // How much
	method string,       // Payment method
	applyDiscount bool,  // Whether to apply discount
) {
	// Clear parameter names make function usage obvious
	fmt.Printf("Processing payment for customer %s\n", customerID)
}