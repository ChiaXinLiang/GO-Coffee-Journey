package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	fmt.Println("=== GoCoffee String Operations ===\n")

	// Basic string operations
	coffeeName := "Cappuccino"
	customerName := "Marcus Chen"

	// String concatenation
	greeting := "Hello, " + customerName + "!"
	order := fmt.Sprintf("One %s for %s", coffeeName, customerName)

	fmt.Println("STRING BASICS:")
	fmt.Println("Greeting:", greeting)
	fmt.Println("Order:", order)

	// String properties
	fmt.Printf("\nSTRING PROPERTIES:")
	fmt.Printf("\nCoffee name length: %d bytes\n", len(coffeeName))
	fmt.Printf("Customer name length: %d bytes\n", len(customerName))

	// Working with Unicode
	menuItem := "Café Latté ☕"
	fmt.Printf("\nUNICODE STRING: %s\n", menuItem)
	fmt.Printf("Byte length: %d\n", len(menuItem))
	fmt.Printf("Rune count: %d\n", utf8.RuneCountInString(menuItem))

	// Iterating over strings
	fmt.Println("\nITERATING BY BYTES:")
	for i := 0; i < len(menuItem); i++ {
		fmt.Printf("%d: %c (%d)\n", i, menuItem[i], menuItem[i])
	}

	fmt.Println("\nITERATING BY RUNES (correct way):")
	for i, r := range menuItem {
		fmt.Printf("%d: %c (%d)\n", i, r, r)
	}

	// String manipulation
	fmt.Println("\nSTRING MANIPULATION:")

	// Case conversion
	fmt.Printf("Uppercase: %s\n", strings.ToUpper(coffeeName))
	fmt.Printf("Lowercase: %s\n", strings.ToLower(coffeeName))
	fmt.Printf("Title case: %s\n", strings.Title(strings.ToLower("ICED LATTE")))

	// Trimming
	messyInput := "  Espresso   "
	fmt.Printf("Before trim: '%s'\n", messyInput)
	fmt.Printf("After trim: '%s'\n", strings.TrimSpace(messyInput))

	// Contains and replacements
	description := "Our coffee is the best coffee in the coffee world!"
	fmt.Printf("\nOriginal: %s\n", description)
	fmt.Printf("Contains 'coffee': %v\n", strings.Contains(description, "coffee"))
	fmt.Printf("Coffee count: %d\n", strings.Count(description, "coffee"))
	fmt.Printf("Replace coffee: %s\n", strings.ReplaceAll(description, "coffee", "☕"))

	// Splitting and joining
	ingredients := "espresso,milk,foam,chocolate"
	parts := strings.Split(ingredients, ",")
	fmt.Printf("\nSplit ingredients: %v\n", parts)
	fmt.Printf("Joined with ' + ': %s\n", strings.Join(parts, " + "))

	// Multi-line strings
	menu := `
    ╔═══════════════════════╗
    ║   GoCoffee Menu       ║
    ╠═══════════════════════╣
    ║ Espresso      $3.00   ║
    ║ Cappuccino    $4.00   ║
    ║ Latte         $4.50   ║
    ╚═══════════════════════╝`

	fmt.Println("\nMULTI-LINE STRING:")
	fmt.Println(menu)
}