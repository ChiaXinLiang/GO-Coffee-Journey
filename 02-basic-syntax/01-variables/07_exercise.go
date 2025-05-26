package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("=== GoCoffee Order System ===")

	// Get customer name
	fmt.Print("Enter customer name: ")
	customerName, _ := reader.ReadString('\n')
	customerName = strings.TrimSpace(customerName)

	// Get coffee type
	fmt.Print("Enter coffee type (Latte/Cappuccino/Americano): ")
	coffeeType, _ := reader.ReadString('\n')
	coffeeType = strings.TrimSpace(coffeeType)

	// Get quantity
	fmt.Print("Enter quantity: ")
	quantityStr, _ := reader.ReadString('\n')
	quantityStr = strings.TrimSpace(quantityStr)
	quantity, err := strconv.Atoi(quantityStr)
	if err != nil {
		quantity = 1
	}

	// Set prices based on coffee type
	var price float64
	switch coffeeType {
	case "Latte":
		price = 4.50
	case "Cappuccino":
		price = 4.00
	case "Americano":
		price = 3.00
	default:
		price = 4.00
	}

	// Calculate total
	total := price * float64(quantity)

	// Display order
	fmt.Println("\n=== Order Summary ===")
	fmt.Printf("Customer: %s\n", customerName)
	fmt.Printf("Order: %d x %s\n", quantity, coffeeType)
	fmt.Printf("Unit price: $%.2f\n", price)
	fmt.Printf("Total: $%.2f\n", total)
}