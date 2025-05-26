package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("=== GoCoffee Number Types Demo ===\n")

	// Integer types for counting
	var (
		// Unsigned integers (no negative values)
		coffeeBags    uint8  = 255                 // 0 to 255
		customers     uint16 = 65535                // 0 to 65,535
		totalSold     uint32 = 4294967295           // 0 to 4,294,967,295
		yearlyRevenue uint64 = 18446744073709551615 // Really big number!

		// Signed integers (can be negative)
		temperature  int8  = -10                    // -128 to 127
		stockChange  int16 = -1000                  // -32,768 to 32,767
		profit       int32 = -2147483648            // About -2 billion to 2 billion
		companyValue int64 = 9223372036854775807    // Really big number!

		// Platform-dependent integers
		dailyOrders int  = 1500 // 32 or 64 bits depending on system
		queueSize   uint = 100  // 32 or 64 bits depending on system
	)

	fmt.Println("UNSIGNED INTEGERS (no negatives):")
	fmt.Printf("Coffee bags in storage (uint8): %d\n", coffeeBags)
	fmt.Printf("Total customers today (uint16): %d\n", customers)
	fmt.Printf("Coffees sold all-time (uint32): %d\n", totalSold)
	fmt.Printf("Yearly revenue in cents (uint64): %d\n", yearlyRevenue)

	fmt.Println("\nSIGNED INTEGERS (can be negative):")
	fmt.Printf("Freezer temperature °C (int8): %d\n", temperature)
	fmt.Printf("Stock change (int16): %d\n", stockChange)
	fmt.Printf("Monthly profit/loss (int32): %d\n", profit)
	fmt.Printf("Company valuation (int64): %d\n", companyValue)

	fmt.Println("\nPLATFORM INTEGERS:")
	fmt.Printf("Daily orders (int): %d\n", dailyOrders)
	fmt.Printf("Queue size (uint): %d\n", queueSize)

	// Floating point for prices and measurements
	var (
		coffeePrice  float32 = 4.99 // 32-bit floating point
		precisePrice float64 = 4.99 // 64-bit floating point (more precise)

		// Demonstrating precision difference
		piFloat32 float32 = math.Pi
		piFloat64 float64 = math.Pi
	)

	fmt.Println("\nFLOATING POINT NUMBERS:")
	fmt.Printf("Coffee price (float32): $%.2f\n", coffeePrice)
	fmt.Printf("Precise price (float64): $%.2f\n", precisePrice)
	fmt.Printf("Pi as float32: %.10f\n", piFloat32)
	fmt.Printf("Pi as float64: %.10f\n", piFloat64)

	// Complex numbers (rarely used in business apps)
	var complexOrder complex64 = 3 + 4i
	fmt.Printf("\nComplex order (just for fun): %v\n", complexOrder)

	// Byte and rune
	var (
		asciiChar   byte = 'A'  // alias for uint8
		unicodeChar rune = '☕' // alias for int32 (Unicode code point)
	)

	fmt.Println("\nCHARACTER TYPES:")
	fmt.Printf("ASCII character (byte): %c (%d)\n", asciiChar, asciiChar)
	fmt.Printf("Unicode character (rune): %c (%d)\n", unicodeChar, unicodeChar)
}