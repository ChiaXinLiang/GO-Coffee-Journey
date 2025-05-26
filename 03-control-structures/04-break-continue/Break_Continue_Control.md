# Chapter 3.4: Break and Continue - Taking Control

[← Previous: For Loops](../03-for-loops/For_Loops_Repetition.md) | [Home](../../Go_Coffee_Journey.md) | [Next: Goto and Labels →](../05-goto-labels/Goto_Labels_Advanced.md)

## Marcus Masters Loop Control

After a particularly hectic morning rush, Marcus noticed Carlos skillfully managing the coffee queue. Sometimes he'd skip certain orders temporarily, other times he'd stop processing entirely when supplies ran low. This gave Marcus an idea about controlling program flow.

"I see you're not just blindly processing every order," Marcus observed. "Sometimes you skip, sometimes you stop entirely."

Carlos nodded. "Exactly! In programming, we call these controls `break` and `continue`. They give us fine-grained control over our loops. Let me show you how they work."

## Understanding Break and Continue

<div class="mermaid">
graph TD
    A[Loop Start] --> B{Check Condition}
    B -->|True| C{Continue?}
    C -->|Yes| A
    C -->|No| D{Break?}
    D -->|Yes| E[Exit Loop]
    D -->|No| F[Execute Code]
    F --> G[Next Iteration]
    G --> B
    B -->|False| E
    
    style A fill:#4A5568,stroke:#E2E8F0,color:#E2E8F0
    style B fill:#805AD5,stroke:#E2E8F0,color:#E2E8F0
    style C fill:#D69E2E,stroke:#E2E8F0,color:#E2E8F0
    style D fill:#E53E3E,stroke:#E2E8F0,color:#E2E8F0
    style E fill:#4A5568,stroke:#E2E8F0,color:#E2E8F0
    style F fill:#2D3748,stroke:#E2E8F0,color:#E2E8F0
    style G fill:#2D3748,stroke:#E2E8F0,color:#E2E8F0
</div>

## The Power of Break

Sarah explained, "The `break` statement immediately exits the current loop. It's like an emergency stop button."

### When to Use Break:
1. **Found what you're looking for** - Stop searching once found
2. **Error condition** - Exit when something goes wrong
3. **Resource exhaustion** - Stop when supplies run out
4. **User cancellation** - Exit on user request

```go
// Example: Stop when coffee beans run out
for orderNum := 1; orderNum <= 100; orderNum++ {
    if coffeeBeans < 18 {
        fmt.Println("Out of coffee beans!")
        break
    }
    // Process order
}
```

## The Flexibility of Continue

"The `continue` statement skips the rest of the current iteration and moves to the next one," Carlos added. "It's like saying 'skip this one, move to the next.'"

### When to Use Continue:
1. **Skip invalid data** - Ignore bad inputs
2. **Filter conditions** - Process only specific items  
3. **Temporary unavailability** - Skip items not ready
4. **Performance optimization** - Skip unnecessary processing

```go
// Example: Skip decaf orders when machine is broken
for _, order := range orders {
    if order.Type == "Decaf" && decafMachineBroken {
        fmt.Println("Skipping decaf order - machine broken")
        continue
    }
    // Process order
}
```

## Break vs Continue Comparison

<div class="mermaid">
graph LR
    A[Loop Running] --> B{Condition Met?}
    B -->|Break| C[Exit Loop Completely]
    B -->|Continue| D[Skip to Next Iteration]
    D --> A
    
    style A fill:#2D3748,stroke:#E2E8F0,color:#E2E8F0
    style B fill:#805AD5,stroke:#E2E8F0,color:#E2E8F0
    style C fill:#E53E3E,stroke:#E2E8F0,color:#E2E8F0
    style D fill:#D69E2E,stroke:#E2E8F0,color:#E2E8F0
</div>

## Try It Yourself!

Practice loop control with these real-world scenarios:

### Example 1: Order Processing with Break
```bash
go run 01_break_basics.go
```

### Example 2: Continue for Filtering
```bash
go run 02_continue_basics.go
```

### Example 3: Nested Loop Control
```bash
go run 03_nested_loop_control.go
```

### Example 4: Search Operations
```bash
go run 04_search_with_break.go
```

### Example 5: Data Validation
```bash
go run 05_validation_with_continue.go
```

### Example 6: Complex Flow Control
```bash
go run 06_complex_flow_control.go
```

### Example 7: Label Break and Continue
```bash
go run 07_labeled_statements.go
```

### Example 8: Performance Optimization
```bash
go run 08_performance_control.go
```

### Example 9: Real-world Application
```bash
go run 09_coffee_shop_simulation.go
```

## Advanced Techniques: Labeled Statements

<div class="mermaid">
graph TD
    A[Outer Loop] --> B[Inner Loop]
    B --> C{Break with Label?}
    C -->|Yes| D[Exit to Labeled Point]
    C -->|No| E[Exit Inner Loop Only]
    
    style A fill:#4A5568,stroke:#E2E8F0,color:#E2E8F0
    style B fill:#2D3748,stroke:#E2E8F0,color:#E2E8F0
    style C fill:#805AD5,stroke:#E2E8F0,color:#E2E8F0
    style D fill:#E53E3E,stroke:#E2E8F0,color:#E2E8F0
    style E fill:#D69E2E,stroke:#E2E8F0,color:#E2E8F0
</div>

Sarah showed an advanced technique: "Sometimes you need to break out of nested loops. Labels help with that."

```go
OuterLoop:
    for floor := 1; floor <= 3; floor++ {
        for table := 1; table <= 10; table++ {
            if foundEmptyTable(floor, table) {
                fmt.Printf("Found table %d on floor %d\n", table, floor)
                break OuterLoop // Breaks both loops
            }
        }
    }
```

## Best Practices

Carlos shared his experience: "Here are the guidelines we follow for clean, maintainable code."

### 1. Use Break Sparingly
```go
// Good: Clear exit condition
for _, item := range inventory {
    if item.Quantity == 0 {
        fmt.Printf("%s is out of stock\n", item.Name)
        break
    }
    process(item)
}

// Avoid: Multiple breaks make flow hard to follow
for {
    if condition1 {
        break
    }
    // ... lots of code ...
    if condition2 {
        break
    }
    // ... more code ...
    if condition3 {
        break
    }
}
```

### 2. Continue for Clarity
```go
// Good: Clear filtering
for _, customer := range queue {
    if customer.Priority != "VIP" {
        continue
    }
    // Process VIP customers
    serveVIPCustomer(customer)
}

// Less clear: Nested if
for _, customer := range queue {
    if customer.Priority == "VIP" {
        // Process VIP customers
        serveVIPCustomer(customer)
    }
}
```

### 3. Avoid Deep Nesting
```go
// Better: Early continue
for _, order := range orders {
    if !order.IsValid() {
        continue
    }
    if !hasIngredients(order) {
        continue
    }
    if !paymentProcessed(order) {
        continue
    }
    
    prepareOrder(order)
}

// Avoid: Deep nesting
for _, order := range orders {
    if order.IsValid() {
        if hasIngredients(order) {
            if paymentProcessed(order) {
                prepareOrder(order)
            }
        }
    }
}
```

## Common Pitfalls

Sarah warned about common mistakes:

### 1. Break in Switch Inside Loop
```go
for i := 0; i < 10; i++ {
    switch i {
    case 5:
        break // This breaks the switch, not the loop!
    }
    fmt.Println(i) // This still executes
}
```

### 2. Unreachable Code
```go
for i := 0; i < 10; i++ {
    fmt.Println(i)
    break
    fmt.Println("Never executes") // Unreachable
}
```

### 3. Forgetting Loop Increment with Continue
```go
i := 0
for i < 10 {
    if i == 5 {
        continue // Infinite loop! i never increments
    }
    i++
}
```

## Marcus's Challenge

"Time for your challenge," Sarah announced. "Create a coffee shop queue management system that uses break and continue effectively. The system should handle VIP customers, skip unavailable items, and stop service at closing time."

**Challenge Requirements:**
1. Process a queue of customers with different priorities
2. Skip orders for unavailable items using `continue`
3. Give VIP customers priority using loop control
4. Stop service at closing time using `break`
5. Handle nested loops for multiple service counters

## What's Next?

Marcus had learned to control loops with precision, making his programs more efficient and readable. The ability to skip iterations and exit loops early opened up new possibilities.

"Tomorrow," Carlos said with a mysterious smile, "we'll explore the controversial `goto` statement and labels. Used wisely, they can solve specific problems elegantly."

## Summary

Today Marcus learned:
- `break` exits the current loop immediately
- `continue` skips to the next iteration
- Labeled statements control nested loops
- Best practices for clean, readable code
- Common pitfalls to avoid

## References

- [Go Spec: Break statements](https://golang.org/ref/spec#Break_statements)
- [Go Spec: Continue statements](https://golang.org/ref/spec#Continue_statements)
- [Effective Go: Control structures](https://golang.org/doc/effective_go#control-structures)

---

[Continue to Goto and Labels →](../05-goto-labels/Goto_Labels_Advanced.md)