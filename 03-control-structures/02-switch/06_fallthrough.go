package main

import "fmt"

func main() {
	fmt.Println("=== GoCoffee Fallthrough Examples ===\n")
	
	// Example 1: Default behavior (no fallthrough)
	fmt.Println("Example 1: Default Switch Behavior")
	fmt.Println("---------------------------------")
	orderSize := "medium"
	
	fmt.Printf("Order size: %s\n", orderSize)
	switch orderSize {
	case "small":
		fmt.Println("• Preparing 8oz cup")
	case "medium":
		fmt.Println("• Preparing 12oz cup")
	case "large":
		fmt.Println("• Preparing 16oz cup")
	}
	fmt.Println("Only one case executes (no fallthrough)\n")
	
	// Example 2: Using fallthrough
	fmt.Println("Example 2: Loyalty Program Benefits")
	fmt.Println("----------------------------------")
	customerLevel := "gold"
	
	fmt.Printf("Customer level: %s\n", customerLevel)
	fmt.Println("Benefits:")
	
	switch customerLevel {
	case "platinum":
		fmt.Println("• 20% discount on all items")
		fmt.Println("• Free birthday drink of any size")
		fallthrough
	case "gold":
		fmt.Println("• 15% discount on drinks")
		fmt.Println("• Free size upgrade")
		fallthrough
	case "silver":
		fmt.Println("• 10% discount on drinks")
		fmt.Println("• Buy 9 get 1 free")
		fallthrough
	case "bronze":
		fmt.Println("• 5% discount on drinks")
		fmt.Println("• Birthday month special")
		fallthrough
	default:
		fmt.Println("• Earn points on every purchase")
		fmt.Println("• Member-only promotions")
	}
	
	// Example 3: Cumulative permissions
	fmt.Println("\n\nExample 3: Staff Access Levels")
	fmt.Println("-----------------------------")
	demonstrateAccessLevel("manager")
	fmt.Println()
	demonstrateAccessLevel("barista")
	
	// Example 4: Order preparation steps
	fmt.Println("\n\nExample 4: Rush Order Processing")
	fmt.Println("--------------------------------")
	processRushOrder("large")
	
	// Example 5: Common mistake with fallthrough
	fmt.Println("\n\nExample 5: Fallthrough Gotcha")
	fmt.Println("----------------------------")
	demonstrateFallthroughGotcha(2)
}

func demonstrateAccessLevel(role string) {
	fmt.Printf("Staff role: %s\n", role)
	fmt.Println("Permissions:")
	
	switch role {
	case "owner":
		fmt.Println("• Access financial reports")
		fmt.Println("• Modify pricing")
		fallthrough
	case "manager":
		fmt.Println("• Create staff schedules")
		fmt.Println("• Process refunds")
		fallthrough
	case "supervisor":
		fmt.Println("• Override discounts")
		fmt.Println("• Handle complaints")
		fallthrough
	case "barista":
		fmt.Println("• Take orders")
		fmt.Println("• Prepare drinks")
		fmt.Println("• Handle cash register")
	case "trainee":
		fmt.Println("• Observe operations")
		fmt.Println("• Assist with cleaning")
	default:
		fmt.Println("• No permissions")
	}
}

func processRushOrder(size string) {
	fmt.Printf("RUSH ORDER: %s coffee\n", size)
	fmt.Println("Speed preparation process:")
	
	rushed := true
	
	switch size {
	case "large":
		if rushed {
			fmt.Println("• Skip foam art")
			fallthrough
		}
		fmt.Println("• Use 16oz cup")
	case "medium":
		if rushed {
			fmt.Println("• Quick pour")
			fallthrough
		}
		fmt.Println("• Use 12oz cup")
	case "small":
		fmt.Println("• Use 8oz cup")
		fmt.Println("• Standard preparation")
	}
	
	fmt.Println("• Mark as RUSH")
	fmt.Println("• Priority queue")
}

func demonstrateFallthroughGotcha(value int) {
	fmt.Printf("Value: %d\n", value)
	
	// This shows why fallthrough can be confusing
	switch value {
	case 1:
		fmt.Println("Case 1")
		fallthrough
	case 2:
		fmt.Println("Case 2")
		// Note: Even though value is 2, if it was 1,
		// it would execute BOTH case 1 and case 2
		fallthrough
	case 3:
		fmt.Println("Case 3")
		// This executes even though value is 2, not 3!
	case 4:
		fmt.Println("Case 4")
		// This doesn't execute (no fallthrough from case 3)
	}
	
	fmt.Println("\n⚠️  Warning: fallthrough executes next case")
	fmt.Println("regardless of its condition!")
}