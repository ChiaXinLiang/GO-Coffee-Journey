package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("=== GoCoffee Performance Control with Break/Continue ===\n")
	
	rand.Seed(time.Now().UnixNano())

	// Example 1: Early exit optimization
	fmt.Println("⚡ Search Performance Comparison")
	fmt.Println("-------------------------------")
	
	// Generate large customer database
	customerCount := 10000
	customers := make([]struct {
		id      int
		name    string
		vip     bool
		balance float64
	}, customerCount)
	
	for i := 0; i < customerCount; i++ {
		customers[i].id = i + 1000
		customers[i].name = fmt.Sprintf("Customer-%d", i)
		customers[i].vip = rand.Float32() < 0.1 // 10% VIP
		customers[i].balance = rand.Float64() * 1000
	}
	
	// Target customer is in the middle
	targetID := 6000
	
	// Method 1: Without break (full scan)
	fmt.Println("\nMethod 1: Full scan without break")
	start := time.Now()
	var found1 interface{}
	comparisons1 := 0
	
	for _, customer := range customers {
		comparisons1++
		if customer.id == targetID {
			found1 = customer
		}
	}
	time1 := time.Since(start)
	
	// Method 2: With break (early exit)
	fmt.Println("Method 2: Early exit with break")
	start = time.Now()
	var found2 interface{}
	comparisons2 := 0
	
	for _, customer := range customers {
		comparisons2++
		if customer.id == targetID {
			found2 = customer
			break
		}
	}
	time2 := time.Since(start)
	
	fmt.Printf("\nResults:")
	fmt.Printf("\n  Full scan: %v, %d comparisons", time1, comparisons1)
	fmt.Printf("\n  Early exit: %v, %d comparisons", time2, comparisons2)
	fmt.Printf("\n  Performance gain: %.2fx faster\n", float64(comparisons1)/float64(comparisons2))

	// Example 2: Filtering optimization
	fmt.Println("\n\n🎯 VIP Processing Optimization")
	fmt.Println("-----------------------------")
	
	// Process only VIP customers
	fmt.Println("\nMethod 1: Nested if statement")
	start = time.Now()
	vipProcessed1 := 0
	
	for _, customer := range customers[:1000] { // Use subset for demo
		if customer.vip {
			if customer.balance > 500 {
				// Process VIP
				vipProcessed1++
			}
		}
	}
	time1 = time.Since(start)
	
	fmt.Println("Method 2: Continue for early filtering")
	start = time.Now()
	vipProcessed2 := 0
	
	for _, customer := range customers[:1000] {
		if !customer.vip {
			continue // Skip non-VIP immediately
		}
		
		if customer.balance <= 500 {
			continue // Skip low balance
		}
		
		// Process VIP
		vipProcessed2++
	}
	time2 = time.Since(start)
	
	fmt.Printf("\nResults:")
	fmt.Printf("\n  Nested if: %v", time1)
	fmt.Printf("\n  Continue: %v", time2)
	fmt.Printf("\n  Both processed: %d VIP customers\n", vipProcessed1)

	// Example 3: Batch processing with threshold
	fmt.Println("\n\n📦 Batch Processing with Limits")
	fmt.Println("-------------------------------")
	
	orders := make([]struct {
		id     string
		amount float64
		risk   int // 1-10
	}, 5000)
	
	for i := 0; i < len(orders); i++ {
		orders[i].id = fmt.Sprintf("ORD-%05d", i)
		orders[i].amount = rand.Float64() * 200
		orders[i].risk = rand.Intn(10) + 1
	}
	
	// Process orders with budget limit
	budget := 10000.0
	maxRisk := 7
	
	fmt.Printf("\nProcessing orders (Budget: $%.2f, Max Risk: %d)\n", budget, maxRisk)
	
	start = time.Now()
	processedAmount := 0.0
	processedCount := 0
	skippedHighRisk := 0
	
	for _, order := range orders {
		// Skip high-risk orders
		if order.risk > maxRisk {
			skippedHighRisk++
			continue
		}
		
		// Check budget
		if processedAmount+order.amount > budget {
			fmt.Printf("\nBudget exhausted at order %s", order.id)
			break
		}
		
		processedAmount += order.amount
		processedCount++
	}
	
	elapsed := time.Since(start)
	
	fmt.Printf("\n\nResults:")
	fmt.Printf("\n  Time: %v", elapsed)
	fmt.Printf("\n  Processed: %d orders", processedCount)
	fmt.Printf("\n  Amount: $%.2f", processedAmount)
	fmt.Printf("\n  Skipped (high risk): %d orders\n", skippedHighRisk)

	// Example 4: Resource-limited processing
	fmt.Println("\n\n☕ Resource-Limited Coffee Making")
	fmt.Println("--------------------------------")
	
	type CoffeeOrder struct {
		drink      string
		size       string
		waterNeeds int
		beansNeeds int
		milkNeeds  int
	}
	
	coffeeQueue := []CoffeeOrder{
		{"Espresso", "S", 30, 7, 0},
		{"Latte", "L", 200, 14, 150},
		{"Americano", "M", 150, 7, 0},
		{"Cappuccino", "M", 120, 14, 80},
		{"Espresso", "S", 30, 7, 0},
		{"Mocha", "L", 200, 14, 100},
		{"Latte", "M", 150, 14, 100},
		{"Espresso", "S", 30, 7, 0},
	}
	
	// Resources
	water := 1000 // ml
	beans := 100  // g
	milk := 500   // ml
	
	fmt.Printf("Starting resources: Water=%dml, Beans=%dg, Milk=%dml\n", water, beans, milk)
	
	completed := 0
	
	for i, order := range coffeeQueue {
		fmt.Printf("\nOrder %d: %s %s", i+1, order.size, order.drink)
		
		// Check resources
		if water < order.waterNeeds {
			fmt.Printf(" - ❌ Not enough water (need %d, have %d)", order.waterNeeds, water)
			continue
		}
		
		if beans < order.beansNeeds {
			fmt.Printf(" - ❌ Not enough beans (need %d, have %d)", order.beansNeeds, beans)
			continue
		}
		
		if milk < order.milkNeeds {
			fmt.Printf(" - ❌ Not enough milk (need %d, have %d)", order.milkNeeds, milk)
			continue
		}
		
		// Make coffee
		water -= order.waterNeeds
		beans -= order.beansNeeds
		milk -= order.milkNeeds
		completed++
		
		fmt.Printf(" - ✓ Complete")
		fmt.Printf("\n  Remaining: Water=%dml, Beans=%dg, Milk=%dml", water, beans, milk)
		
		// Stop if critically low on beans (most important)
		if beans < 7 {
			fmt.Println("\n\n⚠️  Critical: Almost out of beans - stopping service")
			break
		}
	}
	
	fmt.Printf("\n\nCompleted %d/%d orders\n", completed, len(coffeeQueue))

	// Example 5: Timeout simulation
	fmt.Println("\n\n⏱️  Processing with Timeout")
	fmt.Println("--------------------------")
	
	tasks := make([]struct {
		name     string
		duration time.Duration
		priority int
	}, 20)
	
	taskNames := []string{"Grind", "Brew", "Steam", "Pour", "Serve"}
	
	for i := 0; i < len(tasks); i++ {
		tasks[i].name = taskNames[rand.Intn(len(taskNames))]
		tasks[i].duration = time.Duration(rand.Intn(300)+100) * time.Millisecond
		tasks[i].priority = rand.Intn(3) + 1
	}
	
	timeout := 2 * time.Second
	fmt.Printf("Processing tasks (timeout: %v)\n", timeout)
	
	start = time.Now()
	tasksCompleted := 0
	lowPrioritySkipped := 0
	
	for _, task := range tasks {
		// Skip low priority if running out of time
		elapsed := time.Since(start)
		if elapsed > timeout/2 && task.priority == 1 {
			lowPrioritySkipped++
			continue
		}
		
		// Check timeout
		if elapsed+task.duration > timeout {
			fmt.Printf("\n⏱️  Timeout approaching - stopping at task %d", tasksCompleted+1)
			break
		}
		
		// Process task
		fmt.Printf("\nTask %d: %s (P%d, %v)", 
			tasksCompleted+1, task.name, task.priority, task.duration)
		time.Sleep(task.duration)
		tasksCompleted++
		fmt.Print(" ✓")
	}
	
	finalTime := time.Since(start)
	fmt.Printf("\n\nCompleted %d tasks in %v", tasksCompleted, finalTime)
	fmt.Printf("\nSkipped %d low-priority tasks\n", lowPrioritySkipped)
}