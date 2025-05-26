package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("=== GoCoffee Data Validation with Continue ===\n")

	// Example 1: Order validation
	fmt.Println("📝 Order Validation System")
	fmt.Println("-------------------------")
	
	orders := []struct {
		id          string
		customerID  string
		items       []string
		totalAmount float64
		paid        bool
	}{
		{"ORD001", "CUST123", []string{"Latte", "Muffin"}, 8.50, true},
		{"ORD002", "", []string{"Espresso"}, 3.00, true},                    // Invalid: no customer
		{"ORD003", "CUST456", []string{}, 0.00, false},                      // Invalid: no items
		{"ORD004", "CUST789", []string{"Cappuccino"}, -5.00, true},         // Invalid: negative amount
		{"ORD005", "CUST234", []string{"Mocha", "Cookie"}, 9.50, false},    // Invalid: not paid
		{"ORD006", "CUST567", []string{"Americano"}, 4.00, true},
	}
	
	validOrders := 0
	totalRevenue := 0.0
	
	fmt.Println("\nProcessing orders:")
	for _, order := range orders {
		fmt.Printf("\n%s: ", order.id)
		
		// Validate customer ID
		if order.customerID == "" {
			fmt.Print("❌ Missing customer ID")
			continue
		}
		
		// Validate items
		if len(order.items) == 0 {
			fmt.Print("❌ No items in order")
			continue
		}
		
		// Validate amount
		if order.totalAmount <= 0 {
			fmt.Printf("❌ Invalid amount: $%.2f", order.totalAmount)
			continue
		}
		
		// Validate payment
		if !order.paid {
			fmt.Print("❌ Payment not completed")
			continue
		}
		
		// Valid order
		validOrders++
		totalRevenue += order.totalAmount
		fmt.Printf("✓ Valid - %s bought %v for $%.2f", 
			order.customerID, order.items, order.totalAmount)
	}
	
	fmt.Printf("\n\nSummary: %d valid orders, Total revenue: $%.2f\n", 
		validOrders, totalRevenue)

	// Example 2: Email validation
	fmt.Println("\n\n📧 Customer Email Validation")
	fmt.Println("---------------------------")
	
	emails := []string{
		"alice@coffee.com",
		"bob.brown",              // Missing @
		"charlie@",               // Missing domain
		"@example.com",          // Missing local part
		"diana@coffee.shop",
		"eve.evans@go-coffee.io",
		"frank@@double.com",     // Double @
		"grace@coffee",          // Missing TLD
		"henry@valid.co.uk",
	}
	
	validEmails := []string{}
	
	for _, email := range emails {
		fmt.Printf("%-25s: ", email)
		
		// Basic validation checks
		if !strings.Contains(email, "@") {
			fmt.Println("❌ Missing @ symbol")
			continue
		}
		
		parts := strings.Split(email, "@")
		if len(parts) != 2 {
			fmt.Println("❌ Invalid @ usage")
			continue
		}
		
		localPart := parts[0]
		domain := parts[1]
		
		if len(localPart) == 0 {
			fmt.Println("❌ Missing local part")
			continue
		}
		
		if len(domain) == 0 {
			fmt.Println("❌ Missing domain")
			continue
		}
		
		if !strings.Contains(domain, ".") {
			fmt.Println("❌ Invalid domain format")
			continue
		}
		
		// Valid email
		validEmails = append(validEmails, email)
		fmt.Println("✓ Valid")
	}
	
	fmt.Printf("\nValid emails: %v\n", validEmails)

	// Example 3: Password strength validation
	fmt.Println("\n\n🔐 Password Strength Checker")
	fmt.Println("---------------------------")
	
	passwords := []struct {
		user     string
		password string
	}{
		{"alice", "12345"},           // Too weak
		{"bob", "password"},          // No numbers
		{"charlie", "Pass123"},       // Too short
		{"diana", "SecurePass123!"},  // Good
		{"eve", "NoSpecialChar8"},    // No special
		{"frank", "Good@Pass9"},      // Good
		{"grace", "abc"},             // Way too short
	}
	
	fmt.Println("Requirements: 8+ chars, uppercase, lowercase, number, special char\n")
	
	for _, entry := range passwords {
		fmt.Printf("%-10s: %-16s - ", entry.user, entry.password)
		
		pwd := entry.password
		
		// Check length
		if len(pwd) < 8 {
			fmt.Printf("❌ Too short (%d chars)\n", len(pwd))
			continue
		}
		
		// Check for required character types
		hasUpper := false
		hasLower := false
		hasDigit := false
		hasSpecial := false
		
		for _, char := range pwd {
			if unicode.IsUpper(char) {
				hasUpper = true
			}
			if unicode.IsLower(char) {
				hasLower = true
			}
			if unicode.IsDigit(char) {
				hasDigit = true
			}
			if !unicode.IsLetter(char) && !unicode.IsDigit(char) {
				hasSpecial = true
			}
		}
		
		if !hasUpper {
			fmt.Println("❌ No uppercase letter")
			continue
		}
		
		if !hasLower {
			fmt.Println("❌ No lowercase letter")
			continue
		}
		
		if !hasDigit {
			fmt.Println("❌ No number")
			continue
		}
		
		if !hasSpecial {
			fmt.Println("❌ No special character")
			continue
		}
		
		fmt.Println("✓ Strong password!")
	}

	// Example 4: Menu item validation
	fmt.Println("\n\n☕ Menu Item Validation")
	fmt.Println("----------------------")
	
	menuItems := []struct {
		name        string
		category    string
		price       float64
		available   bool
		ingredients []string
	}{
		{"Espresso", "Coffee", 2.50, true, []string{"coffee"}},
		{"Mystery Drink", "Unknown", 5.00, true, []string{"???"}},      // Invalid category
		{"Free Water", "Drinks", 0.00, true, []string{"water"}},        // Invalid price
		{"Unicorn Latte", "Coffee", 15.00, false, []string{"magic"}},   // Not available
		{"Broken Item", "Pastry", 3.50, true, []string{}},              // No ingredients
		{"Cappuccino", "Coffee", 4.00, true, []string{"coffee", "milk"}},
		{"Super Expensive", "Coffee", 100.00, true, []string{"gold"}},  // Too expensive
	}
	
	validCategories := map[string]bool{
		"Coffee": true,
		"Tea":    true,
		"Pastry": true,
		"Drinks": true,
	}
	
	fmt.Println("\nValidating menu items:")
	activeMenuItems := 0
	
	for _, item := range menuItems {
		fmt.Printf("\n%-15s ($%.2f): ", item.name, item.price)
		
		// Validate category
		if !validCategories[item.category] {
			fmt.Printf("❌ Invalid category: %s", item.category)
			continue
		}
		
		// Validate price
		if item.price <= 0 {
			fmt.Print("❌ Invalid price (must be > 0)")
			continue
		}
		
		if item.price > 20 {
			fmt.Printf("❌ Price too high ($%.2f > $20)", item.price)
			continue
		}
		
		// Check availability
		if !item.available {
			fmt.Print("❌ Not available")
			continue
		}
		
		// Validate ingredients
		if len(item.ingredients) == 0 {
			fmt.Print("❌ No ingredients listed")
			continue
		}
		
		// Valid item
		activeMenuItems++
		fmt.Printf("✓ Valid - %s with %v", item.category, item.ingredients)
	}
	
	fmt.Printf("\n\nActive menu items: %d\n", activeMenuItems)

	// Example 5: Transaction validation
	fmt.Println("\n\n💳 Transaction Validation")
	fmt.Println("------------------------")
	
	transactions := []struct {
		id     string
		type_  string
		amount float64
		status string
	}{
		{"TXN001", "payment", 15.50, "completed"},
		{"TXN002", "refund", -8.00, "pending"},     // Refunds should be positive
		{"TXN003", "payment", 0.00, "completed"},   // Zero amount
		{"TXN004", "unknown", 10.00, "completed"},  // Invalid type
		{"TXN005", "payment", 12.75, "failed"},     // Failed transaction
		{"TXN006", "refund", 5.00, "completed"},
		{"TXN007", "payment", 1000.00, "pending"},  // Too large
	}
	
	validTypes := map[string]bool{"payment": true, "refund": true}
	processedAmount := 0.0
	
	for _, txn := range transactions {
		fmt.Printf("%s (%-7s): $%7.2f - ", txn.id, txn.type_, txn.amount)
		
		// Validate transaction type
		if !validTypes[txn.type_] {
			fmt.Printf("❌ Invalid type: %s\n", txn.type_)
			continue
		}
		
		// Validate amount
		if txn.amount <= 0 {
			fmt.Println("❌ Invalid amount")
			continue
		}
		
		if txn.amount > 100 {
			fmt.Println("❌ Amount exceeds limit")
			continue
		}
		
		// Check status
		if txn.status != "completed" {
			fmt.Printf("❌ Not completed (%s)\n", txn.status)
			continue
		}
		
		// Process transaction
		if txn.type_ == "payment" {
			processedAmount += txn.amount
		} else {
			processedAmount -= txn.amount
		}
		
		fmt.Printf("✓ Processed (%s)\n", txn.type_)
	}
	
	fmt.Printf("\nNet processed amount: $%.2f\n", processedAmount)
}