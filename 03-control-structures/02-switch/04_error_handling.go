package main

import (
	"fmt"
	"strings"
)

// Error codes for the coffee shop system
const (
	ErrNone = iota
	ErrOutOfStock
	ErrInvalidItem
	ErrPaymentFailed
	ErrMachineFailure
	ErrStaffUnavailable
	ErrSystemDown
)

func main() {
	fmt.Println("=== GoCoffee Error Handling System ===\n")

	// Simulate various error scenarios
	handleError(ErrNone, "Order completed")
	fmt.Println()
	
	handleError(ErrOutOfStock, "Oat milk")
	fmt.Println()
	
	handleError(ErrInvalidItem, "Unicorn Frappuccino")
	fmt.Println()
	
	handleError(ErrPaymentFailed, "Card declined")
	fmt.Println()
	
	handleError(ErrMachineFailure, "Espresso machine")
	fmt.Println()
	
	handleError(ErrStaffUnavailable, "All baristas busy")
	fmt.Println()
	
	handleError(ErrSystemDown, "POS system offline")
	fmt.Println()
	
	handleError(99, "Unknown error") // Unknown error code
}

func handleError(errorCode int, context string) {
	fmt.Printf("Error Code: %d\n", errorCode)
	fmt.Printf("Context: %s\n", context)
	fmt.Println(strings.Repeat("-", 40))
	
	switch errorCode {
	case ErrNone:
		fmt.Println("✅ No error - Operation successful!")
		return
		
	case ErrOutOfStock:
		fmt.Println("❌ OUT OF STOCK ERROR")
		fmt.Printf("Item '%s' is currently unavailable\n", context)
		fmt.Println("\nRecommended actions:")
		fmt.Println("1. Suggest alternative item")
		fmt.Println("2. Check other locations")
		fmt.Println("3. Add to restock list")
		fmt.Println("4. Notify customer of wait time")
		
	case ErrInvalidItem:
		fmt.Println("❌ INVALID ITEM ERROR")
		fmt.Printf("'%s' is not on our menu\n", context)
		fmt.Println("\nRecommended actions:")
		fmt.Println("1. Show menu to customer")
		fmt.Println("2. Suggest similar items")
		fmt.Println("3. Check for typos")
		
	case ErrPaymentFailed:
		fmt.Println("❌ PAYMENT FAILED ERROR")
		fmt.Printf("Payment issue: %s\n", context)
		fmt.Println("\nRecommended actions:")
		fmt.Println("1. Try payment again")
		fmt.Println("2. Try different payment method")
		fmt.Println("3. Check card details")
		fmt.Println("4. Contact bank if needed")
		
	case ErrMachineFailure:
		fmt.Println("❌ EQUIPMENT FAILURE ERROR")
		fmt.Printf("Equipment problem: %s\n", context)
		fmt.Println("\nRecommended actions:")
		fmt.Println("1. Use backup equipment")
		fmt.Println("2. Call maintenance")
		fmt.Println("3. Offer alternative drinks")
		fmt.Println("4. Update customers on wait time")
		
	case ErrStaffUnavailable:
		fmt.Println("❌ STAFF UNAVAILABLE ERROR")
		fmt.Printf("Staffing issue: %s\n", context)
		fmt.Println("\nRecommended actions:")
		fmt.Println("1. Call additional staff")
		fmt.Println("2. Implement queue system")
		fmt.Println("3. Notify customers of delay")
		fmt.Println("4. Offer self-service options")
		
	case ErrSystemDown:
		fmt.Println("❌ SYSTEM DOWN ERROR")
		fmt.Printf("System failure: %s\n", context)
		fmt.Println("\nRecommended actions:")
		fmt.Println("1. Switch to manual mode")
		fmt.Println("2. Use paper receipts")
		fmt.Println("3. Call IT support")
		fmt.Println("4. Keep manual transaction log")
		
	default:
		fmt.Println("❓ UNKNOWN ERROR")
		fmt.Printf("Unrecognized error code: %d\n", errorCode)
		fmt.Println("\nRecommended actions:")
		fmt.Println("1. Log error details")
		fmt.Println("2. Contact supervisor")
		fmt.Println("3. Document incident")
		fmt.Println("4. Escalate to tech support")
	}
	
	// Log error for all cases except success
	if errorCode != ErrNone {
		logError(errorCode, context)
	}
}

func logError(code int, context string) {
	fmt.Printf("\n📝 Logged: Error %d - %s at %s\n", 
		code, context, "2024-01-15 10:30:45")
}