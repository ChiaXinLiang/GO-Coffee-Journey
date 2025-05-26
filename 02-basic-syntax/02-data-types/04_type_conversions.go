package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("=== GoCoffee Type Conversions ===\n")

	// Number conversions
	var (
		smallCups  int8  = 100
		mediumCups int16 = 200
		largeCups  int32 = 300
	)

	// Calculate total (must convert to same type)
	totalCups := int32(smallCups) + int32(mediumCups) + largeCups
	fmt.Printf("Total cups: %d\n", totalCups)

	// Float to int conversion (truncates decimal)
	price := 4.99
	wholePrice := int(price)
	fmt.Printf("Price: $%.2f, Whole price: $%d (lost $.99!)\n", price, wholePrice)

	// Int to float for calculations
	items := 3
	itemPrice := 4.50
	total := float64(items) * itemPrice
	fmt.Printf("%d items at $%.2f = $%.2f\n", items, itemPrice, total)

	// String conversions
	fmt.Println("\nSTRING CONVERSIONS:")

	// Number to string
	orderNumber := 1234
	orderString := strconv.Itoa(orderNumber)
	fmt.Printf("Order #%s (was int: %d)\n", orderString, orderNumber)

	// String to number
	inputPrice := "4.75"
	convertedPrice, err := strconv.ParseFloat(inputPrice, 64)
	if err == nil {
		fmt.Printf("Converted price: $%.2f\n", convertedPrice)
	}

	// String to int
	quantity := "5"
	convertedQty, err := strconv.Atoi(quantity)
	if err == nil {
		fmt.Printf("Converted quantity: %d\n", convertedQty)
	}

	// Boolean conversions
	fmt.Println("\nBOOLEAN CONVERSIONS:")

	isOpenString := "true"
	isOpen, err := strconv.ParseBool(isOpenString)
	if err == nil {
		fmt.Printf("Shop open: %v\n", isOpen)
	}

	// Format conversions for display
	fmt.Println("\nFORMATTING FOR DISPLAY:")

	revenue := 12345.67
	fmt.Printf("Revenue: $%.2f\n", revenue)
	fmt.Printf("Revenue (no decimals): $%.0f\n", revenue)
	fmt.Printf("Revenue (with commas): $%,.2f\n", revenue)

	// Type assertions (preview - we'll cover interfaces later)
	var anyValue interface{} = "Cappuccino"

	// Safe type assertion
	if coffeeType, ok := anyValue.(string); ok {
		fmt.Printf("\nCoffee type: %s\n", coffeeType)
	}

	// Dangerous conversions to avoid
	fmt.Println("\n⚠️  CONVERSION PITFALLS:")

	// Overflow example
	bigNumber := int16(32767)
	// smallNumber := int8(bigNumber) // This would overflow!
	fmt.Printf("Big number: %d (too big for int8!)\n", bigNumber)

	// Precision loss
	preciseMoney := 10.99
	cents := int(preciseMoney * 100) // Better way to handle money
	fmt.Printf("Money: $%.2f = %d cents\n", preciseMoney, cents)
}