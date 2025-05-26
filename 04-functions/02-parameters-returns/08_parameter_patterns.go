package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	fmt.Println("=== GoCoffee Parameter Patterns ===\n")
	
	// Example 1: Option pattern
	fmt.Println("Example 1: Options Pattern")
	fmt.Println("--------------------------")
	demonstrateOptionsPattern()
	
	// Example 2: Builder pattern
	fmt.Println("\n\nExample 2: Builder Pattern")
	fmt.Println("--------------------------")
	demonstrateBuilderPattern()
	
	// Example 3: Configuration pattern
	fmt.Println("\n\nExample 3: Configuration Pattern")
	fmt.Println("--------------------------------")
	demonstrateConfigPattern()
	
	// Example 4: Method chaining
	fmt.Println("\n\nExample 4: Method Chaining")
	fmt.Println("--------------------------")
	demonstrateMethodChaining()
}

// Example 1: Options pattern for optional parameters
func demonstrateOptionsPattern() {
	// Basic coffee with defaults
	coffee1 := NewCoffee("Latte")
	fmt.Printf("Basic: %+v\n", coffee1)
	
	// Coffee with options
	coffee2 := NewCoffee("Cappuccino",
		WithSize("Large"),
		WithExtraShots(2),
		WithMilkType("Oat"),
		WithTemperature(180),
	)
	fmt.Printf("Custom: %+v\n", coffee2)
	
	// Partial options
	coffee3 := NewCoffee("Espresso",
		WithSize("Small"),
		WithExtraShots(1),
	)
	fmt.Printf("Partial: %+v\n", coffee3)
}

// Coffee represents our coffee order
type Coffee struct {
	Type        string
	Size        string
	Shots       int
	Milk        string
	Temperature int
	Decaf       bool
}

// Option is a function that modifies a Coffee
type Option func(*Coffee)

// NewCoffee creates a coffee with options
func NewCoffee(coffeeType string, options ...Option) *Coffee {
	// Default coffee
	c := &Coffee{
		Type:        coffeeType,
		Size:        "Medium",
		Shots:       1,
		Milk:        "Whole",
		Temperature: 165,
		Decaf:       false,
	}
	
	// Apply all options
	for _, opt := range options {
		opt(c)
	}
	
	return c
}

// Option functions
func WithSize(size string) Option {
	return func(c *Coffee) {
		c.Size = size
	}
}

func WithExtraShots(extra int) Option {
	return func(c *Coffee) {
		c.Shots += extra
	}
}

func WithMilkType(milk string) Option {
	return func(c *Coffee) {
		c.Milk = milk
	}
}

func WithTemperature(temp int) Option {
	return func(c *Coffee) {
		c.Temperature = temp
	}
}

func WithDecaf() Option {
	return func(c *Coffee) {
		c.Decaf = true
	}
}

// Example 2: Builder pattern
func demonstrateBuilderPattern() {
	// Build an order step by step
	order := NewOrderBuilder().
		SetCustomer("Alice").
		AddItem("Latte", 4.50).
		AddItem("Croissant", 3.25).
		AddItem("Muffin", 3.00).
		SetPaymentMethod("Credit Card").
		Build()
	
	fmt.Printf("Order built: %+v\n", order)
	
	// Another order with different items
	order2 := NewOrderBuilder().
		SetCustomer("Bob").
		AddItem("Espresso", 2.50).
		SetTakeout(true).
		ApplyDiscount(0.10).
		Build()
	
	fmt.Printf("Order 2: %+v\n", order2)
}

// OrderBuilder builds orders step by step
type OrderBuilder struct {
	order *Order
}

type Order struct {
	Customer      string
	Items         []OrderItem
	PaymentMethod string
	Takeout       bool
	Discount      float64
	Total         float64
}

type OrderItem struct {
	Name  string
	Price float64
}

func NewOrderBuilder() *OrderBuilder {
	return &OrderBuilder{
		order: &Order{
			Items: []OrderItem{},
		},
	}
}

func (b *OrderBuilder) SetCustomer(name string) *OrderBuilder {
	b.order.Customer = name
	return b
}

func (b *OrderBuilder) AddItem(name string, price float64) *OrderBuilder {
	b.order.Items = append(b.order.Items, OrderItem{Name: name, Price: price})
	return b
}

func (b *OrderBuilder) SetPaymentMethod(method string) *OrderBuilder {
	b.order.PaymentMethod = method
	return b
}

func (b *OrderBuilder) SetTakeout(takeout bool) *OrderBuilder {
	b.order.Takeout = takeout
	return b
}

func (b *OrderBuilder) ApplyDiscount(rate float64) *OrderBuilder {
	b.order.Discount = rate
	return b
}

func (b *OrderBuilder) Build() *Order {
	// Calculate total
	total := 0.0
	for _, item := range b.order.Items {
		total += item.Price
	}
	
	// Apply discount
	if b.order.Discount > 0 {
		total *= (1 - b.order.Discount)
	}
	
	b.order.Total = total
	return b.order
}

// Example 3: Configuration pattern
func demonstrateConfigPattern() {
	// Default config
	shop1 := NewCoffeeShop(DefaultConfig())
	fmt.Printf("Default shop: %+v\n", shop1)
	
	// Custom config
	config := ShopConfig{
		Name:        "GoCoffee Express",
		OpenTime:    "05:00",
		CloseTime:   "23:00",
		MaxCapacity: 30,
		WiFiEnabled: true,
		Features: []string{
			"Drive-through",
			"Mobile ordering",
			"Rewards program",
		},
	}
	
	shop2 := NewCoffeeShop(config)
	fmt.Printf("Custom shop: %+v\n", shop2)
	
	// Config from function
	shop3 := NewCoffeeShop(WeekendConfig())
	fmt.Printf("Weekend shop: %+v\n", shop3)
}

type ShopConfig struct {
	Name        string
	OpenTime    string
	CloseTime   string
	MaxCapacity int
	WiFiEnabled bool
	Features    []string
}

type CoffeeShop struct {
	config ShopConfig
	isOpen bool
}

func DefaultConfig() ShopConfig {
	return ShopConfig{
		Name:        "GoCoffee",
		OpenTime:    "06:00",
		CloseTime:   "22:00",
		MaxCapacity: 50,
		WiFiEnabled: true,
		Features:    []string{"WiFi", "Seating", "Takeout"},
	}
}

func WeekendConfig() ShopConfig {
	config := DefaultConfig()
	config.OpenTime = "07:00"
	config.CloseTime = "23:00"
	config.Features = append(config.Features, "Live music", "Brunch menu")
	return config
}

func NewCoffeeShop(config ShopConfig) *CoffeeShop {
	return &CoffeeShop{
		config: config,
		isOpen: false,
	}
}

// Example 4: Method chaining / Fluent interface
func demonstrateMethodChaining() {
	// Create a receipt with chaining
	receipt := NewReceipt().
		SetOrderNumber("ORD-123").
		SetDate(time.Now()).
		AddLine("Latte", 2, 4.50).
		AddLine("Croissant", 1, 3.25).
		AddLine("Espresso", 1, 2.50).
		ApplyTax(0.08).
		AddTip(2.00).
		Generate()
	
	fmt.Println(receipt)
	
	// Barista workflow with chaining
	drink := NewDrink("Latte").
		GrindBeans(18).
		ExtractEspresso(25).
		SteamMilk(150, 165).
		Pour().
		AddArt("Heart").
		Serve()
	
	fmt.Printf("\nPrepared: %s\n", drink)
}

// Receipt with fluent interface
type Receipt struct {
	orderNumber string
	date        time.Time
	lines       []ReceiptLine
	subtotal    float64
	tax         float64
	tip         float64
	total       float64
}

type ReceiptLine struct {
	Item     string
	Quantity int
	Price    float64
	Total    float64
}

func NewReceipt() *Receipt {
	return &Receipt{
		lines: []ReceiptLine{},
	}
}

func (r *Receipt) SetOrderNumber(num string) *Receipt {
	r.orderNumber = num
	return r
}

func (r *Receipt) SetDate(date time.Time) *Receipt {
	r.date = date
	return r
}

func (r *Receipt) AddLine(item string, qty int, price float64) *Receipt {
	line := ReceiptLine{
		Item:     item,
		Quantity: qty,
		Price:    price,
		Total:    float64(qty) * price,
	}
	r.lines = append(r.lines, line)
	r.subtotal += line.Total
	return r
}

func (r *Receipt) ApplyTax(rate float64) *Receipt {
	r.tax = r.subtotal * rate
	return r
}

func (r *Receipt) AddTip(amount float64) *Receipt {
	r.tip = amount
	return r
}

func (r *Receipt) Generate() string {
	r.total = r.subtotal + r.tax + r.tip
	
	output := fmt.Sprintf("=== RECEIPT %s ===\n", r.orderNumber)
	output += fmt.Sprintf("Date: %s\n\n", r.date.Format("Jan 2, 2006 3:04 PM"))
	
	for _, line := range r.lines {
		output += fmt.Sprintf("%-15s %d × $%.2f = $%.2f\n", 
			line.Item, line.Quantity, line.Price, line.Total)
	}
	
	output += strings.Repeat("-", 40) + "\n"
	output += fmt.Sprintf("%-25s $%.2f\n", "Subtotal:", r.subtotal)
	output += fmt.Sprintf("%-25s $%.2f\n", "Tax:", r.tax)
	
	if r.tip > 0 {
		output += fmt.Sprintf("%-25s $%.2f\n", "Tip:", r.tip)
	}
	
	output += fmt.Sprintf("%-25s $%.2f\n", "Total:", r.total)
	
	return output
}

// Drink preparation with method chaining
type Drink struct {
	name        string
	steps       []string
	temperature int
}

func NewDrink(name string) *Drink {
	return &Drink{
		name:  name,
		steps: []string{},
	}
}

func (d *Drink) GrindBeans(grams int) *Drink {
	d.steps = append(d.steps, fmt.Sprintf("Ground %dg beans", grams))
	return d
}

func (d *Drink) ExtractEspresso(seconds int) *Drink {
	d.steps = append(d.steps, fmt.Sprintf("Extracted espresso for %d seconds", seconds))
	return d
}

func (d *Drink) SteamMilk(ml, temp int) *Drink {
	d.steps = append(d.steps, fmt.Sprintf("Steamed %dml milk to %d°F", ml, temp))
	d.temperature = temp
	return d
}

func (d *Drink) Pour() *Drink {
	d.steps = append(d.steps, "Poured milk into espresso")
	return d
}

func (d *Drink) AddArt(design string) *Drink {
	d.steps = append(d.steps, fmt.Sprintf("Added %s latte art", design))
	return d
}

func (d *Drink) Serve() string {
	preparation := strings.Join(d.steps, " → ")
	return fmt.Sprintf("%s: %s", d.name, preparation)
}

// Additional patterns
func additionalPatterns() {
	// Variadic options
	type Server struct {
		host string
		port int
	}
	
	type ServerOption func(*Server)
	
	NewServer := func(opts ...ServerOption) *Server {
		s := &Server{
			host: "localhost",
			port: 8080,
		}
		for _, opt := range opts {
			opt(s)
		}
		return s
	}
	
	WithHost := func(host string) ServerOption {
		return func(s *Server) {
			s.host = host
		}
	}
	
	WithPort := func(port int) ServerOption {
		return func(s *Server) {
			s.port = port
		}
	}
	
	// Usage
	_ = NewServer(
		WithHost("coffeeshop.com"),
		WithPort(443),
	)
	
	// Functional options with validation
	type ValidatedOption func(*Coffee) error
	
	WithValidatedSize := func(size string) ValidatedOption {
		return func(c *Coffee) error {
			if size != "Small" && size != "Medium" && size != "Large" {
				return fmt.Errorf("invalid size: %s", size)
			}
			c.Size = size
			return nil
		}
	}
}