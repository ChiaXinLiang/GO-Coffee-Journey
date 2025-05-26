package main

import (
	"fmt"
	"math"
	"strings"
	"time"
)

func main() {
	fmt.Println("=== GoCoffee Multiple Return Values ===\n")
	
	// Example 1: Basic multiple returns
	fmt.Println("Example 1: Basic Multiple Returns")
	fmt.Println("---------------------------------")
	demonstrateBasicMultipleReturns()
	
	// Example 2: Common patterns
	fmt.Println("\n\nExample 2: Common Multiple Return Patterns")
	fmt.Println("-----------------------------------------")
	demonstrateCommonPatterns()
	
	// Example 3: Error handling pattern
	fmt.Println("\n\nExample 3: Error Handling with Multiple Returns")
	fmt.Println("-----------------------------------------------")
	demonstrateErrorHandling()
	
	// Example 4: Complex returns
	fmt.Println("\n\nExample 4: Complex Multiple Returns")
	fmt.Println("-----------------------------------")
	demonstrateComplexReturns()
}

// Example 1: Basic multiple returns
func demonstrateBasicMultipleReturns() {
	// Two returns
	quotient, remainder := divide(17, 5)
	fmt.Printf("17 ÷ 5 = %d remainder %d\n", quotient, remainder)
	
	// Three returns
	min, max, avg := getStats([]float64{2.5, 4.0, 3.5, 5.0, 1.5})
	fmt.Printf("Stats: min=%.1f, max=%.1f, avg=%.1f\n", min, max, avg)
	
	// Different types
	name, age, isMember := getCustomerInfo("CUST-001")
	fmt.Printf("Customer: %s, Age: %d, Member: %v\n", name, age, isMember)
	
	// Using only some returns
	total, _ := calculateTotalAndTax(45.50, 0.08)
	fmt.Printf("Total (ignoring tax breakdown): $%.2f\n", total)
}

func divide(a, b int) (int, int) {
	return a / b, a % b
}

func getStats(values []float64) (min, max, avg float64) {
	if len(values) == 0 {
		return 0, 0, 0
	}
	
	min = values[0]
	max = values[0]
	sum := 0.0
	
	for _, v := range values {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
		sum += v
	}
	
	avg = sum / float64(len(values))
	return min, max, avg
}

func getCustomerInfo(id string) (string, int, bool) {
	// Simulated database lookup
	customers := map[string]struct {
		name   string
		age    int
		member bool
	}{
		"CUST-001": {"Alice Smith", 28, true},
		"CUST-002": {"Bob Jones", 35, false},
	}
	
	if customer, exists := customers[id]; exists {
		return customer.name, customer.age, customer.member
	}
	
	return "", 0, false
}

func calculateTotalAndTax(subtotal, taxRate float64) (float64, float64) {
	tax := subtotal * taxRate
	total := subtotal + tax
	return total, tax
}

// Example 2: Common patterns
func demonstrateCommonPatterns() {
	// Pattern 1: Value and existence
	price, exists := getPrice("Latte")
	if exists {
		fmt.Printf("Latte price: $%.2f\n", price)
	} else {
		fmt.Println("Latte price not found")
	}
	
	// Pattern 2: Result and error
	result, err := processOrder("ORD-123", 3)
	if err != nil {
		fmt.Printf("Order error: %v\n", err)
	} else {
		fmt.Printf("Order result: %s\n", result)
	}
	
	// Pattern 3: Value and success flag
	coffee, ok := parseOrderString("Large Latte with Extra Shot")
	if ok {
		fmt.Printf("Parsed order: %+v\n", coffee)
	} else {
		fmt.Println("Failed to parse order")
	}
	
	// Pattern 4: Multiple related values
	width, height, area := getRoomDimensions("Main Dining")
	fmt.Printf("Room dimensions: %.1f × %.1f = %.1f sq ft\n", 
		width, height, area)
}

func getPrice(item string) (float64, bool) {
	prices := map[string]float64{
		"Coffee":     3.00,
		"Latte":      4.50,
		"Cappuccino": 4.00,
		"Espresso":   2.50,
	}
	
	price, exists := prices[item]
	return price, exists
}

func processOrder(orderID string, quantity int) (string, error) {
	if orderID == "" {
		return "", fmt.Errorf("empty order ID")
	}
	
	if quantity <= 0 {
		return "", fmt.Errorf("invalid quantity: %d", quantity)
	}
	
	if quantity > 10 {
		return "", fmt.Errorf("quantity too large: %d", quantity)
	}
	
	// Process order
	result := fmt.Sprintf("Order %s processed: %d items", orderID, quantity)
	return result, nil
}

type CoffeeOrder struct {
	Size   string
	Type   string
	Extras []string
}

func parseOrderString(order string) (CoffeeOrder, bool) {
	parts := strings.Fields(strings.ToLower(order))
	if len(parts) < 2 {
		return CoffeeOrder{}, false
	}
	
	coffee := CoffeeOrder{
		Size: strings.Title(parts[0]),
		Type: strings.Title(parts[1]),
	}
	
	// Look for extras
	if strings.Contains(order, "Extra Shot") {
		coffee.Extras = append(coffee.Extras, "Extra Shot")
	}
	
	return coffee, true
}

func getRoomDimensions(room string) (float64, float64, float64) {
	dimensions := map[string][2]float64{
		"Main Dining":  {20.5, 15.0},
		"Private Room": {12.0, 10.0},
		"Patio":        {25.0, 18.5},
	}
	
	if dims, exists := dimensions[room]; exists {
		width := dims[0]
		height := dims[1]
		area := width * height
		return width, height, area
	}
	
	return 0, 0, 0
}

// Example 3: Error handling
func demonstrateErrorHandling() {
	// Success case
	receipt1, err := generateReceipt("ORD-001", []string{"Latte", "Muffin"}, []float64{4.50, 3.25})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("Receipt generated:")
		fmt.Println(receipt1)
	}
	
	// Error case - mismatched lengths
	receipt2, err := generateReceipt("ORD-002", []string{"Coffee"}, []float64{3.00, 2.50})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println(receipt2)
	}
	
	// Multiple error checks
	coffee, preparationTime, err := prepareCoffeeOrder("Mocha", "Large")
	if err != nil {
		fmt.Printf("Cannot prepare order: %v\n", err)
	} else {
		fmt.Printf("Prepared: %s (took %v)\n", coffee, preparationTime)
	}
	
	// Chained error handling
	processCompleteOrder()
}

func generateReceipt(orderID string, items []string, prices []float64) (string, error) {
	if len(items) != len(prices) {
		return "", fmt.Errorf("items and prices length mismatch")
	}
	
	if len(items) == 0 {
		return "", fmt.Errorf("no items in order")
	}
	
	receipt := fmt.Sprintf("=== RECEIPT %s ===\n", orderID)
	total := 0.0
	
	for i, item := range items {
		receipt += fmt.Sprintf("%-15s $%.2f\n", item, prices[i])
		total += prices[i]
	}
	
	receipt += strings.Repeat("-", 25) + "\n"
	receipt += fmt.Sprintf("%-15s $%.2f\n", "TOTAL:", total)
	
	return receipt, nil
}

func prepareCoffeeOrder(coffeeType, size string) (string, time.Duration, error) {
	// Validate coffee type
	validTypes := map[string]bool{
		"Espresso": true, "Latte": true, "Cappuccino": true,
		"Americano": true, "Mocha": true,
	}
	
	if !validTypes[coffeeType] {
		return "", 0, fmt.Errorf("unknown coffee type: %s", coffeeType)
	}
	
	// Validate size
	validSizes := map[string]bool{
		"Small": true, "Medium": true, "Large": true,
	}
	
	if !validSizes[size] {
		return "", 0, fmt.Errorf("invalid size: %s", size)
	}
	
	// Prepare coffee
	prepTime := 2 * time.Minute
	if size == "Large" {
		prepTime += 30 * time.Second
	}
	
	result := fmt.Sprintf("%s %s", size, coffeeType)
	return result, prepTime, nil
}

func processCompleteOrder() {
	fmt.Println("\n--- Processing Complete Order ---")
	
	// Step 1: Validate order
	orderID, err := validateOrderInput("ORD-789", "Charlie")
	if err != nil {
		fmt.Printf("Validation failed: %v\n", err)
		return
	}
	
	// Step 2: Check inventory
	available, err := checkInventory("Coffee Beans", 20)
	if err != nil {
		fmt.Printf("Inventory check failed: %v\n", err)
		return
	}
	
	if !available {
		fmt.Println("Item not available")
		return
	}
	
	// Step 3: Process payment
	transactionID, remaining, err := processPayment(15.50, 20.00)
	if err != nil {
		fmt.Printf("Payment failed: %v\n", err)
		return
	}
	
	fmt.Printf("✓ Order %s processed\n", orderID)
	fmt.Printf("✓ Transaction: %s\n", transactionID)
	fmt.Printf("✓ Change: $%.2f\n", remaining)
}

func validateOrderInput(orderID, customerName string) (string, error) {
	if orderID == "" {
		return "", fmt.Errorf("order ID required")
	}
	if customerName == "" {
		return "", fmt.Errorf("customer name required")
	}
	return orderID, nil
}

func checkInventory(item string, needed int) (bool, error) {
	inventory := map[string]int{
		"Coffee Beans": 100,
		"Milk":        50,
	}
	
	stock, exists := inventory[item]
	if !exists {
		return false, fmt.Errorf("item not found: %s", item)
	}
	
	return stock >= needed, nil
}

func processPayment(amount, paid float64) (string, float64, error) {
	if amount <= 0 {
		return "", 0, fmt.Errorf("invalid amount: %.2f", amount)
	}
	
	if paid < amount {
		return "", 0, fmt.Errorf("insufficient payment: paid %.2f, need %.2f", 
			paid, amount)
	}
	
	transactionID := fmt.Sprintf("TXN-%d", time.Now().Unix())
	change := paid - amount
	
	return transactionID, change, nil
}

// Example 4: Complex returns
func demonstrateComplexReturns() {
	// Returning structs
	order, summary, metrics := analyzeDay(time.Now())
	fmt.Printf("Today's analysis:\n")
	fmt.Printf("  Best seller: %+v\n", order)
	fmt.Printf("  Summary: %+v\n", summary)
	fmt.Printf("  Metrics: %+v\n", metrics)
	
	// Returning functions
	validator, formatter := getOrderProcessors("coffee")
	
	if validator("Latte") {
		fmt.Printf("Formatted: %s\n", formatter("Latte"))
	}
	
	// Returning channels (preview)
	dataChan, errorChan := startOrderStream()
	fmt.Println("\nOrder stream started (channels returned)")
	close(dataChan)
	close(errorChan)
}

type TopOrder struct {
	Item     string
	Quantity int
	Revenue  float64
}

type DailySummary struct {
	TotalOrders   int
	TotalRevenue  float64
	AverageOrder  float64
}

type PerformanceMetrics struct {
	PeakHour        int
	CustomersServed int
	Efficiency      float64
}

func analyzeDay(date time.Time) (TopOrder, DailySummary, PerformanceMetrics) {
	// Simulated analysis
	top := TopOrder{
		Item:     "Latte",
		Quantity: 47,
		Revenue:  211.50,
	}
	
	summary := DailySummary{
		TotalOrders:  156,
		TotalRevenue: 892.45,
		AverageOrder: 5.72,
	}
	
	metrics := PerformanceMetrics{
		PeakHour:        14, // 2 PM
		CustomersServed: 203,
		Efficiency:      0.87,
	}
	
	return top, summary, metrics
}

func getOrderProcessors(category string) (func(string) bool, func(string) string) {
	validator := func(item string) bool {
		// Simple validation
		return item != ""
	}
	
	formatter := func(item string) string {
		return fmt.Sprintf("[%s] %s", strings.ToUpper(category), item)
	}
	
	return validator, formatter
}

func startOrderStream() (chan string, chan error) {
	dataChan := make(chan string)
	errorChan := make(chan error)
	
	// In real app, would start goroutine to send data
	return dataChan, errorChan
}