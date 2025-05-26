package main

import (
    "fmt"
    "strings"
)

type Order struct {
    ID          int
    CustomerID  string
    Items       []string
    TotalAmount float64
    IsPaid      bool
    IsReady     bool
    IsPriority  bool
}

func main() {
    fmt.Println("=== GoCoffee Order Flow Control ===\n")
    
    // Sample orders
    orders := []Order{
        {1001, "C101", []string{"Latte", "Muffin"}, 8.50, true, false, false},
        {1002, "C102", []string{"Espresso"}, 3.00, false, false, true},
        {1003, "C103", []string{"Cappuccino", "Croissant"}, 7.50, true, true, false},
        {1004, "C104", []string{"Mocha", "Sandwich"}, 12.00, true, false, true},
        {1005, "C105", []string{"Americano"}, 3.50, false, false, false},
    }
    
    // Process each order
    for _, order := range orders {
        processOrderFlow(order)
        fmt.Println()
    }
    
    // Summary statistics
    printOrderStats(orders)
}

func processOrderFlow(order Order) {
    fmt.Printf("Processing Order #%d\n", order.ID)
    fmt.Println(strings.Repeat("-", 30))
    
    // Step 1: Check payment
    if !order.IsPaid {
        fmt.Println("⏸️ Status: Awaiting Payment")
        fmt.Printf("  Amount due: $%.2f\n", order.TotalAmount)
        
        // Priority orders get special handling
        if order.IsPriority {
            fmt.Println("  ⚡ Priority order - Sending payment reminder")
        }
        return // Can't proceed without payment
    }
    
    fmt.Println("✓ Payment confirmed")
    
    // Step 2: Check if ready
    if order.IsReady {
        fmt.Println("✅ Status: Ready for Pickup")
        fmt.Printf("  Items: %s\n", strings.Join(order.Items, ", "))
        
        // Send notification based on order value
        if order.TotalAmount > 10 {
            fmt.Println("  📱 SMS notification sent")
        } else {
            fmt.Println("  🔔 Order number called")
        }
        return
    }
    
    // Step 3: Process order
    fmt.Println("🔄 Status: In Preparation")
    
    // Check order complexity
    if len(order.Items) > 3 {
        fmt.Println("  ⚠️ Complex order - Assigned to senior barista")
    } else if order.IsPriority {
        fmt.Println("  ⚡ Priority queue - Expedited preparation")
    } else {
        fmt.Println("  📋 Added to standard queue")
    }
    
    // Estimate preparation time
    prepTime := estimatePrepTime(order)
    fmt.Printf("  ⏱️ Estimated time: %d minutes\n", prepTime)
    
    // Special handling for large orders
    if order.TotalAmount > 20 {
        fmt.Println("  🎁 Complimentary cookie added")
    }
}

func estimatePrepTime(order Order) int {
    baseTime := 3 // Base prep time in minutes
    
    // Add time for each item
    itemTime := len(order.Items) * 2
    
    // Rush hour adjustment
    rushHourPenalty := 0
    // Simulating rush hour (in real app, would check actual time)
    isRushHour := true
    if isRushHour {
        rushHourPenalty = 5
    }
    
    // Priority bonus
    priorityBonus := 0
    if order.IsPriority {
        priorityBonus = -2 // Negative means faster
    }
    
    totalTime := baseTime + itemTime + rushHourPenalty + priorityBonus
    
    // Ensure minimum time
    if totalTime < 1 {
        totalTime = 1
    }
    
    return totalTime
}

func printOrderStats(orders []Order) {
    fmt.Println("=== Order Statistics ===")
    
    totalOrders := len(orders)
    paidOrders := 0
    readyOrders := 0
    priorityOrders := 0
    totalRevenue := 0.0
    
    for _, order := range orders {
        if order.IsPaid {
            paidOrders++
            totalRevenue += order.TotalAmount
        }
        if order.IsReady {
            readyOrders++
        }
        if order.IsPriority {
            priorityOrders++
        }
    }
    
    pendingPayment := totalOrders - paidOrders
    inPreparation := paidOrders - readyOrders
    
    fmt.Printf("Total orders: %d\n", totalOrders)
    fmt.Printf("Paid orders: %d\n", paidOrders)
    fmt.Printf("Ready for pickup: %d\n", readyOrders)
    fmt.Printf("Priority orders: %d\n", priorityOrders)
    fmt.Printf("Pending payment: %d\n", pendingPayment)
    fmt.Printf("In preparation: %d\n", inPreparation)
    fmt.Printf("Total revenue: $%.2f\n", totalRevenue)
    
    if paidOrders > 0 {
        avgOrderValue := totalRevenue / float64(paidOrders)
        fmt.Printf("Average order value: $%.2f\n", avgOrderValue)
    }
    
    // Performance metrics
    fmt.Println("\nPerformance:")
    
    paymentRate := float64(paidOrders) / float64(totalOrders) * 100
    if paymentRate < 80 {
        fmt.Printf("⚠️ Low payment rate: %.1f%%\n", paymentRate)
    } else {
        fmt.Printf("✓ Good payment rate: %.1f%%\n", paymentRate)
    }
    
    if inPreparation > 5 {
        fmt.Println("⚠️ High number of orders in preparation - need more staff!")
    }
}