package main

import "fmt"

func main() {
	fmt.Println("=== GoCoffee Menu Display System ===\n")

	// Menu categories and items
	categories := []string{"Hot Drinks", "Cold Drinks", "Pastries"}
	
	hotDrinks := []struct {
		name  string
		price float64
	}{
		{"Espresso", 2.50},
		{"Americano", 3.00},
		{"Latte", 4.50},
		{"Cappuccino", 4.00},
	}
	
	coldDrinks := []struct {
		name  string
		price float64
	}{
		{"Iced Coffee", 3.50},
		{"Cold Brew", 4.00},
		{"Frappé", 5.00},
	}
	
	pastries := []struct {
		name  string
		price float64
	}{
		{"Croissant", 3.50},
		{"Muffin", 3.00},
		{"Scone", 3.25},
	}

	// Display menu with nested loops
	fmt.Println("📋 MENU")
	fmt.Println("=" * 40)
	
	for categoryIndex, category := range categories {
		fmt.Printf("\n%s:\n", category)
		fmt.Println("-" * 30)
		
		// Different items for each category
		switch categoryIndex {
		case 0: // Hot Drinks
			for i, item := range hotDrinks {
				fmt.Printf("  %d. %-15s $%.2f\n", i+1, item.name, item.price)
			}
		case 1: // Cold Drinks
			for i, item := range coldDrinks {
				fmt.Printf("  %d. %-15s $%.2f\n", i+1, item.name, item.price)
			}
		case 2: // Pastries
			for i, item := range pastries {
				fmt.Printf("  %d. %-15s $%.2f\n", i+1, item.name, item.price)
			}
		}
	}

	// Size and price matrix
	fmt.Println("\n\n📏 SIZE OPTIONS AND PRICING")
	fmt.Println("=" * 40)
	
	sizes := []string{"Small", "Medium", "Large"}
	sizeMultipliers := []float64{0.8, 1.0, 1.3}
	
	fmt.Printf("%-15s", "Drink")
	for _, size := range sizes {
		fmt.Printf("%-10s", size)
	}
	fmt.Println()
	fmt.Println("-" * 45)
	
	// Nested loop for price matrix
	for _, drink := range hotDrinks[:3] { // First 3 hot drinks
		fmt.Printf("%-15s", drink.name)
		for _, multiplier := range sizeMultipliers {
			price := drink.price * multiplier
			fmt.Printf("$%-9.2f", price)
		}
		fmt.Println()
	}

	// Multiplication table style display
	fmt.Println("\n\n☕ COMBO DEALS")
	fmt.Println("=" * 40)
	fmt.Println("Buy X coffees, get Y% discount:")
	fmt.Println()
	
	quantities := []int{2, 3, 4, 5}
	discounts := []int{5, 10, 15, 20}
	
	// Header
	fmt.Print("Coffees:  ")
	for _, qty := range quantities {
		fmt.Printf("%4d  ", qty)
	}
	fmt.Println("\nDiscount: ")
	
	// Nested loop for discount display
	for i := 0; i < len(quantities); i++ {
		fmt.Printf("          %3d%%  ", discounts[i])
	}
	fmt.Println()

	// Pattern printing
	fmt.Println("\n\n🎨 Coffee Cup Art")
	fmt.Println("=" * 40)
	
	// Triangle pattern
	for i := 1; i <= 5; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("☕ ")
		}
		fmt.Println()
	}
	
	// Rectangle pattern
	fmt.Println("\nCoffee Shop Floor Plan (Tables):")
	for row := 1; row <= 3; row++ {
		for col := 1; col <= 4; col++ {
			tableNum := (row-1)*4 + col
			fmt.Printf("[T%02d] ", tableNum)
		}
		fmt.Println()
	}
}