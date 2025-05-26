package main

import (
	"fmt"
	"strings"
	"time"
)

// Package-level constants
const (
	ShopName    = "GoCoffee"           // Exported (capital letter)
	maxCapacity = 50                   // Unexported (lowercase)
	version     = "1.0.0"              // Unexported
)

// Package-level variables
var (
	TotalSales    float64              // Exported - accessible from other packages
	customerCount int                  // Unexported - only in this package
	isOpen        bool                 // Unexported
)

// Exported function - can be used by other packages
func OpenShop() {
	isOpen = true
	customerCount = 0
	fmt.Printf("%s is now OPEN! 🎉\n", ShopName)
	logEvent("Shop opened")
}

// Exported function with parameters
func ServeCoffee(customerName string, coffeeType string) string {
	if !isOpen {
		return "Sorry, we're closed!"
	}
	
	customerCount++
	price := getCoffeePrice(coffeeType)
	TotalSales += price
	
	logEvent(fmt.Sprintf("Served %s to %s", coffeeType, customerName))
	return fmt.Sprintf("Here's your %s, %s! That'll be $%.2f", 
		coffeeType, customerName, price)
}

// Unexported function - only available within this package
func getCoffeePrice(coffeeType string) float64 {
	prices := map[string]float64{
		"Espresso":   2.50,
		"Latte":      4.50,
		"Cappuccino": 4.00,
		"Americano":  3.00,
	}
	
	if price, exists := prices[coffeeType]; exists {
		return price
	}
	return 3.50 // Default price
}

// Unexported helper function
func logEvent(event string) {
	timestamp := time.Now().Format("15:04:05")
	fmt.Printf("[%s] %s\n", timestamp, event)
}

// Exported function showing shop status
func GetShopStatus() string {
	if !isOpen {
		return fmt.Sprintf("%s is CLOSED", ShopName)
	}
	
	occupancy := float64(customerCount) / float64(maxCapacity) * 100
	return fmt.Sprintf("%s is OPEN - Customers: %d/%d (%.0f%% full)", 
		ShopName, customerCount, maxCapacity, occupancy)
}

// Exported function with validation
func ProcessPayment(amount float64, method string) (bool, string) {
	validMethods := []string{"cash", "card", "mobile"}
	methodValid := false
	
	for _, valid := range validMethods {
		if strings.ToLower(method) == valid {
			methodValid = true
			break
		}
	}
	
	if !methodValid {
		return false, "Invalid payment method"
	}
	
	if amount <= 0 {
		return false, "Invalid amount"
	}
	
	// Process payment (simplified)
	logEvent(fmt.Sprintf("Payment processed: $%.2f via %s", amount, method))
	return true, "Payment successful"
}

// Unexported function for internal calculations
func calculateTax(amount float64) float64 {
	const taxRate = 0.08
	return amount * taxRate
}

// Exported function using unexported functions
func GenerateReceipt(items []string, prices []float64) string {
	if len(items) != len(prices) {
		return "Error: Invalid receipt data"
	}
	
	receipt := formatHeader()
	subtotal := 0.0
	
	for i, item := range items {
		receipt += fmt.Sprintf("%-20s $%6.2f\n", item, prices[i])
		subtotal += prices[i]
	}
	
	tax := calculateTax(subtotal)
	total := subtotal + tax
	
	receipt += formatFooter(subtotal, tax, total)
	return receipt
}

// Unexported formatting helpers
func formatHeader() string {
	return fmt.Sprintf("\n%s RECEIPT\n%s\n", 
		ShopName, strings.Repeat("=", 30))
}

func formatFooter(subtotal, tax, total float64) string {
	footer := strings.Repeat("-", 30) + "\n"
	footer += fmt.Sprintf("%-20s $%6.2f\n", "Subtotal:", subtotal)
	footer += fmt.Sprintf("%-20s $%6.2f\n", "Tax:", tax)
	footer += fmt.Sprintf("%-20s $%6.2f\n", "Total:", total)
	return footer
}

// Exported function demonstrating init pattern
func Initialize() {
	fmt.Printf("Initializing %s v%s...\n", ShopName, version)
	resetDailyTotals()
	loadConfiguration()
	fmt.Println("Initialization complete!")
}

// Unexported initialization helpers
func resetDailyTotals() {
	TotalSales = 0
	customerCount = 0
	logEvent("Daily totals reset")
}

func loadConfiguration() {
	// Simulate loading config
	logEvent("Configuration loaded")
}

// Exported function for closing
func CloseShop() string {
	if !isOpen {
		return "Shop is already closed"
	}
	
	isOpen = false
	summary := fmt.Sprintf("\nClosing %s:\n", ShopName)
	summary += fmt.Sprintf("- Customers served: %d\n", customerCount)
	summary += fmt.Sprintf("- Total sales: $%.2f\n", TotalSales)
	summary += fmt.Sprintf("- Average per customer: $%.2f\n", 
		TotalSales/float64(customerCount))
	
	logEvent("Shop closed")
	return summary
}

// Main function demonstrating package functions
func main() {
	fmt.Println("=== GoCoffee Package-Level Functions ===")
	
	// Initialize the shop
	Initialize()
	
	// Open the shop
	fmt.Println()
	OpenShop()
	
	// Serve some customers
	fmt.Println("\n--- Serving Customers ---")
	fmt.Println(ServeCoffee("Alice", "Latte"))
	fmt.Println(ServeCoffee("Bob", "Espresso"))
	fmt.Println(ServeCoffee("Charlie", "Cappuccino"))
	
	// Check status
	fmt.Println("\n--- Shop Status ---")
	fmt.Println(GetShopStatus())
	
	// Process payments
	fmt.Println("\n--- Processing Payments ---")
	success, msg := ProcessPayment(4.50, "card")
	fmt.Printf("Payment 1: %v - %s\n", success, msg)
	
	success, msg = ProcessPayment(2.50, "bitcoin")
	fmt.Printf("Payment 2: %v - %s\n", success, msg)
	
	// Generate receipt
	fmt.Println("\n--- Receipt Generation ---")
	items := []string{"Latte", "Espresso", "Cappuccino"}
	prices := []float64{4.50, 2.50, 4.00}
	receipt := GenerateReceipt(items, prices)
	fmt.Println(receipt)
	
	// Close shop
	fmt.Println(CloseShop())
	
	// Try to serve after closing
	fmt.Println("\n--- After Closing ---")
	fmt.Println(ServeCoffee("David", "Americano"))
	
	fmt.Println("\n💡 Package Function Guidelines:")
	fmt.Println("• Exported functions start with capital letters")
	fmt.Println("• Unexported functions start with lowercase")
	fmt.Println("• Group related functions together")
	fmt.Println("• Use unexported helpers for internal logic")
	fmt.Println("• Exported functions form your package's API")
	fmt.Println("• Document all exported functions")
}