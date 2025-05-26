package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("=== GoCoffee Continue Statement Examples ===\n")

	// Example 1: Skip processing certain items
	fmt.Println("☕ Processing Orders (Skip Decaf)")
	fmt.Println("---------------------------------")
	
	orders := []string{"Latte", "Decaf Americano", "Espresso", "Decaf Latte", "Cappuccino", "Mocha"}
	
	for i, order := range orders {
		if strings.Contains(order, "Decaf") {
			fmt.Printf("Order %d: %s - Skipping (No decaf today!)\n", i+1, order)
			continue
		}
		
		fmt.Printf("Order %d: %s - ✓ Processing...\n", i+1, order)
		// Process the order
	}

	// Example 2: Skip invalid data
	fmt.Println("\n\n💰 Daily Sales Calculation")
	fmt.Println("--------------------------")
	
	// Sales data with some invalid entries (-1 means invalid)
	salesData := []float64{125.50, 89.25, -1, 156.75, -1, 203.00, 178.50, -1, 95.25}
	
	totalSales := 0.0
	validDays := 0
	
	for day, amount := range salesData {
		if amount < 0 {
			fmt.Printf("Day %d: Invalid data - skipping\n", day+1)
			continue
		}
		
		fmt.Printf("Day %d: $%.2f\n", day+1, amount)
		totalSales += amount
		validDays++
	}
	
	fmt.Printf("\nTotal sales: $%.2f (from %d valid days)\n", totalSales, validDays)
	fmt.Printf("Average daily sales: $%.2f\n", totalSales/float64(validDays))

	// Example 3: Skip even numbers
	fmt.Println("\n\n🎫 Lucky Customer Rewards")
	fmt.Println("-------------------------")
	fmt.Println("Every odd-numbered customer gets a free cookie!")
	
	for customer := 1; customer <= 10; customer++ {
		if customer%2 == 0 {
			fmt.Printf("Customer #%d: Regular service\n", customer)
			continue
		}
		
		fmt.Printf("Customer #%d: 🍪 FREE COOKIE! Lucky customer!\n", customer)
	}

	// Example 4: Quality control with continue
	fmt.Println("\n\n🔍 Coffee Bean Quality Check")
	fmt.Println("----------------------------")
	
	type Bean struct {
		batch   string
		quality int // 1-10 scale
		origin  string
	}
	
	beans := []Bean{
		{"B001", 8, "Colombia"},
		{"B002", 5, "Brazil"},      // Too low quality
		{"B003", 9, "Ethiopia"},
		{"B004", 4, "Vietnam"},      // Too low quality
		{"B005", 7, "Costa Rica"},
		{"B006", 10, "Jamaica"},
		{"B007", 6, "Peru"},         // Borderline
	}
	
	minQuality := 7
	approvedBatches := []string{}
	
	for _, bean := range beans {
		fmt.Printf("Checking batch %s from %s (Quality: %d)... ", 
			bean.batch, bean.origin, bean.quality)
		
		if bean.quality < minQuality {
			fmt.Println("❌ Rejected")
			continue
		}
		
		fmt.Println("✓ Approved")
		approvedBatches = append(approvedBatches, bean.batch)
	}
	
	fmt.Printf("\nApproved batches: %v\n", approvedBatches)

	// Example 5: Skip certain characters
	fmt.Println("\n\n📝 Generating Order Codes")
	fmt.Println("-------------------------")
	fmt.Println("(Skipping vowels for unique codes)")
	
	baseCode := "COFFEESHOP"
	orderCode := ""
	
	for _, char := range baseCode {
		// Skip vowels
		if char == 'A' || char == 'E' || char == 'I' || char == 'O' || char == 'U' {
			continue
		}
		orderCode += string(char)
	}
	
	fmt.Printf("Base: %s\n", baseCode)
	fmt.Printf("Code: %s\n", orderCode)

	// Example 6: Nested loops with continue
	fmt.Println("\n\n📊 Sales Report (Skip Weekends)")
	fmt.Println("--------------------------------")
	
	weeks := 2
	daysPerWeek := 7
	dayNames := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	
	totalWeekdaySales := 0.0
	
	for week := 1; week <= weeks; week++ {
		fmt.Printf("\nWeek %d:\n", week)
		
		for day := 0; day < daysPerWeek; day++ {
			// Skip weekends (Saturday and Sunday)
			if day >= 5 {
				fmt.Printf("  %s: Weekend - Shop closed\n", dayNames[day])
				continue
			}
			
			// Simulate daily sales
			dailySales := float64((week * 100) + (day * 50) + 100)
			totalWeekdaySales += dailySales
			
			fmt.Printf("  %s: $%.2f\n", dayNames[day], dailySales)
		}
	}
	
	fmt.Printf("\nTotal weekday sales: $%.2f\n", totalWeekdaySales)

	// Example 7: Continue with labeled loops
	fmt.Println("\n\n🏢 Finding Available Meeting Rooms")
	fmt.Println("----------------------------------")
	
	// Building floors and rooms
	type Room struct {
		number    int
		capacity  int
		available bool
	}
	
	building := [][]Room{
		{{101, 4, false}, {102, 6, true}, {103, 8, false}},
		{{201, 4, true}, {202, 6, false}, {203, 8, true}},
		{{301, 4, true}, {302, 6, true}, {303, 8, false}},
	}
	
	minCapacity := 6
	
	fmt.Printf("Looking for available rooms with capacity >= %d:\n\n", minCapacity)
	
	for floor, rooms := range building {
		fmt.Printf("Floor %d:\n", floor+1)
		
		for _, room := range rooms {
			if !room.available {
				fmt.Printf("  Room %d: Not available\n", room.number)
				continue
			}
			
			if room.capacity < minCapacity {
				fmt.Printf("  Room %d: Available but too small (capacity: %d)\n", 
					room.number, room.capacity)
				continue
			}
			
			fmt.Printf("  Room %d: ✓ PERFECT! (capacity: %d)\n", 
				room.number, room.capacity)
		}
	}
}