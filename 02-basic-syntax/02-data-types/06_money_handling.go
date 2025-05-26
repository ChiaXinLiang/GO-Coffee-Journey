package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("=== GoCoffee Money Handling ===\n")

	// DON'T: Using float for money (precision issues)
	fmt.Println("❌ WRONG WAY (float problems):")

	price1 := 0.1
	price2 := 0.2
	total := price1 + price2
	fmt.Printf("$0.10 + $0.20 = $%.20f (not exactly 0.30!)\n", total)

	// Compound the error
	wrongTotal := 0.0
	for i := 0; i < 100; i++ {
		wrongTotal += 0.01
	}
	fmt.Printf("100 × $0.01 = $%.20f (not exactly 1.00!)\n", wrongTotal)

	// DO: Use integers for cents
	fmt.Println("\n✅ CORRECT WAY (using cents):")

	// Prices in cents
	espressoPriceCents := 300   // $3.00
	lattePriceCents := 450      // $4.50
	cappuccinoPriceCents := 400 // $4.00

	// Order calculation
	orderItems := []struct {
		name       string
		priceCents int
		quantity   int
	}{
		{"Espresso", espressoPriceCents, 2},
		{"Latte", lattePriceCents, 1},
		{"Cappuccino", cappuccinoPriceCents, 3},
	}

	subtotalCents := 0
	fmt.Println("\nORDER DETAILS:")
	for _, item := range orderItems {
		itemTotal := item.priceCents * item.quantity
		subtotalCents += itemTotal
		fmt.Printf("%d × %s @ $%.2f = $%.2f\n",
			item.quantity,
			item.name,
			float64(item.priceCents)/100,
			float64(itemTotal)/100)
	}

	// Tax calculation (8.5%)
	taxRate := 0.085
	taxCents := int(math.Round(float64(subtotalCents) * taxRate))
	totalCents := subtotalCents + taxCents

	fmt.Printf("\nSubtotal: $%.2f\n", float64(subtotalCents)/100)
	fmt.Printf("Tax (8.5%%): $%.2f\n", float64(taxCents)/100)
	fmt.Printf("Total: $%.2f\n", float64(totalCents)/100)

	// Tip calculation
	fmt.Println("\nTIP CALCULATION:")
	tipPercentages := []float64{0.15, 0.18, 0.20}
	for _, pct := range tipPercentages {
		tipCents := int(math.Round(float64(totalCents) * pct))
		finalCents := totalCents + tipCents
		fmt.Printf("%.0f%% tip: $%.2f (Total: $%.2f)\n",
			pct*100,
			float64(tipCents)/100,
			float64(finalCents)/100)
	}

	// Currency formatting helper
	fmt.Println("\nCURRENCY FORMATTING:")
	amounts := []int{100, 1000, 10000, 100000} // cents
	for _, cents := range amounts {
		fmt.Printf("%d cents = %s\n", cents, formatMoney(cents))
	}
}

// formatMoney converts cents to a formatted dollar string
func formatMoney(cents int) string {
	dollars := float64(cents) / 100
	return fmt.Sprintf("$%.2f", dollars)
}