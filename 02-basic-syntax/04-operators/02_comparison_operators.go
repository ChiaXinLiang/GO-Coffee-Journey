package main

import "fmt"

func main() {
    fmt.Println("=== GoCoffee Comparison Operations ===\n")
    
    // Price comparisons
    lattePrice := 4.50
    cappuccinoPrice := 4.00
    mochaPrice := 5.00
    budget := 5.00
    
    fmt.Println("PRICE COMPARISONS:")
    fmt.Printf("Latte ($%.2f) == Cappuccino ($%.2f): %v\n", 
        lattePrice, cappuccinoPrice, lattePrice == cappuccinoPrice)
    fmt.Printf("Latte ($%.2f) != Mocha ($%.2f): %v\n", 
        lattePrice, mochaPrice, lattePrice != mochaPrice)
    fmt.Printf("Latte ($%.2f) < Budget ($%.2f): %v\n", 
        lattePrice, budget, lattePrice < budget)
    fmt.Printf("Mocha ($%.2f) <= Budget ($%.2f): %v\n", 
        mochaPrice, budget, mochaPrice <= budget)
    
    // Inventory checks
    fmt.Println("\nINVENTORY CHECKS:")
    
    coffeeBags := 25
    milkGallons := 10
    
    minimumCoffee := 20
    minimumMilk := 15
    
    fmt.Printf("Coffee (%d) >= Minimum (%d): %v\n", 
        coffeeBags, minimumCoffee, coffeeBags >= minimumCoffee)
    fmt.Printf("Milk (%d) > Minimum (%d): %v\n", 
        milkGallons, minimumMilk, milkGallons > minimumMilk)
    
    // String comparisons
    fmt.Println("\nSTRING COMPARISONS:")
    
    customer1 := "Alice"
    customer2 := "Bob"
    vipCustomer := "Alice"
    
    fmt.Printf("'%s' == '%s': %v\n", customer1, vipCustomer, customer1 == vipCustomer)
    fmt.Printf("'%s' < '%s' (alphabetically): %v\n", customer1, customer2, customer1 < customer2)
    
    // Time-based comparisons
    fmt.Println("\nTIME COMPARISONS:")
    
    currentHour := 14 // 2 PM
    openingHour := 6
    closingHour := 22
    happyHourStart := 15
    happyHourEnd := 17
    
    isOpen := currentHour >= openingHour && currentHour < closingHour
    isHappyHour := currentHour >= happyHourStart && currentHour < happyHourEnd
    
    fmt.Printf("Current hour: %d:00\n", currentHour)
    fmt.Printf("Is open (6-22): %v\n", isOpen)
    fmt.Printf("Is happy hour (15-17): %v\n", isHappyHour)
    
    // Order validation
    fmt.Println("\nORDER VALIDATION:")
    
    orderTotal := 12.50
    minimumOrder := 5.00
    maximumOrder := 100.00
    
    validOrder := orderTotal >= minimumOrder && orderTotal <= maximumOrder
    fmt.Printf("Order $%.2f valid ($%.2f-$%.2f): %v\n", 
        orderTotal, minimumOrder, maximumOrder, validOrder)
    
    // Loyalty program
    fmt.Println("\nLOYALTY PROGRAM:")
    
    customerOrders := 15
    customerSpending := 250.00
    
    // Different tier requirements
    bronzeOrders := 5
    silverOrders := 10
    goldOrders := 20
    
    bronzeSpending := 100.00
    silverSpending := 200.00
    goldSpending := 500.00
    
    isBronze := customerOrders >= bronzeOrders || customerSpending >= bronzeSpending
    isSilver := customerOrders >= silverOrders || customerSpending >= silverSpending
    isGold := customerOrders >= goldOrders || customerSpending >= goldSpending
    
    fmt.Printf("Customer: %d orders, $%.2f spent\n", customerOrders, customerSpending)
    fmt.Printf("Bronze status: %v\n", isBronze)
    fmt.Printf("Silver status: %v\n", isSilver)
    fmt.Printf("Gold status: %v\n", isGold)
    
    // Practical example: Discount eligibility
    fmt.Println("\nDISCOUNT ELIGIBILITY:")
    
    type Order struct {
        total    float64
        items    int
        isMember bool
    }
    
    orders := []Order{
        {total: 25.00, items: 3, isMember: true},
        {total: 15.00, items: 2, isMember: false},
        {total: 50.00, items: 5, isMember: true},
    }
    
    for i, order := range orders {
        // 10% off for members OR orders over $30 OR 4+ items
        eligible := order.isMember || order.total > 30 || order.items >= 4
        
        fmt.Printf("Order %d: $%.2f, %d items, member:%v → Discount:%v\n", 
            i+1, order.total, order.items, order.isMember, eligible)
    }
}