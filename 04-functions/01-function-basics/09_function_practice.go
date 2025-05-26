package main

import (
	"fmt"
	"strings"
	"time"
)

// This file contains practice exercises for functions
// Try to complete each challenge!

func main() {
	fmt.Println("=== GoCoffee Function Practice ===\n")
	
	// Run all challenges
	challenge1()
	challenge2()
	challenge3()
	challenge4()
	challenge5()
	
	// Final challenge
	finalChallenge()
}

// ============================================
// Challenge 1: Basic Functions
// ============================================

func challenge1() {
	fmt.Println("📝 Challenge 1: Basic Functions")
	fmt.Println(strings.Repeat("-", 40))
	fmt.Println("Create functions for coffee shop greetings\n")
	
	// TODO: Create these functions:
	// 1. greetMorningCustomer(name string) - prints morning greeting
	// 2. greetAfternoonCustomer(name string) - prints afternoon greeting  
	// 3. greetEveningCustomer(name string) - prints evening greeting
	
	// Test your functions
	greetMorningCustomer("Alice")
	greetAfternoonCustomer("Bob")
	greetEveningCustomer("Charlie")
	
	fmt.Println()
}

// Solution 1: Greeting functions
func greetMorningCustomer(name string) {
	fmt.Printf("☀️ Good morning, %s! Perfect day for coffee!\n", name)
}

func greetAfternoonCustomer(name string) {
	fmt.Printf("🌤️ Good afternoon, %s! Need an energy boost?\n", name)
}

func greetEveningCustomer(name string) {
	fmt.Printf("🌙 Good evening, %s! How about a decaf?\n", name)
}

// ============================================
// Challenge 2: Functions with Returns
// ============================================

func challenge2() {
	fmt.Println("📝 Challenge 2: Functions with Returns")
	fmt.Println(strings.Repeat("-", 40))
	fmt.Println("Create price calculation functions\n")
	
	// TODO: Create these functions:
	// 1. getCoffeeBasePrice(coffeeType string) float64
	// 2. getSizeMultiplier(size string) float64
	// 3. calculateFinalPrice(basePrice, multiplier float64) float64
	
	// Test your functions
	coffeeType := "Latte"
	size := "Large"
	
	basePrice := getCoffeeBasePrice(coffeeType)
	multiplier := getSizeMultiplier(size)
	finalPrice := calculateFinalPrice(basePrice, multiplier)
	
	fmt.Printf("%s %s: $%.2f\n", size, coffeeType, finalPrice)
	fmt.Println()
}

// Solution 2: Price calculation functions
func getCoffeeBasePrice(coffeeType string) float64 {
	prices := map[string]float64{
		"Espresso":   2.50,
		"Americano":  3.00,
		"Latte":      4.00,
		"Cappuccino": 3.50,
		"Mocha":      4.50,
	}
	
	if price, exists := prices[coffeeType]; exists {
		return price
	}
	return 3.00 // default price
}

func getSizeMultiplier(size string) float64 {
	switch strings.ToLower(size) {
	case "small":
		return 0.8
	case "medium":
		return 1.0
	case "large":
		return 1.2
	default:
		return 1.0
	}
}

func calculateFinalPrice(basePrice, multiplier float64) float64 {
	return basePrice * multiplier
}

// ============================================
// Challenge 3: Functions with Multiple Tasks
// ============================================

func challenge3() {
	fmt.Println("📝 Challenge 3: Order Processing")
	fmt.Println(strings.Repeat("-", 40))
	fmt.Println("Create a complete order processing system\n")
	
	// TODO: Create these functions:
	// 1. takeOrder(customer string, drink string, size string)
	// 2. prepareOrder(drink string) 
	// 3. serveOrder(customer string, drink string)
	// 4. processPayment(amount float64, payment float64) float64
	
	// Simulate an order
	customer := "David"
	drink := "Cappuccino"
	size := "Medium"
	
	takeOrder(customer, drink, size)
	prepareOrder(drink)
	serveOrder(customer, drink)
	
	// Calculate payment
	price := calculateFinalPrice(getCoffeeBasePrice(drink), getSizeMultiplier(size))
	change := processPayment(price, 10.00)
	
	if change >= 0 {
		fmt.Printf("Change: $%.2f\n", change)
	}
	
	fmt.Println()
}

// Solution 3: Order processing functions
func takeOrder(customer string, drink string, size string) {
	fmt.Printf("📝 Taking order: %s %s for %s\n", size, drink, customer)
}

func prepareOrder(drink string) {
	fmt.Printf("☕ Preparing %s", drink)
	
	// Simulate preparation time
	for i := 0; i < 3; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Print(".")
	}
	fmt.Println(" Ready!")
}

func serveOrder(customer string, drink string) {
	fmt.Printf("🤝 Serving %s to %s\n", drink, customer)
}

func processPayment(amount float64, payment float64) float64 {
	fmt.Printf("💰 Amount due: $%.2f, Payment: $%.2f\n", amount, payment)
	
	if payment < amount {
		fmt.Printf("❌ Insufficient payment! Need $%.2f more\n", amount-payment)
		return -1
	}
	
	return payment - amount
}

// ============================================
// Challenge 4: Helper Functions
// ============================================

func challenge4() {
	fmt.Println("📝 Challenge 4: Helper Functions")
	fmt.Println(strings.Repeat("-", 40))
	fmt.Println("Create helper functions for common tasks\n")
	
	// TODO: Create these helper functions:
	// 1. isWeekend() bool - check if today is weekend
	// 2. formatReceipt(items []string, prices []float64) string
	// 3. applyHappyHourDiscount(price float64, hour int) float64
	
	// Test your functions
	if isWeekend() {
		fmt.Println("🎉 It's the weekend! Special menu available!")
	} else {
		fmt.Println("📅 Weekday menu active")
	}
	
	// Test receipt
	items := []string{"Latte", "Croissant"}
	prices := []float64{4.00, 3.50}
	receipt := formatReceipt(items, prices)
	fmt.Print(receipt)
	
	// Test happy hour
	regularPrice := 5.00
	happyHourPrice := applyHappyHourDiscount(regularPrice, 15) // 3 PM
	fmt.Printf("\nRegular: $%.2f, Happy Hour (3-5 PM): $%.2f\n", 
		regularPrice, happyHourPrice)
	
	fmt.Println()
}

// Solution 4: Helper functions
func isWeekend() bool {
	today := time.Now().Weekday()
	return today == time.Saturday || today == time.Sunday
}

func formatReceipt(items []string, prices []float64) string {
	if len(items) != len(prices) {
		return "Error: Invalid receipt data\n"
	}
	
	receipt := "\n--- Receipt ---\n"
	total := 0.0
	
	for i, item := range items {
		receipt += fmt.Sprintf("%-15s $%.2f\n", item, prices[i])
		total += prices[i]
	}
	
	receipt += strings.Repeat("-", 23) + "\n"
	receipt += fmt.Sprintf("%-15s $%.2f\n", "Total:", total)
	
	return receipt
}

func applyHappyHourDiscount(price float64, hour int) float64 {
	// Happy hour: 3-5 PM (15-17)
	if hour >= 15 && hour < 17 {
		return price * 0.8 // 20% off
	}
	return price
}

// ============================================
// Challenge 5: Function Organization
// ============================================

func challenge5() {
	fmt.Println("📝 Challenge 5: Function Organization")
	fmt.Println(strings.Repeat("-", 40))
	fmt.Println("Organize functions by purpose\n")
	
	// TODO: Create organized function groups:
	// 1. Inventory functions (checkStock, updateStock)
	// 2. Customer functions (addLoyaltyPoints, checkLoyaltyStatus)
	// 3. Reporting functions (dailySummary, topSellingItems)
	
	// Test inventory
	if checkStock("Coffee Beans") {
		fmt.Println("✓ Coffee beans in stock")
	}
	updateStock("Coffee Beans", -10)
	
	// Test customer loyalty
	points := addLoyaltyPoints("Emma", 50)
	status := checkLoyaltyStatus(points)
	fmt.Printf("Emma has %d points (%s)\n", points, status)
	
	// Test reporting
	dailySummary()
	
	fmt.Println()
}

// Solution 5: Organized functions

// --- Inventory Functions ---
func checkStock(item string) bool {
	// Simulated stock check
	stock := map[string]int{
		"Coffee Beans": 100,
		"Milk":         50,
		"Sugar":        200,
	}
	
	if quantity, exists := stock[item]; exists {
		return quantity > 0
	}
	return false
}

func updateStock(item string, change int) {
	// In real app, would update database
	fmt.Printf("📦 Stock update: %s %+d units\n", item, change)
}

// --- Customer Functions ---
func addLoyaltyPoints(customer string, points int) int {
	// Simulated loyalty system
	currentPoints := 150 // Would fetch from database
	newPoints := currentPoints + points
	fmt.Printf("🌟 Added %d points for %s\n", points, customer)
	return newPoints
}

func checkLoyaltyStatus(points int) string {
	switch {
	case points >= 1000:
		return "Platinum"
	case points >= 500:
		return "Gold"
	case points >= 100:
		return "Silver"
	default:
		return "Bronze"
	}
}

// --- Reporting Functions ---
func dailySummary() {
	fmt.Println("\n📊 Daily Summary:")
	fmt.Println("  Orders: 127")
	fmt.Println("  Revenue: $543.75")
	fmt.Println("  Top item: Latte (34 sold)")
}

// ============================================
// Final Challenge: Build a Complete System
// ============================================

func finalChallenge() {
	fmt.Println("🏆 Final Challenge: Coffee Shop Simulator")
	fmt.Println(strings.Repeat("=", 40))
	fmt.Println("Build a complete coffee shop order system!\n")
	
	// TODO: Combine everything you've learned to:
	// 1. Take a customer order
	// 2. Validate the order
	// 3. Calculate the price
	// 4. Process payment
	// 5. Update inventory
	// 6. Add loyalty points
	// 7. Print a receipt
	
	// Simulate a complete transaction
	simulateCompleteOrder()
	
	fmt.Println("\n🎉 Congratulations! You've completed the function basics!")
	fmt.Println("\n💡 Key Takeaways:")
	fmt.Println("• Functions organize code into reusable blocks")
	fmt.Println("• Use descriptive names for functions")
	fmt.Println("• Keep functions focused on one task")
	fmt.Println("• Group related functions together")
	fmt.Println("• Test functions with different inputs")
}

// Final challenge solution
func simulateCompleteOrder() {
	// Customer details
	customer := "Frank"
	drink := "Mocha"
	size := "Large"
	payment := 10.00
	
	fmt.Printf("👤 Customer: %s\n\n", customer)
	
	// Take and validate order
	if !validateDrink(drink) {
		fmt.Println("❌ Invalid drink selection")
		return
	}
	
	takeOrder(customer, drink, size)
	
	// Calculate price
	basePrice := getCoffeeBasePrice(drink)
	sizeMultiplier := getSizeMultiplier(size)
	price := calculateFinalPrice(basePrice, sizeMultiplier)
	
	// Apply discounts
	hour := time.Now().Hour()
	price = applyHappyHourDiscount(price, hour)
	
	fmt.Printf("\n💵 Price: $%.2f\n", price)
	
	// Process payment
	change := processPayment(price, payment)
	if change < 0 {
		return
	}
	
	// Prepare and serve
	prepareOrder(drink)
	serveOrder(customer, drink)
	
	// Update systems
	updateStock("Coffee Beans", -18)
	points := addLoyaltyPoints(customer, int(price*10))
	
	// Print receipt
	fmt.Println("\n" + generateFinalReceipt(customer, drink, size, price, payment, change, points))
}

func validateDrink(drink string) bool {
	validDrinks := []string{"Espresso", "Americano", "Latte", "Cappuccino", "Mocha"}
	for _, valid := range validDrinks {
		if drink == valid {
			return true
		}
	}
	return false
}

func generateFinalReceipt(customer, drink, size string, price, payment, change float64, points int) string {
	receipt := "========== RECEIPT ==========\n"
	receipt += fmt.Sprintf("Customer: %s\n", customer)
	receipt += fmt.Sprintf("Time: %s\n", time.Now().Format("3:04 PM"))
	receipt += strings.Repeat("-", 28) + "\n"
	receipt += fmt.Sprintf("%-15s %s\n", drink, size)
	receipt += fmt.Sprintf("%-15s $%.2f\n", "Price:", price)
	receipt += strings.Repeat("-", 28) + "\n"
	receipt += fmt.Sprintf("%-15s $%.2f\n", "Payment:", payment)
	receipt += fmt.Sprintf("%-15s $%.2f\n", "Change:", change)
	receipt += strings.Repeat("-", 28) + "\n"
	receipt += fmt.Sprintf("Loyalty Points: %d\n", points)
	receipt += "\nThank you for visiting GoCoffee!\n"
	receipt += "============================\n"
	
	return receipt
}