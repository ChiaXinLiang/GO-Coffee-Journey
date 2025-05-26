# Chapter 4.1: Function Basics - Coffee Shop Recipes

[← Previous: Chapter 4 - Functions](../Chapter4_Functions_And_Methods.md) | [Home](../../Go_Coffee_Journey.md) | [Next: Parameters and Returns →](../02-parameters-returns/Parameters_Returns.md)

## Marcus Learns the Art of Recipes

Marcus watched Carlos repeatedly write the same code for making different coffee drinks. Each time an order came in, Carlos would write out all the steps: grind beans, heat water, extract espresso, steam milk...

"This is getting repetitive," Marcus observed.

Sarah nodded. "Time to learn about functions! Think of them as recipes. We write the recipe once, then just call it whenever we need that drink."

## What Are Functions?

<div class="mermaid">
graph LR
    A[Input] --> B[Function]
    B --> C[Process]
    C --> D[Output]
    
    style A fill:#4A5568,stroke:#E2E8F0,color:#E2E8F0
    style B fill:#805AD5,stroke:#E2E8F0,color:#E2E8F0
    style C fill:#2D3748,stroke:#E2E8F0,color:#E2E8F0
    style D fill:#4A5568,stroke:#E2E8F0,color:#E2E8F0
</div>

Functions are reusable blocks of code that:
- Perform specific tasks
- Can accept inputs (parameters)
- Can return outputs (return values)
- Help organize and modularize code
- Make code more maintainable

## Basic Function Syntax

```go
func functionName(parameter1 type1, parameter2 type2) returnType {
    // Function body
    return value
}
```

Sarah showed Marcus the components:
- `func` keyword declares a function
- Function name follows Go naming conventions
- Parameters in parentheses (can be empty)
- Return type after parameters (optional)
- Function body in curly braces

## Your First Functions

<div class="mermaid">
graph TD
    A[main function] --> B[greetCustomer]
    A --> C[makeCoffee]
    A --> D[calculatePrice]
    
    B --> E[Print greeting]
    C --> F[Print recipe]
    D --> G[Return price]
    
    style A fill:#805AD5,stroke:#E2E8F0,color:#E2E8F0
    style B fill:#2D3748,stroke:#E2E8F0,color:#E2E8F0
    style C fill:#2D3748,stroke:#E2E8F0,color:#E2E8F0
    style D fill:#2D3748,stroke:#E2E8F0,color:#E2E8F0
</div>

## Try It Yourself!

Let's start with basic function concepts:

### Example 1: Simple Functions
```bash
go run 01_simple_functions.go
```

### Example 2: Functions with Parameters
```bash
go run 02_with_parameters.go
```

### Example 3: Functions with Returns
```bash
go run 03_with_returns.go
```

### Example 4: Function Scope
```bash
go run 04_function_scope.go
```

### Example 5: Package-Level Functions
```bash
go run 05_package_functions.go
```

### Example 6: Function Organization
```bash
go run 06_function_organization.go
```

### Example 7: Main Function
```bash
go run 07_main_function.go
```

### Example 8: Helper Functions
```bash
go run 08_helper_functions.go
```

### Example 9: Function Practice
```bash
go run 09_function_practice.go
```

## Function Declaration

Carlos explained the different ways to declare functions:

### 1. Basic Function (No Parameters, No Return)
```go
func brewCoffee() {
    fmt.Println("Brewing coffee...")
}
```

### 2. Function with Parameters
```go
func greetCustomer(name string) {
    fmt.Printf("Welcome to GoCoffee, %s!\n", name)
}
```

### 3. Function with Return Value
```go
func calculateTotal(price float64, quantity int) float64 {
    return price * float64(quantity)
}
```

### 4. Function with Multiple Parameters of Same Type
```go
func addSugar(coffee, sugar string) string {
    return coffee + " with " + sugar
}
```

## Function Scope

<div class="mermaid">
graph TD
    A[Package Scope] --> B[Exported Functions]
    A --> C[Unexported Functions]
    
    D[Function Scope] --> E[Parameters]
    D --> F[Local Variables]
    D --> G[Return Values]
    
    style A fill:#805AD5,stroke:#E2E8F0,color:#E2E8F0
    style B fill:#38A169,stroke:#E2E8F0,color:#E2E8F0
    style C fill:#E53E3E,stroke:#E2E8F0,color:#E2E8F0
    style D fill:#805AD5,stroke:#E2E8F0,color:#E2E8F0
</div>

Sarah explained scope rules:

1. **Package-level functions**: Accessible throughout the package
2. **Exported functions**: Start with capital letter, accessible from other packages
3. **Unexported functions**: Start with lowercase, only accessible within package
4. **Local variables**: Only exist within the function

## The main Function

Every Go program starts with the main function:

```go
package main

func main() {
    // Program starts here
}
```

Key points about main:
- Must be in package main
- Takes no arguments
- Returns no values
- Program exits when main returns

## Function Naming Conventions

Carlos shared Go's naming guidelines:

### Good Names
- `calculatePrice` - Clear action
- `isValidOrder` - Boolean check (starts with is/has)
- `GetCustomerName` - Exported function
- `processPayment` - Single responsibility

### Poor Names
- `doStuff` - Too vague
- `calc` - Too abbreviated
- `PROCESS_ORDER` - Not Go style
- `handle_customer_order_and_payment` - Too long

## Common Patterns

Sarah showed common function patterns:

### 1. Guard Clauses
```go
func serveCustomer(age int) {
    if age < 18 {
        fmt.Println("Sorry, coffee only for adults!")
        return
    }
    
    // Continue with service
}
```

### 2. Helper Functions
```go
func makeLatte() {
    espresso := brewEspresso()
    milk := steamMilk()
    return combine(espresso, milk)
}
```

### 3. Validation Functions
```go
func isValidCoffeeSize(size string) bool {
    return size == "small" || size == "medium" || size == "large"
}
```

## Marcus's Challenge

"Your turn," Sarah said. "Create a set of functions for our coffee shop operations. Start simple and build up complexity."

**Challenge Requirements:**
1. Create a function to greet customers
2. Create a function to display the menu
3. Create a function to calculate prices
4. Create a function to process an order
5. Organize functions logically

## Best Practices

Carlos shared function wisdom:
1. **Single Responsibility**: Each function should do one thing
2. **Clear Names**: Function name should describe what it does
3. **Short and Focused**: If it's too long, split it
4. **Consistent Style**: Follow Go conventions
5. **Document Exported Functions**: Add comments for public API

## Common Mistakes

- Forgetting to call the function (defining without using)
- Mismatching function signatures
- Accessing variables outside scope
- Creating functions that are too complex
- Poor naming that doesn't describe purpose

## What's Next?

Marcus had learned the basics of functions. But coffee orders are complex - they need inputs and provide outputs. Next, he would learn about parameters and return values to make functions truly useful.

## Summary

Today Marcus learned:
- Functions organize code into reusable blocks
- Basic function syntax and declaration
- Function scope and visibility rules
- The special main function
- Naming conventions and best practices
- Common function patterns

## References

- [A Tour of Go: Functions](https://tour.golang.org/basics/4)
- [Effective Go: Functions](https://golang.org/doc/effective_go#functions)
- [Go Spec: Function declarations](https://golang.org/ref/spec#Function_declarations)

---

[Continue to Parameters and Returns →](../02-parameters-returns/Parameters_Returns.md)