package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

type MenuOption struct {
    ID          int
    Title       string
    Description string
    Action      func()
    SubMenu     []MenuOption
}

var currentUser = "Marcus"
var isManager = true

func main() {
    fmt.Println("=== GoCoffee Management System ===\n")
    
    mainMenu := []MenuOption{
        {
            ID:          1,
            Title:       "Order Management",
            Description: "Manage customer orders",
            SubMenu: []MenuOption{
                {1, "New Order", "Create a new order", createNewOrder, nil},
                {2, "View Orders", "View pending orders", viewOrders, nil},
                {3, "Update Order Status", "Change order status", updateOrderStatus, nil},
                {4, "Cancel Order", "Cancel an order", cancelOrder, nil},
            },
        },
        {
            ID:          2,
            Title:       "Inventory",
            Description: "Manage inventory",
            SubMenu: []MenuOption{
                {1, "View Stock", "Check current stock levels", viewStock, nil},
                {2, "Add Stock", "Add items to inventory", addStock, nil},
                {3, "Low Stock Alert", "View items running low", lowStockAlert, nil},
            },
        },
        {
            ID:          3,
            Title:       "Reports",
            Description: "View reports and analytics",
            SubMenu: []MenuOption{
                {1, "Daily Sales", "Today's sales report", dailySalesReport, nil},
                {2, "Popular Items", "Best selling items", popularItemsReport, nil},
                {3, "Staff Performance", "Employee metrics", staffPerformance, nil},
            },
        },
        {
            ID:          4,
            Title:       "Settings",
            Description: "System settings",
            Action:      showSettings,
        },
    }
    
    // Add manager-only options
    if isManager {
        mainMenu = append(mainMenu, MenuOption{
            ID:          5,
            Title:       "Manager Tools",
            Description: "Manager-only functions",
            SubMenu: []MenuOption{
                {1, "Employee Management", "Manage staff", manageEmployees, nil},
                {2, "Price Adjustment", "Update menu prices", adjustPrices, nil},
                {3, "Promotional Offers", "Manage promotions", managePromotions, nil},
            },
        })
    }
    
    navigateMenu(mainMenu, "Main Menu")
}

func navigateMenu(menu []MenuOption, title string) {
    reader := bufio.NewReader(os.Stdin)
    
    for {
        // Clear screen (simulated)
        fmt.Println("\n" + strings.Repeat("=", 50))
        fmt.Printf("%s (User: %s)\n", title, currentUser)
        fmt.Println(strings.Repeat("=", 50))
        
        // Display menu options
        for _, option := range menu {
            fmt.Printf("%d. %s", option.ID, option.Title)
            if option.Description != "" {
                fmt.Printf(" - %s", option.Description)
            }
            fmt.Println()
        }
        fmt.Println("0. Back/Exit")
        
        // Get user choice
        fmt.Print("\nEnter your choice: ")
        input, _ := reader.ReadString('\n')
        choice, err := strconv.Atoi(strings.TrimSpace(input))
        
        if err != nil {
            fmt.Println("❌ Invalid input. Please enter a number.")
            fmt.Print("Press Enter to continue...")
            reader.ReadString('\n')
            continue
        }
        
        // Handle choice
        if choice == 0 {
            return
        }
        
        // Find selected option
        var selected *MenuOption
        for i := range menu {
            if menu[i].ID == choice {
                selected = &menu[i]
                break
            }
        }
        
        if selected == nil {
            fmt.Println("❌ Invalid option. Please try again.")
            fmt.Print("Press Enter to continue...")
            reader.ReadString('\n')
            continue
        }
        
        // Execute action or show submenu
        if selected.SubMenu != nil {
            navigateMenu(selected.SubMenu, selected.Title)
        } else if selected.Action != nil {
            fmt.Println("\n" + strings.Repeat("-", 50))
            selected.Action()
            fmt.Print("\nPress Enter to continue...")
            reader.ReadString('\n')
        }
    }
}

// Sample action functions
func createNewOrder() {
    fmt.Println("📝 Creating new order...")
    fmt.Println("Order creation would happen here")
}

func viewOrders() {
    fmt.Println("📋 Pending Orders:")
    fmt.Println("1. Order #1001 - 2 Lattes - Preparing")
    fmt.Println("2. Order #1002 - 1 Cappuccino - Ready")
    fmt.Println("3. Order #1003 - 3 Americanos - New")
}

func updateOrderStatus() {
    fmt.Println("🔄 Update order status functionality")
}

func cancelOrder() {
    fmt.Println("❌ Cancel order functionality")
}

func viewStock() {
    fmt.Println("📦 Current Stock Levels:")
    fmt.Println("Coffee Beans: 45 kg")
    fmt.Println("Milk: 120 liters")
    fmt.Println("Cups: 850 units")
}

func addStock() {
    fmt.Println("➕ Add stock functionality")
}

func lowStockAlert() {
    fmt.Println("⚠️  Low Stock Items:")
    fmt.Println("- Oat Milk: 5 liters (reorder soon)")
    fmt.Println("- Medium Cups: 50 units (critical)")
}

func dailySalesReport() {
    fmt.Println("💰 Daily Sales Report")
    fmt.Println("Total Revenue: $1,234.56")
    fmt.Println("Orders: 87")
    fmt.Println("Average Order: $14.19")
}

func popularItemsReport() {
    fmt.Println("⭐ Popular Items Today:")
    fmt.Println("1. Latte - 45 sold")
    fmt.Println("2. Cappuccino - 32 sold")
    fmt.Println("3. Croissant - 28 sold")
}

func staffPerformance() {
    fmt.Println("👥 Staff Performance")
    fmt.Println("Best Barista: Sarah (45 orders)")
}

func showSettings() {
    fmt.Println("⚙️  System Settings")
    fmt.Printf("Current User: %s\n", currentUser)
    fmt.Printf("Manager Access: %v\n", isManager)
    fmt.Println("Store: Downtown Location")
}

func manageEmployees() {
    fmt.Println("👤 Employee Management")
    fmt.Println("Total Employees: 12")
    fmt.Println("On Duty: 4")
}

func adjustPrices() {
    fmt.Println("💲 Price Adjustment Tool")
    fmt.Println("Current pricing loaded...")
}

func managePromotions() {
    fmt.Println("🎁 Promotional Offers")
    fmt.Println("Active Promotions: 3")
}