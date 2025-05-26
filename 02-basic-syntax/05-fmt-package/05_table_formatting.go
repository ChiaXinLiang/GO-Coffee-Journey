package main

import (
    "fmt"
    "strings"
    "time"
)

type SalesData struct {
    Date     time.Time
    Product  string
    Quantity int
    Revenue  float64
    Cost     float64
}

func main() {
    fmt.Println("=== GoCoffee Table Formatting ===\n")
    
    // Sample sales data
    sales := []SalesData{
        {time.Now().AddDate(0, 0, -6), "Espresso", 45, 135.00, 67.50},
        {time.Now().AddDate(0, 0, -5), "Latte", 62, 279.00, 124.00},
        {time.Now().AddDate(0, 0, -4), "Cappuccino", 38, 152.00, 76.00},
        {time.Now().AddDate(0, 0, -3), "Americano", 29, 101.50, 43.50},
        {time.Now().AddDate(0, 0, -2), "Mocha", 18, 90.00, 54.00},
        {time.Now().AddDate(0, 0, -1), "Macchiato", 25, 125.00, 62.50},
        {time.Now(), "Cold Brew", 41, 184.50, 82.00},
    }
    
    // Different table styles
    printBasicTable(sales)
    fmt.Println()
    printAlignedTable(sales)
    fmt.Println()
    printSummaryReport(sales)
    fmt.Println()
    printASCIIArtTable(sales)
}

func printBasicTable(sales []SalesData) {
    fmt.Println("BASIC TABLE")
    fmt.Println("-----------")
    
    fmt.Println("Date       Product     Qty  Revenue  Cost    Profit")
    for _, s := range sales {
        profit := s.Revenue - s.Cost
        fmt.Printf("%s %-11s %3d  $%6.2f  $%6.2f  $%6.2f\n",
            s.Date.Format("01/02"),
            s.Product,
            s.Quantity,
            s.Revenue,
            s.Cost,
            profit)
    }
}

func printAlignedTable(sales []SalesData) {
    fmt.Println("ALIGNED TABLE WITH FORMATTING")
    fmt.Println("=============================")
    
    // Define column widths
    const (
        dateWidth    = 10
        productWidth = 12
        qtyWidth     = 5
        moneyWidth   = 9
    )
    
    // Header
    fmt.Printf("%-*s %-*s %*s %*s %*s %*s %*s\n",
        dateWidth, "Date",
        productWidth, "Product",
        qtyWidth, "Qty",
        moneyWidth, "Revenue",
        moneyWidth, "Cost",
        moneyWidth, "Profit",
        moneyWidth, "Margin")
    
    // Separator
    fmt.Println(strings.Repeat("-", 
        dateWidth+productWidth+qtyWidth+moneyWidth*4+7))
    
    // Data rows
    totalRevenue, totalCost := 0.0, 0.0
    
    for _, s := range sales {
        profit := s.Revenue - s.Cost
        margin := (profit / s.Revenue) * 100
        
        totalRevenue += s.Revenue
        totalCost += s.Cost
        
        fmt.Printf("%-*s %-*s %*d %*s %*s %*s %*.1f%%\n",
            dateWidth, s.Date.Format("2006-01-02"),
            productWidth, truncate(s.Product, productWidth),
            qtyWidth, s.Quantity,
            moneyWidth, formatMoney(s.Revenue),
            moneyWidth, formatMoney(s.Cost),
            moneyWidth, formatMoney(profit),
            moneyWidth-1, margin)
    }
    
    // Total row
    fmt.Println(strings.Repeat("=", 
        dateWidth+productWidth+qtyWidth+moneyWidth*4+7))
    
    totalProfit := totalRevenue - totalCost
    totalMargin := (totalProfit / totalRevenue) * 100
    
    fmt.Printf("%-*s %*s %*s %*s %*s %*.1f%%\n",
        dateWidth+productWidth+1, "TOTAL",
        qtyWidth, "",
        moneyWidth, formatMoney(totalRevenue),
        moneyWidth, formatMoney(totalCost),
        moneyWidth, formatMoney(totalProfit),
        moneyWidth-1, totalMargin)
}

func printSummaryReport(sales []SalesData) {
    fmt.Println("╔═══════════════════════════════════════╗")
    fmt.Printf("║%s║\n", center("WEEKLY SALES SUMMARY", 39))
    fmt.Println("╠═══════════════════════════════════════╣")
    
    // Calculate totals by product
    productTotals := make(map[string]struct {
        qty     int
        revenue float64
    })
    
    totalQty := 0
    totalRevenue := 0.0
    
    for _, s := range sales {
        total := productTotals[s.Product]
        total.qty += s.Quantity
        total.revenue += s.Revenue
        productTotals[s.Product] = total
        
        totalQty += s.Quantity
        totalRevenue += s.Revenue
    }
    
    // Find best seller
    bestProduct := ""
    bestRevenue := 0.0
    
    for product, total := range productTotals {
        fmt.Printf("║ %-15s %4d units  $%7.2f ║\n",
            product, total.qty, total.revenue)
        
        if total.revenue > bestRevenue {
            bestRevenue = total.revenue
            bestProduct = product
        }
    }
    
    fmt.Println("╟───────────────────────────────────────╢")
    fmt.Printf("║ %-15s %4d units  $%7.2f ║\n",
        "TOTAL", totalQty, totalRevenue)
    fmt.Println("╠═══════════════════════════════════════╣")
    
    avgSale := totalRevenue / float64(totalQty)
    fmt.Printf("║ Average sale:     $%-18.2f ║\n", avgSale)
    fmt.Printf("║ Best seller:      %-19s ║\n", bestProduct)
    fmt.Printf("║ Daily average:    $%-18.2f ║\n", totalRevenue/7)
    
    fmt.Println("╚═══════════════════════════════════════╝")
}

func printASCIIArtTable(sales []SalesData) {
    fmt.Println("VISUAL SALES CHART")
    fmt.Println("==================")
    
    // Find max revenue for scaling
    maxRevenue := 0.0
    for _, s := range sales {
        if s.Revenue > maxRevenue {
            maxRevenue = s.Revenue
        }
    }
    
    // Print bars
    barWidth := 40
    
    for _, s := range sales {
        barLength := int((s.Revenue / maxRevenue) * float64(barWidth))
        
        fmt.Printf("%s %-11s [%s%s] $%.2f\n",
            s.Date.Format("01/02"),
            s.Product,
            strings.Repeat("█", barLength),
            strings.Repeat("░", barWidth-barLength),
            s.Revenue)
    }
    
    fmt.Printf("\n%s %s\n", strings.Repeat(" ", 19),
        "└─ Revenue Scale: $0 to $%.2f", maxRevenue)
}

func formatMoney(amount float64) string {
    return fmt.Sprintf("$%.2f", amount)
}

func truncate(s string, maxLen int) string {
    if len(s) <= maxLen {
        return s
    }
    return s[:maxLen-3] + "..."
}

func center(s string, width int) string {
    padding := (width - len(s)) / 2
    return fmt.Sprintf("%*s%s%*s", padding, "", s, width-len(s)-padding, "")
}