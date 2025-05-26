package main

import "fmt"

// Package-level variables (accessible to all functions in the package)
var shopName = "GoCoffee"
var openHour = 6
var closeHour = 20

// Package-level constant
const taxRate = 0.08

// Function demonstrating local scope
func demonstrateLocalScope() {
	// Local variable - only exists within this function
	greeting := "Welcome to " + shopName
	customerCount := 0
	
	fmt.Println(greeting)
	fmt.Printf("Customers served: %d\n", customerCount)
	
	// Can access package-level variables
	fmt.Printf("We're open from %d:00 to %d:00\n", openHour, closeHour)
}

// Function showing parameter scope
func processOrder(drink string, price float64) {
	// Parameters are local to the function
	// 'drink' and 'price' only exist here
	
	// Local variable
	discount := 0.1
	
	// Calculate using parameters and local variable
	finalPrice := price * (1 - discount)
	
	fmt.Printf("Order: %s\n", drink)
	fmt.Printf("Original price: $%.2f\n", price)
	fmt.Printf("After discount: $%.2f\n", finalPrice)
	
	// This modifies the local copy, not any external variable
	price = finalPrice
	fmt.Printf("Local price variable: $%.2f\n", price)
}

// Function with block scope
func checkInventory(item string) {
	available := false
	quantity := 0
	
	// Block scope in if statement
	if item == "Coffee Beans" {
		// Variables declared here only exist in this block
		minStock := 10
		currentStock := 50
		
		if currentStock > minStock {
			available = true
			quantity = currentStock
		}
		
		fmt.Printf("Minimum stock level: %d\n", minStock)
	}
	// minStock and currentStock don't exist here
	
	if available {
		fmt.Printf("%s: %d units available\n", item, quantity)
	} else {
		fmt.Printf("%s: Out of stock\n", item)
	}
}

// Function showing scope shadowing
func demonstrateShadowing() {
	// Package-level shopName is "GoCoffee"
	fmt.Printf("Package-level shop: %s\n", shopName)
	
	// Local variable shadows package-level variable
	shopName := "GoCoffee Express"
	fmt.Printf("Local shop (shadowed): %s\n", shopName)
	
	// Create a new scope
	{
		// Another level of shadowing
		shopName := "GoCoffee Mobile"
		fmt.Printf("Block-level shop: %s\n", shopName)
	}
	
	// Back to function-level shadow
	fmt.Printf("Function-level shop: %s\n", shopName)
}

// Function that modifies package-level variable
func updateShopHours(newOpen, newClose int) {
	fmt.Printf("Updating hours from %d-%d to %d-%d\n", 
		openHour, closeHour, newOpen, newClose)
	
	// This modifies the package-level variables
	openHour = newOpen
	closeHour = newClose
}

// Function showing loop scope
func dailySalesReport() {
	totalSales := 0.0
	
	// Loop variable scope
	for hour := openHour; hour < closeHour; hour++ {
		// 'hour' only exists within this loop
		hourlySales := float64(hour * 10) // Simulated sales
		totalSales += hourlySales
		
		if hour == 12 {
			// Lunch rush bonus
			lunchBonus := 50.0
			totalSales += lunchBonus
			fmt.Printf("Hour %d: $%.2f (includes lunch bonus: $%.2f)\n", 
				hour, hourlySales+lunchBonus, lunchBonus)
		}
	}
	// 'hour' and 'lunchBonus' don't exist here
	
	fmt.Printf("Total daily sales: $%.2f\n", totalSales)
}

// Nested function calls and scope
func outer() {
	outerVar := "I'm from outer"
	fmt.Printf("Outer function: %s\n", outerVar)
	
	// Call inner function
	inner("Hello from outer")
}

func inner(message string) {
	// Can't access outerVar here - it's not in scope
	innerVar := "I'm from inner"
	fmt.Printf("Inner function: %s, %s\n", message, innerVar)
}

// Function showing scope lifetime
func demonstrateLifetime() {
	fmt.Println("\n--- Demonstrating Variable Lifetime ---")
	
	for i := 0; i < 3; i++ {
		// This variable is created and destroyed each iteration
		iterationMessage := fmt.Sprintf("Iteration %d", i)
		fmt.Println(iterationMessage)
		
		// Demonstrating scope in switch
		switch i {
		case 0:
			first := "First iteration"
			fmt.Printf("  %s\n", first)
		case 1:
			second := "Second iteration"
			fmt.Printf("  %s\n", second)
		}
		// 'first' and 'second' don't exist here
	}
	// 'iterationMessage' and 'i' don't exist here
}

func main() {
	fmt.Println("=== GoCoffee Function Scope ===\n")
	
	// Demonstrate local scope
	fmt.Println("--- Local Scope ---")
	demonstrateLocalScope()
	// Can't access 'greeting' or 'customerCount' here
	
	// Parameter scope
	fmt.Println("\n--- Parameter Scope ---")
	originalPrice := 5.00
	processOrder("Latte", originalPrice)
	fmt.Printf("Original price unchanged: $%.2f\n", originalPrice)
	
	// Block scope
	fmt.Println("\n--- Block Scope ---")
	checkInventory("Coffee Beans")
	checkInventory("Tea")
	
	// Shadowing
	fmt.Println("\n--- Variable Shadowing ---")
	demonstrateShadowing()
	fmt.Printf("Package-level shop unchanged: %s\n", shopName)
	
	// Modifying package-level variables
	fmt.Println("\n--- Package-Level Modifications ---")
	fmt.Printf("Current hours: %d:00 - %d:00\n", openHour, closeHour)
	updateShopHours(7, 21)
	fmt.Printf("New hours: %d:00 - %d:00\n", openHour, closeHour)
	
	// Loop scope
	fmt.Println("\n--- Loop Scope ---")
	dailySalesReport()
	
	// Nested function scope
	fmt.Println("\n--- Nested Function Calls ---")
	outer()
	
	// Variable lifetime
	demonstrateLifetime()
	
	fmt.Println("\n💡 Key Scope Rules:")
	fmt.Println("• Variables are only accessible within their scope")
	fmt.Println("• Inner scopes can access outer scope variables")
	fmt.Println("• Local variables shadow outer variables with same name")
	fmt.Println("• Function parameters are local to the function")
	fmt.Println("• Block statements create new scopes")
	fmt.Println("• Variables are destroyed when their scope ends")
}