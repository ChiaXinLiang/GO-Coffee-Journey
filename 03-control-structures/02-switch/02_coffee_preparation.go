package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("=== GoCoffee Preparation System ===\n")

	// Prepare different coffee types
	prepareCoffee("Espresso")
	fmt.Println()
	
	prepareCoffee("Latte")
	fmt.Println()
	
	prepareCoffee("Cappuccino")
	fmt.Println()
	
	prepareCoffee("Mocha")
	fmt.Println()
	
	prepareCoffee("Tea") // Unknown type
}

func prepareCoffee(coffeeType string) {
	fmt.Printf("Preparing: %s\n", coffeeType)
	fmt.Println(strings.Repeat("-", 30))
	
	switch coffeeType {
	case "Espresso":
		fmt.Println("1. Grind 18g of coffee beans")
		fmt.Println("2. Tamp with 30 pounds of pressure")
		fmt.Println("3. Extract for 25-30 seconds")
		fmt.Println("4. Serve in small cup")
		fmt.Println("☕ Espresso ready!")
		
	case "Americano":
		fmt.Println("1. Prepare double espresso")
		fmt.Println("2. Add 6oz hot water")
		fmt.Println("3. Stir gently")
		fmt.Println("4. Serve in medium cup")
		fmt.Println("☕ Americano ready!")
		
	case "Latte":
		fmt.Println("1. Prepare single espresso")
		fmt.Println("2. Steam 8oz milk to 150°F")
		fmt.Println("3. Pour milk over espresso")
		fmt.Println("4. Create latte art")
		fmt.Println("5. Serve in large cup")
		fmt.Println("☕ Latte ready!")
		
	case "Cappuccino":
		fmt.Println("1. Prepare double espresso")
		fmt.Println("2. Steam 4oz milk to 160°F")
		fmt.Println("3. Create thick foam")
		fmt.Println("4. Pour 1/3 espresso, 1/3 milk, 1/3 foam")
		fmt.Println("5. Dust with cocoa powder")
		fmt.Println("☕ Cappuccino ready!")
		
	case "Mocha":
		fmt.Println("1. Add 2 tbsp chocolate syrup to cup")
		fmt.Println("2. Prepare double espresso")
		fmt.Println("3. Mix espresso with chocolate")
		fmt.Println("4. Steam 6oz milk")
		fmt.Println("5. Pour milk and top with whipped cream")
		fmt.Println("☕ Mocha ready!")
		
	case "Macchiato":
		fmt.Println("1. Prepare double espresso")
		fmt.Println("2. Steam small amount of milk")
		fmt.Println("3. Add just a 'mark' of foam")
		fmt.Println("4. Serve in small cup")
		fmt.Println("☕ Macchiato ready!")
		
	default:
		fmt.Println("❌ Unknown beverage type!")
		fmt.Println("Available options:")
		fmt.Println("- Espresso")
		fmt.Println("- Americano")
		fmt.Println("- Latte")
		fmt.Println("- Cappuccino")
		fmt.Println("- Mocha")
		fmt.Println("- Macchiato")
	}
}