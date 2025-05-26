package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== GoCoffee Daily Operations ===\n")

	// Get current day
	today := time.Now().Weekday()
	fmt.Printf("Today is %s\n\n", today)
	
	// Show operations for each day
	for day := time.Sunday; day <= time.Saturday; day++ {
		showDayOperations(day)
		fmt.Println()
	}
	
	// Current day special operations
	fmt.Println("=== TODAY'S SPECIAL OPERATIONS ===")
	showDayOperations(today)
}

func showDayOperations(day time.Weekday) {
	fmt.Printf("📅 %s Operations\n", day)
	fmt.Println(strings.Repeat("-", 35))
	
	switch day {
	case time.Monday:
		fmt.Println("🌟 Monday - New Week Start")
		fmt.Println("• Open: 6:00 AM")
		fmt.Println("• Special: Motivation Monday - 20% off all coffee")
		fmt.Println("• Staff: Full team (recovery from weekend)")
		fmt.Println("• Focus: Inventory check and restock")
		fmt.Println("• Music: Upbeat jazz")
		fmt.Println("• Expected rush: 7-9 AM (heavy)")
		
	case time.Tuesday:
		fmt.Println("🌮 Tuesday - Steady Day")
		fmt.Println("• Open: 6:00 AM")
		fmt.Println("• Special: Taco Tuesday - Mexican hot chocolate")
		fmt.Println("• Staff: Regular staffing")
		fmt.Println("• Focus: Deep cleaning schedule")
		fmt.Println("• Music: Acoustic covers")
		fmt.Println("• Expected rush: 7-9 AM (moderate)")
		
	case time.Wednesday:
		fmt.Println("🐪 Wednesday - Hump Day")
		fmt.Println("• Open: 6:00 AM")
		fmt.Println("• Special: Double shot Wednesday")
		fmt.Println("• Staff: Regular + 1 trainee")
		fmt.Println("• Focus: Staff training")
		fmt.Println("• Music: Classic rock")
		fmt.Println("• Expected rush: 7-9 AM, 2-3 PM (moderate)")
		
	case time.Thursday:
		fmt.Println("🎭 Thursday - Pre-Friday")
		fmt.Println("• Open: 6:00 AM")
		fmt.Println("• Special: Throwback Thursday - Classic drinks")
		fmt.Println("• Staff: Regular staffing")
		fmt.Println("• Focus: Weekend preparation")
		fmt.Println("• Music: 80s and 90s hits")
		fmt.Println("• Expected rush: 7-9 AM (moderate-heavy)")
		
	case time.Friday:
		fmt.Println("🎉 Friday - TGIF")
		fmt.Println("• Open: 6:00 AM")
		fmt.Println("• Special: Frappé Friday - All cold drinks")
		fmt.Println("• Staff: Full team + extra")
		fmt.Println("• Focus: Quick service")
		fmt.Println("• Music: Feel-good Friday mix")
		fmt.Println("• Expected rush: 7-10 AM (very heavy)")
		fmt.Println("• Note: Extended morning rush!")
		
	case time.Saturday:
		fmt.Println("🌈 Saturday - Weekend Vibes")
		fmt.Println("• Open: 7:00 AM (late start)")
		fmt.Println("• Special: Weekend brunch menu")
		fmt.Println("• Staff: Weekend crew")
		fmt.Println("• Focus: Customer experience")
		fmt.Println("• Music: Chill weekend playlist")
		fmt.Println("• Expected rush: 9-11 AM (brunch rush)")
		fmt.Println("• Events: Live music 2-4 PM")
		
	case time.Sunday:
		fmt.Println("☀️ Sunday - Relaxed Day")
		fmt.Println("• Open: 8:00 AM (latest start)")
		fmt.Println("• Special: Sunday Special - Cinnamon rolls")
		fmt.Println("• Staff: Minimum weekend crew")
		fmt.Println("• Focus: Relaxed atmosphere")
		fmt.Println("• Music: Sunday morning jazz")
		fmt.Println("• Expected rush: 10-12 PM (light)")
		fmt.Println("• Note: Close early at 6 PM")
		
	default:
		fmt.Println("❓ Unknown day")
		fmt.Println("• Standard operations apply")
	}
	
	// Additional day-specific alerts
	switch day {
	case time.Monday, time.Friday:
		fmt.Println("\n⚠️  HIGH TRAFFIC DAY - Extra staff needed!")
	case time.Sunday:
		fmt.Println("\n😌 RELAXED DAY - Minimum staffing okay")
	}
}

// Helper function
var strings = struct {
	Repeat func(string, int) string
}{
	Repeat: func(s string, n int) string {
		result := ""
		for i := 0; i < n; i++ {
			result += s
		}
		return result
	},
}