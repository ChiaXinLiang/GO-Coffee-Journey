package main

import "fmt"

func main() {
	fmt.Println("=== GoCoffee Range Loop Examples ===\n")

	// Range over slice
	fmt.Println("☕ Coffee Menu:")
	coffeeTypes := []string{"Espresso", "Americano", "Latte", "Cappuccino", "Mocha"}
	
	for index, coffee := range coffeeTypes {
		fmt.Printf("%d. %s\n", index+1, coffee)
	}

	// Range with only values (ignore index)
	fmt.Println("\n🌟 Today's Specials:")
	specials := []string{"Caramel Macchiato", "Vanilla Latte", "Hazelnut Cappuccino"}
	
	for _, special := range specials {
		fmt.Printf("  ⭐ %s\n", special)
	}

	// Range over map
	fmt.Println("\n💰 Prices:")
	prices := map[string]float64{
		"Espresso":    2.50,
		"Americano":   3.00,
		"Latte":       4.50,
		"Cappuccino":  4.00,
		"Mocha":       5.00,
	}
	
	for coffee, price := range prices {
		fmt.Printf("%-12s: $%.2f\n", coffee, price)
	}

	// Range with only keys
	fmt.Println("\n📝 Available drinks:")
	for coffee := range prices {
		fmt.Printf("  • %s\n", coffee)
	}

	// Range over string (runes)
	fmt.Println("\n🔤 Spelling out our name:")
	shopName := "GoCoffee"
	
	for index, char := range shopName {
		fmt.Printf("Position %d: %c\n", index, char)
	}

	// Range over array
	fmt.Println("\n📊 Daily sales (last week):")
	dailySales := [7]float64{1250.50, 1180.25, 1320.75, 1405.00, 1520.50, 1850.25, 1680.00}
	days := [7]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	
	totalSales := 0.0
	for i, sales := range dailySales {
		fmt.Printf("%s: $%.2f\n", days[i], sales)
		totalSales += sales
	}
	fmt.Printf("\nTotal weekly sales: $%.2f\n", totalSales)
	fmt.Printf("Average daily sales: $%.2f\n", totalSales/7)

	// Nested range loops
	fmt.Println("\n🍰 Pastry inventory by category:")
	pastryInventory := map[string][]string{
		"Croissants": {"Plain", "Chocolate", "Almond"},
		"Muffins":    {"Blueberry", "Chocolate Chip", "Bran"},
		"Cookies":    {"Chocolate Chip", "Oatmeal", "Sugar"},
	}
	
	for category, items := range pastryInventory {
		fmt.Printf("\n%s:\n", category)
		for i, item := range items {
			fmt.Printf("  %d. %s %s\n", i+1, item, category[:len(category)-1])
		}
	}

	// Range with struct slice
	fmt.Println("\n👥 Staff schedule:")
	
	type Staff struct {
		name  string
		shift string
		hours int
	}
	
	schedule := []Staff{
		{"Marcus", "Morning", 6},
		{"Sarah", "Morning", 8},
		{"Carlos", "Afternoon", 8},
		{"Elena", "Evening", 6},
	}
	
	for _, staff := range schedule {
		fmt.Printf("%-10s: %-10s shift (%d hours)\n", 
			staff.name, staff.shift, staff.hours)
	}

	// Calculate with range
	fmt.Println("\n📈 Customer ratings:")
	ratings := []int{5, 4, 5, 3, 5, 4, 4, 5, 5, 4}
	
	sum := 0
	for _, rating := range ratings {
		sum += rating
		
		// Display stars
		fmt.Print("Rating: ")
		for i := 0; i < rating; i++ {
			fmt.Print("⭐")
		}
		fmt.Println()
	}
	
	average := float64(sum) / float64(len(ratings))
	fmt.Printf("\nAverage rating: %.1f ⭐\n", average)
}