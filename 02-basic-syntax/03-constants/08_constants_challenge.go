package main

import "fmt"

// TODO: Create a comprehensive pricing system using constants
// Requirements:
// 1. Define coffee types with base prices
// 2. Define size modifiers
// 3. Define temperature options
// 4. Define milk options with prices
// 5. Calculate final prices

func main() {
	fmt.Println("=== GoCoffee Constants Challenge ===\n")
	fmt.Println("Create a pricing system with these requirements:")
	fmt.Println("1. Coffee types: Espresso, Latte, Cappuccino, Mocha")
	fmt.Println("2. Sizes: Small (12oz), Medium (16oz), Large (20oz)")
	fmt.Println("3. Temperatures: Hot, Iced, Blended")
	fmt.Println("4. Milk options: Regular, Skim, Soy, Almond, Oat")
	fmt.Println("5. Extra options: Extra shot, Decaf, Sugar-free")

	// Your implementation here...

	// Hint: Use iota for enumerations
	// Hint: Use typed constants for type safety
	// Hint: Group related constants together
}

// Example solution structure:
/*
type CoffeeType int
const (
    Espresso CoffeeType = iota
    Latte
    Cappuccino
    Mocha
)

type Size int
const (
    Small Size = iota
    Medium
    Large
)

// ... continue implementation
*/