package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("=== GoCoffee Labeled Break and Continue ===\n")

	// Example 1: Basic labeled break
	fmt.Println("🏢 Finding Meeting Room")
	fmt.Println("----------------------")
	
	buildings := []string{"Main", "Annex", "Tower"}
	floorsPerBuilding := 3
	roomsPerFloor := 4
	
	// Simulated availability: [building][floor][room]
	availability := [][][]bool{
		// Main building
		{
			{false, false, false, false}, // Floor 1
			{false, true, false, false},  // Floor 2
			{false, false, false, true},  // Floor 3
		},
		// Annex
		{
			{true, false, false, false},  // Floor 1
			{false, false, false, false}, // Floor 2
			{false, false, true, false},  // Floor 3
		},
		// Tower
		{
			{false, false, false, false}, // Floor 1
			{false, false, false, false}, // Floor 2
			{true, true, false, false},   // Floor 3
		},
	}
	
	var foundBuilding, foundFloor, foundRoom int
	found := false
	
BuildingSearch:
	for b, building := range buildings {
		fmt.Printf("\nSearching %s building:\n", building)
		
		for f := 0; f < floorsPerBuilding; f++ {
			fmt.Printf("  Floor %d: ", f+1)
			
			for r := 0; r < roomsPerFloor; r++ {
				if availability[b][f][r] {
					foundBuilding = b
					foundFloor = f
					foundRoom = r
					found = true
					fmt.Printf("Room %d is AVAILABLE!\n", r+1)
					break BuildingSearch // Exit all three loops
				}
				fmt.Printf("Room %d ❌ ", r+1)
			}
			fmt.Println()
		}
	}
	
	if found {
		fmt.Printf("\n✓ Found room: %s Building, Floor %d, Room %d\n",
			buildings[foundBuilding], foundFloor+1, foundRoom+1)
	} else {
		fmt.Println("\n❌ No rooms available!")
	}

	// Example 2: Labeled continue
	fmt.Println("\n\n📋 Order Processing Matrix")
	fmt.Println("-------------------------")
	
	orderTypes := []string{"Coffee", "Pastry", "Sandwich"}
	
	// Orders: [type][items]
	orders := [][]string{
		{"Latte", "Espresso", "Decaf", "Cappuccino"},
		{"Croissant", "Muffin", "Expired-Donut", "Scone"},
		{"BLT", "Out-of-Stock", "Club", "Veggie"},
	}
	
	processedCount := 0
	
OrderType:
	for typeIdx, orderType := range orderTypes {
		fmt.Printf("\n%s Orders:\n", orderType)
		
		validInCategory := false
		
		for itemIdx, item := range orders[typeIdx] {
			fmt.Printf("  Item %d: %s - ", itemIdx+1, item)
			
			// Skip problematic items
			if strings.Contains(item, "Decaf") {
				fmt.Println("Skipped (machine broken)")
				continue
			}
			
			if strings.Contains(item, "Expired") {
				fmt.Println("ALERT: Expired item found!")
				fmt.Printf("  ⚠️  Skipping entire %s category for safety check\n", orderType)
				continue OrderType // Skip to next category
			}
			
			if strings.Contains(item, "Out-of-Stock") {
				fmt.Println("Skipped (not available)")
				continue
			}
			
			// Process valid item
			processedCount++
			validInCategory = true
			fmt.Println("✓ Processed")
		}
		
		if !validInCategory {
			fmt.Printf("  ❌ No valid items in %s category\n", orderType)
		}
	}
	
	fmt.Printf("\nTotal items processed: %d\n", processedCount)

	// Example 3: Multiple labels
	fmt.Println("\n\n🔍 Complex Search Pattern")
	fmt.Println("------------------------")
	
	// Searching for specific coffee blend combination
	regions := []string{"Africa", "Asia", "Americas"}
	countries := [][]string{
		{"Ethiopia", "Kenya", "Rwanda"},
		{"Indonesia", "Vietnam", "India"},
		{"Colombia", "Brazil", "Peru"},
	}
	roastLevels := []string{"Light", "Medium", "Dark"}
	
	targetRegion := "Americas"
	targetCountry := "Colombia"
	targetRoast := "Medium"
	
	fmt.Printf("Looking for: %s roast from %s, %s\n\n", targetRoast, targetCountry, targetRegion)
	
	foundBlend := false
	
RegionLoop:
	for r, region := range regions {
		fmt.Printf("Region: %s\n", region)
		
		if region != targetRegion {
			fmt.Println("  Skipping region...")
			continue RegionLoop
		}
		
	CountryLoop:
		for c, country := range countries[r] {
			fmt.Printf("  Country: %s\n", country)
			
			if country != targetCountry {
				continue CountryLoop
			}
			
			for _, roast := range roastLevels {
				fmt.Printf("    Roast: %s - ", roast)
				
				if roast == targetRoast {
					foundBlend = true
					fmt.Println("✓ FOUND!")
					break RegionLoop
				}
				fmt.Println("not a match")
			}
		}
	}
	
	if foundBlend {
		fmt.Printf("\n✓ Found the perfect blend: %s roast from %s, %s\n", 
			targetRoast, targetCountry, targetRegion)
	}

	// Example 4: Error handling with labeled break
	fmt.Println("\n\n⚙️ Batch Processing with Error Recovery")
	fmt.Println("--------------------------------------")
	
	batches := [][]string{
		{"Process A", "Process B", "Process C"},
		{"Process D", "ERROR", "Process F"},
		{"Process G", "Process H", "Process I"},
	}
	
	maxRetries := 2
	
BatchProcessing:
	for batchNum, batch := range batches {
		fmt.Printf("\nBatch %d:\n", batchNum+1)
		
		for retry := 0; retry <= maxRetries; retry++ {
			if retry > 0 {
				fmt.Printf("  Retry attempt %d:\n", retry)
			}
			
			errorFound := false
			
			for processNum, process := range batch {
				fmt.Printf("    %s: ", process)
				
				if process == "ERROR" {
					fmt.Println("❌ Failed!")
					errorFound = true
					
					if retry == maxRetries {
						fmt.Println("  ⛔ Max retries reached - abandoning batch processing")
						break BatchProcessing
					}
					
					fmt.Println("  ⚠️  Error detected - retrying batch...")
					break // Break inner loop to retry
				}
				
				fmt.Println("✓")
			}
			
			if !errorFound {
				fmt.Printf("  ✅ Batch %d completed successfully!\n", batchNum+1)
				break // Success - move to next batch
			}
		}
	}

	// Example 5: Performance optimization with labels
	fmt.Println("\n\n⚡ Optimized Inventory Search")
	fmt.Println("----------------------------")
	
	// Large inventory structure
	warehouses := []string{"North", "South", "East", "West"}
	sections := []string{"A", "B", "C", "D", "E"}
	shelves := 10
	
	// Item to find
	targetItem := "Premium Colombian Beans"
	itemLocation := map[string]string{
		"Premium Colombian Beans": "South-C-7",
	}
	
	comparisons := 0
	found = false
	
	// Parse target location
	targetLoc := itemLocation[targetItem]
	parts := strings.Split(targetLoc, "-")
	targetWarehouse := parts[0]
	targetSection := parts[1]
	targetShelf := parts[2]
	
	fmt.Printf("Searching for: %s\n", targetItem)
	fmt.Printf("Expected location: %s\n\n", targetLoc)
	
WarehouseSearch:
	for _, warehouse := range warehouses {
		comparisons++
		
		if warehouse != targetWarehouse {
			fmt.Printf("Skipping %s warehouse...\n", warehouse)
			continue WarehouseSearch
		}
		
		fmt.Printf("Entering %s warehouse:\n", warehouse)
		
		for _, section := range sections {
			comparisons++
			
			if section != targetSection {
				continue // Don't need label here
			}
			
			fmt.Printf("  Section %s:\n", section)
			
			for shelf := 1; shelf <= shelves; shelf++ {
				comparisons++
				
				if fmt.Sprintf("%d", shelf) == targetShelf {
					found = true
					fmt.Printf("    Shelf %d: ✓ FOUND %s!\n", shelf, targetItem)
					break WarehouseSearch
				}
			}
		}
	}
	
	fmt.Printf("\nSearch complete: %d comparisons made\n", comparisons)
	fmt.Println("(Optimized search with early exit using labels)")
}