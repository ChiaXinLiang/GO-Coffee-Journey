package main

import (
	"fmt"
	"strings"
)

// Define some custom types for the coffee shop
type Coffee struct {
	name string
	size string
}

type Payment struct {
	amount float64
	method string
}

type Customer struct {
	name   string
	status string
}

func main() {
	fmt.Println("=== GoCoffee Type Switches ===\n")
	
	// Example 1: Basic type switch
	fmt.Println("Example 1: Processing Different Order Types")
	fmt.Println("------------------------------------------")
	
	processItem("Latte")
	processItem(4.50)
	processItem(true)
	processItem(Coffee{name: "Espresso", size: "small"})
	processItem([]string{"Latte", "Muffin"})
	
	// Example 2: Order system with multiple types
	fmt.Println("\n\nExample 2: Complete Order System")
	fmt.Println("--------------------------------")
	
	// Process various order components
	orders := []interface{}{
		Customer{name: "Alice", status: "VIP"},
		Coffee{name: "Cappuccino", size: "large"},
		Payment{amount: 15.50, method: "card"},
		"Add extra shot",
		2, // Quantity
		true, // Rush order
	}
	
	for _, item := range orders {
		handleOrderComponent(item)
	}
	
	// Example 3: Flexible menu system
	fmt.Println("\n\nExample 3: Dynamic Menu Items")
	fmt.Println("-----------------------------")
	
	menuItems := []interface{}{
		"Espresso",
		Coffee{name: "Special Blend", size: "medium"},
		map[string]float64{"small": 2.50, "large": 4.50},
		[]string{"Organic", "Fair Trade", "Single Origin"},
	}
	
	for _, item := range menuItems {
		displayMenuItem(item)
	}
	
	// Example 4: Error handling with types
	fmt.Println("\n\nExample 4: Error Response Handling")
	fmt.Println("---------------------------------")
	
	responses := []interface{}{
		"Success",
		fmt.Errorf("payment declined"),
		404,
		Payment{amount: 10.00, method: "cash"},
		nil,
	}
	
	for _, resp := range responses {
		handleResponse(resp)
	}
}

func processItem(item interface{}) {
	switch v := item.(type) {
	case string:
		fmt.Printf("String order: %s\n", v)
		
	case int:
		fmt.Printf("Quantity: %d items\n", v)
		
	case float64:
		fmt.Printf("Price: $%.2f\n", v)
		
	case bool:
		if v {
			fmt.Println("Rush order: YES")
		} else {
			fmt.Println("Rush order: NO")
		}
		
	case Coffee:
		fmt.Printf("Coffee order: %s %s\n", v.size, v.name)
		
	case []string:
		fmt.Printf("Multiple items: %s\n", strings.Join(v, ", "))
		
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}
}

func handleOrderComponent(component interface{}) {
	switch v := component.(type) {
	case Customer:
		fmt.Printf("\n👤 Customer: %s", v.name)
		if v.status == "VIP" {
			fmt.Print(" ⭐ VIP")
		}
		fmt.Println()
		
	case Coffee:
		fmt.Printf("☕ Ordered: %s %s\n", v.size, v.name)
		price := calculatePrice(v)
		fmt.Printf("   Price: $%.2f\n", price)
		
	case Payment:
		fmt.Printf("💳 Payment: $%.2f via %s\n", v.amount, v.method)
		
	case string:
		fmt.Printf("📝 Special request: %s\n", v)
		
	case int:
		fmt.Printf("🔢 Quantity: %d\n", v)
		
	case bool:
		if v {
			fmt.Println("⚡ RUSH ORDER - Priority service")
		}
		
	default:
		fmt.Printf("❓ Unknown component type: %T\n", v)
	}
}

func displayMenuItem(item interface{}) {
	fmt.Println()
	
	switch v := item.(type) {
	case string:
		fmt.Printf("☕ %s - Classic menu item\n", v)
		
	case Coffee:
		fmt.Printf("☕ %s (%s) - Today's special\n", v.name, v.size)
		
	case map[string]float64:
		fmt.Println("💰 Price list:")
		for size, price := range v {
			fmt.Printf("   %s: $%.2f\n", size, price)
		}
		
	case []string:
		fmt.Printf("🏷️ Features: %s\n", strings.Join(v, ", "))
		
	default:
		fmt.Printf("❓ Unknown menu item type: %T\n", v)
	}
}

func handleResponse(response interface{}) {
	fmt.Print("Response: ")
	
	switch v := response.(type) {
	case string:
		fmt.Printf("✅ %s\n", v)
		
	case error:
		fmt.Printf("❌ Error: %v\n", v)
		
	case int:
		switch v {
		case 200:
			fmt.Println("✅ OK")
		case 404:
			fmt.Println("❌ Not Found")
		case 500:
			fmt.Println("❌ Server Error")
		default:
			fmt.Printf("Code: %d\n", v)
		}
		
	case Payment:
		fmt.Printf("✅ Payment processed: $%.2f\n", v.amount)
		
	case nil:
		fmt.Println("⚠️  No response")
		
	default:
		fmt.Printf("❓ Unexpected response type: %T\n", v)
	}
}

func calculatePrice(coffee Coffee) float64 {
	basePrice := 3.00
	
	switch coffee.size {
	case "small":
		return basePrice * 0.8
	case "medium":
		return basePrice
	case "large":
		return basePrice * 1.3
	default:
		return basePrice
	}
}