package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("=== GoCoffee Goto Error Handling ===\n")
	
	rand.Seed(time.Now().UnixNano())

	// Example 1: Traditional goto error handling pattern
	fmt.Println("Example 1: Goto-based Error Handling")
	fmt.Println("-----------------------------------")
	
	err := processOrderWithGoto()
	if err != nil {
		fmt.Printf("Order failed: %v\n", err)
	}
	
	fmt.Println("\n" + strings.Repeat("-", 50) + "\n")
	
	// Example 2: Better approach without goto
	fmt.Println("Example 2: Better Error Handling (No Goto)")
	fmt.Println("-----------------------------------------")
	
	err = processOrderWithoutGoto()
	if err != nil {
		fmt.Printf("Order failed: %v\n", err)
	}
	
	fmt.Println("\n" + strings.Repeat("-", 50) + "\n")
	
	// Example 3: Complex resource management
	fmt.Println("Example 3: Resource Management Comparison")
	fmt.Println("----------------------------------------")
	
	fmt.Println("\nWith goto:")
	manageResourcesWithGoto()
	
	fmt.Println("\nWithout goto (using defer):")
	manageResourcesWithDefer()
}

// Example 1: Using goto for error handling (NOT RECOMMENDED)
func processOrderWithGoto() error {
	var grinder, machine, steamer *Resource
	var err error
	
	fmt.Println("Processing order with goto error handling...")
	
	// Step 1: Acquire grinder
	grinder, err = acquireResource("grinder")
	if err != nil {
		goto ErrorHandler
	}
	fmt.Println("✓ Grinder acquired")
	
	// Step 2: Acquire espresso machine
	machine, err = acquireResource("espresso machine")
	if err != nil {
		goto ErrorHandler
	}
	fmt.Println("✓ Espresso machine acquired")
	
	// Step 3: Acquire milk steamer
	steamer, err = acquireResource("milk steamer")
	if err != nil {
		goto ErrorHandler
	}
	fmt.Println("✓ Milk steamer acquired")
	
	// Step 4: Make coffee
	err = makeCoffee(grinder, machine, steamer)
	if err != nil {
		goto ErrorHandler
	}
	
	// Success path - release resources
	releaseResource(steamer)
	releaseResource(machine)
	releaseResource(grinder)
	
	fmt.Println("✅ Order completed successfully!")
	return nil
	
ErrorHandler:
	// Cleanup in reverse order
	fmt.Println("\n❌ Error occurred, cleaning up...")
	if steamer != nil {
		releaseResource(steamer)
	}
	if machine != nil {
		releaseResource(machine)
	}
	if grinder != nil {
		releaseResource(grinder)
	}
	return fmt.Errorf("order processing failed: %v", err)
}

// Example 2: Better approach without goto
func processOrderWithoutGoto() error {
	fmt.Println("Processing order with proper error handling...")
	
	// Step 1: Acquire grinder
	grinder, err := acquireResource("grinder")
	if err != nil {
		return fmt.Errorf("order processing failed: %v", err)
	}
	defer releaseResource(grinder)
	fmt.Println("✓ Grinder acquired")
	
	// Step 2: Acquire espresso machine
	machine, err := acquireResource("espresso machine")
	if err != nil {
		return fmt.Errorf("order processing failed: %v", err)
	}
	defer releaseResource(machine)
	fmt.Println("✓ Espresso machine acquired")
	
	// Step 3: Acquire milk steamer
	steamer, err := acquireResource("milk steamer")
	if err != nil {
		return fmt.Errorf("order processing failed: %v", err)
	}
	defer releaseResource(steamer)
	fmt.Println("✓ Milk steamer acquired")
	
	// Step 4: Make coffee
	if err := makeCoffee(grinder, machine, steamer); err != nil {
		return fmt.Errorf("order processing failed: %v", err)
	}
	
	fmt.Println("✅ Order completed successfully!")
	return nil
}

// Example 3: Complex resource management
func manageResourcesWithGoto() {
	var conn1, conn2, conn3 *Connection
	var err error
	
	fmt.Println("Attempting to establish connections...")
	
	// Try to establish connections
	conn1, err = connect("database")
	if err != nil {
		goto Cleanup
	}
	fmt.Printf("  ✓ Connected to database\n")
	
	conn2, err = connect("payment-gateway")
	if err != nil {
		goto Cleanup
	}
	fmt.Printf("  ✓ Connected to payment gateway\n")
	
	conn3, err = connect("inventory-system")
	if err != nil {
		goto Cleanup
	}
	fmt.Printf("  ✓ Connected to inventory system\n")
	
	// Do work
	fmt.Println("  → Processing transaction...")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("  ✅ Transaction complete!")
	
	// Success cleanup
	disconnect(conn3)
	disconnect(conn2)
	disconnect(conn1)
	return
	
Cleanup:
	fmt.Printf("  ❌ Error: %v\n", err)
	fmt.Println("  Cleaning up connections...")
	if conn3 != nil {
		disconnect(conn3)
	}
	if conn2 != nil {
		disconnect(conn2)
	}
	if conn1 != nil {
		disconnect(conn1)
	}
}

func manageResourcesWithDefer() {
	fmt.Println("Attempting to establish connections...")
	
	conn1, err := connect("database")
	if err != nil {
		fmt.Printf("  ❌ Error: %v\n", err)
		return
	}
	defer disconnect(conn1)
	fmt.Printf("  ✓ Connected to database\n")
	
	conn2, err := connect("payment-gateway")
	if err != nil {
		fmt.Printf("  ❌ Error: %v\n", err)
		return
	}
	defer disconnect(conn2)
	fmt.Printf("  ✓ Connected to payment gateway\n")
	
	conn3, err := connect("inventory-system")
	if err != nil {
		fmt.Printf("  ❌ Error: %v\n", err)
		return
	}
	defer disconnect(conn3)
	fmt.Printf("  ✓ Connected to inventory system\n")
	
	// Do work
	fmt.Println("  → Processing transaction...")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("  ✅ Transaction complete!")
}

// Helper types and functions

type Resource struct {
	name string
}

type Connection struct {
	name string
}

func acquireResource(name string) (*Resource, error) {
	// Simulate 30% failure rate
	if rand.Float32() < 0.3 {
		return nil, fmt.Errorf("failed to acquire %s", name)
	}
	return &Resource{name: name}, nil
}

func releaseResource(r *Resource) {
	if r != nil {
		fmt.Printf("  → Released %s\n", r.name)
	}
}

func makeCoffee(grinder, machine, steamer *Resource) error {
	// Simulate 20% failure rate
	if rand.Float32() < 0.2 {
		return errors.New("coffee making failed")
	}
	fmt.Println("☕ Making coffee...")
	return nil
}

func connect(system string) (*Connection, error) {
	// Simulate 25% failure rate
	if rand.Float32() < 0.25 {
		return nil, fmt.Errorf("failed to connect to %s", system)
	}
	return &Connection{name: system}, nil
}

func disconnect(c *Connection) {
	if c != nil {
		fmt.Printf("  → Disconnected from %s\n", c.name)
	}
}

// strings package substitute
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