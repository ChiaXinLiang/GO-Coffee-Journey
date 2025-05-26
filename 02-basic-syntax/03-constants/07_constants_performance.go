package main

import (
	"fmt"
	"time"
)

// Constants are replaced at compile time
const (
	ConstantPrice    = 4.50
	ConstantTaxRate  = 0.085
	ConstantDiscount = 0.10
)

// Variables exist in memory
var (
	variablePrice    = 4.50
	variableTaxRate  = 0.085
	variableDiscount = 0.10
)

func main() {
	fmt.Println("=== Constants vs Variables Performance ===\n")

	iterations := 100_000_000

	// Test with constants
	start := time.Now()
	var totalConst float64
	for i := 0; i < iterations; i++ {
		totalConst = ConstantPrice * (1 + ConstantTaxRate) * (1 - ConstantDiscount)
	}
	constTime := time.Since(start)

	// Test with variables
	start = time.Now()
	var totalVar float64
	for i := 0; i < iterations; i++ {
		totalVar = variablePrice * (1 + variableTaxRate) * (1 - variableDiscount)
	}
	varTime := time.Since(start)

	fmt.Printf("Iterations: %d\n\n", iterations)
	fmt.Printf("Constants time: %v\n", constTime)
	fmt.Printf("Variables time: %v\n", varTime)
	fmt.Printf("Difference: %v\n\n", varTime-constTime)

	// Show that results are the same
	fmt.Printf("Constant result: %.2f\n", totalConst)
	fmt.Printf("Variable result: %.2f\n", totalVar)

	// Constants in conditionals
	const debugMode = false

	// This code might be eliminated by compiler
	if debugMode {
		fmt.Println("Debug: This won't be included in binary")
	}

	// Memory comparison
	fmt.Println("\nMEMORY USAGE:")
	fmt.Println("Constants: No runtime memory (compile-time replacement)")
	fmt.Println("Variables: Stored in memory at runtime")

	// Best practices
	fmt.Println("\nBEST PRACTICES:")
	fmt.Println("✓ Use constants for values that never change")
	fmt.Println("✓ Constants improve performance")
	fmt.Println("✓ Constants prevent accidental modifications")
	fmt.Println("✓ Constants make code more maintainable")
}