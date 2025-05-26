package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	fmt.Println("=== GoCoffee Payment Processing ===\n")

	// Process different payment types
	processPayment("cash", 15.50)
	fmt.Println()
	
	processPayment("credit", 27.35)
	fmt.Println()
	
	processPayment("debit", 8.75)
	fmt.Println()
	
	processPayment("mobile", 12.00)
	fmt.Println()
	
	processPayment("gift_card", 45.25)
	fmt.Println()
	
	processPayment("bitcoin", 100.00) // Unsupported
}

func processPayment(method string, amount float64) {
	fmt.Printf("Processing $%.2f payment\n", amount)
	fmt.Printf("Method: %s\n", method)
	fmt.Println(strings.Repeat("-", 30))
	
	var success bool
	var message string
	
	switch method {
	case "cash":
		// Cash payment processing
		fmt.Println("💵 Cash payment")
		fmt.Println("1. Count bills")
		fmt.Println("2. Open register")
		fmt.Println("3. Make change")
		success = true
		message = "Cash accepted"
		
	case "credit":
		// Credit card processing
		fmt.Println("💳 Credit card payment")
		fmt.Println("1. Swipe/Insert/Tap card")
		fmt.Println("2. Verify with bank")
		fmt.Println("3. Get authorization")
		
		// Simulate processing delay
		time.Sleep(1 * time.Second)
		
		// Simulate 95% success rate
		if amount < 500 {
			success = true
			message = "Credit card approved"
		} else {
			success = false
			message = "Credit limit exceeded"
		}
		
	case "debit":
		// Debit card processing
		fmt.Println("💳 Debit card payment")
		fmt.Println("1. Insert card")
		fmt.Println("2. Enter PIN")
		fmt.Println("3. Check account balance")
		
		// Always succeeds for demo
		success = true
		message = "Debit card approved"
		
	case "mobile":
		// Mobile payment (Apple Pay, Google Pay)
		fmt.Println("📱 Mobile payment")
		fmt.Println("1. Tap phone on reader")
		fmt.Println("2. Authenticate with biometrics")
		fmt.Println("3. Process payment")
		
		success = true
		message = "Mobile payment successful"
		
	case "gift_card":
		// Gift card processing
		fmt.Println("🎁 Gift card payment")
		fmt.Println("1. Scan gift card")
		fmt.Println("2. Check balance")
		
		// Simulate balance check
		balance := 50.00
		if amount <= balance {
			fmt.Printf("3. Deduct $%.2f from card\n", amount)
			fmt.Printf("4. Remaining balance: $%.2f\n", balance-amount)
			success = true
			message = "Gift card accepted"
		} else {
			success = false
			message = fmt.Sprintf("Insufficient balance. Card has $%.2f", balance)
		}
		
	default:
		// Unsupported payment method
		fmt.Println("❌ Unsupported payment method")
		success = false
		message = "Payment method not accepted"
		
		fmt.Println("\nAccepted payment methods:")
		fmt.Println("• Cash")
		fmt.Println("• Credit card")
		fmt.Println("• Debit card")
		fmt.Println("• Mobile payment")
		fmt.Println("• Gift card")
	}
	
	// Print result
	fmt.Println()
	if success {
		fmt.Printf("✅ Success: %s\n", message)
		printReceipt(method, amount)
	} else {
		fmt.Printf("❌ Failed: %s\n", message)
	}
}

func printReceipt(method string, amount float64) {
	fmt.Println("\n--- RECEIPT ---")
	fmt.Printf("Amount: $%.2f\n", amount)
	fmt.Printf("Method: %s\n", strings.Title(strings.Replace(method, "_", " ", -1)))
	fmt.Printf("Time: %s\n", time.Now().Format("15:04:05"))
	fmt.Println("Thank you!")
}