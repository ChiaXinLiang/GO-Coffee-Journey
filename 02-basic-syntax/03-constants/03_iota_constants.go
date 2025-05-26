package main

import "fmt"

// Days of the week using iota
const (
	Sunday = iota // 0
	Monday        // 1
	Tuesday       // 2
	Wednesday     // 3
	Thursday      // 4
	Friday        // 5
	Saturday      // 6
)

// Order status with iota
const (
	OrderPending = iota
	OrderConfirmed
	OrderPreparing
	OrderReady
	OrderDelivered
	OrderCancelled
)

// Bit flags for features
const (
	FeatureNone        = 0
	FeatureWifi        = 1 << iota // 1 << 0 = 1
	FeatureParking                 // 1 << 1 = 2
	FeatureDriveThru               // 1 << 2 = 4
	FeatureOutdoor                 // 1 << 3 = 8
	FeaturePetFriendly             // 1 << 4 = 16
)

// Custom iota patterns
const (
	_          = iota // Skip 0
	KB float64 = 1 << (10 * iota) // 1 << 10 = 1024
	MB                             // 1 << 20
	GB                             // 1 << 30
	TB                             // 1 << 40
)

func main() {
	fmt.Println("=== GoCoffee Iota Examples ===\n")

	// Days of week
	fmt.Println("BUSINESS DAYS:")
	businessDays := []int{Monday, Tuesday, Wednesday, Thursday, Friday}
	for _, day := range businessDays {
		fmt.Printf("Day %d is a business day\n", day)
	}

	// Order status
	fmt.Println("\nORDER STATUS VALUES:")
	fmt.Printf("Pending:   %d\n", OrderPending)
	fmt.Printf("Confirmed: %d\n", OrderConfirmed)
	fmt.Printf("Preparing: %d\n", OrderPreparing)
	fmt.Printf("Ready:     %d\n", OrderReady)
	fmt.Printf("Delivered: %d\n", OrderDelivered)
	fmt.Printf("Cancelled: %d\n", OrderCancelled)

	// Store features using bit flags
	fmt.Println("\nSTORE FEATURES:")
	downtownStore := FeatureWifi | FeatureParking | FeatureOutdoor
	airportStore := FeatureWifi | FeatureDriveThru

	fmt.Printf("Downtown store features: %b\n", downtownStore)
	checkFeatures("Downtown", downtownStore)

	fmt.Printf("\nAirport store features: %b\n", airportStore)
	checkFeatures("Airport", airportStore)

	// File sizes
	fmt.Println("\nFILE SIZE CONSTANTS:")
	fmt.Printf("1 KB = %.0f bytes\n", KB)
	fmt.Printf("1 MB = %.0f bytes\n", MB)
	fmt.Printf("1 GB = %.0f bytes\n", GB)

	imageSize := 2.5 * MB
	fmt.Printf("\nMenu image size: %.0f bytes (%.1f MB)\n", imageSize, imageSize/MB)
}

func checkFeatures(storeName string, features int) {
	fmt.Printf("%s store has:\n", storeName)
	if features&FeatureWifi != 0 {
		fmt.Println("  ✓ WiFi")
	}
	if features&FeatureParking != 0 {
		fmt.Println("  ✓ Parking")
	}
	if features&FeatureDriveThru != 0 {
		fmt.Println("  ✓ Drive-Thru")
	}
	if features&FeatureOutdoor != 0 {
		fmt.Println("  ✓ Outdoor Seating")
	}
	if features&FeaturePetFriendly != 0 {
		fmt.Println("  ✓ Pet Friendly")
	}
}