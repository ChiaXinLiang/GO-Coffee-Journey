package main

import (
    "fmt"
    "strings"
)

type Coffee struct {
    Name  string
    Size  string
    Price float64
}

func main() {
    fmt.Println("=== GoCoffee Format Verbs ===\n")
    
    // General verbs
    fmt.Println("GENERAL VERBS:")
    
    coffee := Coffee{"Latte", "Large", 5.50}
    
    fmt.Printf("%%v  (value):        %v\n", coffee)
    fmt.Printf("%%+v (with fields):  %+v\n", coffee)
    fmt.Printf("%%#v (Go syntax):    %#v\n", coffee)
    fmt.Printf("%%T  (type):         %T\n", coffee)
    fmt.Printf("%%%%  (literal %%):    %%\n")
    
    // Boolean
    fmt.Println("\nBOOLEAN:")
    isOpen := true
    fmt.Printf("%%t: %t\n", isOpen)
    
    // Integers
    fmt.Println("\nINTEGERS:")
    orderNum := 1234
    
    fmt.Printf("%%d (decimal):     %d\n", orderNum)
    fmt.Printf("%%b (binary):      %b\n", orderNum)
    fmt.Printf("%%o (octal):       %o\n", orderNum)
    fmt.Printf("%%x (hex lower):   %x\n", orderNum)
    fmt.Printf("%%X (hex upper):   %X\n", orderNum)
    fmt.Printf("%%c (character):   %c\n", 65) // 'A'
    fmt.Printf("%%U (Unicode):     %U\n", '☕')
    
    // Floating-point
    fmt.Println("\nFLOATING-POINT:")
    price := 12.3456789
    
    fmt.Printf("%%f (decimal):     %f\n", price)
    fmt.Printf("%%.2f (2 decimal): %.2f\n", price)
    fmt.Printf("%%e (scientific):  %e\n", price)
    fmt.Printf("%%E (scientific):  %E\n", price)
    fmt.Printf("%%g (compact):     %g\n", price)
    
    // Strings and slices
    fmt.Println("\nSTRINGS:")
    name := "GoCoffee"
    
    fmt.Printf("%%s (string):      %s\n", name)
    fmt.Printf("%%q (quoted):      %q\n", name)
    fmt.Printf("%%x (hex):         %x\n", name)
    fmt.Printf("%%X (HEX):         %X\n", name)
    
    // Pointers
    fmt.Println("\nPOINTERS:")
    fmt.Printf("%%p (pointer):     %p\n", &coffee)
    
    // Width and precision
    fmt.Println("\nWIDTH AND PRECISION:")
    
    // Width
    fmt.Printf("|%5d|\n", 42)      // Right-aligned
    fmt.Printf("|%-5d|\n", 42)     // Left-aligned
    fmt.Printf("|%05d|\n", 42)     // Zero-padded
    
    // Precision for floats
    pi := 3.14159265359
    fmt.Printf("Default:    %f\n", pi)
    fmt.Printf("%.2f\n", pi)
    fmt.Printf("%.4f\n", pi)
    fmt.Printf("%8.2f\n", pi)   // Width 8, precision 2
    fmt.Printf("%08.2f\n", pi)  // Zero-padded
    
    // Precision for strings
    longName := "Venti Iced Caramel Macchiato"
    fmt.Printf("Full:       %s\n", longName)
    fmt.Printf("Max 10:     %.10s\n", longName)
    fmt.Printf("Width 15:   %15s\n", "Latte")
    fmt.Printf("Left align: %-15s|\n", "Latte")
    
    // Complex formatting
    fmt.Println("\nCOMPLEX FORMATTING:")
    
    items := []struct {
        name  string
        price float64
        qty   int
    }{
        {"Espresso", 3.00, 2},
        {"Cappuccino", 4.50, 1},
        {"Blueberry Muffin", 3.75, 3},
    }
    
    fmt.Printf("%-20s %8s %5s %10s\n", "Item", "Price", "Qty", "Total")
    fmt.Println(strings.Repeat("—", 45))
    
    total := 0.0
    for _, item := range items {
        itemTotal := item.price * float64(item.qty)
        total += itemTotal
        fmt.Printf("%-20s $%7.2f %5d $%9.2f\n", 
            item.name, item.price, item.qty, itemTotal)
    }
    fmt.Println(strings.Repeat("—", 45))
    fmt.Printf("%-20s %8s %5s $%9.2f\n", "TOTAL", "", "", total)
}