// +build !production

package main

import (
    "fmt"
    "math/rand"
    "time"
)

// This file contains debug features that should not be in production

func main() {
    fmt.Println("=== GoCoffee Debug Mode ===")
    fmt.Println("⚠️  Running in DEBUG mode - Not for production use!\n")
    
    // Debug features
    runDiagnostics()
    simulateOrders()
    testPaymentSystem()
}

func runDiagnostics() {
    fmt.Println("🔍 Running System Diagnostics...")
    
    checks := []struct {
        name   string
        status bool
        detail string
    }{
        {"Database Connection", true, "Connected to test DB"},
        {"Payment Gateway", true, "Using sandbox mode"},
        {"Inventory System", false, "Mock data enabled"},
        {"Email Service", true, "Redirecting to test inbox"},
        {"SMS Service", false, "Disabled in debug mode"},
    }
    
    for _, check := range checks {
        if check.status {
            fmt.Printf("✅ %s: %s\n", check.name, check.detail)
        } else {
            fmt.Printf("❌ %s: %s\n", check.name, check.detail)
        }
    }
}

func simulateOrders() {
    fmt.Println("\n🎮 Simulating Customer Orders...")
    
    rand.Seed(time.Now().UnixNano())
    
    items := []string{"Latte", "Cappuccino", "Espresso", "Americano"}
    
    for i := 1; i <= 5; i++ {
        item := items[rand.Intn(len(items))]
        quantity := rand.Intn(3) + 1
        price := float64(rand.Intn(300)+200) / 100.0
        
        fmt.Printf("Order #TEST%03d: %d x %s @ $%.2f\n", 
            i, quantity, item, price)
    }
}

func testPaymentSystem() {
    fmt.Println("\n💳 Testing Payment System...")
    
    testCards := []struct {
        number string
        result string
    }{
        {"4111-1111-1111-1111", "Success"},
        {"4000-0000-0000-0002", "Declined"},
        {"4000-0000-0000-9995", "Insufficient Funds"},
    }
    
    for _, card := range testCards {
        fmt.Printf("Testing card %s: %s\n", card.number, card.result)
    }
    
    fmt.Println("\n📝 Debug Notes:")
    fmt.Println("- All transactions are simulated")
    fmt.Println("- No real charges will occur")
    fmt.Println("- Test data will be cleared daily")
}

// Production build would have a different file:
// +build production
// with actual implementation without debug features