package main

import (
	"fmt"
	"time"
)

// Business rules as constants
const (
	// Minimum values
	MinimumOrderAmount    = 3.00
	MinimumDeliveryAmount = 15.00
	MinimumTipPercentage  = 0

	// Maximum values
	MaximumOrderItems     = 25
	MaximumDeliveryRadius = 5.0 // miles
	MaximumTipPercentage  = 100

	// Business percentages
	TaxRate          = 0.085 // 8.5%
	MemberDiscount   = 0.10  // 10%
	BulkDiscount     = 0.15  // 15% for orders over $50
	EmployeeDiscount = 0.25  // 25%

	// Time limits
	OrderTimeout    = 30 * time.Minute
	DeliveryTimeout = 45 * time.Minute
	PaymentTimeout  = 5 * time.Minute

	// Loyalty points
	PointsPerDollar    = 10
	PointsForFreeDrink = 1000
)

// Store locations
const (
	HeadquartersZip = "98101"
	HeadquartersLat = 47.6062
	HeadquartersLng = -122.3321
)

// API endpoints
const (
	APIVersion = "v1"
	APIBaseURL = "https://api.gocoffee.com/"
	APITimeout = 30 * time.Second
)

// Error messages
const (
	ErrInsufficientFunds = "insufficient funds"
	ErrItemOutOfStock    = "item out of stock"
	ErrStoreClosed       = "store is closed"
	ErrInvalidCoupon     = "invalid coupon code"
)

func main() {
	fmt.Println("=== GoCoffee Business Constants ===\n")

	// Order validation
	orderAmount := 45.00
	orderItems := 3

	fmt.Println("ORDER VALIDATION:")
	if orderAmount < MinimumOrderAmount {
		fmt.Printf("❌ Order too small. Minimum: $%.2f\n", MinimumOrderAmount)
	} else {
		fmt.Printf("✓ Order amount OK: $%.2f\n", orderAmount)
	}

	if orderItems > MaximumOrderItems {
		fmt.Printf("❌ Too many items. Maximum: %d\n", MaximumOrderItems)
	} else {
		fmt.Printf("✓ Item count OK: %d\n", orderItems)
	}

	// Calculate discounts
	fmt.Println("\nDISCOUNT CALCULATION:")

	subtotal := orderAmount
	tax := subtotal * TaxRate

	// Check for bulk discount
	var discount float64
	var discountType string

	if subtotal >= 50 {
		discount = subtotal * BulkDiscount
		discountType = "Bulk"
	}

	total := subtotal - discount + tax

	fmt.Printf("Subtotal:     $%.2f\n", subtotal)
	if discount > 0 {
		fmt.Printf("%s Discount: -$%.2f\n", discountType, discount)
	}
	fmt.Printf("Tax (%.1f%%):   $%.2f\n", TaxRate*100, tax)
	fmt.Printf("Total:        $%.2f\n", total)

	// Loyalty points
	pointsEarned := int(total) * PointsPerDollar
	fmt.Printf("\nPoints earned: %d\n", pointsEarned)
	fmt.Printf("Points needed for free drink: %d\n", PointsForFreeDrink)

	// Delivery check
	fmt.Println("\nDELIVERY CHECK:")
	deliveryDistance := 3.5

	if total < MinimumDeliveryAmount {
		fmt.Printf("❌ Minimum for delivery: $%.2f\n", MinimumDeliveryAmount)
	} else if deliveryDistance > MaximumDeliveryRadius {
		fmt.Printf("❌ Too far for delivery. Max: %.1f miles\n", MaximumDeliveryRadius)
	} else {
		fmt.Printf("✓ Delivery available (%.1f miles)\n", deliveryDistance)
	}

	// API usage
	fmt.Printf("\nAPI ENDPOINT: %s%s/orders\n", APIBaseURL, APIVersion)
	fmt.Printf("API Timeout: %v\n", APITimeout)
}