package main

import (
	"fmt"
	"strings"
)

// Challenge: Create a type-safe menu system
// TODO: Define appropriate types for:
// 1. Menu categories (drinks, food, etc.)
// 2. Dietary restrictions (vegan, gluten-free, etc.)
// 3. Availability status

func main() {
	fmt.Println("=== GoCoffee Type Challenge ===\n")
	fmt.Println("TODO: Implement a type-safe menu system")
	fmt.Println("Requirements:")
	fmt.Println("- Use custom types for categories")
	fmt.Println("- Handle dietary restrictions")
	fmt.Println("- Track item availability")
	fmt.Println("- Calculate prices correctly")

	// Starter code
	type Category string
	const (
		CategoryDrinks   Category = "drinks"
		CategoryFood     Category = "food"
		CategoryDesserts Category = "desserts"
	)

	// Example solution (uncomment to see one way to solve it):
	/*
		// Dietary restrictions as bit flags
		type DietaryFlag uint8
		const (
			DietaryNone       DietaryFlag = 0
			DietaryVegan      DietaryFlag = 1 << iota
			DietaryGlutenFree
			DietaryDairyFree
			DietaryNutFree
		)

		type MenuItem struct {
			ID           string
			Name         string
			Category     Category
			PriceCents   int
			Dietary      DietaryFlag
			Available    bool
			Description  string
		}

		// Example menu items
		menu := []MenuItem{
			{
				ID:          "ESP001",
				Name:        "Espresso",
				Category:    CategoryDrinks,
				PriceCents:  300,
				Dietary:     DietaryVegan | DietaryGlutenFree,
				Available:   true,
				Description: "Rich, bold espresso shot",
			},
			{
				ID:          "CRO001",
				Name:        "Butter Croissant",
				Category:    CategoryFood,
				PriceCents:  350,
				Dietary:     DietaryNone,
				Available:   true,
				Description: "Flaky, buttery croissant",
			},
		}

		// Display menu
		for _, item := range menu {
			fmt.Printf("\n%s - %s\n", item.ID, item.Name)
			fmt.Printf("  Category: %s\n", item.Category)
			fmt.Printf("  Price: $%.2f\n", float64(item.PriceCents)/100)
			fmt.Printf("  Available: %v\n", item.Available)

			// Check dietary flags
			if item.Dietary&DietaryVegan != 0 {
				fmt.Println("  ✓ Vegan")
			}
			if item.Dietary&DietaryGlutenFree != 0 {
				fmt.Println("  ✓ Gluten-Free")
			}
		}
	*/

	// Your implementation here...
	fmt.Println("\n" + strings.Repeat("-", 40))
	fmt.Println("Implement your solution above!")
}