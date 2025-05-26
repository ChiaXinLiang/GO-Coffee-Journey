package main

import (
	"fmt"
	"strings"
)

// Environment constants
const (
	EnvDevelopment = "development"
	EnvStaging     = "staging"
	EnvProduction  = "production"
)

// HTTP status codes we use
const (
	StatusOK                  = 200
	StatusCreated             = 201
	StatusBadRequest          = 400
	StatusUnauthorized        = 401
	StatusNotFound            = 404
	StatusInternalServerError = 500
)

// Database table names
const (
	TableUsers     = "users"
	TableOrders    = "orders"
	TableProducts  = "products"
	TableInventory = "inventory"
)

// Regular expressions
const (
	EmailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	PhoneRegex = `^\+?1?\d{10,14}$`
	ZipRegex   = `^\d{5}(-\d{4})?$`
)

// Configuration keys
const (
	ConfigDBHost   = "DB_HOST"
	ConfigDBPort   = "DB_PORT"
	ConfigDBName   = "DB_NAME"
	ConfigAPIKey   = "API_KEY"
	ConfigLogLevel = "LOG_LEVEL"
)

// Create enum-like behavior
type OrderStatus int

const (
	StatusPending OrderStatus = iota + 1 // Start from 1
	StatusConfirmed
	StatusPreparing
	StatusReady
	StatusDelivered
	StatusCancelled
)

// String representation for OrderStatus
func (s OrderStatus) String() string {
	statuses := []string{
		"Unknown",
		"Pending",
		"Confirmed",
		"Preparing",
		"Ready",
		"Delivered",
		"Cancelled",
	}
	if s < 1 || int(s) >= len(statuses) {
		return statuses[0]
	}
	return statuses[s]
}

func main() {
	fmt.Println("=== GoCoffee Constant Patterns ===\n")

	// Environment detection
	currentEnv := EnvDevelopment
	fmt.Printf("Current environment: %s\n", currentEnv)

	switch currentEnv {
	case EnvDevelopment:
		fmt.Println("- Debug logging enabled")
		fmt.Println("- Using local database")
	case EnvProduction:
		fmt.Println("- Optimized performance")
		fmt.Println("- Using production database")
	}

	// HTTP responses
	fmt.Println("\nHTTP RESPONSE SIMULATION:")
	simulateResponse("/api/orders", StatusOK)
	simulateResponse("/api/orders/999", StatusNotFound)

	// Database queries
	fmt.Println("\nDATABASE QUERIES:")
	fmt.Printf("SELECT * FROM %s WHERE status = 'active'\n", TableUsers)
	fmt.Printf("SELECT * FROM %s WHERE created_at > NOW() - INTERVAL '1 day'\n", TableOrders)

	// Validation examples
	fmt.Println("\nVALIDATION EXAMPLES:")
	testEmail := "marcus@gocoffee.com"
	testPhone := "+1234567890"
	testZip := "98101"

	fmt.Printf("Email '%s' valid: %v\n", testEmail, strings.Contains(testEmail, "@"))
	fmt.Printf("Phone '%s' valid: %v\n", testPhone, len(testPhone) >= 10)
	fmt.Printf("ZIP '%s' valid: %v\n", testZip, len(testZip) == 5)

	// Order status enum
	fmt.Println("\nORDER STATUS ENUM:")
	order := struct {
		ID     int
		Status OrderStatus
	}{
		ID:     1001,
		Status: StatusPreparing,
	}

	fmt.Printf("Order #%d: %s\n", order.ID, order.Status)

	// Progress through statuses
	fmt.Println("\nOrder progression:")
	statuses := []OrderStatus{
		StatusPending,
		StatusConfirmed,
		StatusPreparing,
		StatusReady,
		StatusDelivered,
	}

	for _, status := range statuses {
		fmt.Printf("  → %s\n", status)
	}
}

func simulateResponse(endpoint string, status int) {
	fmt.Printf("GET %s → %d", endpoint, status)
	switch status {
	case StatusOK:
		fmt.Println(" (Success)")
	case StatusNotFound:
		fmt.Println(" (Not Found)")
	case StatusInternalServerError:
		fmt.Println(" (Server Error)")
	default:
		fmt.Println()
	}
}