package main

import (
	"fmt"
	"strings"
	"time"
)

// Menu state constants
const (
	StateMain = iota
	StateDrinks
	StateFood
	StateCheckout
	StateExit
)

// Shopping cart
type Cart struct {
	items  []string
	prices []float64
	total  float64
}

func main() {
	fmt.Println("=== GoCoffee Interactive Menu System ===\n")
	
	// Initialize cart
	cart := &Cart{
		items:  []string{},
		prices: []float64{},
		total:  0.0,
	}
	
	// Start menu system
	state := StateMain
	
	for state != StateExit {
		clearScreen()
		
		switch state {
		case StateMain:
			state = showMainMenu()
			
		case StateDrinks:
			state = showDrinksMenu(cart)
			
		case StateFood:
			state = showFoodMenu(cart)
			
		case StateCheckout:
			state = showCheckout(cart)
			
		default:
			fmt.Println("Unknown state!")
			state = StateMain
		}
		
		if state != StateExit {
			fmt.Println("\nPress Enter to continue...")
			fmt.Scanln()
		}
	}
	
	fmt.Println("\nThank you for visiting GoCoffee! ☕")
}

func showMainMenu() int {
	fmt.Println("🏠 MAIN MENU")
	fmt.Println("============")
	fmt.Println()
	fmt.Println("1. ☕ Drinks Menu")
	fmt.Println("2. 🍰 Food Menu")
	fmt.Println("3. 🛒 Checkout")
	fmt.Println("4. 🚪 Exit")
	fmt.Println()
	
	choice := getMenuChoice(1, 4)
	
	switch choice {
	case 1:
		return StateDrinks
	case 2:
		return StateFood
	case 3:
		return StateCheckout
	case 4:
		return StateExit
	default:
		return StateMain
	}
}

func showDrinksMenu(cart *Cart) int {
	fmt.Println("☕ DRINKS MENU")
	fmt.Println("=============")
	fmt.Println()
	
	drinks := []struct {
		name  string
		price float64
	}{
		{"Espresso", 2.50},
		{"Americano", 3.00},
		{"Latte", 4.50},
		{"Cappuccino", 4.00},
		{"Mocha", 5.00},
		{"Macchiato", 4.25},
	}
	
	// Display drinks
	for i, drink := range drinks {
		fmt.Printf("%d. %-15s $%.2f\n", i+1, drink.name, drink.price)
	}
	fmt.Printf("%d. ← Back to Main Menu\n", len(drinks)+1)
	fmt.Println()
	
	choice := getMenuChoice(1, len(drinks)+1)
	
	// Handle choice
	switch {
	case choice >= 1 && choice <= len(drinks):
		// Add drink to cart
		selectedDrink := drinks[choice-1]
		addToCart(cart, selectedDrink.name, selectedDrink.price)
		
		// Ask for customization
		customizeDrink(cart, len(cart.items)-1)
		
		return StateDrinks // Stay in drinks menu
		
	case choice == len(drinks)+1:
		return StateMain
		
	default:
		return StateDrinks
	}
}

func showFoodMenu(cart *Cart) int {
	fmt.Println("🍰 FOOD MENU")
	fmt.Println("===========")
	fmt.Println()
	
	foods := []struct {
		name  string
		price float64
	}{
		{"Croissant", 3.50},
		{"Muffin", 3.00},
		{"Bagel", 4.00},
		{"Sandwich", 8.50},
		{"Salad", 9.00},
		{"Cookie", 2.50},
	}
	
	// Display food items
	for i, food := range foods {
		fmt.Printf("%d. %-15s $%.2f\n", i+1, food.name, food.price)
	}
	fmt.Printf("%d. ← Back to Main Menu\n", len(foods)+1)
	fmt.Println()
	
	choice := getMenuChoice(1, len(foods)+1)
	
	// Handle choice
	switch {
	case choice >= 1 && choice <= len(foods):
		// Add food to cart
		selectedFood := foods[choice-1]
		addToCart(cart, selectedFood.name, selectedFood.price)
		return StateFood // Stay in food menu
		
	case choice == len(foods)+1:
		return StateMain
		
	default:
		return StateFood
	}
}

func showCheckout(cart *Cart) int {
	fmt.Println("🛒 CHECKOUT")
	fmt.Println("==========")
	fmt.Println()
	
	if len(cart.items) == 0 {
		fmt.Println("Your cart is empty!")
		return StateMain
	}
	
	// Display cart items
	fmt.Println("Your order:")
	fmt.Println(strings.Repeat("-", 35))
	
	for i, item := range cart.items {
		fmt.Printf("%-25s $%.2f\n", item, cart.prices[i])
	}
	
	fmt.Println(strings.Repeat("-", 35))
	fmt.Printf("%-25s $%.2f\n", "Subtotal:", cart.total)
	
	tax := cart.total * 0.08
	fmt.Printf("%-25s $%.2f\n", "Tax (8%):", tax)
	
	total := cart.total + tax
	fmt.Printf("%-25s $%.2f\n", "TOTAL:", total)
	fmt.Println()
	
	fmt.Println("1. 💳 Pay Now")
	fmt.Println("2. ➕ Add More Items")
	fmt.Println("3. ❌ Cancel Order")
	fmt.Println()
	
	choice := getMenuChoice(1, 3)
	
	switch choice {
	case 1:
		processPayment(total)
		// Clear cart after payment
		cart.items = []string{}
		cart.prices = []float64{}
		cart.total = 0.0
		return StateMain
		
	case 2:
		return StateMain
		
	case 3:
		// Clear cart
		cart.items = []string{}
		cart.prices = []float64{}
		cart.total = 0.0
		fmt.Println("\n❌ Order cancelled")
		return StateMain
		
	default:
		return StateCheckout
	}
}

func customizeDrink(cart *Cart, index int) {
	fmt.Println("\n☕ Customize your drink:")
	fmt.Println("1. Regular")
	fmt.Println("2. Add extra shot (+$0.50)")
	fmt.Println("3. Decaf")
	fmt.Println("4. Add flavor syrup (+$0.75)")
	
	choice := getMenuChoice(1, 4)
	
	switch choice {
	case 2:
		cart.items[index] += " + Extra Shot"
		cart.prices[index] += 0.50
		cart.total += 0.50
		fmt.Println("✓ Extra shot added")
		
	case 3:
		cart.items[index] = "Decaf " + cart.items[index]
		fmt.Println("✓ Made decaf")
		
	case 4:
		cart.items[index] += " + Vanilla Syrup"
		cart.prices[index] += 0.75
		cart.total += 0.75
		fmt.Println("✓ Vanilla syrup added")
	}
}

func addToCart(cart *Cart, item string, price float64) {
	cart.items = append(cart.items, item)
	cart.prices = append(cart.prices, price)
	cart.total += price
	
	fmt.Printf("\n✓ Added %s to cart ($%.2f)\n", item, price)
}

func processPayment(total float64) {
	fmt.Println("\n💳 PAYMENT PROCESSING")
	fmt.Println("====================")
	fmt.Printf("Amount due: $%.2f\n", total)
	fmt.Println()
	
	fmt.Println("Select payment method:")
	fmt.Println("1. Cash")
	fmt.Println("2. Credit Card")
	fmt.Println("3. Debit Card")
	fmt.Println("4. Mobile Pay")
	
	choice := getMenuChoice(1, 4)
	
	var method string
	switch choice {
	case 1:
		method = "Cash"
	case 2:
		method = "Credit Card"
	case 3:
		method = "Debit Card"
	case 4:
		method = "Mobile Pay"
	}
	
	fmt.Printf("\nProcessing %s payment...\n", method)
	time.Sleep(2 * time.Second)
	
	fmt.Println("\n✅ Payment successful!")
	fmt.Println("\n--- RECEIPT ---")
	fmt.Printf("Date: %s\n", time.Now().Format("Jan 2, 2006 3:04 PM"))
	fmt.Printf("Method: %s\n", method)
	fmt.Printf("Total: $%.2f\n", total)
	fmt.Println("Thank you!")
	fmt.Println("---------------")
}

func getMenuChoice(min, max int) int {
	// Simulate user input
	// In a real program, you'd read from stdin
	choices := []int{1, 1, 2, 3, 1, 3, 1, 4}
	static choiceIndex int
	
	if choiceIndex < len(choices) {
		choice := choices[choiceIndex]
		choiceIndex++
		fmt.Printf("Your choice: %d\n", choice)
		return choice
	}
	
	// Default to exit
	return max
}

func clearScreen() {
	// Simple clear screen simulation
	fmt.Print("\033[H\033[2J")
	fmt.Println()
}