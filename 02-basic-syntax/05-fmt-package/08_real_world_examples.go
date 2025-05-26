package main

import (
    "fmt"
    "strings"
    "time"
)

func main() {
    fmt.Println("=== GoCoffee Real-World Formatting ===\n")
    
    // Order tracking display
    printOrderTracking()
    
    // Daily report
    printDailyReport()
    
    // Employee schedule
    printEmployeeSchedule()
    
    // Inventory alert
    printInventoryAlert()
}

func printOrderTracking() {
    fmt.Println("ORDER TRACKING SYSTEM")
    fmt.Println("====================\n")
    
    orders := []struct {
        id       int
        customer string
        items    int
        status   string
        time     int // minutes ago
    }{
        {1051, "Sarah M.", 2, "Ready", 0},
        {1050, "John D.", 1, "Preparing", 2},
        {1049, "Lisa K.", 3, "Preparing", 5},
        {1048, "Mike R.", 1, "Completed", 10},
        {1047, "Emma S.", 4, "Completed", 15},
    }
    
    // Status colors
    statusColor := map[string]string{
        "Ready":     Green,
        "Preparing": Yellow,
        "Completed": Blue,
    }
    
    fmt.Printf("%-6s %-15s %-6s %-12s %-10s\n", 
        "Order", "Customer", "Items", "Status", "Time")
    fmt.Println(strings.Repeat("-", 50))
    
    for _, order := range orders {
        timeStr := "Just now"
        if order.time > 0 {
            timeStr = fmt.Sprintf("%d min ago", order.time)
        }
        
        color := statusColor[order.status]
        fmt.Printf("#%-5d %-15s %-6d %s%-12s%s %-10s\n",
            order.id,
            truncate(order.customer, 15),
            order.items,
            color, order.status, Reset,
            timeStr)
    }
    
    fmt.Printf("\n%s● Ready%s  %s● Preparing%s  %s● Completed%s\n\n",
        Green, Reset, Yellow, Reset, Blue, Reset)
}

func printDailyReport() {
    fmt.Println("DAILY SALES REPORT")
    fmt.Println("==================")
    fmt.Printf("Date: %s\n\n", time.Now().Format("Monday, January 2, 2006"))
    
    // Hourly sales
    hours := []string{
        "6-7 AM", "7-8 AM", "8-9 AM", "9-10 AM", "10-11 AM",
        "11-12 PM", "12-1 PM", "1-2 PM", "2-3 PM", "3-4 PM",
    }
    
    sales := []float64{
        45.50, 125.75, 285.25, 195.50, 145.00,
        165.25, 225.75, 189.00, 125.50, 95.25,
    }
    
    maxSale := 0.0
    for _, sale := range sales {
        if sale > maxSale {
            maxSale = sale
        }
    }
    
    fmt.Println("Hourly Sales:")
    for i, hour := range hours {
        sale := sales[i]
        barLen := int(sale / maxSale * 30)
        
        fmt.Printf("%-9s $%6.2f [%s%s]\n",
            hour, sale,
            strings.Repeat("█", barLen),
            strings.Repeat("░", 30-barLen))
    }
    
    // Summary
    total := 0.0
    for _, sale := range sales {
        total += sale
    }
    
    fmt.Printf("\n%-9s $%6.2f\n", "TOTAL:", total)
    fmt.Printf("%-9s $%6.2f\n", "Average:", total/float64(len(sales)))
    fmt.Printf("%-9s $%6.2f (8-9 AM)\n", "Peak:", maxSale)
}

func printEmployeeSchedule() {
    fmt.Println("\nEMPLOYEE SCHEDULE - TODAY")
    fmt.Println("=========================\n")
    
    schedule := []struct {
        name  string
        shift string
        hours string
        role  string
    }{
        {"Sarah M.", "Morning", "6:00 AM - 2:00 PM", "Barista"},
        {"John D.", "Morning", "7:00 AM - 3:00 PM", "Shift Lead"},
        {"Lisa K.", "Afternoon", "12:00 PM - 8:00 PM", "Barista"},
        {"Mike R.", "Afternoon", "2:00 PM - 10:00 PM", "Barista"},
        {"Emma S.", "Closing", "3:00 PM - 11:00 PM", "Shift Lead"},
    }
    
    // Group by shift
    fmt.Printf("%sMORNING SHIFT%s\n", Bold, Reset)
    for _, emp := range schedule {
        if emp.shift == "Morning" {
            fmt.Printf("  %-12s %-20s %s\n", emp.name, emp.hours, emp.role)
        }
    }
    
    fmt.Printf("\n%sAFTERNOON SHIFT%s\n", Bold, Reset)
    for _, emp := range schedule {
        if emp.shift == "Afternoon" {
            fmt.Printf("  %-12s %-20s %s\n", emp.name, emp.hours, emp.role)
        }
    }
    
    fmt.Printf("\n%sCLOSING SHIFT%s\n", Bold, Reset)
    for _, emp := range schedule {
        if emp.shift == "Closing" {
            fmt.Printf("  %-12s %-20s %s\n", emp.name, emp.hours, emp.role)
        }
    }
    
    // Coverage chart
    fmt.Println("\nCOVERAGE:")
    fmt.Println("6AM  8AM  10AM 12PM 2PM  4PM  6PM  8PM  10PM")
    fmt.Println("├────┼────┼────┼────┼────┼────┼────┼────┼────┤")
    
    // Visual representation of shifts
    shifts := []struct {
        name  string
        start int
        end   int
    }{
        {"Sarah", 6, 14},
        {"John", 7, 15},
        {"Lisa", 12, 20},
        {"Mike", 14, 22},
        {"Emma", 15, 23},
    }
    
    for _, shift := range shifts {
        fmt.Printf("%-5s ", shift.name[:5])
        
        for hour := 6; hour < 23; hour++ {
            if hour >= shift.start && hour < shift.end {
                fmt.Print("████")
            } else {
                fmt.Print("    ")
            }
            
            if hour < 22 {
                fmt.Print(" ")
            }
        }
        fmt.Println()
    }
}

func printInventoryAlert() {
    fmt.Println("\n⚠️  INVENTORY ALERTS")
    fmt.Println("===================\n")
    
    inventory := []struct {
        item     string
        current  float64
        unit     string
        minimum  float64
        status   string
    }{
        {"Espresso Beans", 15.5, "kg", 20.0, "Low"},
        {"Whole Milk", 8.0, "gal", 15.0, "Critical"},
        {"Oat Milk", 25.0, "L", 20.0, "OK"},
        {"Sugar", 5.0, "kg", 10.0, "Low"},
        {"Cups (12oz)", 150, "units", 500, "Critical"},
    }
    
    for _, item := range inventory {
        var color string
        var icon string
        
        switch item.status {
        case "OK":
            color = Green
            icon = "✓"
        case "Low":
            color = Yellow
            icon = "⚠"
        case "Critical":
            color = Red
            icon = "✗"
        }
        
        fmt.Printf("%s%s %-20s: %.1f %s (min: %.1f)%s\n",
            color, icon, item.item, item.current, item.unit, 
            item.minimum, Reset)
        
        // Visual bar
        percent := item.current / item.minimum * 100
        barWidth := 20
        filled := int(percent / 100 * float64(barWidth))
        if filled > barWidth {
            filled = barWidth
        }
        
        fmt.Printf("   [%s%s%s%s] %.0f%%\n\n",
            color, strings.Repeat("█", filled), Reset,
            strings.Repeat("░", barWidth-filled),
            percent)
    }
    
    fmt.Printf("%sAction Required:%s Order supplies marked as Critical immediately!\n",
        Bold+Red, Reset)
}

func truncate(s string, max int) string {
    if len(s) <= max {
        return s
    }
    return s[:max-3] + "..."
}

// Color constants
const (
    Reset  = "\033[0m"
    Red    = "\033[31m"
    Green  = "\033[32m"
    Yellow = "\033[33m"
    Blue   = "\033[34m"
    Bold   = "\033[1m"
)