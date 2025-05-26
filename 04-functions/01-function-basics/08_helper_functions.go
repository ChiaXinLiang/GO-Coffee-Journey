package main

import (
	"fmt"
	"strings"
	"time"
	"unicode"
)

// Helper functions are small, focused functions that:
// - Perform specific tasks
// - Are reused by other functions
// - Make code more readable and maintainable
// - Follow the DRY principle (Don't Repeat Yourself)

func main() {
	fmt.Println("=== GoCoffee Helper Functions ===\n")
	
	// Demonstrate various helper function patterns
	demonstrateStringHelpers()
	demonstrateValidationHelpers()
	demonstrateFormattingHelpers()
	demonstrateCalculationHelpers()
	demonstrateConversionHelpers()
	demonstrateDateTimeHelpers()
	
	// Show a complete example
	processCompleteOrder()
}

// ============================================
// String Helper Functions
// ============================================

func demonstrateStringHelpers() {
	fmt.Println("📝 String Helpers")
	fmt.Println(strings.Repeat("-", 30))
	
	// Capitalize names
	name1 := "marcus"
	name2 := "SARAH"
	fmt.Printf("Capitalize: %s → %s\n", name1, capitalize(name1))
	fmt.Printf("Capitalize: %s → %s\n", name2, capitalize(name2))
	
	// Clean input
	input := "  Latte  "
	fmt.Printf("Clean input: '%s' → '%s'\n", input, cleanInput(input))
	
	// Truncate long strings
	longName := "Super Grande Caramel Macchiato with Extra Shot"
	fmt.Printf("Truncate: %s\n", truncateString(longName, 20))
	
	fmt.Println()
}

// String helper: Capitalize first letter
func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	
	// Convert to lowercase first, then capitalize
	s = strings.ToLower(s)
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

// String helper: Clean user input
func cleanInput(s string) string {
	// Trim spaces and normalize
	return strings.TrimSpace(strings.ToLower(s))
}

// String helper: Truncate with ellipsis
func truncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}

// ============================================
// Validation Helper Functions
// ============================================

func demonstrateValidationHelpers() {
	fmt.Println("✅ Validation Helpers")
	fmt.Println(strings.Repeat("-", 30))
	
	// Email validation
	emails := []string{"marcus@coffee.com", "invalid-email", "sarah@", "bob@shop.co"}
	for _, email := range emails {
		fmt.Printf("%-20s: %v\n", email, isValidEmail(email))
	}
	
	// Phone validation
	fmt.Println("\nPhone validation:")
	phones := []string{"123-456-7890", "1234567890", "123-45-67890"}
	for _, phone := range phones {
		fmt.Printf("%-15s: %v\n", phone, isValidPhone(phone))
	}
	
	// Coffee size validation
	fmt.Println("\nSize validation:")
	sizes := []string{"small", "medium", "large", "extra-large"}
	for _, size := range sizes {
		fmt.Printf("%-12s: %v\n", size, isValidSize(size))
	}
	
	fmt.Println()
}

// Validation helper: Simple email check
func isValidEmail(email string) bool {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false
	}
	
	return len(parts[0]) > 0 && 
	       len(parts[1]) > 0 && 
	       strings.Contains(parts[1], ".")
}

// Validation helper: Phone format
func isValidPhone(phone string) bool {
	// Simple check for XXX-XXX-XXXX format
	if len(phone) != 12 {
		return false
	}
	
	return phone[3] == '-' && phone[7] == '-'
}

// Validation helper: Coffee sizes
func isValidSize(size string) bool {
	validSizes := []string{"small", "medium", "large"}
	return contains(validSizes, strings.ToLower(size))
}

// Generic helper: Check if slice contains string
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

// ============================================
// Formatting Helper Functions
// ============================================

func demonstrateFormattingHelpers() {
	fmt.Println("🎨 Formatting Helpers")
	fmt.Println(strings.Repeat("-", 30))
	
	// Format currency
	prices := []float64{4.5, 10.99, 0.5, 1234.56}
	for _, price := range prices {
		fmt.Printf("Price: %s\n", formatCurrency(price))
	}
	
	// Format order number
	fmt.Println("\nOrder numbers:")
	for i := 1; i <= 3; i++ {
		fmt.Printf("Order: %s\n", formatOrderNumber(i))
	}
	
	// Format percentage
	fmt.Println("\nPercentages:")
	values := []float64{0.15, 0.5, 0.999, 1.0}
	for _, val := range values {
		fmt.Printf("%.3f = %s\n", val, formatPercentage(val))
	}
	
	fmt.Println()
}

// Formatting helper: Currency
func formatCurrency(amount float64) string {
	return fmt.Sprintf("$%.2f", amount)
}

// Formatting helper: Order numbers
func formatOrderNumber(num int) string {
	return fmt.Sprintf("ORD-%05d", num)
}

// Formatting helper: Percentage
func formatPercentage(value float64) string {
	return fmt.Sprintf("%.0f%%", value*100)
}

// ============================================
// Calculation Helper Functions
// ============================================

func demonstrateCalculationHelpers() {
	fmt.Println("🧮 Calculation Helpers")
	fmt.Println(strings.Repeat("-", 30))
	
	// Calculate discounts
	price := 10.00
	fmt.Printf("Original price: %s\n", formatCurrency(price))
	fmt.Printf("10%% discount: %s\n", formatCurrency(applyDiscount(price, 0.10)))
	fmt.Printf("VIP discount: %s\n", formatCurrency(applyVIPDiscount(price)))
	
	// Calculate tips
	bill := 25.00
	fmt.Printf("\nBill: %s\n", formatCurrency(bill))
	fmt.Printf("15%% tip: %s\n", formatCurrency(calculateTip(bill, 0.15)))
	fmt.Printf("20%% tip: %s\n", formatCurrency(calculateTip(bill, 0.20)))
	
	// Round to nearest
	fmt.Println("\nRounding:")
	values := []float64{1.234, 5.678, 9.995}
	for _, val := range values {
		fmt.Printf("%.3f → %.2f\n", val, roundToNearest(val, 0.01))
	}
	
	fmt.Println()
}

// Calculation helper: Apply discount
func applyDiscount(price, discountRate float64) float64 {
	return price * (1 - discountRate)
}

// Calculation helper: VIP discount
func applyVIPDiscount(price float64) float64 {
	if price > 20 {
		return applyDiscount(price, 0.15) // 15% for large orders
	}
	return applyDiscount(price, 0.10) // 10% standard
}

// Calculation helper: Calculate tip
func calculateTip(amount, tipRate float64) float64 {
	return amount * tipRate
}

// Calculation helper: Round to nearest
func roundToNearest(value, nearest float64) float64 {
	return float64(int(value/nearest+0.5)) * nearest
}

// ============================================
// Conversion Helper Functions
// ============================================

func demonstrateConversionHelpers() {
	fmt.Println("🔄 Conversion Helpers")
	fmt.Println(strings.Repeat("-", 30))
	
	// Temperature conversions
	celsius := 93.0
	fmt.Printf("Coffee temp: %.1f°C = %.1f°F\n", 
		celsius, celsiusToFahrenheit(celsius))
	
	// Size conversions
	fmt.Println("\nSize conversions:")
	sizes := []string{"S", "M", "L"}
	for _, size := range sizes {
		fmt.Printf("%s → %s (%doz)\n", 
			size, sizeCodeToName(size), sizeToOunces(size))
	}
	
	// Time conversions
	fmt.Println("\nTime conversions:")
	minutes := []int{90, 150, 45}
	for _, min := range minutes {
		fmt.Printf("%d minutes = %s\n", min, minutesToReadable(min))
	}
	
	fmt.Println()
}

// Conversion helper: Celsius to Fahrenheit
func celsiusToFahrenheit(c float64) float64 {
	return c*9/5 + 32
}

// Conversion helper: Size code to name
func sizeCodeToName(code string) string {
	sizeMap := map[string]string{
		"S": "Small",
		"M": "Medium",
		"L": "Large",
	}
	
	if name, exists := sizeMap[code]; exists {
		return name
	}
	return "Unknown"
}

// Conversion helper: Size to ounces
func sizeToOunces(size string) int {
	switch size {
	case "S":
		return 8
	case "M":
		return 12
	case "L":
		return 16
	default:
		return 0
	}
}

// Conversion helper: Minutes to readable format
func minutesToReadable(minutes int) string {
	if minutes < 60 {
		return fmt.Sprintf("%d minutes", minutes)
	}
	
	hours := minutes / 60
	mins := minutes % 60
	
	if mins == 0 {
		return fmt.Sprintf("%d hour(s)", hours)
	}
	return fmt.Sprintf("%d hour(s) %d minutes", hours, mins)
}

// ============================================
// Date/Time Helper Functions
// ============================================

func demonstrateDateTimeHelpers() {
	fmt.Println("🕐 Date/Time Helpers")
	fmt.Println(strings.Repeat("-", 30))
	
	now := time.Now()
	
	// Format times
	fmt.Printf("Current time: %s\n", formatTime(now))
	fmt.Printf("Order time: %s\n", formatOrderTime(now))
	
	// Business hours
	fmt.Printf("Is open now: %v\n", isBusinessHours(now))
	
	// Time until closing
	fmt.Printf("Time until closing: %s\n", timeUntilClosing(now))
	
	fmt.Println()
}

// DateTime helper: Format time
func formatTime(t time.Time) string {
	return t.Format("3:04 PM")
}

// DateTime helper: Format order time
func formatOrderTime(t time.Time) string {
	return t.Format("Jan 2, 2006 at 3:04 PM")
}

// DateTime helper: Check business hours
func isBusinessHours(t time.Time) bool {
	hour := t.Hour()
	dayOfWeek := t.Weekday()
	
	// Closed on Sundays
	if dayOfWeek == time.Sunday {
		return false
	}
	
	// Mon-Sat: 6 AM to 8 PM
	return hour >= 6 && hour < 20
}

// DateTime helper: Time until closing
func timeUntilClosing(t time.Time) string {
	if !isBusinessHours(t) {
		return "Currently closed"
	}
	
	closingHour := 20 // 8 PM
	currentHour := t.Hour()
	hoursLeft := closingHour - currentHour
	
	if hoursLeft <= 1 {
		return "Closing soon!"
	}
	return fmt.Sprintf("%d hours", hoursLeft)
}

// ============================================
// Complete Example Using Multiple Helpers
// ============================================

func processCompleteOrder() {
	fmt.Println("🎯 Complete Order Example")
	fmt.Println(strings.Repeat("=", 40))
	
	// Customer info
	customerName := "  MARCUS  "
	customerEmail := "marcus@coffee.com"
	
	// Order details
	orderTime := time.Now()
	orderNumber := 42
	drinkType := "Latte"
	size := "L"
	basePrice := 4.50
	
	// Process order using helpers
	fmt.Println("Processing order...\n")
	
	// Clean and validate input
	name := capitalize(cleanInput(customerName))
	fmt.Printf("Customer: %s\n", name)
	
	if isValidEmail(customerEmail) {
		fmt.Printf("Email: %s ✓\n", customerEmail)
	}
	
	// Format order details
	fmt.Printf("Order: %s\n", formatOrderNumber(orderNumber))
	fmt.Printf("Time: %s\n", formatOrderTime(orderTime))
	
	// Calculate pricing
	sizeMultiplier := 1.3 // Large size
	price := basePrice * sizeMultiplier
	
	if strings.ToLower(name) == "marcus" {
		price = applyVIPDiscount(price)
		fmt.Println("VIP discount applied!")
	}
	
	// Display final details
	fmt.Printf("\nDrink: %s %s (%d oz)\n", 
		sizeCodeToName(size), drinkType, sizeToOunces(size))
	fmt.Printf("Price: %s\n", formatCurrency(price))
	
	// Check if open
	if isBusinessHours(orderTime) {
		fmt.Printf("\n✅ Order confirmed! %s\n", timeUntilClosing(orderTime))
	} else {
		fmt.Println("\n❌ Sorry, we're closed!")
	}
	
	fmt.Println("\n💡 Helper Functions Summary:")
	fmt.Println("• Keep helpers small and focused")
	fmt.Println("• Name them clearly")
	fmt.Println("• Make them reusable")
	fmt.Println("• Test edge cases")
	fmt.Println("• Group related helpers together")
}