package main

import (
    "fmt"
    "strings"
)

type MenuItem struct {
    Category    string
    Name        string
    Description string
    Prices      map[string]float64 // size -> price
    Available   bool
    Popular     bool
    New         bool
}

func main() {
    fmt.Println("=== GoCoffee Menu Display ===\n")
    
    menu := []MenuItem{
        {
            Category:    "Hot Drinks",
            Name:        "Signature Espresso",
            Description: "Rich, bold espresso with hints of chocolate",
            Prices:      map[string]float64{"Solo": 2.50, "Doppio": 3.50},
            Available:   true,
            Popular:     true,
        },
        {
            Category:    "Hot Drinks",
            Name:        "Vanilla Latte",
            Description: "Smooth espresso with steamed milk and vanilla",
            Prices:      map[string]float64{"Small": 4.00, "Medium": 4.75, "Large": 5.50},
            Available:   true,
            New:         true,
        },
        {
            Category:    "Cold Drinks",
            Name:        "Iced Americano",
            Description: "Espresso over ice with cold water",
            Prices:      map[string]float64{"Small": 3.50, "Medium": 4.25, "Large": 5.00},
            Available:   true,
        },
        {
            Category:    "Cold Drinks",
            Name:        "Cold Brew",
            Description: "Smooth, less acidic coffee steeped for 20 hours",
            Prices:      map[string]float64{"Medium": 4.50, "Large": 5.25},
            Available:   false,
        },
        {
            Category:    "Food",
            Name:        "Avocado Toast",
            Description: "Multigrain bread with fresh avocado, tomatoes, and herbs",
            Prices:      map[string]float64{"Regular": 8.50},
            Available:   true,
            Popular:     true,
        },
    }
    
    // Display different menu styles
    printClassicMenu(menu)
    fmt.Println()
    printCompactMenu(menu)
    fmt.Println()
    printDigitalMenu(menu)
}

func printClassicMenu(menu []MenuItem) {
    fmt.Println("╔════════════════════════════════════════════════════════════╗")
    fmt.Printf("║%s║\n", center("☕ GOCOFFEE MENU ☕", 60))
    fmt.Println("╚════════════════════════════════════════════════════════════╝")
    
    // Group by category
    categories := make(map[string][]MenuItem)
    for _, item := range menu {
        categories[item.Category] = append(categories[item.Category], item)
    }
    
    // Print each category
    for _, category := range []string{"Hot Drinks", "Cold Drinks", "Food"} {
        items, exists := categories[category]
        if !exists {
            continue
        }
        
        fmt.Printf("\n【 %s 】\n", strings.ToUpper(category))
        fmt.Println(strings.Repeat("─", 60))
        
        for _, item := range items {
            // Name with badges
            nameLine := item.Name
            if item.Popular {
                nameLine += " ⭐"
            }
            if item.New {
                nameLine += " 🆕"
            }
            if !item.Available {
                nameLine += " (Sold Out)"
            }
            
            fmt.Printf("\n%-40s\n", nameLine)
            
            // Description
            fmt.Printf("  %s\n", item.Description)
            
            // Prices
            fmt.Print("  ")
            first := true
            for size, price := range item.Prices {
                if !first {
                    fmt.Print(" | ")
                }
                fmt.Printf("%s: $%.2f", size, price)
                first = false
            }
            fmt.Println()
        }
    }
    
    fmt.Println("\n" + strings.Repeat("─", 60))
    fmt.Println("⭐ Popular  🆕 New Item")
}

func printCompactMenu(menu []MenuItem) {
    fmt.Println("GOCOFFEE EXPRESS MENU")
    fmt.Println("====================")
    
    lastCategory := ""
    for _, item := range menu {
        if item.Category != lastCategory {
            fmt.Printf("\n%s:\n", item.Category)
            lastCategory = item.Category
        }
        
        // Skip unavailable items in compact view
        if !item.Available {
            continue
        }
        
        // Find min and max prices
        minPrice, maxPrice := 999.99, 0.0
        for _, price := range item.Prices {
            if price < minPrice {
                minPrice = price
            }
            if price > maxPrice {
                maxPrice = price
            }
        }
        
        priceStr := fmt.Sprintf("$%.2f", minPrice)
        if minPrice != maxPrice {
            priceStr = fmt.Sprintf("$%.2f-%.2f", minPrice, maxPrice)
        }
        
        fmt.Printf("  %-25s %10s", item.Name, priceStr)
        
        if item.Popular {
            fmt.Print(" ★")
        }
        if item.New {
            fmt.Print(" NEW")
        }
        fmt.Println()
    }
}

func printDigitalMenu(menu []MenuItem) {
    // Digital board style with columns
    fmt.Println("┌─────────────────────────┬─────────┬─────────┬─────────┐")
    fmt.Println("│      GOCOFFEE MENU      │  SMALL  │ MEDIUM  │  LARGE  │")
    fmt.Println("├─────────────────────────┼─────────┼─────────┼─────────┤")
    
    for i, item := range menu {
        if !item.Available {
            continue
        }
        
        name := item.Name
        if item.New {
            name = "» " + name
        }
        
        // Get prices for standard sizes
        small := getPriceForSize(item.Prices, "Small", "Solo", "")
        medium := getPriceForSize(item.Prices, "Medium", "Doppio", "Regular")
        large := getPriceForSize(item.Prices, "Large", "", "")
        
        fmt.Printf("│ %-23s │", name)
        printPriceCell(small)
        fmt.Print(" │")
        printPriceCell(medium)
        fmt.Print(" │")
        printPriceCell(large)
        fmt.Println(" │")
        
        if i < len(menu)-1 {
            fmt.Println("├─────────────────────────┼─────────┼─────────┼─────────┤")
        }
    }
    
    fmt.Println("└─────────────────────────┴─────────┴─────────┴─────────┘")
}

func getPriceForSize(prices map[string]float64, sizes ...string) float64 {
    for _, size := range sizes {
        if price, exists := prices[size]; exists {
            return price
        }
    }
    return 0
}

func printPriceCell(price float64) {
    if price > 0 {
        fmt.Printf(" %7.2f", price)
    } else {
        fmt.Print("    -   ")
    }
}

func center(text string, width int) string {
    padding := (width - len(text)) / 2
    return fmt.Sprintf("%*s%s%*s", padding, "", text, width-len(text)-padding, "")
}