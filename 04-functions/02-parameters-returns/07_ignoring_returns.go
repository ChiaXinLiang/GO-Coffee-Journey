package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	fmt.Println("=== GoCoffee Ignoring Return Values ===\n")
	
	// Example 1: The blank identifier
	fmt.Println("Example 1: Using the Blank Identifier (_)")
	fmt.Println("-----------------------------------------")
	demonstrateBlankIdentifier()
	
	// Example 2: When to ignore returns
	fmt.Println("\n\nExample 2: When It's OK to Ignore Returns")
	fmt.Println("-----------------------------------------")
	demonstrateWhenToIgnore()
	
	// Example 3: When NOT to ignore returns
	fmt.Println("\n\nExample 3: When NOT to Ignore Returns")
	fmt.Println("-------------------------------------")
	demonstrateWhenNotToIgnore()
	
	// Example 4: Common patterns
	fmt.Println("\n\nExample 4: Common Patterns")
	fmt.Println("--------------------------")
	demonstrateCommonPatterns()
}

// Example 1: Blank identifier basics
func demonstrateBlankIdentifier() {
	// Ignoring one return value
	result, _ := divideNumbers(10, 3)
	fmt.Printf("10 ÷ 3 = %d (remainder ignored)\n", result)
	
	// Ignoring multiple values
	_, remainder := divideNumbers(10, 3)
	fmt.Printf("10 ÷ 3 remainder = %d (quotient ignored)\n", remainder)
	
	// Ignoring error (BE CAREFUL!)
	price, _ := getItemPrice("Coffee")
	fmt.Printf("Coffee price: $%.2f (error ignored - dangerous!)\n", price)
	
	// Ignoring all returns
	_ = performSideEffect("Processing order...")
	
	// Multiple blank identifiers
	_, _, z := get3DCoordinates()
	fmt.Printf("Z coordinate: %.2f (X and Y ignored)\n", z)
}

func divideNumbers(a, b int) (quotient, remainder int) {
	if b == 0 {
		return 0, 0
	}
	return a / b, a % b
}

func getItemPrice(item string) (float64, error) {
	prices := map[string]float64{
		"Coffee": 3.00,
		"Tea":    2.50,
		"Juice":  4.00,
	}
	
	price, exists := prices[item]
	if !exists {
		return 0, fmt.Errorf("item not found: %s", item)
	}
	
	return price, nil
}

func performSideEffect(message string) string {
	// Function has side effects (prints) but also returns
	fmt.Printf("Side effect: %s\n", message)
	return "Done"
}

func get3DCoordinates() (x, y, z float64) {
	return 10.5, 20.3, 15.8
}

// Example 2: When it's OK to ignore
func demonstrateWhenToIgnore() {
	// OK: Ignoring map lookup bool when providing default
	inventory := map[string]int{
		"Coffee": 100,
		"Tea":    50,
	}
	
	// We don't care if it exists, we'll use 0 as default
	coffeeStock, _ := inventory["Coffee"]
	fmt.Printf("Coffee stock: %d\n", coffeeStock)
	
	// OK: Ignoring error when we have fallback
	name, _ := getCustomerName("CUST-999")
	if name == "" {
		name = "Guest"
	}
	fmt.Printf("Customer: %s\n", name)
	
	// OK: Printf returns bytes written, rarely needed
	bytesWritten, _ := fmt.Printf("Welcome to GoCoffee!\n")
	_ = bytesWritten // Acknowledge we're ignoring it
	
	// OK: Type assertion with fallback
	var data interface{} = "Hello"
	str, _ := data.(string)
	fmt.Printf("String value: %s\n", str)
	
	// OK: Channel receive when we don't need the boolean
	ch := make(chan string, 1)
	ch <- "Order ready"
	msg, _ := <-ch // Ignoring closed status
	fmt.Printf("Message: %s\n", msg)
}

func getCustomerName(id string) (string, error) {
	if id == "CUST-001" {
		return "Alice", nil
	}
	return "", fmt.Errorf("customer not found")
}

// Example 3: When NOT to ignore
func demonstrateWhenNotToIgnore() {
	// BAD: Ignoring important errors
	fmt.Println("\n--- BAD: Ignoring Critical Errors ---")
	badExample1()
	
	// BAD: Ignoring validation results
	fmt.Println("\n--- BAD: Ignoring Validation ---")
	badExample2()
	
	// GOOD: Proper error handling
	fmt.Println("\n--- GOOD: Proper Error Handling ---")
	goodExample()
}

func badExample1() {
	// DON'T DO THIS - Ignoring file operation errors
	content, _ := readConfigFile("config.json")
	// What if file doesn't exist? content will be empty!
	fmt.Printf("Config: %s\n", content)
	
	// DON'T DO THIS - Ignoring payment errors
	_, _ = processPayment(100.00, "invalid-card")
	fmt.Println("Payment processed!") // But was it really?
}

func badExample2() {
	// DON'T DO THIS - Ignoring validation
	_, _ = validateEmail("not-an-email")
	fmt.Println("Email valid!") // No, it's not!
	
	// DON'T DO THIS - Ignoring important boolean
	value, ok := getImportantValue()
	_ = ok // Ignoring whether we got a valid value
	fmt.Printf("Important value: %d\n", value) // Might be zero!
}

func goodExample() {
	// GOOD: Always check errors that matter
	content, err := readConfigFile("config.json")
	if err != nil {
		log.Printf("Config error: %v", err)
		// Use defaults or exit
		return
	}
	fmt.Printf("Config loaded: %s\n", content)
	
	// GOOD: Check payment results
	txnID, err := processPayment(100.00, "valid-card")
	if err != nil {
		fmt.Printf("Payment failed: %v\n", err)
		return
	}
	fmt.Printf("Payment successful: %s\n", txnID)
	
	// GOOD: Validate properly
	isValid, err := validateEmail("user@example.com")
	if err != nil {
		fmt.Printf("Validation error: %v\n", err)
		return
	}
	if !isValid {
		fmt.Println("Email format invalid")
		return
	}
	fmt.Println("Email is valid!")
}

func readConfigFile(filename string) (string, error) {
	if filename != "config.json" {
		return "", fmt.Errorf("file not found: %s", filename)
	}
	return `{"shop": "GoCoffee"}`, nil
}

func processPayment(amount float64, card string) (string, error) {
	if !strings.Contains(card, "valid") {
		return "", fmt.Errorf("invalid card")
	}
	return "TXN-12345", nil
}

func validateEmail(email string) (bool, error) {
	if email == "" {
		return false, fmt.Errorf("email cannot be empty")
	}
	
	isValid := strings.Contains(email, "@")
	return isValid, nil
}

func getImportantValue() (int, bool) {
	// Returns value and whether it's valid
	return 0, false
}

// Example 4: Common patterns
func demonstrateCommonPatterns() {
	// Pattern 1: Checking only error
	checkOnlyError()
	
	// Pattern 2: Checking only existence
	checkOnlyExistence()
	
	// Pattern 3: Side-effect functions
	sideEffectFunctions()
	
	// Pattern 4: Loop patterns
	loopPatterns()
}

func checkOnlyError() {
	fmt.Println("\n--- Pattern: Checking Only Error ---")
	
	// Common pattern: only care about error
	if _, err := validateOrder("ORD-123", 5); err != nil {
		fmt.Printf("Order validation failed: %v\n", err)
		return
	}
	fmt.Println("Order is valid!")
	
	// Writing to file - only care if it succeeded
	if _, err := writeToLog("Customer arrived"); err != nil {
		fmt.Printf("Log error: %v\n", err)
	}
}

func validateOrder(orderID string, quantity int) (bool, error) {
	if orderID == "" {
		return false, fmt.Errorf("empty order ID")
	}
	if quantity <= 0 {
		return false, fmt.Errorf("invalid quantity")
	}
	return true, nil
}

func writeToLog(message string) (int, error) {
	// Simulate writing to log
	return len(message), nil
}

func checkOnlyExistence() {
	fmt.Println("\n--- Pattern: Checking Only Existence ---")
	
	menu := map[string]float64{
		"Coffee":     3.00,
		"Tea":        2.50,
		"Cappuccino": 4.00,
	}
	
	// Only care if item exists
	if _, exists := menu["Latte"]; !exists {
		fmt.Println("Latte not on menu")
	}
	
	// Type assertion - only care about success
	var data interface{} = 42
	if _, ok := data.(int); ok {
		fmt.Println("Data is an integer")
	}
}

func sideEffectFunctions() {
	fmt.Println("\n--- Pattern: Side-Effect Functions ---")
	
	// fmt.Println returns values we rarely need
	fmt.Println("This returns (int, error) but we ignore both")
	
	// Explicit ignore for clarity
	_ = initializeSystem()
	
	// Defer with return values
	defer func() {
		_ = cleanup()
	}()
}

func initializeSystem() error {
	fmt.Println("System initialized")
	return nil
}

func cleanup() error {
	fmt.Println("Cleanup completed")
	return nil
}

func loopPatterns() {
	fmt.Println("\n--- Pattern: Loops and Ranges ---")
	
	items := []string{"Coffee", "Tea", "Juice"}
	
	// Ignoring index
	fmt.Println("Items:")
	for _, item := range items {
		fmt.Printf("  - %s\n", item)
	}
	
	// Ignoring value
	fmt.Println("\nCounting items:")
	count := 0
	for range items {
		count++
	}
	fmt.Printf("Total: %d\n", count)
	
	// Map iteration - ignoring key or value
	prices := map[string]float64{
		"Coffee": 3.00,
		"Tea":    2.50,
	}
	
	fmt.Println("\nJust prices:")
	for _, price := range prices {
		fmt.Printf("  $%.2f\n", price)
	}
}

// Best practices summary
func blankIdentifierBestPractices() {
	// Rule 1: Never ignore errors that could affect program correctness
	
	// Rule 2: Be explicit about ignoring values
	_ = fmt.Sprintf("Explicitly ignoring return")
	
	// Rule 3: Document why you're ignoring a value if not obvious
	_, _ = fmt.Printf("Output text") // Don't need bytes written
	
	// Rule 4: Consider if ignoring indicates design issue
	// Maybe function returns too much?
	
	// Rule 5: Use linters to catch ignored errors
	// errcheck, golangci-lint can help
}