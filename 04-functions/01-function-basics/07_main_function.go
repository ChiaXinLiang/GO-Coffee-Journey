package main

import (
	"fmt"
	"os"
	"strings"
)

// The main function is special in Go:
// 1. It's the entry point of the program
// 2. It must be in package main
// 3. It takes no arguments
// 4. It returns no values

// Package initialization happens before main
var shopStatus = "Preparing to open"

// init functions run before main
func init() {
	fmt.Println("🔧 init() function running...")
	fmt.Println("   - This runs before main()")
	fmt.Println("   - Used for package initialization")
	shopStatus = "Initialized"
}

// Another init function (you can have multiple)
func init() {
	fmt.Println("   - Second init() function")
	fmt.Println("   - init functions run in order")
	fmt.Println()
}

// Main function - program entry point
func main() {
	fmt.Println("=== GoCoffee Main Function ===\n")
	
	fmt.Println("🚀 main() function started")
	fmt.Printf("   Shop status: %s\n\n", shopStatus)
	
	// Demonstrate program flow
	demonstrateProgramFlow()
	
	// Show command line arguments
	showCommandLineInfo()
	
	// Demonstrate exit behaviors
	demonstrateExitBehaviors()
	
	// Main function patterns
	showMainPatterns()
	
	// Program will exit when main returns
	fmt.Println("\n✅ main() function completing...")
	fmt.Println("   Program will exit after main returns")
}

func demonstrateProgramFlow() {
	fmt.Println("📋 Program Execution Flow:")
	fmt.Println("1. Package variables initialized")
	fmt.Println("2. init() functions run (in order)")
	fmt.Println("3. main() function starts")
	fmt.Println("4. Program executes")
	fmt.Println("5. main() returns")
	fmt.Println("6. Program exits\n")
}

func showCommandLineInfo() {
	fmt.Println("💻 Command Line Information:")
	
	// os.Args contains command line arguments
	fmt.Printf("Program name: %s\n", os.Args[0])
	fmt.Printf("Number of arguments: %d\n", len(os.Args))
	
	if len(os.Args) > 1 {
		fmt.Println("Arguments provided:")
		for i, arg := range os.Args[1:] {
			fmt.Printf("   Arg %d: %s\n", i+1, arg)
		}
	} else {
		fmt.Println("No command line arguments provided")
		fmt.Println("Try running: go run 07_main_function.go coffee latte")
	}
	fmt.Println()
}

func demonstrateExitBehaviors() {
	fmt.Println("🚪 Exit Behaviors:")
	
	// Different ways a program can exit
	fmt.Println("1. Normal return from main() - exit code 0")
	fmt.Println("2. os.Exit(code) - immediate exit with code")
	fmt.Println("3. panic() - runtime error (we'll avoid this)")
	fmt.Println("4. Fatal error - out of memory, etc.")
	
	// Example of checking before exit
	if checkShopStatus() {
		fmt.Println("✓ Shop systems check passed")
	} else {
		fmt.Println("✗ Shop systems check failed")
		// In real code, might use os.Exit(1) here
	}
	fmt.Println()
}

func checkShopStatus() bool {
	// Simulate system check
	return true
}

func showMainPatterns() {
	fmt.Println("📚 Common Main Function Patterns:\n")
	
	// Pattern 1: Simple CLI
	pattern1()
	
	// Pattern 2: With setup and cleanup
	pattern2()
	
	// Pattern 3: With error handling
	pattern3()
	
	// Pattern 4: Delegating to run function
	pattern4()
}

// Pattern 1: Simple main
func pattern1() {
	fmt.Println("Pattern 1: Simple Main")
	fmt.Println(strings.Repeat("-", 30))
	fmt.Println(`func main() {
    fmt.Println("Hello, World!")
}`)
	fmt.Println()
}

// Pattern 2: Setup and cleanup
func pattern2() {
	fmt.Println("Pattern 2: With Setup/Cleanup")
	fmt.Println(strings.Repeat("-", 30))
	fmt.Println(`func main() {
    setup()
    defer cleanup()
    
    runApplication()
}`)
	fmt.Println()
}

// Pattern 3: Error handling
func pattern3() {
	fmt.Println("Pattern 3: Error Handling")
	fmt.Println(strings.Repeat("-", 30))
	fmt.Println(`func main() {
    if err := run(); err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        os.Exit(1)
    }
}`)
	fmt.Println()
}

// Pattern 4: Delegating pattern
func pattern4() {
	fmt.Println("Pattern 4: Delegating to Run")
	fmt.Println(strings.Repeat("-", 30))
	fmt.Println(`func main() {
    app := NewApplication()
    if err := app.Run(); err != nil {
        log.Fatal(err)
    }
}`)
	fmt.Println()
}

// This would be in a real application
func simulateApplication() {
	fmt.Println("🏃 Application Running...")
	
	// Step 1: Initialize
	fmt.Println("   Initializing...")
	
	// Step 2: Process
	fmt.Println("   Processing...")
	
	// Step 3: Complete
	fmt.Println("   ✓ Complete!")
}

// Demonstrate what NOT to do in main
func showMainAntiPatterns() {
	fmt.Println("\n⚠️  Main Function Anti-Patterns:")
	fmt.Println("• Don't put all logic in main()")
	fmt.Println("• Don't ignore errors")
	fmt.Println("• Don't forget cleanup")
	fmt.Println("• Don't make main() too complex")
	fmt.Println("• Don't use global state excessively")
}

// Note: These functions won't be called because they're after main returns
// This demonstrates that code after main() returns is never executed

func neverCalled() {
	fmt.Println("This will never be printed!")
}

// Summary information (called from main)
func init() {
	// This init provides summary at the very beginning
	fmt.Println("\n💡 Key Points about main():")
	fmt.Println("• Entry point of every Go program")
	fmt.Println("• Must be in package main")
	fmt.Println("• No parameters or return values")
	fmt.Println("• Program exits when main returns")
	fmt.Println("• Exit code 0 means success")
	fmt.Println()
}