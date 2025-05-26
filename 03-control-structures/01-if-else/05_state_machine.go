package main

import (
    "fmt"
    "time"
)

// OrderState represents the current state of an order
type OrderState string

const (
    StateNew        OrderState = "NEW"
    StatePaid       OrderState = "PAID"
    StatePreparing  OrderState = "PREPARING"
    StateReady      OrderState = "READY"
    StateDelivered  OrderState = "DELIVERED"
    StateCancelled  OrderState = "CANCELLED"
)

type CoffeeOrder struct {
    ID         int
    State      OrderState
    Items      []string
    Total      float64
    CreatedAt  time.Time
    PaidAt     *time.Time
    ReadyAt    *time.Time
    DeliveredAt *time.Time
}

func main() {
    fmt.Println("=== GoCoffee Order State Machine ===\n")
    
    // Create a new order
    order := &CoffeeOrder{
        ID:        2001,
        State:     StateNew,
        Items:     []string{"Latte", "Croissant"},
        Total:     8.00,
        CreatedAt: time.Now(),
    }
    
    // Simulate order lifecycle
    fmt.Printf("Order #%d created\n", order.ID)
    displayOrderState(order)
    
    // Process payment
    fmt.Println("\n💳 Processing payment...")
    if err := processPayment(order); err != nil {
        fmt.Printf("Error: %s\n", err)
        return
    }
    displayOrderState(order)
    
    // Start preparation
    fmt.Println("\n👨‍🍳 Starting preparation...")
    if err := startPreparation(order); err != nil {
        fmt.Printf("Error: %s\n", err)
        return
    }
    displayOrderState(order)
    
    // Mark as ready
    fmt.Println("\n✅ Order ready...")
    if err := markAsReady(order); err != nil {
        fmt.Printf("Error: %s\n", err)
        return
    }
    displayOrderState(order)
    
    // Deliver order
    fmt.Println("\n📦 Delivering order...")
    if err := deliverOrder(order); err != nil {
        fmt.Printf("Error: %s\n", err)
        return
    }
    displayOrderState(order)
    
    // Try invalid transition
    fmt.Println("\n❌ Attempting invalid transition...")
    if err := processPayment(order); err != nil {
        fmt.Printf("Error: %s\n", err)
    }
}

func processPayment(order *CoffeeOrder) error {
    // Only NEW orders can be paid
    if order.State != StateNew {
        return fmt.Errorf("cannot process payment: order is %s", order.State)
    }
    
    now := time.Now()
    order.PaidAt = &now
    order.State = StatePaid
    
    return nil
}

func startPreparation(order *CoffeeOrder) error {
    // Only PAID orders can start preparation
    if order.State != StatePaid {
        return fmt.Errorf("cannot start preparation: order is %s", order.State)
    }
    
    order.State = StatePreparing
    
    return nil
}

func markAsReady(order *CoffeeOrder) error {
    // Only PREPARING orders can be marked ready
    if order.State != StatePreparing {
        return fmt.Errorf("cannot mark as ready: order is %s", order.State)
    }
    
    now := time.Now()
    order.ReadyAt = &now
    order.State = StateReady
    
    return nil
}

func deliverOrder(order *CoffeeOrder) error {
    // Only READY orders can be delivered
    if order.State != StateReady {
        return fmt.Errorf("cannot deliver: order is %s", order.State)
    }
    
    now := time.Now()
    order.DeliveredAt = &now
    order.State = StateDelivered
    
    return nil
}

func cancelOrder(order *CoffeeOrder) error {
    // Can cancel if not already delivered or cancelled
    if order.State == StateDelivered {
        return fmt.Errorf("cannot cancel delivered order")
    }
    
    if order.State == StateCancelled {
        return fmt.Errorf("order already cancelled")
    }
    
    order.State = StateCancelled
    
    return nil
}

func displayOrderState(order *CoffeeOrder) {
    fmt.Printf("\nOrder #%d Status:\n", order.ID)
    fmt.Printf("State: %s\n", order.State)
    fmt.Printf("Items: %v\n", order.Items)
    fmt.Printf("Total: $%.2f\n", order.Total)
    
    // Show timeline
    fmt.Println("\nTimeline:")
    fmt.Printf("Created:   %s\n", order.CreatedAt.Format("3:04:05 PM"))
    
    if order.PaidAt != nil {
        fmt.Printf("Paid:      %s\n", order.PaidAt.Format("3:04:05 PM"))
    }
    
    if order.ReadyAt != nil {
        fmt.Printf("Ready:     %s\n", order.ReadyAt.Format("3:04:05 PM"))
    }
    
    if order.DeliveredAt != nil {
        fmt.Printf("Delivered: %s\n", order.DeliveredAt.Format("3:04:05 PM"))
    }
    
    // Show allowed transitions
    fmt.Print("\nNext actions: ")
    switch order.State {
    case StateNew:
        fmt.Println("Pay, Cancel")
    case StatePaid:
        fmt.Println("Start Preparation, Cancel")
    case StatePreparing:
        fmt.Println("Mark Ready, Cancel")
    case StateReady:
        fmt.Println("Deliver")
    case StateDelivered:
        fmt.Println("None (Order Complete)")
    case StateCancelled:
        fmt.Println("None (Order Cancelled)")
    }
}