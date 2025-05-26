package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

func main() {
	fmt.Println("=== GoCoffee: Legitimate Goto Use Cases ===\n")
	
	fmt.Println("While goto should be avoided in most cases,")
	fmt.Println("here are some scenarios where it MIGHT be justified:\n")
	
	rand.Seed(time.Now().UnixNano())
	
	// Example 1: Resource cleanup in system programming
	fmt.Println("Example 1: Resource Cleanup Pattern")
	fmt.Println("----------------------------------")
	complexResourceManagement()
	
	fmt.Println("\n" + strings.Repeat("=", 50) + "\n")
	
	// Example 2: Parser/Lexer implementation
	fmt.Println("Example 2: Simple Parser State Machine")
	fmt.Println("-------------------------------------")
	parseOrderFormat("COFFEE:Latte,SIZE:Large,EXTRAS:ExtraShot")
	
	fmt.Println("\n" + strings.Repeat("=", 50) + "\n")
	
	// Example 3: Performance-critical inner loops
	fmt.Println("Example 3: Performance-Critical Code")
	fmt.Println("-----------------------------------")
	demonstratePerformanceCritical()
	
	fmt.Println("\n" + strings.Repeat("=", 50) + "\n")
	
	// Example 4: Generated code
	fmt.Println("Example 4: Generated Code Pattern")
	fmt.Println("--------------------------------")
	demonstrateGeneratedCode()
}

// Example 1: Complex resource management
// In real code, prefer defer, but this shows a legitimate pattern
func complexResourceManagement() {
	fmt.Println("Simulating complex resource acquisition...\n")
	
	var (
		file1, file2, file3 *os.File
		conn1, conn2        io.Closer
		err                 error
	)
	
	// Simulate opening resources
	fmt.Print("1. Opening order log... ")
	file1, err = simulateFileOpen("orders.log")
	if err != nil {
		fmt.Printf("❌ %v\n", err)
		goto Cleanup
	}
	fmt.Println("✓")
	
	fmt.Print("2. Opening inventory database... ")
	file2, err = simulateFileOpen("inventory.db")
	if err != nil {
		fmt.Printf("❌ %v\n", err)
		goto Cleanup
	}
	fmt.Println("✓")
	
	fmt.Print("3. Connecting to payment system... ")
	conn1, err = simulateConnection("payment")
	if err != nil {
		fmt.Printf("❌ %v\n", err)
		goto Cleanup
	}
	fmt.Println("✓")
	
	fmt.Print("4. Opening transaction log... ")
	file3, err = simulateFileOpen("transactions.log")
	if err != nil {
		fmt.Printf("❌ %v\n", err)
		goto Cleanup
	}
	fmt.Println("✓")
	
	fmt.Print("5. Connecting to notification service... ")
	conn2, err = simulateConnection("notifications")
	if err != nil {
		fmt.Printf("❌ %v\n", err)
		goto Cleanup
	}
	fmt.Println("✓")
	
	// Do work
	fmt.Println("\n✅ All resources acquired successfully!")
	fmt.Println("Processing orders...")
	time.Sleep(500 * time.Millisecond)
	
	// Success path also goes to cleanup
	err = nil
	
Cleanup:
	// Clean up in reverse order
	fmt.Println("\nCleaning up resources:")
	
	if conn2 != nil {
		conn2.Close()
		fmt.Println("  → Closed notification connection")
	}
	if file3 != nil {
		file3.Close()
		fmt.Println("  → Closed transaction log")
	}
	if conn1 != nil {
		conn1.Close()
		fmt.Println("  → Closed payment connection")
	}
	if file2 != nil {
		file2.Close()
		fmt.Println("  → Closed inventory database")
	}
	if file1 != nil {
		file1.Close()
		fmt.Println("  → Closed order log")
	}
	
	if err != nil {
		fmt.Printf("\n❌ Operation failed: %v\n", err)
	} else {
		fmt.Println("\n✅ Operation completed successfully")
	}
	
	fmt.Println("\nNOTE: In real Go code, use defer for cleanup!")
}

// Example 2: Simple parser using goto for state transitions
func parseOrderFormat(input string) {
	fmt.Printf("Parsing order string: %s\n\n", input)
	
	var (
		pos     int
		key     string
		value   string
		order   = make(map[string]string)
	)
	
	if len(input) == 0 {
		goto Error
	}
	
ReadKey:
	key = ""
	for pos < len(input) && input[pos] != ':' {
		if input[pos] == ',' {
			goto Error
		}
		key += string(input[pos])
		pos++
	}
	
	if pos >= len(input) {
		goto Error
	}
	
	pos++ // Skip ':'
	goto ReadValue
	
ReadValue:
	value = ""
	for pos < len(input) && input[pos] != ',' {
		value += string(input[pos])
		pos++
	}
	
	order[key] = value
	fmt.Printf("  Parsed: %s = %s\n", key, value)
	
	if pos < len(input) {
		pos++ // Skip ','
		goto ReadKey
	}
	
	goto Success
	
Error:
	fmt.Println("❌ Parse error!")
	fmt.Printf("  Position: %d\n", pos)
	if pos < len(input) {
		fmt.Printf("  Near: '%s'\n", input[pos:])
	}
	return
	
Success:
	fmt.Println("\n✅ Parse successful!")
	fmt.Println("Order details:")
	for k, v := range order {
		fmt.Printf("  %s: %s\n", k, v)
	}
}

// Example 3: Performance-critical code
func demonstratePerformanceCritical() {
	// Simulate a performance-critical operation
	data := make([]int, 1000000)
	for i := range data {
		data[i] = rand.Intn(100)
	}
	
	fmt.Println("Processing large dataset...")
	fmt.Println("(In very rare cases, goto might improve performance)\n")
	
	// Version with goto (might be slightly faster in some cases)
	start := time.Now()
	sum1 := 0
	i := 0
	
ProcessLoop:
	if i >= len(data) {
		goto Done
	}
	
	if data[i] < 0 {
		goto Skip
	}
	
	if data[i] > 50 {
		sum1 += data[i] * 2
	} else {
		sum1 += data[i]
	}
	
Skip:
	i++
	goto ProcessLoop
	
Done:
	time1 := time.Since(start)
	
	// Normal version
	start = time.Now()
	sum2 := 0
	for _, val := range data {
		if val < 0 {
			continue
		}
		if val > 50 {
			sum2 += val * 2
		} else {
			sum2 += val
		}
	}
	time2 := time.Since(start)
	
	fmt.Printf("Goto version:   %v (sum: %d)\n", time1, sum1)
	fmt.Printf("Normal version: %v (sum: %d)\n", time2, sum2)
	fmt.Println("\nNOTE: Modern compilers optimize both similarly!")
}

// Example 4: Generated code pattern
func demonstrateGeneratedCode() {
	fmt.Println("Generated code (like from parser generators)")
	fmt.Println("might use goto for state machines:\n")
	
	// Simulate generated code for a simple state machine
	input := "BREW_COFFEE"
	state := 0
	output := ""
	
	fmt.Printf("Input: %s\n", input)
	fmt.Println("Processing states:")
	
	// This simulates what a code generator might produce
State0:
	fmt.Printf("  State 0: Initial\n")
	if len(input) > 0 && input[0] == 'B' {
		state = 1
		input = input[1:]
		goto State1
	}
	goto ErrorState
	
State1:
	fmt.Printf("  State 1: Read 'B'\n")
	if len(input) > 0 && input[0] == 'R' {
		state = 2
		input = input[1:]
		goto State2
	}
	goto ErrorState
	
State2:
	fmt.Printf("  State 2: Read 'BR'\n")
	if len(input) > 3 && input[0:3] == "EW_" {
		state = 3
		input = input[3:]
		output = "BREW_"
		goto State3
	}
	goto ErrorState
	
State3:
	fmt.Printf("  State 3: Read 'BREW_'\n")
	if len(input) >= 6 && input[0:6] == "COFFEE" {
		output += "COFFEE"
		goto AcceptState
	}
	goto ErrorState
	
AcceptState:
	fmt.Printf("  Accept State: Recognized '%s'\n", output)
	fmt.Println("\n✅ Valid command recognized!")
	fmt.Println("→ Starting coffee brew cycle...")
	return
	
ErrorState:
	fmt.Printf("  Error State: Unrecognized input\n")
	fmt.Println("\n❌ Invalid command!")
	return
}

// Helper functions

func simulateFileOpen(name string) (*os.File, error) {
	if rand.Float32() < 0.2 {
		return nil, fmt.Errorf("failed to open %s", name)
	}
	return &os.File{}, nil
}

func simulateConnection(service string) (io.Closer, error) {
	if rand.Float32() < 0.3 {
		return nil, fmt.Errorf("failed to connect to %s", service)
	}
	return &dummyCloser{}, nil
}

type dummyCloser struct{}

func (d *dummyCloser) Close() error {
	return nil
}

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