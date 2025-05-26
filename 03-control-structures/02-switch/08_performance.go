package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== GoCoffee Switch Performance ===\n")
	
	// Compare if-else chains vs switch statements
	fmt.Println("Comparing if-else chains vs switch statements")
	fmt.Println("for coffee order processing...\n")
	
	// Test data
	coffeeTypes := []string{
		"Espresso", "Latte", "Cappuccino", "Americano", "Mocha",
		"Macchiato", "Flat White", "Cortado", "Gibraltar", "Breve",
	}
	
	// Generate test orders
	numOrders := 1000000
	orders := make([]string, numOrders)
	for i := 0; i < numOrders; i++ {
		orders[i] = coffeeTypes[i%len(coffeeTypes)]
	}
	
	// Test 1: If-else chain
	fmt.Println("Test 1: Processing with if-else chain")
	start := time.Now()
	
	priceTotal1 := 0.0
	for _, order := range orders {
		priceTotal1 += getPriceIfElse(order)
	}
	
	ifElseTime := time.Since(start)
	fmt.Printf("Time: %v\n", ifElseTime)
	fmt.Printf("Total: $%.2f\n\n", priceTotal1)
	
	// Test 2: Switch statement
	fmt.Println("Test 2: Processing with switch")
	start = time.Now()
	
	priceTotal2 := 0.0
	for _, order := range orders {
		priceTotal2 += getPriceSwitch(order)
	}
	
	switchTime := time.Since(start)
	fmt.Printf("Time: %v\n", switchTime)
	fmt.Printf("Total: $%.2f\n\n", priceTotal2)
	
	// Test 3: Map lookup
	fmt.Println("Test 3: Processing with map lookup")
	start = time.Now()
	
	priceTotal3 := 0.0
	for _, order := range orders {
		priceTotal3 += getPriceMap(order)
	}
	
	mapTime := time.Since(start)
	fmt.Printf("Time: %v\n", mapTime)
	fmt.Printf("Total: $%.2f\n\n", priceTotal3)
	
	// Results comparison
	fmt.Println("Performance Comparison")
	fmt.Println("---------------------")
	fmt.Printf("If-else chain: %v (baseline)\n", ifElseTime)
	fmt.Printf("Switch:        %v (%.2fx)\n", switchTime, 
		float64(ifElseTime)/float64(switchTime))
	fmt.Printf("Map lookup:    %v (%.2fx)\n", mapTime,
		float64(ifElseTime)/float64(mapTime))
	
	// Test 4: First match vs last match
	fmt.Println("\n\nTest 4: Position Impact")
	fmt.Println("-----------------------")
	testPositionImpact()
	
	// Test 5: Type switch performance
	fmt.Println("\n\nTest 5: Type Switch Performance")
	fmt.Println("-------------------------------")
	testTypeSwitchPerformance()
}

// If-else implementation
func getPriceIfElse(coffee string) float64 {
	if coffee == "Espresso" {
		return 2.50
	} else if coffee == "Latte" {
		return 4.50
	} else if coffee == "Cappuccino" {
		return 4.00
	} else if coffee == "Americano" {
		return 3.00
	} else if coffee == "Mocha" {
		return 5.00
	} else if coffee == "Macchiato" {
		return 4.25
	} else if coffee == "Flat White" {
		return 4.75
	} else if coffee == "Cortado" {
		return 3.75
	} else if coffee == "Gibraltar" {
		return 3.50
	} else if coffee == "Breve" {
		return 5.50
	}
	return 3.00
}

// Switch implementation
func getPriceSwitch(coffee string) float64 {
	switch coffee {
	case "Espresso":
		return 2.50
	case "Latte":
		return 4.50
	case "Cappuccino":
		return 4.00
	case "Americano":
		return 3.00
	case "Mocha":
		return 5.00
	case "Macchiato":
		return 4.25
	case "Flat White":
		return 4.75
	case "Cortado":
		return 3.75
	case "Gibraltar":
		return 3.50
	case "Breve":
		return 5.50
	default:
		return 3.00
	}
}

// Map lookup implementation
var priceMap = map[string]float64{
	"Espresso":    2.50,
	"Latte":       4.50,
	"Cappuccino":  4.00,
	"Americano":   3.00,
	"Mocha":       5.00,
	"Macchiato":   4.25,
	"Flat White":  4.75,
	"Cortado":     3.75,
	"Gibraltar":   3.50,
	"Breve":       5.50,
}

func getPriceMap(coffee string) float64 {
	if price, ok := priceMap[coffee]; ok {
		return price
	}
	return 3.00
}

// Test position impact
func testPositionImpact() {
	iterations := 1000000
	
	// Test first position
	start := time.Now()
	for i := 0; i < iterations; i++ {
		_ = getPriceSwitch("Espresso") // First case
	}
	firstTime := time.Since(start)
	
	// Test middle position
	start = time.Now()
	for i := 0; i < iterations; i++ {
		_ = getPriceSwitch("Mocha") // Middle case
	}
	middleTime := time.Since(start)
	
	// Test last position
	start = time.Now()
	for i := 0; i < iterations; i++ {
		_ = getPriceSwitch("Breve") // Last case
	}
	lastTime := time.Since(start)
	
	fmt.Printf("First case:  %v\n", firstTime)
	fmt.Printf("Middle case: %v (%.2fx slower)\n", middleTime,
		float64(middleTime)/float64(firstTime))
	fmt.Printf("Last case:   %v (%.2fx slower)\n", lastTime,
		float64(lastTime)/float64(firstTime))
}

// Test type switch performance
func testTypeSwitchPerformance() {
	iterations := 1000000
	
	// Create test data with different types
	var data []interface{}
	for i := 0; i < iterations; i++ {
		switch i % 4 {
		case 0:
			data = append(data, "Latte")
		case 1:
			data = append(data, 4.50)
		case 2:
			data = append(data, true)
		case 3:
			data = append(data, 42)
		}
	}
	
	// Test type switch
	start := time.Now()
	count := 0
	
	for _, item := range data {
		switch v := item.(type) {
		case string:
			if v != "" {
				count++
			}
		case float64:
			if v > 0 {
				count++
			}
		case bool:
			if v {
				count++
			}
		case int:
			if v > 0 {
				count++
			}
		}
	}
	
	elapsed := time.Since(start)
	fmt.Printf("Type switch: %v for %d items\n", elapsed, len(data))
	fmt.Printf("Average: %v per item\n", elapsed/time.Duration(len(data)))
	fmt.Printf("Matched: %d items\n", count)
}