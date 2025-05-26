package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("=== GoCoffee Continue Statement Basics ===\n")

	// Example 1: Skip specific items
	fmt.Println("☕ Processing Orders (No Decaf Today)")
	fmt.Println("------------------------------------")
	
	orders := []string{
		"Espresso",
		"Decaf Latte",
		"Cappuccino", 
		"Decaf Americano",
		"Mocha",
		"Decaf Espresso",
		"Macchiato",
	}
	
	servedCount := 0
	
	for i, order := range orders {
		fmt.Printf("Order %d: %s - ", i+1, order)
		
		if strings.Contains(order, "Decaf") {
			fmt.Println("❌ SKIPPED (Decaf machine broken)")
			continue
		}
		
		servedCount++
		fmt.Println("✓ Prepared")
	}
	
	fmt.Printf("\nServed %d out of %d orders\n", servedCount, len(orders))

	// Example 2: Continue with validation
	fmt.Println("\n\n💳 Payment Processing")
	fmt.Println("--------------------")
	
	payments := []struct {
		customer string
		amount   float64
		valid    bool
	}{
		{"Alice", 12.50, true},
		{"Bob", -5.00, false},      // Invalid: negative
		{"Charlie", 8.75, true},
		{"Diana", 0.00, false},      // Invalid: zero
		{"Eve", 15.25, true},
		{"Frank", 999.99, false},    // Invalid: too high
	}
	
	totalRevenue := 0.0
	
	for _, payment := range payments {
		fmt.Printf("%-10s: $%6.2f - ", payment.customer, payment.amount)
		
		// Skip invalid payments
		if !payment.valid {
			fmt.Println("⚠️  SKIPPED (Invalid payment)")
			continue
		}
		
		if payment.amount <= 0 {
			fmt.Println("⚠️  SKIPPED (Invalid amount)")
			continue
		}
		
		if payment.amount > 100 {
			fmt.Println("⚠️  SKIPPED (Amount too high)")
			continue
		}
		
		totalRevenue += payment.amount
		fmt.Println("✓ Processed")
	}
	
	fmt.Printf("\nTotal revenue: $%.2f\n", totalRevenue)

	// Example 3: Continue with modulo (every nth item)
	fmt.Println("\n\n🎁 Customer Rewards (Every 3rd Customer)")
	fmt.Println("---------------------------------------")
	
	for customer := 1; customer <= 15; customer++ {
		fmt.Printf("Customer #%2d: ", customer)
		
		if customer%3 != 0 {
			fmt.Println("Regular service")
			continue
		}
		
		fmt.Println("🎉 FREE COOKIE! (Every 3rd customer)")
	}

	// Example 4: Continue with multiple conditions
	fmt.Println("\n\n📊 Staff Performance Review")
	fmt.Println("--------------------------")
	
	staff := []struct {
		name       string
		shifts     int
		rating     float64
		probation  bool
	}{
		{"Marcus", 20, 4.5, false},
		{"Sarah", 18, 4.8, false},
		{"Tom", 5, 3.2, true},      // Skip: probation
		{"Elena", 22, 4.6, false},
		{"Jack", 3, 2.8, false},    // Skip: too few shifts
		{"Linda", 19, 3.9, false},
		{"Mike", 15, 2.5, false},   // Skip: low rating
	}
	
	fmt.Println("Eligible for bonus (>10 shifts, >3.5 rating, not on probation):\n")
	
	for _, employee := range staff {
		// Skip if on probation
		if employee.probation {
			fmt.Printf("%-10s: On probation - SKIPPED\n", employee.name)
			continue
		}
		
		// Skip if too few shifts
		if employee.shifts < 10 {
			fmt.Printf("%-10s: Only %d shifts - SKIPPED\n", employee.name, employee.shifts)
			continue
		}
		
		// Skip if rating too low
		if employee.rating < 3.5 {
			fmt.Printf("%-10s: Rating %.1f - SKIPPED\n", employee.name, employee.rating)
			continue
		}
		
		// Eligible for bonus
		fmt.Printf("%-10s: ✓ ELIGIBLE (%.1f rating, %d shifts)\n", 
			employee.name, employee.rating, employee.shifts)
	}

	// Example 5: Continue in string processing
	fmt.Println("\n\n🔤 Menu Code Generator")
	fmt.Println("---------------------")
	fmt.Println("(Remove vowels and spaces for item codes)")
	
	menuItems := []string{
		"Espresso Shot",
		"Cafe Latte",
		"Iced Americano",
		"Mocha Frappuccino",
	}
	
	for _, item := range menuItems {
		fmt.Printf("%-20s → ", item)
		
		code := ""
		for _, char := range strings.ToUpper(item) {
			// Skip vowels
			if char == 'A' || char == 'E' || char == 'I' || char == 'O' || char == 'U' {
				continue
			}
			
			// Skip spaces
			if char == ' ' {
				continue
			}
			
			code += string(char)
		}
		
		fmt.Printf("Code: %s\n", code)
	}
}