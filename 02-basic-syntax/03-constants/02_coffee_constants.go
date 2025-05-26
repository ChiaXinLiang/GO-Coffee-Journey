package main

import (
	"fmt"
	"strings"
)

// Coffee sizes
const (
	SizeSmall  = "small"
	SizeMedium = "medium"
	SizeLarge  = "large"
)

// Size multipliers for pricing
const (
	SmallMultiplier  = 1.0
	MediumMultiplier = 1.3
	LargeMultiplier  = 1.6
)

// Base prices (in dollars)
const (
	EspressoBase   = 3.00
	AmericanoBase  = 3.50
	LatteBase      = 4.50
	CappuccinoBase = 4.00
	MochaBase      = 5.00
)

// Add-on prices
const (
	ExtraShotPrice    = 0.75
	DecafPrice        = 0.00
	SyrupPrice        = 0.50
	WhippedCreamPrice = 0.50
	OatMilkPrice      = 0.65
	SoyMilkPrice      = 0.60
)

func main() {
	fmt.Println("=== GoCoffee Price Calculator ===\n")

	// Calculate prices for different sizes
	coffeeTypes := []struct {
		name      string
		basePrice float64
	}{
		{"Espresso", EspressoBase},
		{"Americano", AmericanoBase},
		{"Latte", LatteBase},
		{"Cappuccino", CappuccinoBase},
		{"Mocha", MochaBase},
	}

	fmt.Println("MENU PRICES:")
	fmt.Printf("%-12s %8s %8s %8s\n", "Coffee", "Small", "Medium", "Large")
	fmt.Println(strings.Repeat("-", 40))

	for _, coffee := range coffeeTypes {
		smallPrice := coffee.basePrice * SmallMultiplier
		mediumPrice := coffee.basePrice * MediumMultiplier
		largePrice := coffee.basePrice * LargeMultiplier

		fmt.Printf("%-12s $%6.2f $%6.2f $%6.2f\n",
			coffee.name, smallPrice, mediumPrice, largePrice)
	}

	fmt.Println("\nADD-ONS:")
	fmt.Printf("Extra Shot:    $%.2f\n", ExtraShotPrice)
	fmt.Printf("Syrup:         $%.2f\n", SyrupPrice)
	fmt.Printf("Whipped Cream: $%.2f\n", WhippedCreamPrice)
	fmt.Printf("Oat Milk:      $%.2f\n", OatMilkPrice)
	fmt.Printf("Soy Milk:      $%.2f\n", SoyMilkPrice)

	// Example order calculation
	fmt.Println("\nSAMPLE ORDER:")
	orderTotal := LatteBase*LargeMultiplier + ExtraShotPrice + OatMilkPrice
	fmt.Printf("Large Latte + Extra Shot + Oat Milk = $%.2f\n", orderTotal)
}