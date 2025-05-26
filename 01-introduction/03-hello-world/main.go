package main

import (
    "fmt"
    "os"
    "time"
)

func main() {
    // Get current hour
    hour := time.Now().Hour()
    
    // Determine greeting based on time
    var greeting string
    switch {
    case hour < 12:
        greeting = "Good morning"
    case hour < 17:
        greeting = "Good afternoon"
    default:
        greeting = "Good evening"
    }
    
    // Check for customer name
    if len(os.Args) > 1 {
        name := os.Args[1]
        fmt.Printf("%s, %s! Welcome to GoCoffee ☕\n", greeting, name)
    } else {
        fmt.Printf("%s! Welcome to GoCoffee ☕\n", greeting)
    }
}