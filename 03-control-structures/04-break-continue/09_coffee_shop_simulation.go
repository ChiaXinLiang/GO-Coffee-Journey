package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("=== GoCoffee Shop Simulation ===")
	fmt.Println("Real-world break and continue usage\n")
	
	rand.Seed(time.Now().UnixNano())

	// Initialize shop
	type Customer struct {
		name       string
		vip        bool
		orderType  string
		patience   int // seconds willing to wait
		allergies  []string
	}
	
	type ShopState struct {
		hour          int
		isOpen        bool
		coffeeBeans   int    // grams
		milk          int    // ml
		staffEnergy   int    // percentage
		revenue       float64
		servedCount   int
		vipServed     int
		lostCustomers int
	}
	
	shop := ShopState{
		hour:        6,  // Open at 6 AM
		isOpen:      true,
		coffeeBeans: 5000,
		milk:        10000,
		staffEnergy: 100,
		revenue:     0,
	}
	
	// Generate customer queue
	customers := []Customer{
		{"Alice", false, "Latte", 60, nil},
		{"Bob (VIP)", true, "Espresso", 120, nil},
		{"Charlie", false, "Decaf Latte", 45, nil},
		{"Diana", false, "Cappuccino", 90, []string{"lactose"}},
		{"Eve (VIP)", true, "Mocha", 150, nil},
		{"Frank", false, "Americano", 30, nil},
		{"Grace", false, "Latte", 60, nil},
		{"Henry (VIP)", true, "Macchiato", 180, nil},
		{"Iris", false, "Espresso", 45, nil},
		{"Jack", false, "Cappuccino", 75, nil},
	}
	
	// Add more random customers
	names := []string{"Kelly", "Liam", "Maya", "Noah", "Olivia", "Paul", "Quinn", "Ruby", "Sam", "Tara"}
	drinks := []string{"Latte", "Espresso", "Cappuccino", "Americano", "Mocha"}
	
	for _, name := range names {
		isVIP := rand.Float32() < 0.2
		if isVIP {
			name += " (VIP)"
		}
		customers = append(customers, Customer{
			name:      name,
			vip:       isVIP,
			orderType: drinks[rand.Intn(len(drinks))],
			patience:  rand.Intn(120) + 30,
			allergies: nil,
		})
	}
	
	fmt.Println("☕ COFFEE SHOP OPENS")
	fmt.Println("===================")
	fmt.Printf("Starting resources: %dg beans, %dml milk\n", shop.coffeeBeans, shop.milk)
	fmt.Printf("Queue size: %d customers\n\n", len(customers))
	
	// Main service loop
	customerIndex := 0
	
ShopHours:
	for shop.hour <= 20 { // Close at 8 PM
		fmt.Printf("\n⏰ %02d:00 - Staff energy: %d%%\n", shop.hour, shop.staffEnergy)
		fmt.Println("------------------------")
		
		// Check if shop should close early
		if shop.staffEnergy < 10 {
			fmt.Println("❌ Staff too exhausted - closing early!")
			break ShopHours
		}
		
		if shop.coffeeBeans < 50 && shop.milk < 200 {
			fmt.Println("❌ Out of supplies - closing early!")
			break ShopHours
		}
		
		// Process customers for this hour
		customersThisHour := 0
		maxPerHour := 8
		
	HourlyService:
		for customersThisHour < maxPerHour && customerIndex < len(customers) {
			customer := customers[customerIndex]
			customerIndex++
			
			fmt.Printf("\n👤 %s orders %s", customer.name, customer.orderType)
			
			// VIP priority check
			if !customer.vip && shop.hour >= 12 && shop.hour <= 14 {
				// During lunch rush, prioritize VIPs
				vipAhead := false
				for j := customerIndex; j < len(customers) && j < customerIndex+3; j++ {
					if customers[j].vip {
						vipAhead = true
						break
					}
				}
				
				if vipAhead {
					fmt.Print(" - Yielding to VIP customer")
					customerIndex-- // Put customer back
					continue HourlyService
				}
			}
			
			// Check patience
			waitTime := customersThisHour * 15 // 15 seconds per order
			if waitTime > customer.patience {
				fmt.Printf(" - Left! (wait %ds > patience %ds)", waitTime, customer.patience)
				shop.lostCustomers++
				continue HourlyService
			}
			
			// Check for unavailable items
			if customer.orderType == "Decaf Latte" {
				fmt.Print(" - ❌ Decaf not available")
				continue HourlyService
			}
			
			// Check allergies
			if len(customer.allergies) > 0 {
				for _, allergy := range customer.allergies {
					if allergy == "lactose" && 
					   (customer.orderType == "Latte" || 
					    customer.orderType == "Cappuccino" || 
					    customer.orderType == "Mocha") {
						fmt.Printf(" - ⚠️  Contains %s (customer allergic)", allergy)
						continue HourlyService
					}
				}
			}
			
			// Check resources
			beansNeeded := 18
			milkNeeded := 0
			price := 0.0
			
			switch customer.orderType {
			case "Espresso":
				price = 2.50
			case "Americano":
				price = 3.00
			case "Latte":
				milkNeeded = 150
				price = 4.50
			case "Cappuccino":
				milkNeeded = 100
				price = 4.00
			case "Mocha":
				milkNeeded = 120
				beansNeeded = 20
				price = 5.00
			case "Macchiato":
				milkNeeded = 50
				price = 4.25
			}
			
			if shop.coffeeBeans < beansNeeded {
				fmt.Printf(" - ❌ Not enough beans (need %d, have %d)", 
					beansNeeded, shop.coffeeBeans)
				continue HourlyService
			}
			
			if shop.milk < milkNeeded {
				fmt.Printf(" - ❌ Not enough milk (need %d, have %d)", 
					milkNeeded, shop.milk)
				continue HourlyService
			}
			
			// Make the coffee!
			shop.coffeeBeans -= beansNeeded
			shop.milk -= milkNeeded
			shop.revenue += price
			shop.servedCount++
			if customer.vip {
				shop.vipServed++
				shop.revenue += price * 0.1 // VIP bonus
			}
			
			fmt.Printf(" - ✓ Served ($%.2f)", price)
			
			customersThisHour++
			shop.staffEnergy -= 2
			
			// Random events
			if rand.Float32() < 0.1 {
				tip := rand.Float64() * 2 + 0.5
				shop.revenue += tip
				fmt.Printf(" + $%.2f tip!", tip)
			}
			
			time.Sleep(200 * time.Millisecond) // Simulation delay
		}
		
		// Hourly summary
		fmt.Printf("\nHourly summary: Served %d customers\n", customersThisHour)
		
		// Rest period between rush hours
		if shop.hour == 14 {
			fmt.Println("\n☕ AFTERNOON BREAK - Staff resting...")
			shop.staffEnergy += 20
			if shop.staffEnergy > 100 {
				shop.staffEnergy = 100
			}
		}
		
		// Restock check
		if shop.hour == 15 && (shop.coffeeBeans < 1000 || shop.milk < 2000) {
			fmt.Println("\n📦 RESTOCKING SUPPLIES")
			shop.coffeeBeans += 2000
			shop.milk += 5000
			fmt.Printf("New levels: %dg beans, %dml milk\n", shop.coffeeBeans, shop.milk)
		}
		
		shop.hour++
		
		// Check if all customers served
		if customerIndex >= len(customers) {
			fmt.Println("\n✅ All customers in queue have been served!")
			break ShopHours
		}
	}
	
	// End of day summary
	fmt.Println("\n\n🌙 CLOSING TIME - DAILY SUMMARY")
	fmt.Println("================================")
	fmt.Printf("Total customers served: %d\n", shop.servedCount)
	fmt.Printf("VIP customers served: %d\n", shop.vipServed)
	fmt.Printf("Customers lost: %d\n", shop.lostCustomers)
	fmt.Printf("Total revenue: $%.2f\n", shop.revenue)
	fmt.Printf("Average per customer: $%.2f\n", shop.revenue/float64(shop.servedCount))
	fmt.Printf("Final resources: %dg beans, %dml milk\n", shop.coffeeBeans, shop.milk)
	fmt.Printf("Staff energy: %d%%\n", shop.staffEnergy)
	
	// Performance rating
	rating := 5.0
	if shop.lostCustomers > 5 {
		rating -= 1.0
	}
	if shop.servedCount < 15 {
		rating -= 0.5
	}
	if shop.revenue < 100 {
		rating -= 0.5
	}
	if shop.vipServed < 3 {
		rating -= 0.5
	}
	
	fmt.Printf("\n⭐ Daily performance rating: %.1f/5.0\n", rating)
	
	if rating >= 4.5 {
		fmt.Println("🎉 Excellent day!")
	} else if rating >= 3.5 {
		fmt.Println("👍 Good day!")
	} else {
		fmt.Println("📈 Room for improvement")
	}
}