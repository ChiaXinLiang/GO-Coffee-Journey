package main

import (
	"fmt"
	"strings"
	"log"
)

// === GOOD: Clear use cases for variadic functions ===

// 1. When the number of inputs naturally varies
func CombineOrders(orders ...string) string {
	if len(orders) == 0 {
		return "No orders"
	}
	return fmt.Sprintf("Combined order: %s", strings.Join(orders, ", "))
}

// 2. For optional configuration
type Config func(*Server)

type Server struct {
	port    int
	verbose bool
	maxConn int
}

func NewServer(configs ...Config) *Server {
	s := &Server{
		port:    8080, // defaults
		maxConn: 100,
	}
	
	for _, cfg := range configs {
		cfg(s)
	}
	
	return s
}

func WithPort(port int) Config {
	return func(s *Server) {
		s.port = port
	}
}

func WithVerbose() Config {
	return func(s *Server) {
		s.verbose = true
	}
}

// 3. For aggregation operations
func MergeErrors(errs ...error) error {
	var nonNil []error
	for _, err := range errs {
		if err != nil {
			nonNil = append(nonNil, err)
		}
	}
	
	if len(nonNil) == 0 {
		return nil
	}
	
	if len(nonNil) == 1 {
		return nonNil[0]
	}
	
	messages := make([]string, len(nonNil))
	for i, err := range nonNil {
		messages[i] = err.Error()
	}
	
	return fmt.Errorf("multiple errors: %s", strings.Join(messages, "; "))
}

// === BAD: Poor uses of variadic functions ===

// BAD: Fixed number of parameters
// func AddTwo(nums ...int) int {  // Don't do this!
//     if len(nums) != 2 {
//         panic("exactly 2 numbers required")
//     }
//     return nums[0] + nums[1]
// }
// GOOD: Use regular parameters
func AddTwo(a, b int) int {
	return a + b
}

// BAD: Type confusion with interface{}
// func ProcessData(data ...interface{}) {  // Avoid when possible
//     // Complex type checking needed
// }
// GOOD: Use specific types or generics
func ProcessOrders(orders ...Order) {
	for _, order := range orders {
		// Type-safe processing
		fmt.Printf("Processing order %s\n", order.ID)
	}
}

// === Performance Considerations ===

// Variadic functions allocate a new slice
func BenchmarkExample() {
	// This allocates a new slice each call
	expensiveOperation := func(values ...int) int {
		sum := 0
		for _, v := range values {
			sum += v
		}
		return sum
	}
	
	// For hot paths, consider accepting a slice
	cheaperOperation := func(values []int) int {
		sum := 0
		for _, v := range values {
			sum += v
		}
		return sum
	}
	
	// Example usage
	nums := []int{1, 2, 3, 4, 5}
	
	// Allocates new slice
	_ = expensiveOperation(nums...)
	
	// No allocation
	_ = cheaperOperation(nums)
}

// === Documentation Best Practices ===

// AppendToLog appends messages to the log file.
// If no messages are provided, this is a no-op.
// Each message is written on a separate line.
//
// Example:
//   AppendToLog("Server started")
//   AppendToLog("Request received", "Processing", "Complete")
func AppendToLog(messages ...string) {
	if len(messages) == 0 {
		return // Document this behavior!
	}
	
	for _, msg := range messages {
		log.Println(msg)
	}
}

// === Validation Patterns ===

// ValidatedSum returns the sum of positive numbers only.
// Negative numbers are ignored with a warning.
// Returns the sum and count of valid numbers.
func ValidatedSum(numbers ...int) (sum int, validCount int) {
	for _, n := range numbers {
		if n < 0 {
			fmt.Printf("Warning: ignoring negative number %d\n", n)
			continue
		}
		sum += n
		validCount++
	}
	return
}

// === Common Patterns ===

// 1. Chain of responsibility
type Handler func(string) bool

func ProcessWithHandlers(input string, handlers ...Handler) bool {
	for _, h := range handlers {
		if h(input) {
			return true // handled
		}
	}
	return false // not handled
}

// 2. Pipeline pattern
type Transform func(string) string

func Pipeline(input string, transforms ...Transform) string {
	result := input
	for _, t := range transforms {
		result = t(result)
	}
	return result
}

// 3. Event system
type EventHandler func(Event)

type Event struct {
	Type string
	Data interface{}
}

func TriggerEvent(event Event, handlers ...EventHandler) {
	for _, h := range handlers {
		h(event)
	}
}

func main() {
	fmt.Println("=== Variadic Functions Best Practices ===")
	fmt.Println()

	// Good: Natural variable count
	fmt.Println(CombineOrders("Coffee", "Muffin", "Juice"))
	fmt.Println(CombineOrders()) // Works with zero args
	fmt.Println()

	// Good: Optional configuration
	server := NewServer(
		WithPort(3000),
		WithVerbose(),
	)
	fmt.Printf("Server config: port=%d, verbose=%v\n", server.port, server.verbose)
	fmt.Println()

	// Good: Error aggregation
	err := MergeErrors(
		nil,
		fmt.Errorf("connection failed"),
		nil,
		fmt.Errorf("timeout"),
	)
	fmt.Println("Merged error:", err)
	fmt.Println()

	// Validation example
	sum, count := ValidatedSum(10, -5, 20, -3, 15)
	fmt.Printf("Sum of %d valid numbers: %d\n", count, sum)
	fmt.Println()

	// Pipeline example
	result := Pipeline("hello world",
		strings.ToUpper,
		func(s string) string { return s + "!" },
		func(s string) string { return "[" + s + "]" },
	)
	fmt.Println("Pipeline result:", result)
}

// Best Practices Summary:
// 1. Use variadic when count naturally varies
// 2. Provide good defaults for zero arguments
// 3. Document behavior clearly
// 4. Consider performance for hot paths
// 5. Validate inputs when necessary
// 6. Keep type safety when possible
// 7. Use for optional parameters/configuration
// 8. Don't use when parameter count is fixed