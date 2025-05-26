package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"time"
)

// Domain models
type Customer struct {
	ID           string
	Name         string
	Email        string
	LoyaltyTier  string
	TotalSpent   float64
	JoinDate     time.Time
}

type MenuItem struct {
	ID          string
	Name        string
	Category    string
	Price       float64
	Available   bool
	Ingredients []string
}

type OrderItem struct {
	MenuItem     MenuItem
	Quantity     int
	Customizations []string
	Price        float64
}

type Order struct {
	ID          string
	CustomerID  string
	Items       []OrderItem
	Status      string
	OrderTime   time.Time
	Subtotal    float64
	Tax         float64
	Discount    float64
	Total       float64
}

type Transaction struct {
	ID            string
	OrderID       string
	Amount        float64
	Method        string
	Status        string
	ProcessedTime time.Time
}

func main() {
	fmt.Println("=== GoCoffee Real-World Examples ===\n")
	
	// Example 1: Customer management
	fmt.Println("Example 1: Customer Management System")
	fmt.Println("------------------------------------")
	demonstrateCustomerManagement()
	
	// Example 2: Order processing
	fmt.Println("\n\nExample 2: Order Processing System")
	fmt.Println("---------------------------------")
	demonstrateOrderProcessing()
	
	// Example 3: Inventory management
	fmt.Println("\n\nExample 3: Inventory Management")
	fmt.Println("-------------------------------")
	demonstrateInventoryManagement()
	
	// Example 4: Analytics and reporting
	fmt.Println("\n\nExample 4: Analytics & Reporting")
	fmt.Println("--------------------------------")
	demonstrateAnalytics()
}

// Example 1: Customer management
func demonstrateCustomerManagement() {
	// Create customer database
	customers := initializeCustomers()
	
	// Find customer by ID
	customer, err := findCustomerByID(customers, "CUST-002")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Found customer: %s\n", customer.Name)
	}
	
	// Update loyalty tier
	updated, err := updateLoyaltyTier(customer, 500.00)
	if err != nil {
		fmt.Printf("Error updating tier: %v\n", err)
	} else {
		fmt.Printf("Updated tier: %s → %s\n", customer.LoyaltyTier, updated.LoyaltyTier)
	}
	
	// Get customer benefits
	benefits, discount := getCustomerBenefits(updated)
	fmt.Printf("Benefits: %v (%.0f%% discount)\n", benefits, discount*100)
	
	// Search customers
	results := searchCustomers(customers, "john", "email")
	fmt.Printf("\nSearch results for 'john':\n")
	for _, c := range results {
		fmt.Printf("  - %s (%s)\n", c.Name, c.Email)
	}
}

func initializeCustomers() []Customer {
	return []Customer{
		{
			ID:          "CUST-001",
			Name:        "Alice Smith",
			Email:       "alice@email.com",
			LoyaltyTier: "Silver",
			TotalSpent:  150.00,
			JoinDate:    time.Now().AddDate(0, -6, 0),
		},
		{
			ID:          "CUST-002",
			Name:        "Bob Johnson",
			Email:       "bob.j@email.com",
			LoyaltyTier: "Bronze",
			TotalSpent:  75.00,
			JoinDate:    time.Now().AddDate(0, -2, 0),
		},
		{
			ID:          "CUST-003",
			Name:        "John Doe",
			Email:       "john.doe@email.com",
			LoyaltyTier: "Gold",
			TotalSpent:  850.00,
			JoinDate:    time.Now().AddDate(-1, 0, 0),
		},
	}
}

func findCustomerByID(customers []Customer, id string) (*Customer, error) {
	for i := range customers {
		if customers[i].ID == id {
			return &customers[i], nil
		}
	}
	return nil, fmt.Errorf("customer not found: %s", id)
}

func updateLoyaltyTier(customer *Customer, newSpending float64) (*Customer, error) {
	if customer == nil {
		return nil, fmt.Errorf("customer cannot be nil")
	}
	
	// Update total spent
	customer.TotalSpent += newSpending
	
	// Determine new tier
	oldTier := customer.LoyaltyTier
	
	switch {
	case customer.TotalSpent >= 1000:
		customer.LoyaltyTier = "Platinum"
	case customer.TotalSpent >= 500:
		customer.LoyaltyTier = "Gold"
	case customer.TotalSpent >= 200:
		customer.LoyaltyTier = "Silver"
	default:
		customer.LoyaltyTier = "Bronze"
	}
	
	if oldTier != customer.LoyaltyTier {
		fmt.Printf("  🎉 Congratulations! Upgraded to %s tier!\n", customer.LoyaltyTier)
	}
	
	return customer, nil
}

func getCustomerBenefits(customer *Customer) ([]string, float64) {
	if customer == nil {
		return []string{}, 0
	}
	
	baseBenefits := []string{"Birthday reward", "Member pricing"}
	var discount float64
	
	switch customer.LoyaltyTier {
	case "Platinum":
		discount = 0.20
		benefits := append(baseBenefits,
			"20% discount",
			"Free size upgrades",
			"Priority service",
			"Exclusive events",
		)
		return benefits, discount
		
	case "Gold":
		discount = 0.15
		benefits := append(baseBenefits,
			"15% discount",
			"Free size upgrade on Fridays",
			"Early access to new items",
		)
		return benefits, discount
		
	case "Silver":
		discount = 0.10
		benefits := append(baseBenefits,
			"10% discount",
			"Double points on weekends",
		)
		return benefits, discount
		
	default: // Bronze
		discount = 0.05
		benefits := append(baseBenefits, "5% discount")
		return benefits, discount
	}
}

func searchCustomers(customers []Customer, query, field string) []Customer {
	var results []Customer
	query = strings.ToLower(query)
	
	for _, customer := range customers {
		var match bool
		
		switch field {
		case "name":
			match = strings.Contains(strings.ToLower(customer.Name), query)
		case "email":
			match = strings.Contains(strings.ToLower(customer.Email), query)
		case "all":
			match = strings.Contains(strings.ToLower(customer.Name), query) ||
				strings.Contains(strings.ToLower(customer.Email), query)
		default:
			match = strings.Contains(strings.ToLower(customer.Name), query)
		}
		
		if match {
			results = append(results, customer)
		}
	}
	
	return results
}

// Example 2: Order processing
func demonstrateOrderProcessing() {
	// Create menu
	menu := createMenu()
	
	// Create new order
	order, err := createOrder("CUST-001", menu)
	if err != nil {
		fmt.Printf("Error creating order: %v\n", err)
		return
	}
	
	fmt.Printf("Order created: %s\n", order.ID)
	displayOrderSummary(order)
	
	// Process payment
	transaction, err := processOrderPayment(order, "credit", "4242-4242-4242-4242")
	if err != nil {
		fmt.Printf("Payment failed: %v\n", err)
		return
	}
	
	fmt.Printf("\nPayment processed: %s\n", transaction.ID)
	
	// Update order status
	order.Status = "Preparing"
	estimatedTime := estimatePreparationTime(order)
	fmt.Printf("Estimated preparation time: %v\n", estimatedTime)
}

func createMenu() []MenuItem {
	return []MenuItem{
		{
			ID:          "MENU-001",
			Name:        "Espresso",
			Category:    "Coffee",
			Price:       2.50,
			Available:   true,
			Ingredients: []string{"Coffee beans"},
		},
		{
			ID:          "MENU-002",
			Name:        "Latte",
			Category:    "Coffee",
			Price:       4.50,
			Available:   true,
			Ingredients: []string{"Coffee beans", "Milk"},
		},
		{
			ID:          "MENU-003",
			Name:        "Croissant",
			Category:    "Pastry",
			Price:       3.25,
			Available:   true,
			Ingredients: []string{"Flour", "Butter", "Yeast"},
		},
	}
}

func createOrder(customerID string, menu []MenuItem) (*Order, error) {
	if customerID == "" {
		return nil, fmt.Errorf("customer ID required")
	}
	
	order := &Order{
		ID:         fmt.Sprintf("ORD-%d", time.Now().Unix()),
		CustomerID: customerID,
		Items:      []OrderItem{},
		Status:     "New",
		OrderTime:  time.Now(),
	}
	
	// Add items to order
	// Latte with customization
	if item, found := findMenuItem(menu, "MENU-002"); found {
		orderItem := OrderItem{
			MenuItem:       item,
			Quantity:       2,
			Customizations: []string{"Extra hot", "Oat milk"},
			Price:          item.Price,
		}
		order.Items = append(order.Items, orderItem)
	}
	
	// Croissant
	if item, found := findMenuItem(menu, "MENU-003"); found {
		orderItem := OrderItem{
			MenuItem: item,
			Quantity: 1,
			Price:    item.Price,
		}
		order.Items = append(order.Items, orderItem)
	}
	
	// Calculate totals
	order.Subtotal = calculateSubtotal(order.Items)
	order.Tax = order.Subtotal * 0.08
	order.Discount = 0 // Could apply customer discount here
	order.Total = order.Subtotal + order.Tax - order.Discount
	
	return order, nil
}

func findMenuItem(menu []MenuItem, id string) (MenuItem, bool) {
	for _, item := range menu {
		if item.ID == id {
			return item, true
		}
	}
	return MenuItem{}, false
}

func calculateSubtotal(items []OrderItem) float64 {
	total := 0.0
	for _, item := range items {
		total += item.Price * float64(item.Quantity)
	}
	return total
}

func displayOrderSummary(order *Order) {
	fmt.Println("\n--- ORDER SUMMARY ---")
	for _, item := range order.Items {
		fmt.Printf("%d × %s @ $%.2f = $%.2f\n",
			item.Quantity, item.MenuItem.Name, item.Price,
			item.Price*float64(item.Quantity))
		
		if len(item.Customizations) > 0 {
			fmt.Printf("   Customizations: %s\n", strings.Join(item.Customizations, ", "))
		}
	}
	
	fmt.Println(strings.Repeat("-", 30))
	fmt.Printf("Subtotal:  $%.2f\n", order.Subtotal)
	fmt.Printf("Tax:       $%.2f\n", order.Tax)
	if order.Discount > 0 {
		fmt.Printf("Discount: -$%.2f\n", order.Discount)
	}
	fmt.Printf("Total:     $%.2f\n", order.Total)
}

func processOrderPayment(order *Order, method, details string) (*Transaction, error) {
	if order == nil {
		return nil, fmt.Errorf("order cannot be nil")
	}
	
	if order.Total <= 0 {
		return nil, fmt.Errorf("invalid order total: %.2f", order.Total)
	}
	
	// Validate payment method
	validMethods := map[string]bool{
		"cash":   true,
		"credit": true,
		"debit":  true,
		"mobile": true,
	}
	
	if !validMethods[method] {
		return nil, fmt.Errorf("invalid payment method: %s", method)
	}
	
	// Create transaction
	transaction := &Transaction{
		ID:            fmt.Sprintf("TXN-%d", time.Now().Unix()),
		OrderID:       order.ID,
		Amount:        order.Total,
		Method:        method,
		Status:        "Processing",
		ProcessedTime: time.Now(),
	}
	
	// Simulate payment processing
	time.Sleep(500 * time.Millisecond)
	
	// 90% success rate for demo
	if time.Now().Unix()%10 < 9 {
		transaction.Status = "Completed"
		order.Status = "Paid"
	} else {
		transaction.Status = "Failed"
		return transaction, fmt.Errorf("payment declined")
	}
	
	return transaction, nil
}

func estimatePreparationTime(order *Order) time.Duration {
	if order == nil || len(order.Items) == 0 {
		return 0
	}
	
	// Base time per item
	baseTime := 2 * time.Minute
	
	// Additional time for customizations
	customizationTime := 30 * time.Second
	
	totalTime := time.Duration(0)
	
	for _, item := range order.Items {
		itemTime := baseTime * time.Duration(item.Quantity)
		itemTime += customizationTime * time.Duration(len(item.Customizations))
		
		// Coffee takes longer than pastries
		if item.MenuItem.Category == "Coffee" {
			itemTime += 1 * time.Minute
		}
		
		totalTime += itemTime
	}
	
	// Round to nearest minute
	return totalTime.Round(time.Minute)
}

// Example 3: Inventory management
func demonstrateInventoryManagement() {
	inventory := initializeInventory()
	
	// Check stock levels
	lowStock := checkLowStock(inventory, 20)
	if len(lowStock) > 0 {
		fmt.Println("⚠️  Low stock items:")
		for _, item := range lowStock {
			fmt.Printf("  - %s: %d units (min: %d)\n", 
				item.Name, item.Current, item.Minimum)
		}
	}
	
	// Use inventory for order
	itemsNeeded := map[string]int{
		"Coffee beans": 50,
		"Milk":        10,
		"Sugar":       5,
	}
	
	canFulfill, missing := checkInventoryAvailability(inventory, itemsNeeded)
	if canFulfill {
		fmt.Println("\n✅ Can fulfill order")
		updated := deductInventory(inventory, itemsNeeded)
		fmt.Println("Inventory updated:")
		for _, item := range updated {
			if item.Current < item.Optimal {
				fmt.Printf("  - %s: %d units (below optimal: %d)\n",
					item.Name, item.Current, item.Optimal)
			}
		}
	} else {
		fmt.Println("\n❌ Cannot fulfill order. Missing:")
		for item, needed := range missing {
			fmt.Printf("  - %s: need %d more units\n", item, needed)
		}
	}
}

type InventoryItem struct {
	ID       string
	Name     string
	Unit     string
	Current  int
	Minimum  int
	Optimal  int
	Supplier string
}

func initializeInventory() []InventoryItem {
	return []InventoryItem{
		{
			ID:       "INV-001",
			Name:     "Coffee beans",
			Unit:     "grams",
			Current:  5000,
			Minimum:  1000,
			Optimal:  8000,
			Supplier: "Premium Roasters",
		},
		{
			ID:       "INV-002",
			Name:     "Milk",
			Unit:     "liters",
			Current:  15,
			Minimum:  10,
			Optimal:  50,
			Supplier: "Local Dairy",
		},
		{
			ID:       "INV-003",
			Name:     "Sugar",
			Unit:     "grams",
			Current:  3000,
			Minimum:  500,
			Optimal:  5000,
			Supplier: "Sweet Supplies",
		},
	}
}

func checkLowStock(inventory []InventoryItem, threshold int) []InventoryItem {
	var lowStock []InventoryItem
	
	for _, item := range inventory {
		percentageRemaining := float64(item.Current) / float64(item.Optimal) * 100
		if percentageRemaining < float64(threshold) {
			lowStock = append(lowStock, item)
		}
	}
	
	return lowStock
}

func checkInventoryAvailability(inventory []InventoryItem, needed map[string]int) (bool, map[string]int) {
	missing := make(map[string]int)
	
	for itemName, amountNeeded := range needed {
		found := false
		
		for _, invItem := range inventory {
			if invItem.Name == itemName {
				found = true
				if invItem.Current < amountNeeded {
					missing[itemName] = amountNeeded - invItem.Current
				}
				break
			}
		}
		
		if !found {
			missing[itemName] = amountNeeded
		}
	}
	
	return len(missing) == 0, missing
}

func deductInventory(inventory []InventoryItem, used map[string]int) []InventoryItem {
	updated := make([]InventoryItem, len(inventory))
	copy(updated, inventory)
	
	for i := range updated {
		if amount, exists := used[updated[i].Name]; exists {
			updated[i].Current -= amount
			if updated[i].Current < 0 {
				updated[i].Current = 0
			}
		}
	}
	
	return updated
}

// Example 4: Analytics and reporting
func demonstrateAnalytics() {
	// Generate sample data
	orders := generateSampleOrders()
	
	// Daily summary
	summary := generateDailySummary(orders, time.Now())
	fmt.Printf("Daily Summary for %s:\n", summary.Date.Format("Jan 2, 2006"))
	fmt.Printf("  Orders: %d\n", summary.TotalOrders)
	fmt.Printf("  Revenue: $%.2f\n", summary.Revenue)
	fmt.Printf("  Average order: $%.2f\n", summary.AverageOrder)
	fmt.Printf("  Peak hour: %d:00\n", summary.PeakHour)
	
	// Top products
	fmt.Println("\n📊 Top Products:")
	topProducts := getTopProducts(orders, 3)
	for i, product := range topProducts {
		fmt.Printf("  %d. %s (%d sold, $%.2f revenue)\n",
			i+1, product.Name, product.Quantity, product.Revenue)
	}
	
	// Customer insights
	insights := analyzeCustomerBehavior(orders)
	fmt.Println("\n👥 Customer Insights:")
	fmt.Printf("  Unique customers: %d\n", insights.UniqueCustomers)
	fmt.Printf("  Repeat customers: %d (%.1f%%)\n", 
		insights.RepeatCustomers,
		float64(insights.RepeatCustomers)/float64(insights.UniqueCustomers)*100)
	fmt.Printf("  Most popular customization: %s\n", insights.PopularCustomization)
}

type DailySummary struct {
	Date         time.Time
	TotalOrders  int
	Revenue      float64
	AverageOrder float64
	PeakHour     int
}

type ProductSummary struct {
	Name     string
	Quantity int
	Revenue  float64
}

type CustomerInsights struct {
	UniqueCustomers      int
	RepeatCustomers      int
	PopularCustomization string
	AverageOrderValue    float64
}

func generateSampleOrders() []Order {
	// Generate realistic sample data
	var orders []Order
	
	// Simulate a day's worth of orders
	baseTime := time.Now().Truncate(24 * time.Hour).Add(6 * time.Hour) // 6 AM
	
	for hour := 0; hour < 16; hour++ { // 6 AM to 10 PM
		// More orders during peak hours
		orderCount := 5
		if hour >= 7 && hour <= 9 { // Morning rush
			orderCount = 15
		} else if hour >= 12 && hour <= 13 { // Lunch
			orderCount = 10
		}
		
		for i := 0; i < orderCount; i++ {
			order := Order{
				ID:         fmt.Sprintf("ORD-%d-%d", hour, i),
				CustomerID: fmt.Sprintf("CUST-%03d", (hour*10+i)%50),
				OrderTime:  baseTime.Add(time.Duration(hour) * time.Hour),
				Status:     "Completed",
			}
			
			// Random items
			if i%3 == 0 {
				order.Items = append(order.Items, OrderItem{
					MenuItem: MenuItem{Name: "Latte", Price: 4.50},
					Quantity: 1 + i%2,
				})
			}
			if i%2 == 0 {
				order.Items = append(order.Items, OrderItem{
					MenuItem: MenuItem{Name: "Espresso", Price: 2.50},
					Quantity: 1,
				})
			}
			if i%4 == 0 {
				order.Items = append(order.Items, OrderItem{
					MenuItem:       MenuItem{Name: "Cappuccino", Price: 4.00},
					Quantity:       1,
					Customizations: []string{"Extra foam"},
				})
			}
			
			order.Subtotal = calculateSubtotal(order.Items)
			order.Tax = order.Subtotal * 0.08
			order.Total = order.Subtotal + order.Tax
			
			orders = append(orders, order)
		}
	}
	
	return orders
}

func generateDailySummary(orders []Order, date time.Time) DailySummary {
	summary := DailySummary{
		Date: date.Truncate(24 * time.Hour),
	}
	
	hourCounts := make(map[int]int)
	
	for _, order := range orders {
		if order.OrderTime.Truncate(24*time.Hour) == summary.Date {
			summary.TotalOrders++
			summary.Revenue += order.Total
			
			hour := order.OrderTime.Hour()
			hourCounts[hour]++
		}
	}
	
	if summary.TotalOrders > 0 {
		summary.AverageOrder = summary.Revenue / float64(summary.TotalOrders)
	}
	
	// Find peak hour
	maxCount := 0
	for hour, count := range hourCounts {
		if count > maxCount {
			maxCount = count
			summary.PeakHour = hour
		}
	}
	
	return summary
}

func getTopProducts(orders []Order, limit int) []ProductSummary {
	productMap := make(map[string]*ProductSummary)
	
	for _, order := range orders {
		for _, item := range order.Items {
			key := item.MenuItem.Name
			if _, exists := productMap[key]; !exists {
				productMap[key] = &ProductSummary{Name: key}
			}
			
			productMap[key].Quantity += item.Quantity
			productMap[key].Revenue += item.Price * float64(item.Quantity)
		}
	}
	
	// Convert to slice and sort
	var products []ProductSummary
	for _, p := range productMap {
		products = append(products, *p)
	}
	
	sort.Slice(products, func(i, j int) bool {
		return products[i].Revenue > products[j].Revenue
	})
	
	if len(products) > limit {
		products = products[:limit]
	}
	
	return products
}

func analyzeCustomerBehavior(orders []Order) CustomerInsights {
	insights := CustomerInsights{}
	
	customerOrders := make(map[string]int)
	customizations := make(map[string]int)
	totalValue := 0.0
	
	for _, order := range orders {
		customerOrders[order.CustomerID]++
		totalValue += order.Total
		
		for _, item := range order.Items {
			for _, custom := range item.Customizations {
				customizations[custom]++
			}
		}
	}
	
	insights.UniqueCustomers = len(customerOrders)
	
	for _, count := range customerOrders {
		if count > 1 {
			insights.RepeatCustomers++
		}
	}
	
	if len(orders) > 0 {
		insights.AverageOrderValue = totalValue / float64(len(orders))
	}
	
	// Find most popular customization
	maxCount := 0
	for custom, count := range customizations {
		if count > maxCount {
			maxCount = count
			insights.PopularCustomization = custom
		}
	}
	
	if insights.PopularCustomization == "" {
		insights.PopularCustomization = "None"
	}
	
	return insights
}