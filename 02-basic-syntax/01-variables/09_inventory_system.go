package main

import (
	"fmt"
	"time"
)

// Package-level constants for our coffee shop
const storeName = "GoCoffee"

func main() {
	// Initialize inventory
	var (
		// Coffee beans (in kg)
		brazilianBeans = 50.5
		colombianBeans = 30.0
		ethiopianBeans = 25.5

		// Milk products (in liters)
		wholeMilk = 100.0
		skimMilk  = 50.0
		oatMilk   = 75.0

		// Other supplies
		sugarPackets = 1000
		cupsSmall    = 500
		cupsMedium   = 300
		cupsLarge    = 200
	)

	// Daily usage tracking
	var (
		beansUsedToday = 5.5
		milkUsedToday  = 25.0
		cupsUsedToday  = 150
	)

	// Calculate remaining supplies
	totalBeans := brazilianBeans + colombianBeans + ethiopianBeans
	totalMilk := wholeMilk + skimMilk + oatMilk
	totalCups := cupsSmall + cupsMedium + cupsLarge

	remainingBeans := totalBeans - beansUsedToday
	remainingMilk := totalMilk - milkUsedToday
	remainingCups := totalCups - cupsUsedToday

	// Generate report
	reportTime := time.Now()

	fmt.Println("=====================================")
	fmt.Printf("     %s Inventory Report\n", storeName)
	fmt.Println("=====================================")
	fmt.Printf("Date: %s\n", reportTime.Format("2006-01-02"))
	fmt.Printf("Time: %s\n", reportTime.Format("15:04:05"))
	fmt.Println("-------------------------------------")

	fmt.Println("COFFEE BEANS:")
	fmt.Printf("  Brazilian:  %.1f kg\n", brazilianBeans)
	fmt.Printf("  Colombian:  %.1f kg\n", colombianBeans)
	fmt.Printf("  Ethiopian:  %.1f kg\n", ethiopianBeans)
	fmt.Printf("  Total:      %.1f kg\n", totalBeans)

	fmt.Println("\nMILK PRODUCTS:")
	fmt.Printf("  Whole Milk: %.1f L\n", wholeMilk)
	fmt.Printf("  Skim Milk:  %.1f L\n", skimMilk)
	fmt.Printf("  Oat Milk:   %.1f L\n", oatMilk)
	fmt.Printf("  Total:      %.1f L\n", totalMilk)

	fmt.Println("\nCUPS:")
	fmt.Printf("  Small:      %d\n", cupsSmall)
	fmt.Printf("  Medium:     %d\n", cupsMedium)
	fmt.Printf("  Large:      %d\n", cupsLarge)
	fmt.Printf("  Total:      %d\n", totalCups)

	fmt.Println("\nOTHER SUPPLIES:")
	fmt.Printf("  Sugar packets: %d\n", sugarPackets)

	fmt.Println("\n-------------------------------------")
	fmt.Println("TODAY'S USAGE:")
	fmt.Printf("  Beans used:     %.1f kg\n", beansUsedToday)
	fmt.Printf("  Milk used:      %.1f L\n", milkUsedToday)
	fmt.Printf("  Cups used:      %d\n", cupsUsedToday)

	fmt.Println("\nREMAINING AFTER TODAY:")
	fmt.Printf("  Beans:          %.1f kg\n", remainingBeans)
	fmt.Printf("  Milk:           %.1f L\n", remainingMilk)
	fmt.Printf("  Cups:           %d\n", remainingCups)

	// Alert if running low
	fmt.Println("\n-------------------------------------")
	fmt.Println("ALERTS:")
	if remainingBeans < 20 {
		fmt.Println("  ⚠️  Low on coffee beans!")
	}
	if remainingMilk < 50 {
		fmt.Println("  ⚠️  Low on milk!")
	}
	if remainingCups < 100 {
		fmt.Println("  ⚠️  Low on cups!")
	}
	if remainingBeans >= 20 && remainingMilk >= 50 && remainingCups >= 100 {
		fmt.Println("  ✅ All supplies adequate")
	}

	fmt.Println("=====================================")
}