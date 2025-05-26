package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("=== GoCoffee: Goto Performance Analysis ===\n")
	
	fmt.Println("Testing if goto provides any performance benefits")
	fmt.Println("(Spoiler: Modern compilers optimize well!)\n")
	
	rand.Seed(time.Now().UnixNano())
	
	// Generate test data
	size := 10000000
	data := make([]int, size)
	for i := range data {
		data[i] = rand.Intn(1000)
	}
	
	// Test 1: Loop performance
	fmt.Println("Test 1: Loop Performance")
	fmt.Println("------------------------")
	testLoopPerformance(data)
	
	fmt.Println("\n" + strings.Repeat("=", 50) + "\n")
	
	// Test 2: Error handling performance
	fmt.Println("Test 2: Error Handling Performance")
	fmt.Println("---------------------------------")
	testErrorHandling()
	
	fmt.Println("\n" + strings.Repeat("=", 50) + "\n")
	
	// Test 3: State machine performance
	fmt.Println("Test 3: State Machine Performance")
	fmt.Println("--------------------------------")
	testStateMachine()
	
	fmt.Println("\n" + strings.Repeat("=", 50) + "\n")
	
	// Test 4: Branch prediction impact
	fmt.Println("Test 4: Branch Prediction Impact")
	fmt.Println("-------------------------------")
	testBranchPrediction()
}

// Test 1: Compare loop implementations
func testLoopPerformance(data []int) {
	// Goto-based loop
	start := time.Now()
	sum1 := 0
	i := 0
	
LoopStart:
	if i >= len(data) {
		goto LoopEnd
	}
	sum1 += data[i]
	i++
	goto LoopStart
	
LoopEnd:
	gotoTime := time.Since(start)
	
	// Standard for loop
	start = time.Now()
	sum2 := 0
	for i := 0; i < len(data); i++ {
		sum2 += data[i]
	}
	forTime := time.Since(start)
	
	// Range loop
	start = time.Now()
	sum3 := 0
	for _, v := range data {
		sum3 += v
	}
	rangeTime := time.Since(start)
	
	fmt.Printf("Results (all sums should match: %d)\n", sum1)
	fmt.Printf("  Goto loop:  %12v (sum: %d)\n", gotoTime, sum1)
	fmt.Printf("  For loop:   %12v (sum: %d)\n", forTime, sum2)
	fmt.Printf("  Range loop: %12v (sum: %d)\n", rangeTime, sum3)
	
	// Analysis
	fmt.Println("\nAnalysis:")
	if gotoTime < forTime {
		improvement := float64(forTime-gotoTime) / float64(forTime) * 100
		fmt.Printf("  Goto was %.2f%% faster than for loop\n", improvement)
	} else {
		overhead := float64(gotoTime-forTime) / float64(forTime) * 100
		fmt.Printf("  Goto was %.2f%% slower than for loop\n", overhead)
	}
}

// Test 2: Error handling patterns
func testErrorHandling() {
	iterations := 1000000
	
	// Goto-based error handling
	start := time.Now()
	successCount1 := 0
	
	for i := 0; i < iterations; i++ {
		result := processWithGoto(i)
		if result {
			successCount1++
		}
	}
	gotoTime := time.Since(start)
	
	// Traditional error handling
	start = time.Now()
	successCount2 := 0
	
	for i := 0; i < iterations; i++ {
		result := processWithoutGoto(i)
		if result {
			successCount2++
		}
	}
	traditionalTime := time.Since(start)
	
	fmt.Printf("Results over %d iterations:\n", iterations)
	fmt.Printf("  Goto pattern:        %12v (success: %d)\n", gotoTime, successCount1)
	fmt.Printf("  Traditional pattern: %12v (success: %d)\n", traditionalTime, successCount2)
	
	fmt.Println("\nAnalysis:")
	diff := float64(gotoTime-traditionalTime) / float64(traditionalTime) * 100
	if diff < 0 {
		fmt.Printf("  Goto was %.2f%% faster\n", -diff)
	} else {
		fmt.Printf("  Goto was %.2f%% slower\n", diff)
	}
}

func processWithGoto(n int) bool {
	var step1, step2, step3 bool
	
	// Step 1
	if n%10 == 0 {
		goto Failure
	}
	step1 = true
	
	// Step 2
	if n%100 == 0 {
		goto Failure
	}
	step2 = true
	
	// Step 3
	if n%1000 == 0 {
		goto Failure
	}
	step3 = true
	
	return true
	
Failure:
	_ = step1
	_ = step2
	_ = step3
	return false
}

func processWithoutGoto(n int) bool {
	// Step 1
	if n%10 == 0 {
		return false
	}
	
	// Step 2
	if n%100 == 0 {
		return false
	}
	
	// Step 3
	if n%1000 == 0 {
		return false
	}
	
	return true
}

// Test 3: State machine implementations
func testStateMachine() {
	iterations := 100000
	
	// Goto-based state machine
	start := time.Now()
	result1 := 0
	
	for i := 0; i < iterations; i++ {
		result1 += runStateMachineGoto(i % 4)
	}
	gotoTime := time.Since(start)
	
	// Switch-based state machine
	start = time.Now()
	result2 := 0
	
	for i := 0; i < iterations; i++ {
		result2 += runStateMachineSwitch(i % 4)
	}
	switchTime := time.Since(start)
	
	fmt.Printf("State machine results (%d iterations):\n", iterations)
	fmt.Printf("  Goto-based:   %12v (result: %d)\n", gotoTime, result1)
	fmt.Printf("  Switch-based: %12v (result: %d)\n", switchTime, result2)
	
	fmt.Println("\nAnalysis:")
	diff := float64(gotoTime-switchTime) / float64(switchTime) * 100
	if diff < 0 {
		fmt.Printf("  Goto was %.2f%% faster\n", -diff)
	} else {
		fmt.Printf("  Goto was %.2f%% slower\n", diff)
	}
	fmt.Println("  → Modern compilers optimize both similarly!")
}

func runStateMachineGoto(input int) int {
	value := 0
	
	switch input {
	case 0:
		goto State0
	case 1:
		goto State1
	case 2:
		goto State2
	default:
		goto State3
	}
	
State0:
	value += 10
	goto State1
	
State1:
	value += 20
	goto State2
	
State2:
	value += 30
	goto State3
	
State3:
	value += 40
	goto End
	
End:
	return value
}

func runStateMachineSwitch(input int) int {
	state := input
	value := 0
	
	for {
		switch state {
		case 0:
			value += 10
			state = 1
		case 1:
			value += 20
			state = 2
		case 2:
			value += 30
			state = 3
		case 3:
			value += 40
			return value
		default:
			value += 40
			return value
		}
	}
}

// Test 4: Branch prediction
func testBranchPrediction() {
	size := 1000000
	
	// Predictable pattern
	predictableData := make([]int, size)
	for i := range predictableData {
		if i%2 == 0 {
			predictableData[i] = 1
		} else {
			predictableData[i] = 0
		}
	}
	
	// Random pattern
	randomData := make([]int, size)
	for i := range randomData {
		randomData[i] = rand.Intn(2)
	}
	
	fmt.Println("Testing branch prediction impact...\n")
	
	// Test with predictable data
	fmt.Println("Predictable data pattern (0,1,0,1,...):")
	testBranchingPattern(predictableData)
	
	fmt.Println("\nRandom data pattern:")
	testBranchingPattern(randomData)
	
	fmt.Println("\nConclusion:")
	fmt.Println("→ Goto doesn't magically improve branch prediction")
	fmt.Println("→ CPU branch predictors work the same for both patterns")
}

func testBranchingPattern(data []int) {
	// Goto version
	start := time.Now()
	sum1 := 0
	for i := 0; i < len(data); i++ {
		if data[i] == 0 {
			goto SkipAdd
		}
		sum1 += data[i]
	SkipAdd:
	}
	gotoTime := time.Since(start)
	
	// Continue version
	start = time.Now()
	sum2 := 0
	for i := 0; i < len(data); i++ {
		if data[i] == 0 {
			continue
		}
		sum2 += data[i]
	}
	continueTime := time.Since(start)
	
	fmt.Printf("  Goto:     %12v (sum: %d)\n", gotoTime, sum1)
	fmt.Printf("  Continue: %12v (sum: %d)\n", continueTime, sum2)
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