package main

import (
    "fmt"
    "strings"
)

// ANSI color codes
const (
    Reset  = "\033[0m"
    Red    = "\033[31m"
    Green  = "\033[32m"
    Yellow = "\033[33m"
    Blue   = "\033[34m"
    Purple = "\033[35m"
    Cyan   = "\033[36m"
    White  = "\033[37m"
    
    BgRed    = "\033[41m"
    BgGreen  = "\033[42m"
    BgYellow = "\033[43m"
    BgBlue   = "\033[44m"
    
    Bold      = "\033[1m"
    Underline = "\033[4m"
    Blink     = "\033[5m"
)

func main() {
    fmt.Println("=== GoCoffee Colored Output ===\n")
    
    // Basic colors
    fmt.Println("BASIC COLORS:")
    fmt.Printf("%sRed text%s\n", Red, Reset)
    fmt.Printf("%sGreen text%s\n", Green, Reset)
    fmt.Printf("%sYellow text%s\n", Yellow, Reset)
    fmt.Printf("%sBlue text%s\n", Blue, Reset)
    fmt.Printf("%sPurple text%s\n", Purple, Reset)
    fmt.Printf("%sCyan text%s\n", Cyan, Reset)
    
    // Styled text
    fmt.Println("\nSTYLED TEXT:")
    fmt.Printf("%sBold text%s\n", Bold, Reset)
    fmt.Printf("%sUnderlined text%s\n", Underline, Reset)
    fmt.Printf("%s%sBold and Red%s\n", Bold+Red, Reset)
    
    // Status messages
    fmt.Println("\nSTATUS MESSAGES:")
    printSuccess("✓ Order completed successfully!")
    printError("✗ Payment failed - please try again")
    printWarning("⚠ Low inventory on Espresso beans")
    printInfo("ℹ Store closes at 10 PM today")
    
    // Colored menu
    fmt.Println("\nCOLORED MENU:")
    printColoredMenu()
    
    // Progress bar
    fmt.Println("\nPROGRESS BAR:")
    for i := 0; i <= 100; i += 10 {
        printProgress("Brewing coffee", i, 100)
        // Simulate work
        for j := 0; j < 50000000; j++ {}
    }
    fmt.Println()
    
    // Colored receipt
    fmt.Println("\nCOLORED RECEIPT:")
    printColoredReceipt()
}

func printSuccess(msg string) {
    fmt.Printf("%s%s%s\n", Green, msg, Reset)
}

func printError(msg string) {
    fmt.Printf("%s%s%s\n", Red, msg, Reset)
}

func printWarning(msg string) {
    fmt.Printf("%s%s%s\n", Yellow, msg, Reset)
}

func printInfo(msg string) {
    fmt.Printf("%s%s%s\n", Cyan, msg, Reset)
}

func printColoredMenu() {
    fmt.Printf("\n%s%s╔════════════════════════════════╗%s\n", Bold, Blue, Reset)
    fmt.Printf("%s%s║      ☕ GOCOFFEE MENU ☕       ║%s\n", Bold, Blue, Reset)
    fmt.Printf("%s%s╚════════════════════════════════╝%s\n", Bold, Blue, Reset)
    
    // Hot drinks
    fmt.Printf("\n%s%sHOT DRINKS:%s\n", Bold, Red, Reset)
    fmt.Printf("  %sEspresso%s .............. $3.00\n", Yellow, Reset)
    fmt.Printf("  %sLatte%s ................. $4.50\n", Yellow, Reset)
    fmt.Printf("  %sCappuccino%s ............ $4.00\n", Yellow, Reset)
    
    // Cold drinks
    fmt.Printf("\n%s%sCOLD DRINKS:%s\n", Bold, Cyan, Reset)
    fmt.Printf("  %sIced Coffee%s ........... $3.50\n", Cyan, Reset)
    fmt.Printf("  %sCold Brew%s ............. $4.00\n", Cyan, Reset)
    
    // Special
    fmt.Printf("\n%s%s⭐ TODAY'S SPECIAL:%s\n", Bold+Blink, Green, Reset)
    fmt.Printf("  %s%sCaramel Macchiato%s ..... %s$3.99%s %s(Save $1!)%s\n", 
        Bold, Yellow, Reset, Green, Reset, Red, Reset)
}

func printProgress(task string, current, total int) {
    percent := int(float64(current) / float64(total) * 100)
    barWidth := 30
    filled := int(float64(barWidth) * float64(current) / float64(total))
    
    // Clear line
    fmt.Printf("\r%s", strings.Repeat(" ", 60))
    
    // Print progress
    fmt.Printf("\r%s: [", task)
    
    // Filled part
    if percent < 50 {
        fmt.Printf("%s%s%s", Red, strings.Repeat("█", filled), Reset)
    } else if percent < 80 {
        fmt.Printf("%s%s%s", Yellow, strings.Repeat("█", filled), Reset)
    } else {
        fmt.Printf("%s%s%s", Green, strings.Repeat("█", filled), Reset)
    }
    
    // Empty part
    fmt.Printf("%s] %d%%", strings.Repeat("░", barWidth-filled), percent)
}

func printColoredReceipt() {
    width := 35
    
    // Header
    fmt.Printf("%s%s%s\n", BgBlue+White+Bold, 
        center("GOCOFFEE", width), Reset)
    
    fmt.Println(strings.Repeat("─", width))
    
    // Items
    fmt.Printf("2 × %sLatte%s ............ %s$9.00%s\n", 
        Yellow, Reset, Green, Reset)
    fmt.Printf("1 × %sCroissant%s ........ %s$3.50%s\n", 
        Yellow, Reset, Green, Reset)
    
    fmt.Println(strings.Repeat("─", width))
    
    // Totals
    fmt.Printf("Subtotal: ............ %s$12.50%s\n", Green, Reset)
    fmt.Printf("Tax: ................. %s$1.06%s\n", Red, Reset)
    
    fmt.Printf("%s%s%s\n", Bold, strings.Repeat("═", width), Reset)
    
    fmt.Printf("%s%sTOTAL: ............... $13.56%s\n", 
        Bold+Green, Reset)
    
    fmt.Printf("\n%s%sThank you for your visit!%s\n", 
        Bold+Cyan, Reset)
}

func center(s string, width int) string {
    padding := (width - len(s)) / 2
    return fmt.Sprintf("%*s%s%*s", padding, "", s, width-len(s)-padding, "")
}