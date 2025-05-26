package main

import (
	"fmt"
	"strings"
)

// Function returning a single value
func calculatePrice(basePrice float64, size string) float64 {
	switch size {
	case "small":
		return basePrice * 0.8
	case "medium":
		return basePrice
	case "large":
		return basePrice * 1.3
	default:
		return basePrice
	}
}

// Function returning a string
func getCoffeeDescription(coffeeType string) string {
	switch coffeeType {
	case "Espresso":
		return "Strong, concentrated coffee"
	case "Latte":
		return "Espresso with steamed milk"
	case "Cappuccino":
		return "Espresso with steamed milk foam"
	case "Americano":
		return "Espresso diluted with hot water"
	default:
		return "Unknown coffee type"
	}
}

// Function returning a boolean
func isInStock(item string, quantity int) bool {
	// Simulated inventory
	stock := map[string]int{
		"Coffee Beans": 50,
		"Milk":        30,
		"Sugar":       100,
		"Cups":        200,
	}
	
	available, exists := stock[item]
	if !exists {
		return false
	}
	
	return available >= quantity
}

// Function with early return
func validateOrder(drink string, size string) string {
	// Check if drink is valid
	validDrinks := []string{"Espresso", "Latte", "Cappuccino", "Americano"}
	drinkValid := false
	
	for _, valid := range validDrinks {
		if drink == valid {
			drinkValid = true
			break
		}
	}
	
	if !drinkValid {
		return "Invalid drink selection"
	}
	
	// Check if size is valid
	if size != "small" && size != "medium" && size != "large" {
		return "Invalid size selection"
	}
	
	// Check special combinations
	if drink == "Espresso" && size == "large" {
		return "Espresso only available in small or medium"
	}
	
	return "Order valid"
}

// Function performing calculation and returning result
func calculateChange(paid float64, cost float64) float64 {
	change := paid - cost
	if change < 0 {
		return 0 // Can't have negative change
	}
	return change
}

// Function that builds and returns a value
func createReceipt(items []string, prices []float64) string {
	if len(items) != len(prices) {
		return "Error: Items and prices mismatch"
	}
	
	receipt := "=== GoCoffee Receipt ===\n"
	total := 0.0
	
	for i := 0; i < len(items); i++ {
		receipt += fmt.Sprintf("%-20s $%.2f\n", items[i], prices[i])
		total += prices[i]
	}
	
	receipt += strings.Repeat("-", 28) + "\n"
	receipt += fmt.Sprintf("%-20s $%.2f\n", "Total:", total)
	
	return receipt
}

// Function with conditional returns
func getDiscount(customerType string, purchaseAmount float64) float64 {
	if customerType == "VIP" {
		if purchaseAmount > 50 {
			return 0.20 // 20% discount
		}
		return 0.10 // 10% discount
	}
	
	if customerType == "Regular" && purchaseAmount > 100 {
		return 0.05 // 5% discount
	}
	
	return 0.0 // No discount
}

// Function returning computed value
func estimateWaitTime(ordersAhead int, complexity string) int {
	baseTime := 2 // minutes per order
	
	switch complexity {
	case "simple":
		return ordersAhead * baseTime
	case "medium":
		return ordersAhead * (baseTime + 1)
	case "complex":
		return ordersAhead * (baseTime + 2)
	default:
		return ordersAhead * baseTime
	}
}

func main() {
	fmt.Println("=== GoCoffee Functions with Returns ===\n")
	
	// Using return values
	fmt.Println("--- Price Calculations ---")
	espressoPrice := calculatePrice(3.00, "small")
	lattePrice := calculatePrice(4.50, "large")
	fmt.Printf("Small Espresso: $%.2f\n", espressoPrice)
	fmt.Printf("Large Latte: $%.2f\n", lattePrice)
	
	// String returns
	fmt.Println("\n--- Coffee Descriptions ---")
	desc1 := getCoffeeDescription("Latte")
	desc2 := getCoffeeDescription("Espresso")
	fmt.Printf("Latte: %s\n", desc1)
	fmt.Printf("Espresso: %s\n", desc2)
	
	// Boolean returns
	fmt.Println("\n--- Stock Check ---")
	if isInStock("Coffee Beans", 10) {
		fmt.Println("✓ Coffee beans available")
	} else {
		fmt.Println("✗ Coffee beans out of stock")
	}
	
	// Using return value in condition
	if isInStock("Tea", 5) {
		fmt.Println("✓ Tea available")
	} else {
		fmt.Println("✗ Tea not available")
	}
	
	// Validation with early returns
	fmt.Println("\n--- Order Validation ---")
	result1 := validateOrder("Latte", "medium")
	result2 := validateOrder("Espresso", "large")
	result3 := validateOrder("Tea", "small")
	fmt.Printf("Latte medium: %s\n", result1)
	fmt.Printf("Espresso large: %s\n", result2)
	fmt.Printf("Tea small: %s\n", result3)
	
	// Calculations
	fmt.Println("\n--- Change Calculation ---")
	change := calculateChange(20.00, 15.75)
	fmt.Printf("Paid: $20.00, Cost: $15.75, Change: $%.2f\n", change)
	
	// Building complex returns
	fmt.Println("\n--- Receipt Generation ---")
	items := []string{"Latte", "Croissant", "Espresso"}
	prices := []float64{4.50, 3.25, 2.50}
	receipt := createReceipt(items, prices)
	fmt.Println(receipt)
	
	// Conditional returns
	fmt.Println("--- Discount Calculation ---")
	discount1 := getDiscount("VIP", 75.00)
	discount2 := getDiscount("Regular", 150.00)
	discount3 := getDiscount("Regular", 25.00)
	fmt.Printf("VIP ($75): %.0f%% off\n", discount1*100)
	fmt.Printf("Regular ($150): %.0f%% off\n", discount2*100)
	fmt.Printf("Regular ($25): %.0f%% off\n", discount3*100)
	
	// Using returns in expressions
	fmt.Println("\n--- Wait Time Estimation ---")
	wait := estimateWaitTime(3, "complex")
	fmt.Printf("3 complex orders ahead: %d minutes wait\n", wait)
	
	// Combining function returns
	fmt.Println("\n--- Combined Calculations ---")
	orderPrice := calculatePrice(4.00, "large")
	customerDiscount := getDiscount("VIP", orderPrice)
	finalPrice := orderPrice * (1 - customerDiscount)
	fmt.Printf("Original: $%.2f, After VIP discount: $%.2f\n", orderPrice, finalPrice)
	
	fmt.Println("\n💡 Key Points:")
	fmt.Println("• Functions can return values of any type")
	fmt.Println("• Use return to send a value back to caller")
	fmt.Println("• Return immediately exits the function")
	fmt.Println("• Return values can be used in expressions")
	fmt.Println("• Always handle the returned value appropriately")
}