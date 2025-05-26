package main

import "fmt"

func main() {
	fmt.Println("=== GoCoffee Boolean Logic ===\n")

	// Shop status
	isOpen := true
	isBusy := false
	hasWifi := true

	// Inventory status
	hasCoffee := true
	hasMilk := true
	hasDecaf := false

	// Customer preferences
	wantsLargeCoffee := true
	wantsSugar := false
	isMember := true

	fmt.Println("SHOP STATUS:")
	fmt.Printf("Is open: %v\n", isOpen)
	fmt.Printf("Is busy: %v\n", isBusy)
	fmt.Printf("Has WiFi: %v\n", hasWifi)

	// Boolean operations
	fmt.Println("\nBOOLEAN OPERATIONS:")

	// AND operation
	canServeCustomer := isOpen && !isBusy
	fmt.Printf("Can serve customer (open AND not busy): %v\n", canServeCustomer)

	// OR operation
	needsRestock := !hasCoffee || !hasMilk
	fmt.Printf("Needs restock (no coffee OR no milk): %v\n", needsRestock)

	// Complex conditions
	canMakeLatte := hasCoffee && hasMilk && isOpen
	fmt.Printf("Can make latte: %v\n", canMakeLatte)

	// Customer eligibility
	getsDiscount := isMember && wantsLargeCoffee
	fmt.Printf("Gets member discount: %v\n", getsDiscount)

	// Comparison operations
	price := 4.50
	budget := 5.00
	minimumOrder := 3.00

	fmt.Println("\nCOMPARISON OPERATIONS:")
	fmt.Printf("Price ($%.2f) <= Budget ($%.2f): %v\n",
		price, budget, price <= budget)
	fmt.Printf("Price ($%.2f) >= Minimum ($%.2f): %v\n",
		price, minimumOrder, price >= minimumOrder)

	// Boolean as conditions
	fmt.Println("\nCONDITIONAL LOGIC:")
	if isOpen && canMakeLatte {
		fmt.Println("✓ Ready to serve lattes!")
	}

	if needsRestock {
		fmt.Println("⚠ Need to restock supplies!")
	}

	// Boolean functions
	fmt.Println("\nBOOLEAN FUNCTIONS:")
	fmt.Printf("Shop ready: %v\n", isShopReady(isOpen, hasCoffee, hasMilk))
	fmt.Printf("VIP customer: %v\n", isVIPCustomer(isMember, 15))

	// Unused variable fixes
	if hasDecaf {
		fmt.Println("We have decaf available")
	} else {
		fmt.Println("No decaf available")
	}

	if wantsSugar {
		fmt.Println("Customer wants sugar")
	}
}

// Boolean-returning functions
func isShopReady(open, coffee, milk bool) bool {
	return open && coffee && milk
}

func isVIPCustomer(member bool, ordersThisMonth int) bool {
	return member && ordersThisMonth >= 10
}