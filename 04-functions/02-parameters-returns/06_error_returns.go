package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// Custom error types
type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("validation error: %s - %s", e.Field, e.Message)
}

type BusinessError struct {
	Code    string
	Message string
	Details map[string]interface{}
}

func (e BusinessError) Error() string {
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

func main() {
	fmt.Println("=== GoCoffee Error Returns ===\n")
	
	// Example 1: Basic error handling
	fmt.Println("Example 1: Basic Error Handling")
	fmt.Println("-------------------------------")
	demonstrateBasicErrors()
	
	// Example 2: Error patterns
	fmt.Println("\n\nExample 2: Common Error Patterns")
	fmt.Println("--------------------------------")
	demonstrateErrorPatterns()
	
	// Example 3: Custom errors
	fmt.Println("\n\nExample 3: Custom Error Types")
	fmt.Println("----------------------------")
	demonstrateCustomErrors()
	
	// Example 4: Error handling best practices
	fmt.Println("\n\nExample 4: Error Handling Best Practices")
	fmt.Println("----------------------------------------")
	demonstrateBestPractices()
}

// Example 1: Basic error handling
func demonstrateBasicErrors() {
	// Success case
	result, err := makeCoffee("Latte", "Medium")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Success: %s\n", result)
	}
	
	// Error case
	result, err = makeCoffee("Unknown", "Medium")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Success: %s\n", result)
	}
	
	// Multiple error checks
	processOrderWithErrors()
}

func makeCoffee(coffeeType, size string) (string, error) {
	// Validate coffee type
	validTypes := []string{"Espresso", "Latte", "Cappuccino", "Americano"}
	valid := false
	
	for _, vt := range validTypes {
		if coffeeType == vt {
			valid = true
			break
		}
	}
	
	if !valid {
		return "", fmt.Errorf("invalid coffee type: %s", coffeeType)
	}
	
	// Validate size
	if size != "Small" && size != "Medium" && size != "Large" {
		return "", fmt.Errorf("invalid size: %s", size)
	}
	
	// Make coffee
	result := fmt.Sprintf("One %s %s ready!", size, coffeeType)
	return result, nil
}

func processOrderWithErrors() {
	fmt.Println("\n--- Processing order with multiple checks ---")
	
	// Step 1: Validate input
	orderID, err := validateOrderID("ORD-")
	if err != nil {
		fmt.Printf("Step 1 failed: %v\n", err)
		return
	}
	
	// Step 2: Check stock
	err = checkStock("Coffee Beans", 50)
	if err != nil {
		fmt.Printf("Step 2 failed: %v\n", err)
		return
	}
	
	// Step 3: Process payment
	txnID, err := processPayment(15.50, "credit")
	if err != nil {
		fmt.Printf("Step 3 failed: %v\n", err)
		return
	}
	
	fmt.Printf("✓ Order %s processed, transaction: %s\n", orderID, txnID)
}

func validateOrderID(id string) (string, error) {
	if id == "" {
		return "", errors.New("order ID cannot be empty")
	}
	
	if len(id) < 5 {
		return "", errors.New("order ID too short")
	}
	
	if !strings.HasPrefix(id, "ORD-") {
		return "", errors.New("order ID must start with ORD-")
	}
	
	return id, nil
}

func checkStock(item string, needed int) error {
	stock := map[string]int{
		"Coffee Beans": 100,
		"Milk":        200,
		"Cups":        500,
	}
	
	available, exists := stock[item]
	if !exists {
		return fmt.Errorf("item not found: %s", item)
	}
	
	if available < needed {
		return fmt.Errorf("insufficient stock: need %d, have %d", needed, available)
	}
	
	return nil
}

func processPayment(amount float64, method string) (string, error) {
	if amount <= 0 {
		return "", fmt.Errorf("invalid amount: %.2f", amount)
	}
	
	validMethods := map[string]bool{
		"cash":   true,
		"credit": true,
		"debit":  true,
	}
	
	if !validMethods[method] {
		return "", fmt.Errorf("invalid payment method: %s", method)
	}
	
	// Generate transaction ID
	txnID := fmt.Sprintf("TXN-%d", time.Now().Unix())
	return txnID, nil
}

// Example 2: Common error patterns
func demonstrateErrorPatterns() {
	// Pattern 1: Sentinel errors
	result, err := findCustomer("CUST-999")
	if err == ErrCustomerNotFound {
		fmt.Println("Customer not found - creating new account")
	} else if err != nil {
		fmt.Printf("Database error: %v\n", err)
	} else {
		fmt.Printf("Found customer: %s\n", result)
	}
	
	// Pattern 2: Error wrapping
	err = performComplexOperation()
	if err != nil {
		fmt.Printf("Operation failed: %v\n", err)
	}
	
	// Pattern 3: Multiple error sources
	results, errs := processMultipleItems([]string{"Coffee", "Tea", "InvalidItem", "Juice"})
	fmt.Printf("\nProcessed %d items successfully\n", len(results))
	if len(errs) > 0 {
		fmt.Println("Errors encountered:")
		for i, err := range errs {
			fmt.Printf("  %d: %v\n", i+1, err)
		}
	}
}

var ErrCustomerNotFound = errors.New("customer not found")
var ErrDatabaseConnection = errors.New("database connection failed")

func findCustomer(id string) (string, error) {
	customers := map[string]string{
		"CUST-001": "Alice Smith",
		"CUST-002": "Bob Jones",
	}
	
	name, exists := customers[id]
	if !exists {
		return "", ErrCustomerNotFound
	}
	
	return name, nil
}

func performComplexOperation() error {
	// Simulate multiple steps that could fail
	if err := step1(); err != nil {
		return fmt.Errorf("step 1 failed: %w", err)
	}
	
	if err := step2(); err != nil {
		return fmt.Errorf("step 2 failed: %w", err)
	}
	
	if err := step3(); err != nil {
		return fmt.Errorf("step 3 failed: %w", err)
	}
	
	return nil
}

func step1() error {
	return nil
}

func step2() error {
	return errors.New("network timeout")
}

func step3() error {
	return nil
}

func processMultipleItems(items []string) ([]string, []error) {
	var results []string
	var errors []error
	
	for _, item := range items {
		if strings.Contains(item, "Invalid") {
			errors = append(errors, fmt.Errorf("invalid item: %s", item))
			continue
		}
		
		results = append(results, fmt.Sprintf("Processed: %s", item))
	}
	
	return results, errors
}

// Example 3: Custom errors
func demonstrateCustomErrors() {
	// Validation error
	err := validateOrder(Order{
		CustomerName: "",
		Items:        []string{"Coffee"},
		Total:        -5.00,
	})
	
	if err != nil {
		if vErr, ok := err.(ValidationError); ok {
			fmt.Printf("Validation failed on field '%s': %s\n", vErr.Field, vErr.Message)
		} else {
			fmt.Printf("Error: %v\n", err)
		}
	}
	
	// Business error
	err = processBusinessRule("VIP", 1000.00)
	if err != nil {
		if bErr, ok := err.(BusinessError); ok {
			fmt.Printf("Business rule violation [%s]: %s\n", bErr.Code, bErr.Message)
			if bErr.Details != nil {
				fmt.Printf("Details: %v\n", bErr.Details)
			}
		}
	}
	
	// Chain of errors
	err = performOrderWorkflow("ORD-123")
	if err != nil {
		fmt.Printf("Workflow error: %v\n", err)
	}
}

type Order struct {
	CustomerName string
	Items        []string
	Total        float64
}

func validateOrder(order Order) error {
	if order.CustomerName == "" {
		return ValidationError{
			Field:   "CustomerName",
			Message: "customer name is required",
		}
	}
	
	if len(order.Items) == 0 {
		return ValidationError{
			Field:   "Items",
			Message: "order must contain at least one item",
		}
	}
	
	if order.Total <= 0 {
		return ValidationError{
			Field:   "Total",
			Message: "order total must be positive",
		}
	}
	
	return nil
}

func processBusinessRule(customerType string, orderAmount float64) error {
	if customerType == "VIP" && orderAmount > 500 {
		return BusinessError{
			Code:    "LIMIT_EXCEEDED",
			Message: "VIP order limit exceeded",
			Details: map[string]interface{}{
				"customer_type": customerType,
				"order_amount":  orderAmount,
				"max_allowed":   500.00,
			},
		}
	}
	
	return nil
}

func performOrderWorkflow(orderID string) error {
	// Simulate workflow with potential errors
	if err := validateWorkflowInput(orderID); err != nil {
		return fmt.Errorf("workflow validation failed: %w", err)
	}
	
	if err := checkInventoryForWorkflow(); err != nil {
		return fmt.Errorf("inventory check failed: %w", err)
	}
	
	if err := processWorkflowPayment(); err != nil {
		return fmt.Errorf("payment processing failed: %w", err)
	}
	
	return nil
}

func validateWorkflowInput(orderID string) error {
	if orderID == "" {
		return errors.New("order ID required")
	}
	return nil
}

func checkInventoryForWorkflow() error {
	// Simulate inventory error
	return errors.New("insufficient coffee beans")
}

func processWorkflowPayment() error {
	return nil
}

// Example 4: Best practices
func demonstrateBestPractices() {
	// Good: Check errors immediately
	goodErrorHandling()
	
	// Good: Provide context
	contextualErrors()
	
	// Good: Handle different error types
	handleErrorTypes()
	
	// Good: Aggregate errors
	aggregateErrors()
}

func goodErrorHandling() {
	fmt.Println("\n--- Good: Check errors immediately ---")
	
	// Always check errors right after function call
	result, err := makeCoffee("Latte", "Large")
	if err != nil {
		// Handle error
		fmt.Printf("Cannot make coffee: %v\n", err)
		return
	}
	
	// Use result only after checking error
	fmt.Printf("Coffee ready: %s\n", result)
}

func contextualErrors() {
	fmt.Println("\n--- Good: Provide error context ---")
	
	orderID := "ORD-789"
	customerID := "CUST-123"
	
	err := processContextualOrder(orderID, customerID)
	if err != nil {
		// Error message includes context
		fmt.Printf("Failed to process order: %v\n", err)
	}
}

func processContextualOrder(orderID, customerID string) error {
	// Add context to errors
	if err := validateCustomer(customerID); err != nil {
		return fmt.Errorf("order %s: invalid customer %s: %w", orderID, customerID, err)
	}
	
	if err := checkOrderLimit(orderID); err != nil {
		return fmt.Errorf("order %s: limit check failed: %w", orderID, err)
	}
	
	return nil
}

func validateCustomer(id string) error {
	if id == "" {
		return errors.New("empty customer ID")
	}
	return nil
}

func checkOrderLimit(orderID string) error {
	// Simulate limit exceeded
	return errors.New("daily order limit exceeded")
}

func handleErrorTypes() {
	fmt.Println("\n--- Good: Handle different error types ---")
	
	err := complexOperation()
	
	switch {
	case err == nil:
		fmt.Println("Success!")
		
	case errors.Is(err, ErrCustomerNotFound):
		fmt.Println("Customer not found - please register")
		
	case errors.Is(err, ErrDatabaseConnection):
		fmt.Println("Database issue - please try again")
		
	default:
		fmt.Printf("Unexpected error: %v\n", err)
	}
}

func complexOperation() error {
	// Simulate different error scenarios
	return fmt.Errorf("wrapped error: %w", ErrCustomerNotFound)
}

func aggregateErrors() {
	fmt.Println("\n--- Good: Aggregate multiple errors ---")
	
	items := []string{"Coffee", "BadItem1", "Tea", "BadItem2"}
	errors := processItemsWithErrors(items)
	
	if len(errors) > 0 {
		fmt.Printf("Processing completed with %d errors:\n", len(errors))
		for i, err := range errors {
			fmt.Printf("  Error %d: %v\n", i+1, err)
		}
	} else {
		fmt.Println("All items processed successfully!")
	}
}

func processItemsWithErrors(items []string) []error {
	var errs []error
	
	for _, item := range items {
		if strings.HasPrefix(item, "Bad") {
			errs = append(errs, fmt.Errorf("cannot process item: %s", item))
		}
	}
	
	return errs
}