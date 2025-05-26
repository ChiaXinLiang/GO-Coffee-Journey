package main

import (
	"fmt"
	"time"
	"strings"
	"sort"
)

// === Coffee Shop Order Management System ===

type OrderStatus string

const (
	StatusPending    OrderStatus = "pending"
	StatusProcessing OrderStatus = "processing"
	StatusReady      OrderStatus = "ready"
	StatusCompleted  OrderStatus = "completed"
)

type OrderItem struct {
	Name     string
	Price    float64
	Quantity int
	Options  []string
}

type Order struct {
	ID         string
	CustomerID string
	Items      []OrderItem
	Status     OrderStatus
	CreatedAt  time.Time
	Notes      []string
}

// OrderSystem manages coffee shop orders
type OrderSystem struct {
	orders map[string]*Order
	nextID int
}

func NewOrderSystem() *OrderSystem {
	return &OrderSystem{
		orders: make(map[string]*Order),
		nextID: 1000,
	}
}

// CreateOrder with variadic items - natural use case
func (os *OrderSystem) CreateOrder(customerID string, items ...OrderItem) *Order {
	if len(items) == 0 {
		return nil // No empty orders
	}
	
	orderID := fmt.Sprintf("ORD-%d", os.nextID)
	os.nextID++
	
	order := &Order{
		ID:         orderID,
		CustomerID: customerID,
		Items:      items,
		Status:     StatusPending,
		CreatedAt:  time.Now(),
		Notes:      []string{},
	}
	
	os.orders[orderID] = order
	return order
}

// AddNotes to an order - variadic for flexibility
func (os *OrderSystem) AddNotes(orderID string, notes ...string) error {
	order, exists := os.orders[orderID]
	if !exists {
		return fmt.Errorf("order %s not found", orderID)
	}
	
	order.Notes = append(order.Notes, notes...)
	return nil
}

// === Notification System ===

type NotificationType string

const (
	NotifyEmail NotificationType = "email"
	NotifySMS   NotificationType = "sms"
	NotifyPush  NotificationType = "push"
)

type Notification struct {
	Type    NotificationType
	Message string
	Target  string
}

// NotificationService handles multiple notification channels
type NotificationService struct {
	handlers map[NotificationType]func(Notification) error
}

func NewNotificationService() *NotificationService {
	ns := &NotificationService{
		handlers: make(map[NotificationType]func(Notification) error),
	}
	
	// Register default handlers
	ns.handlers[NotifyEmail] = func(n Notification) error {
		fmt.Printf("📧 Email to %s: %s\n", n.Target, n.Message)
		return nil
	}
	
	ns.handlers[NotifySMS] = func(n Notification) error {
		fmt.Printf("📱 SMS to %s: %s\n", n.Target, n.Message)
		return nil
	}
	
	return ns
}

// SendNotifications to multiple channels - perfect for variadic
func (ns *NotificationService) SendNotifications(message string, targets ...Notification) []error {
	var errors []error
	
	for _, target := range targets {
		if handler, exists := ns.handlers[target.Type]; exists {
			target.Message = message
			if err := handler(target); err != nil {
				errors = append(errors, err)
			}
		}
	}
	
	return errors
}

// === Menu Builder with Options ===

type MenuItem struct {
	Name        string
	Category    string
	BasePrice   float64
	Available   bool
	Options     map[string]float64
	Description string
}

type MenuOption func(*MenuItem)

// Menu building with variadic options pattern
func NewMenuItem(name string, category string, price float64, options ...MenuOption) *MenuItem {
	item := &MenuItem{
		Name:      name,
		Category:  category,
		BasePrice: price,
		Available: true,
		Options:   make(map[string]float64),
	}
	
	for _, opt := range options {
		opt(item)
	}
	
	return item
}

func WithDescription(desc string) MenuOption {
	return func(mi *MenuItem) {
		mi.Description = desc
	}
}

func WithOptions(opts ...string) MenuOption {
	return func(mi *MenuItem) {
		for _, opt := range opts {
			mi.Options[opt] = 0.50 // default addon price
		}
	}
}

func WithCustomOptions(opts map[string]float64) MenuOption {
	return func(mi *MenuItem) {
		for k, v := range opts {
			mi.Options[k] = v
		}
	}
}

func Unavailable() MenuOption {
	return func(mi *MenuItem) {
		mi.Available = false
	}
}

// === Analytics and Reporting ===

type SalesReport struct {
	Period    string
	TotalSales float64
	ItemsSold  map[string]int
	TopItems   []string
}

// GenerateReport with flexible filtering
func GenerateReport(orders []*Order, filters ...func(*Order) bool) *SalesReport {
	report := &SalesReport{
		Period:    time.Now().Format("2006-01-02"),
		ItemsSold: make(map[string]int),
	}
	
	// Apply all filters
	filteredOrders := orders
	for _, filter := range filters {
		var temp []*Order
		for _, order := range filteredOrders {
			if filter(order) {
				temp = append(temp, order)
			}
		}
		filteredOrders = temp
	}
	
	// Calculate stats
	for _, order := range filteredOrders {
		for _, item := range order.Items {
			report.ItemsSold[item.Name] += item.Quantity
			report.TotalSales += item.Price * float64(item.Quantity)
		}
	}
	
	// Find top items
	type itemCount struct {
		name  string
		count int
	}
	
	var counts []itemCount
	for name, count := range report.ItemsSold {
		counts = append(counts, itemCount{name, count})
	}
	
	sort.Slice(counts, func(i, j int) bool {
		return counts[i].count > counts[j].count
	})
	
	for i := 0; i < 3 && i < len(counts); i++ {
		report.TopItems = append(report.TopItems, counts[i].name)
	}
	
	return report
}

// Filter functions
func CompletedOnly(o *Order) bool {
	return o.Status == StatusCompleted
}

func TodayOnly(o *Order) bool {
	return o.CreatedAt.Day() == time.Now().Day()
}

func CustomerFilter(customerID string) func(*Order) bool {
	return func(o *Order) bool {
		return o.CustomerID == customerID
	}
}

func main() {
	fmt.Println("=== Real World Variadic Examples ===")
	fmt.Println()

	// Initialize systems
	orderSystem := NewOrderSystem()
	notifyService := NewNotificationService()
	
	// Create orders with variadic items
	order1 := orderSystem.CreateOrder("CUST-001",
		OrderItem{Name: "Latte", Price: 4.50, Quantity: 2},
		OrderItem{Name: "Croissant", Price: 3.25, Quantity: 1},
		OrderItem{Name: "Espresso", Price: 2.75, Quantity: 1, 
			Options: []string{"Extra Shot", "Oat Milk"}},
	)
	
	fmt.Printf("Created order: %s\n", order1.ID)
	
	// Add notes using variadic
	orderSystem.AddNotes(order1.ID, 
		"Customer allergic to nuts",
		"Extra hot",
		"To go",
	)
	
	// Send notifications to multiple channels
	fmt.Println("\nSending notifications:")
	notifyService.SendNotifications(
		fmt.Sprintf("Order %s is ready!", order1.ID),
		Notification{Type: NotifyEmail, Target: "customer@email.com"},
		Notification{Type: NotifySMS, Target: "+1234567890"},
	)
	
	// Build menu with options
	fmt.Println("\nBuilding menu:")
	
	latte := NewMenuItem("Latte", "Coffee", 4.50,
		WithDescription("Smooth espresso with steamed milk"),
		WithOptions("Extra Shot", "Decaf", "Sugar Free"),
		WithCustomOptions(map[string]float64{
			"Oat Milk": 0.75,
			"Soy Milk": 0.50,
		}),
	)
	
	seasonal := NewMenuItem("Pumpkin Spice Latte", "Seasonal", 5.50,
		WithDescription("Fall favorite with real pumpkin"),
		Unavailable(), // Out of season
	)
	
	fmt.Printf("Menu Item: %s ($%.2f) - %s\n", 
		latte.Name, latte.BasePrice, latte.Description)
	fmt.Printf("Options: %v\n", latte.Options)
	fmt.Printf("Seasonal Item: %s (Available: %v)\n", 
		seasonal.Name, seasonal.Available)
	
	// Generate reports with filters
	fmt.Println("\nGenerating reports:")
	
	// Create more orders for reporting
	order2 := orderSystem.CreateOrder("CUST-002",
		OrderItem{Name: "Cappuccino", Price: 4.25, Quantity: 1},
		OrderItem{Name: "Muffin", Price: 3.50, Quantity: 2},
	)
	order2.Status = StatusCompleted
	
	order3 := orderSystem.CreateOrder("CUST-001",
		OrderItem{Name: "Latte", Price: 4.50, Quantity: 1},
	)
	order3.Status = StatusCompleted
	
	// Get all orders
	var allOrders []*Order
	for _, order := range orderSystem.orders {
		allOrders = append(allOrders, order)
	}
	
	// Generate report with multiple filters
	report := GenerateReport(allOrders, 
		CompletedOnly,
		CustomerFilter("CUST-001"),
	)
	
	fmt.Printf("Sales Report - %s\n", report.Period)
	fmt.Printf("Total Sales: $%.2f\n", report.TotalSales)
	fmt.Printf("Items Sold: %v\n", report.ItemsSold)
	
	// Batch operations example
	fmt.Println("\nBatch operations:")
	
	updateOrders := func(status OrderStatus, orderIDs ...string) {
		for _, id := range orderIDs {
			if order, exists := orderSystem.orders[id]; exists {
				order.Status = status
				fmt.Printf("Updated %s to %s\n", id, status)
			}
		}
	}
	
	updateOrders(StatusProcessing, order1.ID, order3.ID)
}

// Real-world benefits demonstrated:
// 1. Natural API for variable items (orders, notifications)
// 2. Flexible configuration (menu items with options)
// 3. Composable filters for reporting
// 4. Batch operations on multiple entities
// 5. Clean, readable code that matches business logic