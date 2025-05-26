package main

import "fmt"

func main() {
	fmt.Println("=== GoCoffee Inventory Tracking System ===\n")

	// Working with arrays using index
	inventory := [5]string{"Coffee Beans", "Milk", "Sugar", "Cups", "Stirrers"}
	quantities := [5]int{50, 30, 100, 200, 500}
	units := [5]string{"kg", "liters", "kg", "pieces", "pieces"}

	fmt.Println("Current Inventory:")
	fmt.Println("------------------")
	
	for i := 0; i < len(inventory); i++ {
		fmt.Printf("%-15s: %4d %s\n", inventory[i], quantities[i], units[i])
	}

	// Checking low stock items
	fmt.Println("\n🚨 Low Stock Alert (less than 50 units):")
	fmt.Println("----------------------------------------")
	
	lowStockCount := 0
	for i := 0; i < len(quantities); i++ {
		if quantities[i] < 50 {
			fmt.Printf("⚠️  %-15s: Only %d %s left!\n", 
				inventory[i], quantities[i], units[i])
			lowStockCount++
		}
	}
	
	if lowStockCount == 0 {
		fmt.Println("✓ All items well stocked!")
	}

	// Daily usage simulation
	fmt.Println("\n📊 Daily Usage Simulation:")
	fmt.Println("--------------------------")
	
	dailyUsage := [5]int{5, 10, 2, 50, 100}
	
	for day := 1; day <= 7; day++ {
		fmt.Printf("\nDay %d:\n", day)
		
		for i := 0; i < len(inventory); i++ {
			quantities[i] -= dailyUsage[i]
			
			if quantities[i] <= 0 {
				fmt.Printf("  ❌ %s: OUT OF STOCK!\n", inventory[i])
				quantities[i] = 0
			} else if quantities[i] < 20 {
				fmt.Printf("  ⚠️  %s: %d %s (LOW)\n", 
					inventory[i], quantities[i], units[i])
			} else {
				fmt.Printf("  ✓ %s: %d %s\n", 
					inventory[i], quantities[i], units[i])
			}
		}
	}

	// Restock needed calculation
	fmt.Println("\n📦 Restock Requirements:")
	fmt.Println("------------------------")
	
	minStock := [5]int{40, 25, 80, 150, 400}
	
	for i := 0; i < len(inventory); i++ {
		if quantities[i] < minStock[i] {
			needed := minStock[i] - quantities[i]
			fmt.Printf("Order %d %s of %s\n", needed, units[i], inventory[i])
		}
	}
}