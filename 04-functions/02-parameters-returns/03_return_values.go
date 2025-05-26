package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println("=== GoCoffee Return Values ===\n")
	
	// Example 1: Single return values
	fmt.Println("Example 1: Single Return Values")
	fmt.Println("-------------------------------")
	demonstrateSingleReturns()
	
	// Example 2: Using returned values
	fmt.Println("\n\nExample 2: Using Returned Values")
	fmt.Println("--------------------------------")
	demonstrateUsingReturns()
	
	// Example 3: Return statements
	fmt.Println("\n\nExample 3: Return Statement Patterns")
	fmt.Println("-----------------------------------")
	demonstrateReturnPatterns()
	
	// Example 4: Zero values
	fmt.Println("\n\nExample 4: Returning Zero Values")
	fmt.Println("--------------------------------")
	demonstrateZeroValues()
}

// Example 1: Different return types
func demonstrateSingleReturns() {
	// String return
	greeting := getGreeting("morning")
	fmt.Printf("Greeting: %s\n", greeting)
	
	// Numeric returns
	price := calculatePrice("Latte", "Large")
	fmt.Printf("Price: $%.2f\n", price)
	
	count := getCustomerCount()
	fmt.Printf("Customers today: %d\n", count)
	
	// Boolean return
	if isHappyHour() {
		fmt.Println("Happy Hour: YES! 20% off!")
	} else {
		fmt.Println("Happy Hour: No")
	}
	
	// Slice return
	specials := getTodaysSpecials()
	fmt.Printf("Today's specials: %v\n", specials)
	
	// Struct return
	coffee := getCoffeeOfTheDay()
	fmt.Printf("Coffee of the day: %+v\n", coffee)
}

func getGreeting(timeOfDay string) string {
	switch timeOfDay {
	case "morning":
		return "Good morning! Fresh coffee awaits!"
	case "afternoon":
		return "Good afternoon! Need an energy boost?"
	case "evening":
		return "Good evening! How about a decaf?"
	default:
		return "Welcome to GoCoffee!"
	}
}

func calculatePrice(coffeeType, size string) float64 {
	basePrice := 3.00
	
	// Adjust for coffee type
	switch coffeeType {
	case "Espresso":
		basePrice = 2.50
	case "Latte", "Cappuccino":
		basePrice = 4.00
	case "Mocha":
		basePrice = 4.50
	}
	
	// Adjust for size
	switch size {
	case "Small":
		return basePrice * 0.8
	case "Medium":
		return basePrice
	case "Large":
		return basePrice * 1.3
	default:
		return basePrice
	}
}

func getCustomerCount() int {
	// In real app, would query database
	return 42
}

func isHappyHour() bool {
	// In real app, would check current time
	// Happy hour: 3-5 PM
	return true // Simulated
}

func getTodaysSpecials() []string {
	return []string{
		"Caramel Macchiato",
		"Blueberry Muffin",
		"Avocado Toast",
	}
}

type Coffee struct {
	Name        string
	Origin      string
	RoastLevel  string
	FlavorNotes []string
}

func getCoffeeOfTheDay() Coffee {
	return Coffee{
		Name:       "Ethiopian Yirgacheffe",
		Origin:     "Ethiopia",
		RoastLevel: "Light",
		FlavorNotes: []string{"Floral", "Citrus", "Tea-like"},
	}
}

// Example 2: Using returned values
func demonstrateUsingReturns() {
	// Direct use
	fmt.Printf("Welcome message: %s\n", generateWelcome("Sarah"))
	
	// Store and use
	total := calculateOrderTotal([]float64{4.50, 3.25, 2.00})
	fmt.Printf("Order total: $%.2f\n", total)
	
	// Use in expressions
	discountedPrice := applyDiscount(10.00, 0.15)
	finalPrice := discountedPrice + calculateTax(discountedPrice)
	fmt.Printf("Final price after discount and tax: $%.2f\n", finalPrice)
	
	// Use in conditions
	if hasEnoughStock("Coffee Beans", 5) {
		fmt.Println("✓ Sufficient stock for order")
	} else {
		fmt.Println("✗ Insufficient stock")
	}
	
	// Chain function calls
	formatted := formatCurrency(
		roundToNearestCent(
			calculatePrice("Latte", "Medium"),
		),
	)
	fmt.Printf("Formatted price: %s\n", formatted)
}

func generateWelcome(name string) string {
	return fmt.Sprintf("Welcome to GoCoffee, %s!", name)
}

func calculateOrderTotal(prices []float64) float64 {
	total := 0.0
	for _, price := range prices {
		total += price
	}
	return total
}

func applyDiscount(price, discountRate float64) float64 {
	return price * (1 - discountRate)
}

func calculateTax(amount float64) float64 {
	const taxRate = 0.08
	return amount * taxRate
}

func hasEnoughStock(item string, needed int) bool {
	// Simulated stock check
	stock := map[string]int{
		"Coffee Beans": 100,
		"Milk":        50,
		"Cups":        200,
	}
	
	available, exists := stock[item]
	return exists && available >= needed
}

func roundToNearestCent(amount float64) float64 {
	return math.Round(amount*100) / 100
}

func formatCurrency(amount float64) string {
	return fmt.Sprintf("$%.2f", amount)
}

// Example 3: Return patterns
func demonstrateReturnPatterns() {
	// Early returns
	result1 := validateAndProcess("", 5)
	fmt.Printf("Empty name result: %s\n", result1)
	
	result2 := validateAndProcess("Alice", -1)
	fmt.Printf("Invalid quantity result: %s\n", result2)
	
	result3 := validateAndProcess("Bob", 3)
	fmt.Printf("Valid order result: %s\n", result3)
	
	// Multiple return points
	grade := getCustomerGrade(1500)
	fmt.Printf("Customer grade (1500 points): %s\n", grade)
	
	// Computed returns
	stats := calculateStats([]float64{10, 20, 30, 40, 50})
	fmt.Printf("Stats: Sum=%.0f, Avg=%.1f, Count=%d\n", 
		stats["sum"], stats["avg"], int(stats["count"]))
}

func validateAndProcess(name string, quantity int) string {
	// Early return for validation
	if name == "" {
		return "Error: Name required"
	}
	
	if quantity <= 0 {
		return "Error: Invalid quantity"
	}
	
	if quantity > 10 {
		return "Error: Quantity too large"
	}
	
	// Process valid order
	return fmt.Sprintf("Order processed for %s: %d items", name, quantity)
}

func getCustomerGrade(points int) string {
	// Multiple return points based on conditions
	if points >= 5000 {
		return "Platinum"
	}
	
	if points >= 2000 {
		return "Gold"
	}
	
	if points >= 1000 {
		return "Silver"
	}
	
	if points >= 500 {
		return "Bronze"
	}
	
	return "Regular"
}

func calculateStats(values []float64) map[string]float64 {
	if len(values) == 0 {
		return map[string]float64{
			"sum":   0,
			"avg":   0,
			"count": 0,
		}
	}
	
	sum := 0.0
	for _, v := range values {
		sum += v
	}
	
	return map[string]float64{
		"sum":   sum,
		"avg":   sum / float64(len(values)),
		"count": float64(len(values)),
	}
}

// Example 4: Zero values
func demonstrateZeroValues() {
	// String zero value
	emptyGreeting := getGreetingForTime(-1)
	fmt.Printf("Invalid time greeting: '%s' (empty string)\n", emptyGreeting)
	
	// Numeric zero values
	noPrice := getPriceForUnknownItem("Mystery Drink")
	fmt.Printf("Unknown item price: $%.2f (zero)\n", noPrice)
	
	// Boolean zero value
	closed := isOpenAt(25) // Invalid hour
	fmt.Printf("Open at hour 25: %v (false)\n", closed)
	
	// Slice zero value
	noItems := getItemsForCategory("Unknown")
	fmt.Printf("Unknown category items: %v (nil slice)\n", noItems)
	
	// Struct zero value
	noCustomer := findCustomer("INVALID-ID")
	fmt.Printf("Invalid customer: %+v (zero struct)\n", noCustomer)
}

func getGreetingForTime(hour int) string {
	if hour < 0 || hour > 23 {
		return "" // Return zero value for string
	}
	
	if hour < 12 {
		return "Good morning!"
	} else if hour < 17 {
		return "Good afternoon!"
	}
	return "Good evening!"
}

func getPriceForUnknownItem(item string) float64 {
	prices := map[string]float64{
		"Coffee": 3.00,
		"Tea":    2.50,
		"Juice":  4.00,
	}
	
	// Returns zero if not found
	return prices[item]
}

func isOpenAt(hour int) bool {
	if hour < 0 || hour > 23 {
		return false // Zero value for bool
	}
	
	// Open 6 AM to 10 PM
	return hour >= 6 && hour <= 22
}

func getItemsForCategory(category string) []string {
	menu := map[string][]string{
		"Hot Drinks":  {"Coffee", "Tea", "Hot Chocolate"},
		"Cold Drinks": {"Iced Coffee", "Juice", "Smoothie"},
		"Food":        {"Croissant", "Muffin", "Sandwich"},
	}
	
	// Returns nil if category not found
	return menu[category]
}

type Customer struct {
	ID     string
	Name   string
	Email  string
	Points int
}

func findCustomer(id string) Customer {
	customers := map[string]Customer{
		"CUST-001": {ID: "CUST-001", Name: "Alice", Email: "alice@email.com", Points: 150},
		"CUST-002": {ID: "CUST-002", Name: "Bob", Email: "bob@email.com", Points: 200},
	}
	
	// Returns zero value Customer if not found
	return customers[id]
}

// Best practices for returns
func returnBestPractices() {
	// Good: Clear what function returns
	func getCustomerName(id string) string {
		return "Alice"
	}
	
	// Good: Consistent return points
	func calculateDiscount(amount float64, customerType string) float64 {
		switch customerType {
		case "VIP":
			return amount * 0.20
		case "Regular":
			return amount * 0.10
		default:
			return 0.0
		}
	}
	
	// Good: Meaningful zero values
	func findItem(id string) (Item, bool) {
		// Return zero Item and false if not found
		return Item{}, false
	}
}

type Item struct {
	ID    string
	Name  string
	Price float64
}