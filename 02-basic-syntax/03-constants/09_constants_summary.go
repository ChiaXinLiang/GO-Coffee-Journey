package main

import (
	"fmt"
	"strings"
	"time"
)

// GoCoffee configuration constants
const (
	// Company
	CompanyName   = "GoCoffee"
	CompanySlogan = "Code & Coffee"

	// Locations
	NumLocations = 5
	Headquarters = "Seattle"

	// Operating hours
	WeekdayOpen  = 6  // 6 AM
	WeekdayClose = 22 // 10 PM
	WeekendOpen  = 7  // 7 AM
	WeekendClose = 23 // 11 PM

	// Pricing
	BaseEspressoPrice = 3.00
	TaxRate           = 0.085

	// Limits
	MaxItemsPerOrder = 20
	MinOrderAmount   = 3.00

	// Discounts
	MemberDiscount   = 0.10
	StudentDiscount  = 0.15
	EmployeeDiscount = 0.25
)

// Menu categories using iota
const (
	CategoryHotDrinks = iota
	CategoryColdDrinks
	CategoryFood
	CategoryDesserts
)

// Store features as bit flags
const (
	FeatureWifi           = 1 << iota // 1
	FeatureDriveThru                  // 2
	FeatureOutdoorSeating             // 4
	FeatureMobileOrder                // 8
	FeatureLiveMusic                  // 16
)

func main() {
	fmt.Printf("Welcome to %s - %s\n", CompanyName, CompanySlogan)
	fmt.Printf("Serving great coffee at %d locations!\n\n", NumLocations)

	// Check if store is open
	currentHour := time.Now().Hour()
	isWeekend := time.Now().Weekday() == time.Saturday ||
		time.Now().Weekday() == time.Sunday

	var openHour, closeHour int
	if isWeekend {
		openHour, closeHour = WeekendOpen, WeekendClose
	} else {
		openHour, closeHour = WeekdayOpen, WeekdayClose
	}

	if currentHour >= openHour && currentHour < closeHour {
		fmt.Println("🟢 We're OPEN!")
	} else {
		fmt.Println("🔴 We're CLOSED")
	}
	fmt.Printf("Hours today: %d:00 - %d:00\n\n", openHour, closeHour)

	// Calculate a sample order
	fmt.Println("SAMPLE ORDER CALCULATION:")

	// 2 Large Lattes
	quantity := 2
	sizeMultiplier := 1.5              // Large
	latteBase := BaseEspressoPrice + 1.50 // Espresso + milk

	subtotal := float64(quantity) * latteBase * sizeMultiplier
	memberDiscountAmount := subtotal * MemberDiscount
	afterDiscount := subtotal - memberDiscountAmount
	tax := afterDiscount * TaxRate
	total := afterDiscount + tax

	fmt.Printf("2 Large Lattes @ $%.2f\n", latteBase*sizeMultiplier)
	fmt.Printf("Subtotal:        $%.2f\n", subtotal)
	fmt.Printf("Member Discount: -$%.2f\n", memberDiscountAmount)
	fmt.Printf("After Discount:  $%.2f\n", afterDiscount)
	fmt.Printf("Tax (%.1f%%):     $%.2f\n", TaxRate*100, tax)
	fmt.Printf("Total:           $%.2f\n", total)

	// Store features
	fmt.Println("\nDOWNTOWN STORE FEATURES:")
	downtownFeatures := FeatureWifi | FeatureMobileOrder | FeatureOutdoorSeating

	checkFeature := func(name string, flag int) {
		if downtownFeatures&flag != 0 {
			fmt.Printf("✓ %s\n", name)
		} else {
			fmt.Printf("✗ %s\n", name)
		}
	}

	checkFeature("WiFi Available", FeatureWifi)
	checkFeature("Drive-Thru", FeatureDriveThru)
	checkFeature("Outdoor Seating", FeatureOutdoorSeating)
	checkFeature("Mobile Ordering", FeatureMobileOrder)
	checkFeature("Live Music Events", FeatureLiveMusic)

	fmt.Println("\n" + strings.Repeat("=", 40))
	fmt.Println("Constants make our system reliable!")
	fmt.Println("Prices, hours, and rules never change accidentally.")
}