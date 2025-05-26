package main

import "fmt"

// Simple function with no parameters and no return value
func greetShop() {
	fmt.Println("☕ Welcome to GoCoffee!")
	fmt.Println("   The best coffee in town!")
}

// Another simple function
func displayOpenHours() {
	fmt.Println("\n📍 Open Hours:")
	fmt.Println("   Monday - Friday: 6:00 AM - 8:00 PM")
	fmt.Println("   Saturday - Sunday: 7:00 AM - 9:00 PM")
}

// Function that calls other functions
func showShopInfo() {
	greetShop()
	displayOpenHours()
	showSpecials()
}

// Function showing today's specials
func showSpecials() {
	fmt.Println("\n✨ Today's Specials:")
	fmt.Println("   • Caramel Macchiato - 20% off")
	fmt.Println("   • All pastries - Buy 2 get 1 free")
}

// Function demonstrating function flow
func morningRoutine() {
	fmt.Println("\n🌅 Morning Routine Starting...")
	
	turnOnLights()
	startMachines()
	brewFirstCoffee()
	openShop()
	
	fmt.Println("✅ Ready for customers!")
}

func turnOnLights() {
	fmt.Println("   💡 Turning on lights...")
}

func startMachines() {
	fmt.Println("   ⚙️  Starting coffee machines...")
}

func brewFirstCoffee() {
	fmt.Println("   ☕ Brewing first pot of coffee...")
}

func openShop() {
	fmt.Println("   🚪 Opening shop doors...")
}

// Main function - entry point of the program
func main() {
	fmt.Println("=== GoCoffee Simple Functions ===\n")
	
	// Calling our functions
	greetShop()
	
	// Functions can be called multiple times
	fmt.Println("\n--- First Call ---")
	displayOpenHours()
	
	fmt.Println("\n--- Second Call ---")
	displayOpenHours()
	
	// Calling a function that calls other functions
	fmt.Println("\n--- Shop Information ---")
	showShopInfo()
	
	// Demonstrating a sequence of function calls
	morningRoutine()
	
	// Functions help organize code
	fmt.Println("\n💡 Key Points:")
	fmt.Println("• Functions group related code together")
	fmt.Println("• Functions can be called multiple times")
	fmt.Println("• Functions can call other functions")
	fmt.Println("• Functions make code more organized and readable")
}