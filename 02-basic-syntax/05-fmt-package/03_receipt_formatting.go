package main

import (
    "fmt"
    "strings"
    "time"
)

type OrderItem struct {
    Name     string
    Size     string
    Price    float64
    Quantity int
    Mods     []string
}

type Receipt struct {
    OrderNumber int
    Items       []OrderItem
    Subtotal    float64
    Tax         float64
    Tip         float64
    Total       float64
    PaymentType string
    Cashier     string
}

func main() {
    fmt.Println("=== GoCoffee Receipt Formatting ===\n")
    
    // Create sample order
    receipt := Receipt{
        OrderNumber: 1047,
        Items: []OrderItem{
            {
                Name:     "Cappuccino",
                Size:     "Large",
                Price:    5.50,
                Quantity: 2,
                Mods:     []string{"Oat Milk", "Extra Shot"},
            },
            {
                Name:     "Vanilla Latte",
                Size:     "Medium",
                Price:    4.75,
                Quantity: 1,
                Mods:     []string{"Decaf", "2 Pumps Vanilla"},
            },
            {
                Name:     "Chocolate Croissant",
                Price:    3.50,
                Quantity: 2,
            },
        },
        PaymentType: "Credit Card",
        Cashier:     "Sarah M.",
    }
    
    // Calculate totals
    calculateTotals(&receipt)
    
    // Print different receipt styles
    printBasicReceipt(receipt)
    fmt.Println()
    printFancyReceipt(receipt)
    fmt.Println()
    printMinimalReceipt(receipt)
}

func calculateTotals(r *Receipt) {
    r.Subtotal = 0
    for _, item := range r.Items {
        r.Subtotal += item.Price * float64(item.Quantity)
    }
    r.Tax = r.Subtotal * 0.085
    r.Tip = r.Subtotal * 0.18
    r.Total = r.Subtotal + r.Tax + r.Tip
}

func printBasicReceipt(r Receipt) {
    width := 40
    
    fmt.Println(strings.Repeat("=", width))
    fmt.Printf("%s\n", center("GOCOFFEE", width))
    fmt.Printf("%s\n", center("Downtown Location", width))
    fmt.Printf("%s\n", center("123 Coffee St, Seattle WA", width))
    fmt.Println(strings.Repeat("=", width))
    
    fmt.Printf("Order #%d\n", r.OrderNumber)
    fmt.Printf("Date: %s\n", time.Now().Format("Jan 2, 2006 3:04 PM"))
    fmt.Printf("Cashier: %s\n", r.Cashier)
    fmt.Println(strings.Repeat("-", width))
    
    // Items
    for _, item := range r.Items {
        // Item line
        if item.Size != "" {
            fmt.Printf("%d %s %s\n", item.Quantity, item.Size, item.Name)
        } else {
            fmt.Printf("%d %s\n", item.Quantity, item.Name)
        }
        
        // Modifications
        for _, mod := range item.Mods {
            fmt.Printf("  + %s\n", mod)
        }
        
        // Price
        itemTotal := item.Price * float64(item.Quantity)
        fmt.Printf("%s$%.2f\n", strings.Repeat(" ", width-7), itemTotal)
    }
    
    fmt.Println(strings.Repeat("-", width))
    
    // Totals
    fmt.Printf("%-*s $%6.2f\n", width-9, "Subtotal:", r.Subtotal)
    fmt.Printf("%-*s $%6.2f\n", width-9, "Tax:", r.Tax)
    fmt.Printf("%-*s $%6.2f\n", width-9, "Tip (18%):", r.Tip)
    fmt.Println(strings.Repeat("=", width))
    fmt.Printf("%-*s $%6.2f\n", width-9, "TOTAL:", r.Total)
    fmt.Println(strings.Repeat("=", width))
    
    fmt.Printf("\nPayment: %s\n", r.PaymentType)
    fmt.Printf("\n%s\n", center("Thank You!", width))
    fmt.Printf("%s\n", center("Visit us at gocoffee.com", width))
}

func printFancyReceipt(r Receipt) {
    width := 48
    
    // Header
    fmt.Println("╔" + strings.Repeat("═", width-2) + "╗")
    fmt.Printf("║%s║\n", center("☕ GOCOFFEE ☕", width-2))
    fmt.Printf("║%s║\n", center("Premium Coffee Experience", width-2))
    fmt.Println("╠" + strings.Repeat("═", width-2) + "╣")
    
    // Order info
    fmt.Printf("║ Order #%-*d║\n", width-10, r.OrderNumber)
    fmt.Printf("║ %s%*s║\n", time.Now().Format("Jan 2, 15:04"), width-15, "")
    fmt.Println("╟" + strings.Repeat("─", width-2) + "╢")
    
    // Items
    for _, item := range r.Items {
        name := fmt.Sprintf("%d× %s", item.Quantity, item.Name)
        if item.Size != "" {
            name = fmt.Sprintf("%d× %s %s", item.Quantity, item.Size, item.Name)
        }
        
        price := fmt.Sprintf("$%.2f", item.Price*float64(item.Quantity))
        spacing := width - len(name) - len(price) - 4
        fmt.Printf("║ %s%*s%s ║\n", name, spacing, "", price)
        
        for _, mod := range item.Mods {
            modText := "  ⤷ " + mod
            fmt.Printf("║ %-*s ║\n", width-4, modText)
        }
    }
    
    // Footer
    fmt.Println("╟" + strings.Repeat("─", width-2) + "╢")
    printReceiptLine("Subtotal", fmt.Sprintf("$%.2f", r.Subtotal), width)
    printReceiptLine("Tax", fmt.Sprintf("$%.2f", r.Tax), width)
    printReceiptLine("Tip", fmt.Sprintf("$%.2f", r.Tip), width)
    fmt.Println("╟" + strings.Repeat("─", width-2) + "╢")
    printReceiptLine("TOTAL", fmt.Sprintf("$%.2f", r.Total), width)
    fmt.Println("╚" + strings.Repeat("═", width-2) + "╝")
}

func printMinimalReceipt(r Receipt) {
    fmt.Printf("GoCoffee #%d\n\n", r.OrderNumber)
    
    for _, item := range r.Items {
        fmt.Printf("%dx %s", item.Quantity, item.Name)
        if len(item.Mods) > 0 {
            fmt.Printf(" [%s]", strings.Join(item.Mods, ", "))
        }
        fmt.Printf(" ... $%.2f\n", item.Price*float64(item.Quantity))
    }
    
    fmt.Printf("\nTotal: $%.2f (incl. tax & tip)\n", r.Total)
    fmt.Println("Thank you!")
}

func center(s string, width int) string {
    if len(s) >= width {
        return s
    }
    padding := (width - len(s)) / 2
    return fmt.Sprintf("%*s%s%*s", padding, "", s, width-len(s)-padding, "")
}

func printReceiptLine(label, value string, width int) {
    spacing := width - len(label) - len(value) - 4
    fmt.Printf("║ %s%*s%s ║\n", label, spacing, "", value)
}