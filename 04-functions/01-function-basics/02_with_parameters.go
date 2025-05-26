package main

import (
	"fmt"
	"strings"
)

// Function with one parameter
func greetCustomer(name string) {
	fmt.Printf("Hello, %s! Welcome to GoCoffee!\n", name)
}

// Function with multiple parameters
func makeCoffee(coffeeType string, size string) {
	fmt.Printf("🎯 Making one %s %s\n", size, coffeeType)
}

// Function with parameters of same type
func mixIngredients(ingredient1, ingredient2 string) {
	fmt.Printf("   Mixing %s with %s\n", ingredient1, ingredient2)
}

// Function with different parameter types
func prepareOrder(drink string, quantity int, temperature float64) {
	fmt.Printf("\n📝 Order Details:\n")
	fmt.Printf("   Drink: %s\n", drink)
	fmt.Printf("   Quantity: %d\n", quantity)
	fmt.Printf("   Temperature: %.1f°C\n", temperature)
}

// Function that modifies its parameters (but only locally!)
func customizeCoffee(coffee string, addSugar bool, sugarCount int) {
	result := coffee
	
	if addSugar && sugarCount > 0 {
		result += fmt.Sprintf(" with %d sugar", sugarCount)
		if sugarCount > 1 {
			result = strings.Replace(result, "sugar", "sugars", 1)
		}
	}
	
	fmt.Printf("☕ Prepared: %s\n", result)
}

// Function demonstrating parameter usage
func calculateBill(itemPrice float64, quantity int, taxRate float64) {
	subtotal := itemPrice * float64(quantity)
	tax := subtotal * taxRate
	total := subtotal + tax
	
	fmt.Printf("\n💵 Bill Calculation:\n")
	fmt.Printf("   Item Price: $%.2f\n", itemPrice)
	fmt.Printf("   Quantity: %d\n", quantity)
	fmt.Printf("   Subtotal: $%.2f\n", subtotal)
	fmt.Printf("   Tax (%.1f%%): $%.2f\n", taxRate*100, tax)
	fmt.Printf("   Total: $%.2f\n", total)
}

// Function with descriptive parameter names
func serveCoffee(customerName string, tableNumber int, coffeeType string) {
	fmt.Printf("\n🏃 Service Alert:\n")
	fmt.Printf("   Delivering %s to %s at table %d\n", 
		coffeeType, customerName, tableNumber)
}

// Function showing parameter evaluation
func processPayment(amount float64, paymentMethod string) {
	fmt.Printf("\n💳 Processing Payment:\n")
	fmt.Printf("   Amount: $%.2f\n", amount)
	fmt.Printf("   Method: %s\n", paymentMethod)
	
	// Parameters are evaluated when function is called
	if paymentMethod == "cash" {
		fmt.Println("   → Handling cash payment")
	} else if paymentMethod == "card" {
		fmt.Println("   → Processing card payment")
	} else {
		fmt.Println("   → Unknown payment method")
	}
}

func main() {
	fmt.Println("=== GoCoffee Functions with Parameters ===\n")
	
	// Calling function with one parameter
	greetCustomer("Marcus")
	greetCustomer("Sarah")
	
	// Calling function with multiple parameters
	fmt.Println("\n--- Coffee Orders ---")
	makeCoffee("Latte", "Large")
	makeCoffee("Espresso", "Small")
	makeCoffee("Cappuccino", "Medium")
	
	// Parameters of same type
	fmt.Println("\n--- Mixing Ingredients ---")
	mixIngredients("espresso", "steamed milk")
	mixIngredients("coffee", "chocolate")
	
	// Different parameter types
	prepareOrder("Mocha", 2, 65.5)
	
	// Boolean and numeric parameters
	fmt.Println("\n--- Coffee Customization ---")
	customizeCoffee("Latte", true, 2)
	customizeCoffee("Espresso", false, 0)
	customizeCoffee("Cappuccino", true, 1)
	
	// Complex calculations
	calculateBill(4.50, 3, 0.08)
	
	// Descriptive parameters
	serveCoffee("Alice", 5, "Caramel Macchiato")
	
	// Different payment methods
	processPayment(12.75, "card")
	processPayment(8.50, "cash")
	
	// Parameters are expressions
	fmt.Println("\n--- Parameter Expressions ---")
	basePrice := 3.50
	quantity := 2
	makeCoffee("House Blend", "Medium")
	calculateBill(basePrice*1.1, quantity+1, 0.08)
	
	fmt.Println("\n💡 Key Points:")
	fmt.Println("• Parameters allow functions to accept input")
	fmt.Println("• Multiple parameters are separated by commas")
	fmt.Println("• Parameters of same type can share type declaration")
	fmt.Println("• Parameters are passed by value (copied)")
	fmt.Println("• Use descriptive parameter names")
}