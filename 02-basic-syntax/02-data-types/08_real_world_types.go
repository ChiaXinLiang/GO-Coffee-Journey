package main

import (
	"fmt"
	"time"
)

// Define custom types for our coffee shop
type (
	// Money in cents to avoid float issues
	Money int

	// Coffee sizes
	Size string

	// Temperature in Fahrenheit
	Temperature int

	// Percentage as int (e.g., 15 = 15%)
	Percentage int
)

// Size constants
const (
	SizeSmall  Size = "small"
	SizeMedium Size = "medium"
	SizeLarge  Size = "large"
)

// Coffee order structure
type Order struct {
	ID           int
	CustomerName string
	Items        []OrderItem
	CreatedAt    time.Time
	IsPaid       bool
	TipPercent   Percentage
}

type OrderItem struct {
	Name     string
	Size     Size
	Price    Money
	Quantity int
	Options  []string
}

func main() {
	fmt.Println("=== GoCoffee Real-World Type Usage ===\n")

	// Create an order
	order := Order{
		ID:           1001,
		CustomerName: "Marcus",
		CreatedAt:    time.Now(),
		IsPaid:       false,
		TipPercent:   18,
		Items: []OrderItem{
			{
				Name:     "Latte",
				Size:     SizeLarge,
				Price:    Money(550), // $5.50
				Quantity: 1,
				Options:  []string{"extra shot", "oat milk"},
			},
			{
				Name:     "Croissant",
				Size:     "", // Not applicable
				Price:    Money(350), // $3.50
				Quantity: 2,
				Options:  []string{"warmed"},
			},
		},
	}

	// Display order
	displayOrder(order)

	// Calculate totals
	subtotal := calculateSubtotal(order)
	tax := calculateTax(subtotal, 8.5)
	tip := calculateTip(subtotal+tax, order.TipPercent)
	total := subtotal + tax + tip

	fmt.Println("\nPAYMENT SUMMARY:")
	fmt.Printf("Subtotal:  %s\n", formatMoney(subtotal))
	fmt.Printf("Tax:       %s\n", formatMoney(tax))
	fmt.Printf("Tip (%d%%): %s\n", order.TipPercent, formatMoney(tip))
	fmt.Printf("Total:     %s\n", formatMoney(total))

	// Update order status
	order.IsPaid = true
	fmt.Printf("\nPayment status: %v\n", order.IsPaid)

	// Type validation example
	fmt.Println("\nTYPE VALIDATION:")
	validateSize(SizeLarge)
	validateSize("extra-large") // Invalid size
}

func displayOrder(order Order) {
	fmt.Printf("Order #%d for %s\n", order.ID, order.CustomerName)
	fmt.Printf("Time: %s\n", order.CreatedAt.Format("3:04 PM"))
	fmt.Println("\nITEMS:")

	for _, item := range order.Items {
		fmt.Printf("- %d × %s", item.Quantity, item.Name)
		if item.Size != "" {
			fmt.Printf(" (%s)", item.Size)
		}
		fmt.Printf(" @ %s", formatMoney(item.Price))
		if len(item.Options) > 0 {
			fmt.Printf(" [%v]", item.Options)
		}
		fmt.Println()
	}
}

func calculateSubtotal(order Order) Money {
	var total Money
	for _, item := range order.Items {
		total += item.Price * Money(item.Quantity)
	}
	return total
}

func calculateTax(amount Money, rate float64) Money {
	return Money(float64(amount) * rate / 100)
}

func calculateTip(amount Money, percent Percentage) Money {
	return Money(float64(amount) * float64(percent) / 100)
}

func formatMoney(cents Money) string {
	return fmt.Sprintf("$%.2f", float64(cents)/100)
}

func validateSize(size Size) {
	switch size {
	case SizeSmall, SizeMedium, SizeLarge:
		fmt.Printf("✓ Valid size: %s\n", size)
	default:
		fmt.Printf("✗ Invalid size: %s\n", size)
	}
}