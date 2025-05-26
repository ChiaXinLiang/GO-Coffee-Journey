package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== GoCoffee Loop Performance Comparison ===\n")

	// Test data
	size := 1000000
	coffeeOrders := make([]string, size)
	for i := 0; i < size; i++ {
		coffeeOrders[i] = fmt.Sprintf("Order-%d", i)
	}

	// Example 1: Traditional for loop vs range loop
	fmt.Println("📊 Traditional For Loop vs Range Loop")
	fmt.Println("-------------------------------------")
	
	// Traditional for loop
	start := time.Now()
	count1 := 0
	for i := 0; i < len(coffeeOrders); i++ {
		if len(coffeeOrders[i]) > 0 {
			count1++
		}
	}
	traditionalTime := time.Since(start)
	
	// Range loop
	start = time.Now()
	count2 := 0
	for _, order := range coffeeOrders {
		if len(order) > 0 {
			count2++
		}
	}
	rangeTime := time.Since(start)
	
	fmt.Printf("Traditional loop: %v (processed %d items)\n", traditionalTime, count1)
	fmt.Printf("Range loop:       %v (processed %d items)\n", rangeTime, count2)
	fmt.Printf("Difference:       %v\n", rangeTime-traditionalTime)

	// Example 2: Pre-calculated length vs length in condition
	fmt.Println("\n\n📏 Pre-calculated Length vs Dynamic Length Check")
	fmt.Println("-----------------------------------------------")
	
	testSlice := make([]int, 100000)
	
	// Length in condition
	start = time.Now()
	sum1 := 0
	for i := 0; i < len(testSlice); i++ {
		sum1 += i
	}
	dynamicTime := time.Since(start)
	
	// Pre-calculated length
	start = time.Now()
	sum2 := 0
	length := len(testSlice)
	for i := 0; i < length; i++ {
		sum2 += i
	}
	precalcTime := time.Since(start)
	
	fmt.Printf("Dynamic length check: %v\n", dynamicTime)
	fmt.Printf("Pre-calculated:       %v\n", precalcTime)
	fmt.Printf("Difference:           %v\n", dynamicTime-precalcTime)

	// Example 3: String concatenation in loops
	fmt.Println("\n\n🔤 String Building Performance")
	fmt.Println("------------------------------")
	
	orderCount := 10000
	
	// Bad: String concatenation
	start = time.Now()
	result1 := ""
	for i := 0; i < orderCount; i++ {
		result1 += fmt.Sprintf("Order #%d, ", i)
	}
	concatTime := time.Since(start)
	
	// Good: Using byte slice
	start = time.Now()
	result2 := make([]byte, 0, orderCount*15) // Pre-allocate approximate size
	for i := 0; i < orderCount; i++ {
		result2 = append(result2, fmt.Sprintf("Order #%d, ", i)...)
	}
	sliceTime := time.Since(start)
	
	fmt.Printf("String concatenation: %v\n", concatTime)
	fmt.Printf("Byte slice append:    %v\n", sliceTime)
	fmt.Printf("Improvement:          %.2fx faster\n", float64(concatTime)/float64(sliceTime))

	// Example 4: Break early vs full iteration
	fmt.Println("\n\n🔍 Early Break vs Full Iteration")
	fmt.Println("--------------------------------")
	
	// Full iteration
	start = time.Now()
	found1 := false
	targetOrder := "Order-500000"
	for _, order := range coffeeOrders {
		if order == targetOrder {
			found1 = true
		}
	}
	fullTime := time.Since(start)
	
	// Early break
	start = time.Now()
	found2 := false
	for _, order := range coffeeOrders {
		if order == targetOrder {
			found2 = true
			break
		}
	}
	breakTime := time.Since(start)
	
	fmt.Printf("Full iteration:  %v (found: %v)\n", fullTime, found1)
	fmt.Printf("Early break:     %v (found: %v)\n", breakTime, found2)
	fmt.Printf("Improvement:     %.2fx faster\n", float64(fullTime)/float64(breakTime))

	// Example 5: Nested loops optimization
	fmt.Println("\n\n🔄 Nested Loops Optimization")
	fmt.Println("----------------------------")
	
	rows := 1000
	cols := 1000
	
	// Bad: Column-major order (poor cache locality)
	start = time.Now()
	sum3 := 0
	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, cols)
	}
	
	for j := 0; j < cols; j++ {
		for i := 0; i < rows; i++ {
			sum3 += matrix[i][j]
		}
	}
	colMajorTime := time.Since(start)
	
	// Good: Row-major order (better cache locality)
	start = time.Now()
	sum4 := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			sum4 += matrix[i][j]
		}
	}
	rowMajorTime := time.Since(start)
	
	fmt.Printf("Column-major order: %v\n", colMajorTime)
	fmt.Printf("Row-major order:    %v\n", rowMajorTime)
	fmt.Printf("Improvement:        %.2fx faster\n", float64(colMajorTime)/float64(rowMajorTime))

	// Example 6: Map iteration vs slice iteration
	fmt.Println("\n\n🗺️  Map vs Slice Iteration")
	fmt.Println("--------------------------")
	
	// Create map and slice with same data
	orderMap := make(map[int]string)
	orderSlice := make([]string, 100000)
	
	for i := 0; i < 100000; i++ {
		order := fmt.Sprintf("Coffee-Order-%d", i)
		orderMap[i] = order
		orderSlice[i] = order
	}
	
	// Map iteration
	start = time.Now()
	mapCount := 0
	for _, order := range orderMap {
		if len(order) > 10 {
			mapCount++
		}
	}
	mapTime := time.Since(start)
	
	// Slice iteration
	start = time.Now()
	sliceCount := 0
	for _, order := range orderSlice {
		if len(order) > 10 {
			sliceCount++
		}
	}
	sliceIterTime := time.Since(start)
	
	fmt.Printf("Map iteration:   %v (count: %d)\n", mapTime, mapCount)
	fmt.Printf("Slice iteration: %v (count: %d)\n", sliceIterTime, sliceCount)
	fmt.Printf("Note: Maps are unordered, slices maintain order\n")

	// Summary tips
	fmt.Println("\n\n💡 Performance Tips:")
	fmt.Println("--------------------")
	fmt.Println("1. Use range loops when you need both index and value")
	fmt.Println("2. Pre-calculate lengths for large loops")
	fmt.Println("3. Break early when searching")
	fmt.Println("4. Use byte slices for string building")
	fmt.Println("5. Consider cache locality in nested loops")
	fmt.Println("6. Choose the right data structure for your use case")
}