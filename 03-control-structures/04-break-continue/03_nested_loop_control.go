package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== GoCoffee Nested Loop Control ===\n")

	// Example 1: Break in nested loops (only breaks inner)
	fmt.Println("🏢 Finding First Available Table")
	fmt.Println("--------------------------------")
	
	floors := 3
	tablesPerFloor := 5
	
	// Simulated occupancy (true = occupied)
	occupancy := [][]bool{
		{true, true, true, true, true},    // Floor 1: full
		{true, false, true, true, false},   // Floor 2: some available
		{false, false, true, false, true},  // Floor 3: some available
	}
	
	for floor := 0; floor < floors; floor++ {
		fmt.Printf("\nFloor %d: ", floor+1)
		foundOnFloor := false
		
		for table := 0; table < tablesPerFloor; table++ {
			if !occupancy[floor][table] {
				fmt.Printf("Table %d is FREE! ", table+1)
				foundOnFloor = true
				break // Only breaks inner loop
			}
		}
		
		if !foundOnFloor {
			fmt.Print("All tables occupied")
		}
	}

	// Example 2: Labeled break for nested loops
	fmt.Println("\n\n\n🔍 Finding Perfect Table (Window + Available)")
	fmt.Println("--------------------------------------------")
	
	// Window tables are 1 and 5 on each floor
	windowTables := map[int]bool{1: true, 5: true}
	
	foundPerfectTable := false
	var perfectFloor, perfectTable int
	
SearchLoop:
	for floor := 0; floor < floors; floor++ {
		for table := 0; table < tablesPerFloor; table++ {
			tableNum := table + 1
			
			// Check if available AND by window
			if !occupancy[floor][table] && windowTables[tableNum] {
				foundPerfectTable = true
				perfectFloor = floor + 1
				perfectTable = tableNum
				break SearchLoop // Breaks both loops
			}
		}
	}
	
	if foundPerfectTable {
		fmt.Printf("✓ Found perfect table: Table %d on Floor %d (Window seat!)\n", 
			perfectTable, perfectFloor)
	} else {
		fmt.Println("❌ No window tables available")
	}

	// Example 3: Continue in nested loops
	fmt.Println("\n\n📊 Daily Sales Report (Skip Lunch Hours)")
	fmt.Println("---------------------------------------")
	
	days := []string{"Monday", "Tuesday", "Wednesday"}
	hours := []int{9, 10, 11, 12, 13, 14, 15, 16, 17}
	lunchHours := map[int]bool{12: true, 13: true}
	
	for _, day := range days {
		fmt.Printf("\n%s:\n", day)
		dailyTotal := 0.0
		
		for _, hour := range hours {
			// Skip lunch hours
			if lunchHours[hour] {
				fmt.Printf("  %02d:00 - Lunch break\n", hour)
				continue
			}
			
			// Simulate sales
			sales := float64(hour * 10) // Simple calculation
			dailyTotal += sales
			
			fmt.Printf("  %02d:00 - $%.2f\n", hour, sales)
		}
		
		fmt.Printf("  Daily Total: $%.2f\n", dailyTotal)
	}

	// Example 4: Labeled continue
	fmt.Println("\n\n🍰 Inventory Check (Skip Expired Items)")
	fmt.Println("--------------------------------------")
	
	categories := []string{"Pastries", "Sandwiches", "Salads"}
	
	inventory := [][]struct {
		name    string
		expired bool
		count   int
	}{
		// Pastries
		{{"Croissant", false, 12}, {"Muffin", true, 5}, {"Scone", false, 8}},
		// Sandwiches  
		{{"BLT", false, 6}, {"Club", false, 4}, {"Veggie", true, 3}},
		// Salads
		{{"Caesar", true, 2}, {"Greek", false, 5}, {"Garden", false, 7}},
	}
	
CategoryLoop:
	for catIdx, category := range categories {
		fmt.Printf("\n%s:\n", category)
		validItems := 0
		
		for _, item := range inventory[catIdx] {
			if item.expired {
				fmt.Printf("  %-12s: EXPIRED - Skipping\n", item.name)
				continue
			}
			
			// Skip category if too few valid items
			if validItems == 0 && item.count < 5 {
				fmt.Printf("  Low inventory in %s - Skipping category\n", category)
				continue CategoryLoop
			}
			
			validItems++
			fmt.Printf("  %-12s: %d units\n", item.name, item.count)
		}
	}

	// Example 5: Complex control flow
	fmt.Println("\n\n⚙️  Coffee Machine Maintenance Check")
	fmt.Println("----------------------------------")
	
	machines := []string{"Espresso-1", "Espresso-2", "Drip-1"}
	checks := []string{"Water Level", "Bean Level", "Temperature", "Pressure"}
	
	// Simulated check results
	checkResults := map[string]map[string]bool{
		"Espresso-1": {"Water Level": true, "Bean Level": false, "Temperature": true, "Pressure": true},
		"Espresso-2": {"Water Level": true, "Bean Level": true, "Temperature": false, "Pressure": false},
		"Drip-1":     {"Water Level": false, "Bean Level": true, "Temperature": true, "Pressure": true},
	}
	
MachineLoop:
	for _, machine := range machines {
		fmt.Printf("\nChecking %s:\n", machine)
		failedChecks := 0
		
		for _, check := range checks {
			result := checkResults[machine][check]
			
			if !result {
				failedChecks++
				fmt.Printf("  %-15s: ❌ FAILED\n", check)
				
				// Critical failure - stop checking this machine
				if check == "Temperature" || failedChecks >= 2 {
					fmt.Printf("  ⚠️  Critical failure - Machine needs service!\n")
					continue MachineLoop
				}
				continue
			}
			
			fmt.Printf("  %-15s: ✓ OK\n", check)
		}
		
		if failedChecks == 0 {
			fmt.Printf("  ✅ Machine fully operational!\n")
		}
		
		time.Sleep(500 * time.Millisecond)
	}
	
	fmt.Println("\n✓ Maintenance check complete!")
}