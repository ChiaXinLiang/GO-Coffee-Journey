package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	fmt.Println("=== GoCoffee Named Return Values ===\n")
	
	// Example 1: Basic named returns
	fmt.Println("Example 1: Basic Named Returns")
	fmt.Println("------------------------------")
	demonstrateBasicNamedReturns()
	
	// Example 2: Named returns with initialization
	fmt.Println("\n\nExample 2: Named Returns with Initialization")
	fmt.Println("--------------------------------------------")
	demonstrateInitializedReturns()
	
	// Example 3: Naked returns
	fmt.Println("\n\nExample 3: Naked Returns")
	fmt.Println("------------------------")
	demonstrateNakedReturns()
	
	// Example 4: Named returns best practices
	fmt.Println("\n\nExample 4: Named Returns Best Practices")
	fmt.Println("---------------------------------------")
	demonstrateBestPractices()
}

// Example 1: Basic named returns
func demonstrateBasicNamedReturns() {
	// Simple named returns
	width, height := getCoffeeCupDimensions("medium")
	fmt.Printf("Medium cup: %.1fcm × %.1fcm\n", width, height)
	
	// Multiple named returns
	sub, tax, total := calculateBillWithTax(25.50)
	fmt.Printf("Bill: Subtotal=$%.2f, Tax=$%.2f, Total=$%.2f\n", 
		sub, tax, total)
	
	// Named returns as documentation
	name, role, since := getEmployeeInfo("EMP-001")
	fmt.Printf("Employee: %s, %s (since %s)\n", name, role, since)
}

// Named returns make the function signature self-documenting
func getCoffeeCupDimensions(size string) (width, height float64) {
	switch size {
	case "small":
		width = 5.0
		height = 8.0
	case "medium":
		width = 6.0
		height = 10.0
	case "large":
		width = 7.0
		height = 12.0
	default:
		width = 6.0
		height = 10.0
	}
	return // returns width and height
}

func calculateBillWithTax(amount float64) (subtotal, tax, total float64) {
	subtotal = amount
	tax = subtotal * 0.08
	total = subtotal + tax
	return // returns all three values
}

func getEmployeeInfo(id string) (name, position string, startDate time.Time) {
	// Simulated employee lookup
	employees := map[string]struct {
		name     string
		position string
		started  time.Time
	}{
		"EMP-001": {
			name:     "Sarah Johnson",
			position: "Head Barista",
			started:  time.Date(2020, 1, 15, 0, 0, 0, 0, time.UTC),
		},
	}
	
	if emp, exists := employees[id]; exists {
		name = emp.name
		position = emp.position
		startDate = emp.started
	}
	
	return // returns all named values
}

// Example 2: Named returns with initialization
func demonstrateInitializedReturns() {
	// Default values
	coffee := getDefaultCoffee("")
	fmt.Printf("Empty input coffee: %+v\n", coffee)
	
	coffee = getDefaultCoffee("Latte")
	fmt.Printf("Latte details: %+v\n", coffee)
	
	// Error handling with named returns
	result, err := processOrderWithDefaults("", 0)
	fmt.Printf("Invalid order: result='%s', error=%v\n", result, err)
	
	result, err = processOrderWithDefaults("Espresso", 2)
	fmt.Printf("Valid order: result='%s', error=%v\n", result, err)
}

type Coffee struct {
	Type        string
	Size        string
	Temperature int
}

// Named returns are initialized to zero values
func getDefaultCoffee(coffeeType string) (coffee Coffee) {
	// coffee is already initialized to zero value Coffee{}
	
	if coffeeType == "" {
		// Return zero value Coffee
		return
	}
	
	coffee.Type = coffeeType
	coffee.Size = "Medium"      // Default size
	coffee.Temperature = 165    // Default temp
	
	return // returns the modified coffee
}

func processOrderWithDefaults(item string, quantity int) (result string, err error) {
	// result is "" and err is nil by default
	
	if item == "" {
		err = fmt.Errorf("item cannot be empty")
		return // returns "" and error
	}
	
	if quantity <= 0 {
		err = fmt.Errorf("quantity must be positive")
		return
	}
	
	result = fmt.Sprintf("Order processed: %d × %s", quantity, item)
	return // returns result and nil
}

// Example 3: Naked returns
func demonstrateNakedReturns() {
	// Naked return in simple function
	sum, count, avg := calculateAverage(10, 20, 30, 40, 50)
	fmt.Printf("Stats: sum=%.0f, count=%d, average=%.1f\n", sum, count, avg)
	
	// Naked return with conditions
	status := checkStoreStatus(14) // 2 PM
	fmt.Printf("Store status at 2 PM: %s\n", status)
	
	// Naked return in error handling
	conn, err := connectToDatabase("coffee_shop")
	if err != nil {
		fmt.Printf("Connection error: %v\n", err)
	} else {
		fmt.Printf("Connected: %v\n", conn)
	}
}

func calculateAverage(values ...float64) (sum float64, count int, average float64) {
	count = len(values)
	
	if count == 0 {
		return // returns 0, 0, 0
	}
	
	for _, v := range values {
		sum += v
	}
	
	average = sum / float64(count)
	return // naked return of all named values
}

func checkStoreStatus(hour int) (status string) {
	switch {
	case hour < 6:
		status = "Closed - Too early"
		return
	case hour >= 6 && hour < 10:
		status = "Open - Morning rush"
		return
	case hour >= 10 && hour < 14:
		status = "Open - Normal hours"
		return
	case hour >= 14 && hour < 17:
		status = "Open - Afternoon"
		return
	case hour >= 17 && hour <= 22:
		status = "Open - Evening"
		return
	default:
		status = "Closed - After hours"
		return
	}
}

type DBConnection struct {
	Host     string
	Database string
	Connected bool
}

func connectToDatabase(dbName string) (conn *DBConnection, err error) {
	// Validate input
	if dbName == "" {
		err = fmt.Errorf("database name required")
		return // returns nil, error
	}
	
	// Simulate connection
	conn = &DBConnection{
		Host:     "localhost",
		Database: dbName,
	}
	
	// Simulate connection failure
	if dbName == "unavailable" {
		err = fmt.Errorf("cannot connect to %s", dbName)
		conn = nil
		return
	}
	
	conn.Connected = true
	return // returns conn, nil
}

// Example 4: Best practices
func demonstrateBestPractices() {
	// Good: Named returns for documentation
	lat, lon, elev := getShopCoordinates()
	fmt.Printf("Shop location: %.4f°N, %.4f°W, %.0fm elevation\n", lat, lon, elev)
	
	// Good: Named returns for complex calculations
	stats := calculateDailyStats()
	fmt.Printf("Daily stats: %+v\n", stats)
	
	// When NOT to use named returns
	demonstrateWhenNotToUse()
}

// Good: Named returns clarify what each value represents
func getShopCoordinates() (latitude, longitude, elevation float64) {
	latitude = 37.7749
	longitude = -122.4194
	elevation = 16.0
	return
}

type DailyStats struct {
	Revenue    float64
	Orders     int
	AvgOrder   float64
	PeakHour   int
}

// Good: Named returns for complex functions with multiple steps
func calculateDailyStats() (stats DailyStats) {
	// Initialize
	stats.Orders = 0
	stats.Revenue = 0
	
	// Simulate processing orders
	orders := []float64{12.50, 8.75, 15.00, 9.25, 22.00}
	
	for _, amount := range orders {
		stats.Orders++
		stats.Revenue += amount
	}
	
	if stats.Orders > 0 {
		stats.AvgOrder = stats.Revenue / float64(stats.Orders)
	}
	
	stats.PeakHour = 14 // 2 PM
	
	return // Clear what we're returning
}

func demonstrateWhenNotToUse() {
	fmt.Println("\n--- When NOT to use named returns ---")
	
	// Bad: Overuse in simple functions
	badExample := func() (result string) {
		result = "Don't do this for simple returns"
		return
	}
	
	// Good: Direct return for simple functions  
	goodExample := func() string {
		return "This is clearer for simple cases"
	}
	
	fmt.Println("Bad:", badExample())
	fmt.Println("Good:", goodExample())
	
	// Bad: Confusing naked returns
	confusing := func(x int) (y int) {
		if x > 0 {
			y = x * 2
			return // What are we returning?
		}
		return // Easy to miss this returns 0
	}
	
	// Good: Explicit returns
	clear := func(x int) int {
		if x > 0 {
			return x * 2
		}
		return 0 // Clear we return 0
	}
	
	fmt.Printf("Confusing result: %d\n", confusing(5))
	fmt.Printf("Clear result: %d\n", clear(5))
}

// Guidelines for named returns
func namedReturnGuidelines() {
	// 1. Use for documentation when return values aren't obvious
	func getDimensions() (width, height, depth float64) {
		return 10, 20, 30
	}
	
	// 2. Use in functions with multiple return points
	func findItem(id string) (item string, found bool) {
		if id == "" {
			return // returns "", false
		}
		
		// ... search logic ...
		
		if id == "123" {
			item = "Coffee"
			found = true
			return
		}
		
		return // returns "", false
	}
	
	// 3. Avoid in short functions where it adds no value
	func double(x int) int { // Better than (result int)
		return x * 2
	}
	
	// 4. Be careful with naked returns in long functions
	func longFunction() (result string, err error) {
		// ... lots of code ...
		
		// Naked return here is hard to track
		return // What values are we returning?
	}
}