package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("=== GoCoffee: Real-World Goto Scenarios ===\n")
	
	fmt.Println("This example demonstrates scenarios where goto")
	fmt.Println("might appear in real code and better alternatives.\n")
	
	rand.Seed(time.Now().UnixNano())
	
	// Scenario 1: Legacy code migration
	fmt.Println("Scenario 1: Legacy Code Pattern")
	fmt.Println("-------------------------------")
	fmt.Println("(Sometimes found in ported C code)\n")
	
	legacyPaymentProcessing(45.50)
	
	fmt.Println("\n" + strings.Repeat("=", 50) + "\n")
	
	// Scenario 2: Generated parser code
	fmt.Println("Scenario 2: Parser/Scanner Code")
	fmt.Println("-------------------------------")
	fmt.Println("(Common in generated lexers/parsers)\n")
	
	tokens := scanCoffeeOrder("LARGE_LATTE_EXTRA_SHOT")
	fmt.Printf("Tokens found: %v\n", tokens)
	
	fmt.Println("\n" + strings.Repeat("=", 50) + "\n")
	
	// Scenario 3: Low-level system code
	fmt.Println("Scenario 3: System-Level Code")
	fmt.Println("----------------------------")
	fmt.Println("(Initialization with multiple failure points)\n")
	
	initializeCoffeeSystem()
	
	fmt.Println("\n" + strings.Repeat("=", 50) + "\n")
	
	// Scenario 4: Performance-critical hot path
	fmt.Println("Scenario 4: Hot Path Optimization")
	fmt.Println("--------------------------------")
	fmt.Println("(Very rare, usually premature optimization)\n")
	
	processHighVolumeOrders()
	
	fmt.Println("\n" + strings.Repeat("=", 50) + "\n")
	
	// Final recommendations
	showRecommendations()
}

// Scenario 1: Legacy code pattern (often from C ports)
func legacyPaymentProcessing(amount float64) {
	var (
		cardValid     bool
		balanceSufficient bool
		fraudCheck    bool
		processed     bool
		errorCode     int
	)
	
	fmt.Printf("Processing payment of $%.2f\n", amount)
	
	// Simulate legacy C-style code with goto
	cardValid = rand.Float32() > 0.1
	if !cardValid {
		errorCode = 1001
		goto HandleError
	}
	fmt.Println("  ✓ Card validated")
	
	balanceSufficient = rand.Float32() > 0.2
	if !balanceSufficient {
		errorCode = 2001
		goto HandleError
	}
	fmt.Println("  ✓ Balance checked")
	
	fraudCheck = rand.Float32() > 0.05
	if !fraudCheck {
		errorCode = 3001
		goto HandleError
	}
	fmt.Println("  ✓ Fraud check passed")
	
	// Process payment
	processed = true
	fmt.Println("  ✓ Payment processed")
	goto Success
	
HandleError:
	fmt.Printf("  ❌ Error code: %d\n", errorCode)
	switch errorCode {
	case 1001:
		fmt.Println("  → Invalid card")
	case 2001:
		fmt.Println("  → Insufficient balance")
	case 3001:
		fmt.Println("  → Fraud detected")
	}
	goto Cleanup
	
Success:
	fmt.Println("  ✅ Transaction successful!")
	
Cleanup:
	// Common cleanup code
	fmt.Println("  → Closing connections...")
	_ = processed // Use variable
	
	fmt.Println("\n💡 Better approach: Use error types and early returns!")
}

// Scenario 2: Generated parser code pattern
func scanCoffeeOrder(input string) []string {
	var tokens []string
	pos := 0
	
	fmt.Printf("Scanning input: %s\n", input)
	
	// Simplified scanner with goto (like generated code)
ScanStart:
	if pos >= len(input) {
		goto ScanEnd
	}
	
	// Check for size tokens
	if pos+5 <= len(input) && input[pos:pos+5] == "LARGE" {
		tokens = append(tokens, "SIZE:LARGE")
		pos += 5
		goto CheckUnderscore
	}
	
	if pos+5 <= len(input) && input[pos:pos+5] == "SMALL" {
		tokens = append(tokens, "SIZE:SMALL")
		pos += 5
		goto CheckUnderscore
	}
	
	// Check for drink tokens
	if pos+5 <= len(input) && input[pos:pos+5] == "LATTE" {
		tokens = append(tokens, "DRINK:LATTE")
		pos += 5
		goto CheckUnderscore
	}
	
	if pos+8 <= len(input) && input[pos:pos+8] == "ESPRESSO" {
		tokens = append(tokens, "DRINK:ESPRESSO")
		pos += 8
		goto CheckUnderscore
	}
	
	// Check for extras
	if pos+10 <= len(input) && input[pos:pos+10] == "EXTRA_SHOT" {
		tokens = append(tokens, "EXTRA:SHOT")
		pos += 10
		goto CheckUnderscore
	}
	
	// Unknown token
	pos++
	goto ScanStart
	
CheckUnderscore:
	if pos < len(input) && input[pos] == '_' {
		pos++
	}
	goto ScanStart
	
ScanEnd:
	fmt.Printf("  Tokens: %v\n", tokens)
	return tokens
}

// Scenario 3: System initialization with multiple resources
func initializeCoffeeSystem() {
	type Component struct {
		name   string
		weight int
	}
	
	components := []Component{
		{"Config Loader", 10},
		{"Database", 30},
		{"Cache", 20},
		{"Message Queue", 25},
		{"API Server", 40},
	}
	
	initialized := make([]bool, len(components))
	var initError error
	
	fmt.Println("Initializing coffee shop system...")
	
	// Initialize components with weighted failure chance
	for i, comp := range components {
		fmt.Printf("  Initializing %s... ", comp.name)
		
		// Simulate initialization with failure chance
		if rand.Intn(100) < comp.weight {
			initError = fmt.Errorf("failed to initialize %s", comp.name)
			goto RollbackInit
		}
		
		initialized[i] = true
		fmt.Println("✓")
		time.Sleep(200 * time.Millisecond)
	}
	
	fmt.Println("\n✅ System initialized successfully!")
	return
	
RollbackInit:
	fmt.Printf("❌ %v\n", initError)
	fmt.Println("\nRolling back initialization:")
	
	// Cleanup in reverse order
	for i := len(components) - 1; i >= 0; i-- {
		if initialized[i] {
			fmt.Printf("  → Shutting down %s\n", components[i].name)
			time.Sleep(100 * time.Millisecond)
		}
	}
	
	fmt.Println("\n❌ System initialization failed!")
	fmt.Println("\n💡 Better: Use defer or cleanup functions!")
}

// Scenario 4: Performance-critical code (rarely justified)
func processHighVolumeOrders() {
	orders := make([]int, 1000000)
	for i := range orders {
		orders[i] = rand.Intn(5) // 0-4 representing order types
	}
	
	fmt.Println("Processing high-volume orders...")
	fmt.Println("(Comparing goto vs switch performance)\n")
	
	// Version 1: Using goto (might be marginally faster)
	start := time.Now()
	stats1 := processOrdersGoto(orders)
	time1 := time.Since(start)
	
	// Version 2: Using switch (cleaner)
	start = time.Now()
	stats2 := processOrdersSwitch(orders)
	time2 := time.Since(start)
	
	fmt.Printf("Goto version:   %v\n", time1)
	fmt.Printf("  Stats: %v\n", stats1)
	fmt.Printf("\nSwitch version: %v\n", time2)
	fmt.Printf("  Stats: %v\n", stats2)
	
	diff := float64(time1-time2) / float64(time2) * 100
	fmt.Printf("\nPerformance difference: %.2f%%\n", diff)
	fmt.Println("\n💡 The difference is negligible! Use cleaner code!")
}

func processOrdersGoto(orders []int) map[string]int {
	stats := map[string]int{
		"espresso": 0,
		"latte":    0,
		"mocha":    0,
		"americano": 0,
		"other":    0,
	}
	
	for _, order := range orders {
		switch order {
		case 0:
			goto Espresso
		case 1:
			goto Latte
		case 2:
			goto Mocha
		case 3:
			goto Americano
		default:
			goto Other
		}
		
	Espresso:
		stats["espresso"]++
		continue
		
	Latte:
		stats["latte"]++
		continue
		
	Mocha:
		stats["mocha"]++
		continue
		
	Americano:
		stats["americano"]++
		continue
		
	Other:
		stats["other"]++
	}
	
	return stats
}

func processOrdersSwitch(orders []int) map[string]int {
	stats := map[string]int{
		"espresso":  0,
		"latte":     0,
		"mocha":     0,
		"americano": 0,
		"other":     0,
	}
	
	for _, order := range orders {
		switch order {
		case 0:
			stats["espresso"]++
		case 1:
			stats["latte"]++
		case 2:
			stats["mocha"]++
		case 3:
			stats["americano"]++
		default:
			stats["other"]++
		}
	}
	
	return stats
}

// Show final recommendations
func showRecommendations() {
	fmt.Println("📚 GOTO RECOMMENDATIONS")
	fmt.Println("======================\n")
	
	recommendations := []struct {
		category string
		advice   string
	}{
		{
			"General Rule",
			"Avoid goto in 99.9% of cases",
		},
		{
			"Error Handling",
			"Use defer, error returns, or panic/recover",
		},
		{
			"Loop Control",
			"Use break, continue, or labeled statements",
		},
		{
			"State Machines",
			"Use switch statements or state tables",
		},
		{
			"Resource Cleanup",
			"Use defer for guaranteed cleanup",
		},
		{
			"Performance",
			"Profile first! Goto rarely improves performance",
		},
		{
			"Generated Code",
			"Acceptable in machine-generated code only",
		},
		{
			"Legacy Code",
			"Refactor gradually to remove goto",
		},
	}
	
	for _, rec := range recommendations {
		fmt.Printf("%-15s: %s\n", rec.category, rec.advice)
	}
	
	fmt.Println("\n🎯 Remember: Clean, maintainable code > micro-optimizations")
	fmt.Println("   If you think you need goto, think again!")
	fmt.Println("   If you still think you need goto, ask for code review!")
}

// Helper functions
var strings = struct {
	Repeat func(string, int) string
}{
	Repeat: func(s string, n int) string {
		result := ""
		for i := 0; i < n; i++ {
			result += s
		}
		return result
	},
}