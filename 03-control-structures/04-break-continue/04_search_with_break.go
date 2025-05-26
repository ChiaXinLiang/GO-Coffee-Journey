package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	fmt.Println("=== GoCoffee Search Operations with Break ===\n")

	// Example 1: Linear search
	fmt.Println("🔍 Finding Customer Order")
	fmt.Println("------------------------")
	
	orders := []struct {
		id       string
		customer string
		drink    string
		status   string
	}{
		{"ORD-001", "Alice", "Latte", "preparing"},
		{"ORD-002", "Bob", "Espresso", "ready"},
		{"ORD-003", "Charlie", "Cappuccino", "preparing"},
		{"ORD-004", "Diana", "Mocha", "ready"},
		{"ORD-005", "Eve", "Americano", "preparing"},
	}
	
	searchID := "ORD-003"
	found := false
	var foundOrder struct {
		id       string
		customer string
		drink    string
		status   string
	}
	
	startTime := time.Now()
	comparisons := 0
	
	for _, order := range orders {
		comparisons++
		fmt.Printf("Checking %s... ", order.id)
		
		if order.id == searchID {
			found = true
			foundOrder = order
			fmt.Println("✓ FOUND!")
			break
		}
		fmt.Println("not a match")
	}
	
	searchTime := time.Since(startTime)
	
	if found {
		fmt.Printf("\nOrder Details:\n")
		fmt.Printf("  Customer: %s\n", foundOrder.customer)
		fmt.Printf("  Drink: %s\n", foundOrder.drink)
		fmt.Printf("  Status: %s\n", foundOrder.status)
		fmt.Printf("\nSearch stats: %d comparisons in %v\n", comparisons, searchTime)
	}

	// Example 2: Finding first available item
	fmt.Println("\n\n☕ Finding First Available Barista")
	fmt.Println("---------------------------------")
	
	baristas := []struct {
		name      string
		available bool
		specialty string
	}{
		{"Carlos", false, "Espresso"},
		{"Maria", false, "Latte Art"},
		{"John", true, "Cold Brew"},
		{"Sarah", true, "Cappuccino"},
		{"Mike", false, "Americano"},
	}
	
	var assignedBarista string
	
	for _, barista := range baristas {
		fmt.Printf("Checking %s... ", barista.name)
		
		if !barista.available {
			fmt.Println("busy")
			continue
		}
		
		assignedBarista = barista.name
		fmt.Printf("available! (Specialty: %s)\n", barista.specialty)
		break
	}
	
	if assignedBarista != "" {
		fmt.Printf("\n✓ Assigned to: %s\n", assignedBarista)
	} else {
		fmt.Println("\n❌ No baristas available!")
	}

	// Example 3: Multi-criteria search
	fmt.Println("\n\n🎯 Finding Perfect Coffee Bean")
	fmt.Println("-----------------------------")
	
	coffeeBeans := []struct {
		origin    string
		roast     string
		price     float64
		rating    float64
		organic   bool
		available int
	}{
		{"Colombia", "Medium", 12.99, 4.2, false, 50},
		{"Ethiopia", "Light", 15.99, 4.5, true, 0},
		{"Brazil", "Dark", 10.99, 3.8, false, 100},
		{"Kenya", "Medium", 14.99, 4.7, true, 25},
		{"Peru", "Medium", 13.99, 4.4, true, 75},
	}
	
	// Search criteria
	maxPrice := 14.00
	minRating := 4.0
	preferOrganic := true
	minStock := 20
	
	fmt.Printf("Search criteria:\n")
	fmt.Printf("  Max price: $%.2f\n", maxPrice)
	fmt.Printf("  Min rating: %.1f\n", minRating)
	fmt.Printf("  Organic: %v\n", preferOrganic)
	fmt.Printf("  Min stock: %d kg\n\n", minStock)
	
	perfectBean := ""
	
	for _, bean := range coffeeBeans {
		fmt.Printf("Checking %s %s Roast... ", bean.origin, bean.roast)
		
		// Check all criteria
		if bean.price > maxPrice {
			fmt.Printf("too expensive ($%.2f)\n", bean.price)
			continue
		}
		
		if bean.rating < minRating {
			fmt.Printf("rating too low (%.1f)\n", bean.rating)
			continue
		}
		
		if preferOrganic && !bean.organic {
			fmt.Println("not organic")
			continue
		}
		
		if bean.available < minStock {
			fmt.Printf("low stock (%d kg)\n", bean.available)
			continue
		}
		
		// Found perfect match!
		perfectBean = bean.origin
		fmt.Println("✓ PERFECT MATCH!")
		break
	}
	
	if perfectBean != "" {
		fmt.Printf("\n🎉 Selected: %s coffee beans\n", perfectBean)
	} else {
		fmt.Println("\n😞 No beans match all criteria")
	}

	// Example 4: Search with early termination conditions
	fmt.Println("\n\n📋 Recipe Search System")
	fmt.Println("----------------------")
	
	recipes := []struct {
		name       string
		difficulty string
		time       int // minutes
		popular    bool
	}{
		{"Classic Espresso", "Easy", 2, true},
		{"Vanilla Latte", "Medium", 5, true},
		{"Caramel Macchiato", "Hard", 8, true},
		{"Flat White", "Medium", 4, false},
		{"Affogato", "Easy", 3, false},
		{"Turkish Coffee", "Hard", 10, false},
	}
	
	// Search for quick and easy recipe
	fmt.Println("Looking for: Quick (<5 min) and Easy recipe\n")
	
	foundRecipe := ""
	searchAttempts := 0
	maxAttempts := 3 // Give up after 3 attempts
	
	for _, recipe := range recipes {
		searchAttempts++
		fmt.Printf("Attempt %d: %s... ", searchAttempts, recipe.name)
		
		if recipe.difficulty != "Easy" {
			fmt.Println("too difficult")
			if searchAttempts >= maxAttempts {
				fmt.Println("\n⏱️  Search timeout - too many attempts")
				break
			}
			continue
		}
		
		if recipe.time >= 5 {
			fmt.Printf("too slow (%d min)\n", recipe.time)
			if searchAttempts >= maxAttempts {
				fmt.Println("\n⏱️  Search timeout - too many attempts")
				break
			}
			continue
		}
		
		foundRecipe = recipe.name
		fmt.Printf("✓ Perfect! (%d min, %s)\n", recipe.time, recipe.difficulty)
		break
	}
	
	if foundRecipe != "" {
		fmt.Printf("\n✓ Recommended: %s\n", foundRecipe)
	}

	// Example 5: Pattern matching search
	fmt.Println("\n\n🔤 Customer Name Search")
	fmt.Println("----------------------")
	
	customers := []string{
		"Alice Anderson",
		"Bob Brown",
		"Charlie Chen",
		"Diana Davis",
		"Eve Evans",
		"Frank Fisher",
		"Grace Green",
		"Henry Harris",
	}
	
	searchPattern := "Harris"
	fmt.Printf("Searching for names containing: '%s'\n\n", searchPattern)
	
	matchFound := false
	var matchedCustomer string
	
	for i, customer := range customers {
		fmt.Printf("%d. Checking '%s'... ", i+1, customer)
		
		if strings.Contains(customer, searchPattern) {
			matchFound = true
			matchedCustomer = customer
			fmt.Println("✓ MATCH!")
			break
		}
		
		fmt.Println("no match")
		
		// Add small delay for visual effect
		time.Sleep(200 * time.Millisecond)
	}
	
	if matchFound {
		fmt.Printf("\n✓ Found customer: %s\n", matchedCustomer)
	} else {
		fmt.Printf("\n❌ No customer found with pattern '%s'\n", searchPattern)
	}
}