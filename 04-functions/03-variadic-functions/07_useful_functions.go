package main

import (
	"fmt"
	"strings"
	"math"
)

// === Common Utility Functions ===

// Max returns the largest of the given values
func Max(values ...float64) (float64, error) {
	if len(values) == 0 {
		return 0, fmt.Errorf("Max requires at least one value")
	}
	
	max := values[0]
	for _, v := range values[1:] {
		if v > max {
			max = v
		}
	}
	return max, nil
}

// Min returns the smallest of the given values
func Min(values ...float64) (float64, error) {
	if len(values) == 0 {
		return 0, fmt.Errorf("Min requires at least one value")
	}
	
	min := values[0]
	for _, v := range values[1:] {
		if v < min {
			min = v
		}
	}
	return min, nil
}

// Sum adds all values together
func Sum(values ...float64) float64 {
	var total float64
	for _, v := range values {
		total += v
	}
	return total
}

// Average calculates the mean of values
func Average(values ...float64) (float64, error) {
	if len(values) == 0 {
		return 0, fmt.Errorf("Average requires at least one value")
	}
	
	return Sum(values...) / float64(len(values)), nil
}

// === String Utilities ===

// Join concatenates strings with a separator (more flexible than strings.Join)
func Join(sep string, parts ...string) string {
	return strings.Join(parts, sep)
}

// Contains checks if any of the values matches the target
func Contains(target string, values ...string) bool {
	for _, v := range values {
		if v == target {
			return true
		}
	}
	return false
}

// === Coffee Shop Specific Functions ===

// CalculateOrderTotal with optional discounts and fees
func CalculateOrderTotal(basePrice float64, modifiers ...float64) float64 {
	total := basePrice
	
	for _, mod := range modifiers {
		total += mod
	}
	
	// Ensure non-negative total
	if total < 0 {
		return 0
	}
	
	return total
}

// LogOrder creates a formatted order log entry
func LogOrder(orderID string, items ...string) string {
	var log strings.Builder
	
	log.WriteString(fmt.Sprintf("Order #%s", orderID))
	
	if len(items) > 0 {
		log.WriteString(" contains: ")
		log.WriteString(strings.Join(items, ", "))
	} else {
		log.WriteString(" (empty order)")
	}
	
	return log.String()
}

// ValidateOrder checks if all required items are present
func ValidateOrder(required []string, provided ...string) (bool, []string) {
	missing := []string{}
	
	for _, req := range required {
		if !Contains(req, provided...) {
			missing = append(missing, req)
		}
	}
	
	return len(missing) == 0, missing
}

// === Builder Pattern with Variadic ===

type Coffee struct {
	Base      string
	Additions []string
	Size      string
	Price     float64
}

// NewCoffee creates a coffee with optional additions
func NewCoffee(base string, additions ...string) *Coffee {
	coffee := &Coffee{
		Base:      base,
		Additions: additions,
		Size:      "medium", // default
		Price:     3.50,     // base price
	}
	
	// Add $0.50 for each addition
	coffee.Price += float64(len(additions)) * 0.50
	
	return coffee
}

// === Functional Options Pattern ===

type OrderOption func(*Order)

type Order struct {
	ID         string
	Items      []string
	TotalPrice float64
	Priority   bool
	Notes      string
}

func NewOrder(id string, options ...OrderOption) *Order {
	order := &Order{
		ID:    id,
		Items: []string{},
	}
	
	// Apply all options
	for _, opt := range options {
		opt(order)
	}
	
	return order
}

func WithItems(items ...string) OrderOption {
	return func(o *Order) {
		o.Items = append(o.Items, items...)
	}
}

func WithPriority() OrderOption {
	return func(o *Order) {
		o.Priority = true
	}
}

func WithNotes(notes string) OrderOption {
	return func(o *Order) {
		o.Notes = notes
	}
}

func main() {
	fmt.Println("=== Useful Variadic Function Patterns ===")
	fmt.Println()

	// Math utilities
	prices := []float64{3.50, 4.25, 5.00, 2.75}
	
	if max, err := Max(prices...); err == nil {
		fmt.Printf("Highest price: $%.2f\n", max)
	}
	
	if avg, err := Average(prices...); err == nil {
		fmt.Printf("Average price: $%.2f\n", avg)
	}
	
	fmt.Printf("Total revenue: $%.2f\n", Sum(prices...))
	fmt.Println()

	// String utilities
	fmt.Println("Menu items:", Join(" | ", "Espresso", "Latte", "Cappuccino"))
	fmt.Println("Has Latte?", Contains("Latte", "Espresso", "Latte", "Cappuccino"))
	fmt.Println()

	// Order calculations
	total := CalculateOrderTotal(5.00, 0.75, -1.00, 0.50) // base + additions - discount + tip
	fmt.Printf("Order total: $%.2f\n", total)
	fmt.Println()

	// Order logging
	log := LogOrder("12345", "Latte", "Croissant", "Orange Juice")
	fmt.Println(log)
	fmt.Println()

	// Order validation
	required := []string{"payment", "items", "customer"}
	valid, missing := ValidateOrder(required, "payment", "customer")
	if !valid {
		fmt.Println("Missing required fields:", missing)
	}
	fmt.Println()

	// Coffee builder
	latte := NewCoffee("Espresso", "Steamed Milk", "Foam")
	fmt.Printf("%s coffee with %v: $%.2f\n", latte.Base, latte.Additions, latte.Price)
	fmt.Println()

	// Functional options
	order := NewOrder("ORD-789",
		WithItems("Cappuccino", "Bagel"),
		WithPriority(),
		WithNotes("Extra hot, no sugar"),
	)
	
	fmt.Printf("Order %s:\n", order.ID)
	fmt.Printf("  Items: %v\n", order.Items)
	fmt.Printf("  Priority: %v\n", order.Priority)
	fmt.Printf("  Notes: %s\n", order.Notes)
}

// Variadic functions enable:
// 1. Flexible APIs that grow with needs
// 2. Clean function calls without slice creation
// 3. Optional parameters without overloading
// 4. Builder and option patterns
// 5. Utilities that work with any number of inputs