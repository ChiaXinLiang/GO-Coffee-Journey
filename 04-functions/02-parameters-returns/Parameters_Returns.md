# Parameters & Returns: Input Ingredients, Output Coffee ☕↔️

"Think of parameters as ingredients going INTO our coffee machine," Sarah explains, "and return values as what comes OUT!"

## Function Parameters

### Basic Parameters

```go
// Single parameter
func grindCoffee(beans string) {
    fmt.Printf("Grinding %s beans...\n", beans)
}

// Multiple parameters of same type
func brewCoffee(beans, grind string) {  // Both are strings
    fmt.Printf("Brewing %s coffee with %s grind\n", beans, grind)
}

// Multiple parameters of different types
func makeEspresso(beans string, shots int, temperature float64) {
    fmt.Printf("Making %d shots of %s at %.1f°C\n", shots, beans, temperature)
}
```

### Parameter Passing

"Go passes parameters by value," Sarah notes. "Let's see what that means:"

```go
// Pass by value - original doesn't change
func addSugar(coffee string, spoons int) string {
    // This modifies the local copy, not the original
    return fmt.Sprintf("%s with %d spoons of sugar", coffee, spoons)
}

// Pass by value with structs
type CoffeeOrder struct {
    Type  string
    Size  string
    Price float64
}

func applyMemberDiscount(order CoffeeOrder) CoffeeOrder {
    order.Price *= 0.9  // 10% off
    return order
}

func main() {
    original := CoffeeOrder{Type: "Latte", Size: "Large", Price: 5.00}
    discounted := applyMemberDiscount(original)
    
    fmt.Printf("Original: $%.2f\n", original.Price)    // Still $5.00
    fmt.Printf("Discounted: $%.2f\n", discounted.Price) // Now $4.50
}
```

### Pointer Parameters

"Sometimes we need to modify the original," Sarah explains:

```go
// Using pointers to modify the original
func applyLoyaltyPoints(order *CoffeeOrder, points int) {
    discount := float64(points) * 0.01  // 1 cent per point
    order.Price -= discount  // This modifies the original!
}

// Comparing approaches
func addWhippedCream(order CoffeeOrder) CoffeeOrder {
    order.Price += 0.50
    return order  // Must return and reassign
}

func addWhippedCreamDirect(order *CoffeeOrder) {
    order.Price += 0.50  // Modifies original directly
}

func main() {
    order := CoffeeOrder{Type: "Mocha", Size: "Medium", Price: 4.50}
    
    // Method 1: Return new value
    order = addWhippedCream(order)
    
    // Method 2: Modify directly
    addWhippedCreamDirect(&order)  // Note the & to pass address
}
```

## Return Values

### Single Returns

```go
// Simple return
func calculateTax(amount float64) float64 {
    return amount * 0.08
}

// Early returns for validation
func validateCoffeeSize(size string) bool {
    if size == "" {
        return false
    }
    
    validSizes := []string{"small", "medium", "large"}
    for _, valid := range validSizes {
        if size == valid {
            return true
        }
    }
    
    return false
}
```

### Multiple Returns

"Go's multiple returns are perfect for coffee operations!" Sarah shows:

```go
// Return value and success indicator
func getCoffeePrice(coffeeType string) (float64, bool) {
    prices := map[string]float64{
        "espresso":   3.00,
        "latte":      4.50,
        "cappuccino": 4.00,
    }
    
    price, exists := prices[coffeeType]
    return price, exists
}

// Return value and error
func prepareDrink(order CoffeeOrder) (string, error) {
    if order.Type == "" {
        return "", fmt.Errorf("coffee type cannot be empty")
    }
    
    if order.Size == "" {
        return "", fmt.Errorf("size must be specified")
    }
    
    drink := fmt.Sprintf("%s %s", order.Size, order.Type)
    return drink, nil
}

// Using multiple returns
func main() {
    // Check both returns
    price, ok := getCoffeePrice("latte")
    if ok {
        fmt.Printf("Latte costs $%.2f\n", price)
    }
    
    // Handle errors
    drink, err := prepareDrink(CoffeeOrder{Type: "Espresso", Size: "Double"})
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    fmt.Printf("Prepared: %s\n", drink)
}
```

### Named Returns

"Named returns make intent clear," Sarah demonstrates:

```go
// Named returns for clarity
func calculateOrderTotal(items []CoffeeOrder) (subtotal, tax, total float64) {
    // Named returns are pre-declared and zero-initialized
    
    for _, item := range items {
        subtotal += item.Price
    }
    
    tax = subtotal * 0.08
    total = subtotal + tax
    
    return  // Naked return - returns all named values
}

// Named returns with early exit
func processPayment(amount, payment float64) (change float64, success bool) {
    if payment < amount {
        // change is already 0.00, success is already false
        return
    }
    
    change = payment - amount
    success = true
    return
}

// Be careful with named returns!
func confusingExample() (result int) {
    result = 10
    // This creates a new local variable, doesn't update the return!
    result := 20  // Shadow variable - avoid this!
    return        // Returns 10, not 20
}
```

## Advanced Parameter Patterns

### Functional Options Pattern

"For complex configurations, we use this pattern," Sarah shows:

```go
// Coffee with many optional parameters
type Coffee struct {
    Type        string
    Size        string
    Temperature float64
    Milk        string
    Shots       int
    Decaf       bool
}

// Option is a function that modifies a Coffee
type Option func(*Coffee)

// Option functions
func WithSize(size string) Option {
    return func(c *Coffee) {
        c.Size = size
    }
}

func WithExtraShots(shots int) Option {
    return func(c *Coffee) {
        c.Shots += shots
    }
}

func WithMilk(milk string) Option {
    return func(c *Coffee) {
        c.Milk = milk
    }
}

func Decaf() Option {
    return func(c *Coffee) {
        c.Decaf = true
    }
}

// Constructor using options
func NewCoffee(coffeeType string, opts ...Option) *Coffee {
    // Default coffee
    c := &Coffee{
        Type:        coffeeType,
        Size:        "medium",
        Temperature: 65.0,
        Milk:        "whole",
        Shots:       1,
        Decaf:       false,
    }
    
    // Apply all options
    for _, opt := range opts {
        opt(c)
    }
    
    return c
}

func main() {
    // Clean, readable coffee creation
    latte := NewCoffee("latte",
        WithSize("large"),
        WithMilk("oat"),
        WithExtraShots(1),
    )
    
    decafCapp := NewCoffee("cappuccino",
        Decaf(),
        WithSize("small"),
    )
}
```

### Parameter Validation

"Always validate your inputs," Sarah emphasizes:

```go
// Validation with detailed errors
type ValidationError struct {
    Field   string
    Message string
}

func (e ValidationError) Error() string {
    return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

func createOrder(customerName string, items []CoffeeOrder) (*Order, error) {
    // Validate customer name
    if customerName == "" {
        return nil, ValidationError{
            Field:   "customerName",
            Message: "cannot be empty",
        }
    }
    
    // Validate items
    if len(items) == 0 {
        return nil, ValidationError{
            Field:   "items",
            Message: "order must contain at least one item",
        }
    }
    
    // Validate each item
    for i, item := range items {
        if item.Type == "" {
            return nil, ValidationError{
                Field:   fmt.Sprintf("items[%d].Type", i),
                Message: "coffee type cannot be empty",
            }
        }
        
        if item.Price < 0 {
            return nil, ValidationError{
                Field:   fmt.Sprintf("items[%d].Price", i),
                Message: "price cannot be negative",
            }
        }
    }
    
    // Create order...
    return &Order{
        Customer: customerName,
        Items:    items,
    }, nil
}
```

## Real-World Example: Complete Order System

```go
package main

import (
    "fmt"
    "time"
)

// Order represents a complete coffee order
type Order struct {
    ID         string
    Customer   string
    Items      []OrderItem
    TotalPrice float64
    CreatedAt  time.Time
    Status     string
}

// OrderItem represents a single item in the order
type OrderItem struct {
    Coffee   Coffee
    Quantity int
    Subtotal float64
}

// ProcessOrderInput contains all order parameters
type ProcessOrderInput struct {
    CustomerName  string
    Items         []OrderItem
    PaymentMethod string
    LoyaltyCard   *string  // Optional
}

// ProcessOrderOutput contains all return values
type ProcessOrderOutput struct {
    Order        *Order
    Receipt      string
    LoyaltyPoints int
}

// processOrder handles the complete order flow
func processOrder(input ProcessOrderInput) (*ProcessOrderOutput, error) {
    // Validate input
    if err := validateOrderInput(input); err != nil {
        return nil, fmt.Errorf("validation failed: %w", err)
    }
    
    // Calculate totals
    subtotal, tax, total := calculateTotals(input.Items)
    
    // Apply loyalty discount if card provided
    if input.LoyaltyCard != nil {
        discount := calculateLoyaltyDiscount(*input.LoyaltyCard, total)
        total -= discount
    }
    
    // Create order
    order := &Order{
        ID:         generateOrderID(),
        Customer:   input.CustomerName,
        Items:      input.Items,
        TotalPrice: total,
        CreatedAt:  time.Now(),
        Status:     "pending",
    }
    
    // Process payment
    if err := processPayment(order, input.PaymentMethod); err != nil {
        return nil, fmt.Errorf("payment failed: %w", err)
    }
    
    // Generate output
    output := &ProcessOrderOutput{
        Order:         order,
        Receipt:       generateReceipt(order, subtotal, tax),
        LoyaltyPoints: calculatePoints(total),
    }
    
    return output, nil
}

// Helper functions with clear parameters and returns
func validateOrderInput(input ProcessOrderInput) error {
    if input.CustomerName == "" {
        return fmt.Errorf("customer name required")
    }
    if len(input.Items) == 0 {
        return fmt.Errorf("order must contain items")
    }
    return nil
}

func calculateTotals(items []OrderItem) (subtotal, tax, total float64) {
    for _, item := range items {
        subtotal += item.Subtotal
    }
    tax = subtotal * 0.08
    total = subtotal + tax
    return
}

func calculateLoyaltyDiscount(cardNumber string, amount float64) float64 {
    // Mock implementation
    return amount * 0.05  // 5% discount
}

func processPayment(order *Order, method string) error {
    // Mock implementation
    fmt.Printf("Processing %s payment for $%.2f\n", method, order.TotalPrice)
    order.Status = "paid"
    return nil
}

func generateOrderID() string {
    return fmt.Sprintf("ORD-%d", time.Now().Unix())
}

func calculatePoints(amount float64) int {
    return int(amount * 10)  // 10 points per dollar
}

func generateReceipt(order *Order, subtotal, tax float64) string {
    return fmt.Sprintf(`
GoCoffee Receipt
================
Order: %s
Customer: %s
Items: %d
Subtotal: $%.2f
Tax: $%.2f
Total: $%.2f
================
Thank you!
`, order.ID, order.Customer, len(order.Items), subtotal, tax, order.TotalPrice)
}
```

## Best Practices

Sarah shares her parameter and return wisdom:

### Parameters
1. **Limit parameter count** - More than 3-4? Use a struct
2. **Order matters** - Required params first, optional last
3. **Be consistent** - Similar functions should have similar signatures
4. **Validate early** - Check parameters at the start

### Returns
1. **Error last** - Always return error as the last value
2. **Be explicit** - Named returns for complex functions
3. **Zero values** - Ensure meaningful zeros on error
4. **Document returns** - Especially for multiple returns

## Common Patterns

```go
// Pattern 1: Optional parameters with defaults
func makeLatte(size string, options ...string) Coffee {
    coffee := Coffee{Type: "latte", Size: size}
    
    // Process options
    for _, opt := range options {
        switch opt {
        case "decaf":
            coffee.Decaf = true
        case "extra-hot":
            coffee.Temperature = 75.0
        }
    }
    
    return coffee
}

// Pattern 2: Builder pattern for complex objects
type CoffeeBuilder struct {
    coffee Coffee
}

func (b *CoffeeBuilder) WithSize(size string) *CoffeeBuilder {
    b.coffee.Size = size
    return b
}

func (b *CoffeeBuilder) WithMilk(milk string) *CoffeeBuilder {
    b.coffee.Milk = milk
    return b
}

func (b *CoffeeBuilder) Build() Coffee {
    return b.coffee
}

// Pattern 3: Result type for complex returns
type Result struct {
    Value interface{}
    Error error
    Warnings []string
}
```

## Try It Yourself!

```go
// Challenge 1: Write a function that takes coffee preferences
// and returns a customized coffee order with price

// Challenge 2: Create a function that processes multiple
// payment methods and returns change and receipt

// Challenge 3: Design a function that validates a complete
// order and returns detailed error information

// Challenge 4: Build a function that calculates discounts
// based on multiple factors (loyalty, time of day, quantity)
```

Continue to [Variadic Functions](../03-variadic-functions/Variadic_Functions.md) →

---

*"Good parameters are like a well-designed coffee machine interface - intuitive inputs, predictable outputs!"*