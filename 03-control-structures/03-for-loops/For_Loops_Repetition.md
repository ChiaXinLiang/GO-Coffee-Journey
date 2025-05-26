# Chapter 3.3: For Loops - The Daily Grind

[← Previous: Switch Statements](../02-switch/Switch_Statements.md) | [Home](../../Go_Coffee_Journey.md) | [Next: Break and Continue →](../04-break-continue/Break_Continue_Control.md)

## Marcus Discovers the Power of Repetition

The morning rush was particularly intense. Marcus watched Carlos methodically brewing coffee after coffee, following the same precise steps each time. The repetitive nature of the work suddenly made him realize something important about programming.

"You know what I noticed?" Marcus said during a brief lull. "We do the same tasks over and over. There must be a way to handle repetition in Go."

Carlos smiled. "Ah, you're ready for loops! In Go, we have one loop construct that's incredibly versatile: the `for` loop. Let me show you how it mirrors our daily grind."

## Understanding For Loops

<div class="mermaid">
graph TD
    A[Start Loop] --> B{Condition True?}
    B -->|Yes| C[Execute Body]
    C --> D[Update Counter]
    D --> B
    B -->|No| E[Exit Loop]
    
    style A fill:#4A5568,stroke:#E2E8F0,color:#E2E8F0
    style B fill:#805AD5,stroke:#E2E8F0,color:#E2E8F0
    style C fill:#2D3748,stroke:#E2E8F0,color:#E2E8F0
    style D fill:#2D3748,stroke:#E2E8F0,color:#E2E8F0
    style E fill:#4A5568,stroke:#E2E8F0,color:#E2E8F0
</div>

## The Three Faces of For Loops

Sarah joined them with her laptop. "Go's `for` loop is unique because it replaces `for`, `while`, and `do-while` from other languages. Let me show you the three main patterns."

### 1. Traditional For Loop
```go
// Classic for loop with initialization, condition, and post statement
for i := 0; i < 5; i++ {
    fmt.Printf("Brewing coffee #%d\n", i+1)
}
```

### 2. While-Style Loop
```go
// For loop with just a condition (like while in other languages)
cupsRemaining := 10
for cupsRemaining > 0 {
    fmt.Printf("Cups remaining: %d\n", cupsRemaining)
    cupsRemaining--
}
```

### 3. Infinite Loop
```go
// Infinite loop (with break to exit)
for {
    fmt.Println("Taking orders...")
    // Use break to exit when needed
    if shopClosed {
        break
    }
}
```

## Real Coffee Shop Applications

<div class="mermaid">
graph LR
    A[Morning Tasks] --> B[Brew Coffee]
    B --> C[Take Orders]
    C --> D[Process Payments]
    D --> E{More Customers?}
    E -->|Yes| C
    E -->|No| F[Clean Up]
    
    style A fill:#4A5568,stroke:#E2E8F0,color:#E2E8F0
    style B fill:#2D3748,stroke:#E2E8F0,color:#E2E8F0
    style C fill:#2D3748,stroke:#E2E8F0,color:#E2E8F0
    style D fill:#2D3748,stroke:#E2E8F0,color:#E2E8F0
    style E fill:#805AD5,stroke:#E2E8F0,color:#E2E8F0
    style F fill:#4A5568,stroke:#E2E8F0,color:#E2E8F0
</div>

## Try It Yourself!

Let's practice with real coffee shop scenarios:

### Example 1: Morning Brewing Routine
```bash
go run 01_basic_for_loop.go
```

### Example 2: Processing Customer Queue
```bash
go run 02_while_style_loop.go
```

### Example 3: Inventory Tracking
```bash
go run 03_loop_with_arrays.go
```

### Example 4: Nested Loops for Menu Display
```bash
go run 04_nested_loops.go
```

### Example 5: Range Loops with Slices
```bash
go run 05_range_loops.go
```

### Example 6: Loop Control with Break
```bash
go run 06_loop_with_break.go
```

### Example 7: Loop Control with Continue
```bash
go run 07_loop_with_continue.go
```

### Example 8: Infinite Loop for Service
```bash
go run 08_infinite_loop.go
```

### Example 9: Performance Comparison
```bash
go run 09_loop_performance.go
```

## Common Loop Patterns

Carlos shared his experience: "Here are the loop patterns we use most often in the coffee shop system."

### 1. Counting Loop
```go
// Count from 1 to 10
for i := 1; i <= 10; i++ {
    fmt.Printf("Customer #%d\n", i)
}
```

### 2. Reverse Loop
```go
// Count down from 10 to 1
for i := 10; i > 0; i-- {
    fmt.Printf("T-minus %d\n", i)
}
```

### 3. Step Loop
```go
// Count by 2s
for i := 0; i < 20; i += 2 {
    fmt.Printf("Table %d\n", i)
}
```

### 4. Range Loop
```go
// Iterate over slice
coffees := []string{"Espresso", "Latte", "Cappuccino"}
for index, coffee := range coffees {
    fmt.Printf("%d: %s\n", index, coffee)
}
```

## Loop Best Practices

<div class="mermaid">
graph TD
    A[Loop Best Practices] --> B[Clear Exit Conditions]
    A --> C[Avoid Infinite Loops]
    A --> D[Use Range When Possible]
    A --> E[Keep Loop Bodies Simple]
    A --> F[Consider Performance]
    
    style A fill:#805AD5,stroke:#E2E8F0,color:#E2E8F0
    style B fill:#2D3748,stroke:#E2E8F0,color:#E2E8F0
    style C fill:#2D3748,stroke:#E2E8F0,color:#E2E8F0
    style D fill:#2D3748,stroke:#E2E8F0,color:#E2E8F0
    style E fill:#2D3748,stroke:#E2E8F0,color:#E2E8F0
    style F fill:#2D3748,stroke:#E2E8F0,color:#E2E8F0
</div>

## Common Pitfalls

Sarah warned about common mistakes: "Watch out for these loop gotchas!"

### 1. Off-by-One Errors
```go
// Wrong: misses the last element
for i := 0; i < len(items)-1; i++ {
    // Process items[i]
}

// Correct: processes all elements
for i := 0; i < len(items); i++ {
    // Process items[i]
}
```

### 2. Modifying Loop Variable
```go
// Dangerous: can cause unexpected behavior
for i := 0; i < 10; i++ {
    if someCondition {
        i = 5  // Don't do this!
    }
}
```

### 3. Infinite Loops
```go
// Always ensure loops can exit
for {
    // Make sure there's a break condition!
    if done {
        break
    }
}
```

## Marcus's Challenge

"Now for your challenge," Sarah said. "Create a coffee shop simulation that runs through a typical day. Use different loop types to handle various scenarios like brewing coffee, serving customers, and tracking inventory."

**Challenge**: Create a program that:
1. Uses a for loop to brew morning coffee (10 cups)
2. Uses a while-style loop to serve customers until closing
3. Uses a range loop to display the menu
4. Uses nested loops to track hourly sales
5. Implements proper loop control with break and continue

## What's Next?

Marcus had mastered the art of repetition in Go. The for loop's versatility impressed him - one construct that could handle all repetitive tasks in the coffee shop.

"Tomorrow," Carlos said, "we'll explore how to control these loops even more precisely with break and continue statements. You'll learn how to exit loops early and skip iterations when needed."

## Summary

Today Marcus learned:
- Go has only one loop construct: `for`
- Three main patterns: traditional, while-style, and infinite
- Range loops for iterating over collections
- Loop best practices and common pitfalls
- Real-world applications in the coffee shop

## References

- [Go Spec: For statements](https://golang.org/ref/spec#For_statements)
- [Effective Go: For](https://golang.org/doc/effective_go#for)
- [Go by Example: For](https://gobyexample.com/for)

---

[Continue to Break and Continue →](../04-break-continue/Break_Continue_Control.md)