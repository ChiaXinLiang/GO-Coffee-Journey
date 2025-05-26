# Chapter 3: Control Structures - The Flow of Coffee

[← Previous: Chapter 2 - Basic Syntax](../02-basic-syntax/Chapter2_Basic_Syntax.md) | [Home](../Go_Coffee_Journey.md) | [Next: Chapter 4 - Functions →](../04-functions/Chapter4_Functions_And_Methods.md)

## Marcus Takes Control

After mastering Go's basic syntax, Marcus was ready to learn how to control the flow of his programs. Sarah gathered the team for an important lesson.

"Think of control structures like the decision-making process in our coffee shop," Sarah explained. "Every day, we make choices: which coffee to brew, how many times to steam the milk, when to stop taking orders. Go's control structures let us model these decisions in code."

## Chapter Overview

In this chapter, Marcus learns how to:
- Make decisions with if-else statements
- Handle multiple choices efficiently with switch
- Repeat tasks using for loops
- Control loop execution with break and continue
- Understand the controversial goto statement

<div class="mermaid">
graph TD
    A[Control Structures] --> B[If-Else]
    A --> C[Switch]
    A --> D[For Loops]
    A --> E[Break/Continue]
    A --> F[Goto/Labels]
    
    B --> G[Simple Conditions]
    B --> H[Complex Logic]
    
    C --> I[Value Matching]
    C --> J[Type Switching]
    
    D --> K[Traditional]
    D --> L[While-style]
    D --> M[Infinite]
    D --> N[Range]
    
    style A fill:#805AD5,stroke:#E2E8F0,color:#E2E8F0
    style B fill:#2D3748,stroke:#E2E8F0,color:#E2E8F0
    style C fill:#2D3748,stroke:#E2E8F0,color:#E2E8F0
    style D fill:#2D3748,stroke:#E2E8F0,color:#E2E8F0
    style E fill:#2D3748,stroke:#E2E8F0,color:#E2E8F0
    style F fill:#2D3748,stroke:#E2E8F0,color:#E2E8F0
</div>

## 3.1 If-Else Decisions

Marcus starts with the fundamental decision-making structure. Like deciding whether a customer wants sugar in their coffee, if-else statements handle binary choices and complex conditions.

[**Start Learning →**](01-if-else/If_Else_Decisions.md)

Key concepts:
- Basic if statements
- If-else chains
- Nested conditions
- Short variable declarations in conditions
- Best practices for readable conditions

## 3.2 Switch Statements

When the coffee menu grew, Marcus needed a better way to handle multiple options. Switch statements provide elegant solutions for multi-way decisions.

[**Continue Learning →**](02-switch/Switch_Statements.md)

Key concepts:
- Basic switch syntax
- Multiple case values
- Default cases
- Fallthrough behavior
- Type switches
- Expression switches

## 3.3 For Loops - The Daily Grind

The repetitive nature of coffee making taught Marcus about loops. Go's versatile for loop handles all repetition needs.

[**Continue Learning →**](03-for-loops/For_Loops_Repetition.md)

Key concepts:
- Traditional for loops
- While-style loops
- Infinite loops
- Range loops
- Loop performance
- Common patterns

## 3.4 Break and Continue

Sometimes you need to exit early or skip iterations. Marcus learns fine-grained loop control.

[**Continue Learning →**](04-break-continue/Break_Continue_Control.md)

Key concepts:
- Break statements
- Continue statements
- Labeled breaks and continues
- Nested loop control
- Best practices

## 3.5 Goto and Labels (Advanced)

The controversial goto statement - Marcus learns why it exists and why to avoid it.

[**Continue Learning →**](05-goto-labels/Goto_Labels_Advanced.md)

Key concepts:
- Goto syntax
- Go's restrictions
- Why goto is dangerous
- Better alternatives
- Legitimate use cases (rare)

## Real-World Applications

Throughout this chapter, Marcus applies control structures to real coffee shop scenarios:
- Order processing logic
- Inventory management
- Customer queue handling
- Payment processing
- Daily operation cycles

## Best Practices

Sarah's control structure wisdom:
1. **Keep conditions simple** - Complex conditions are hard to debug
2. **Prefer switch over long if-else chains** - More readable and efficient
3. **Use early returns** - Reduce nesting by handling edge cases first
4. **One loop, one purpose** - Don't try to do too much in a single loop
5. **Avoid goto** - There's almost always a better way

## Common Pitfalls

Carlos warns about typical mistakes:
- Forgetting that switch doesn't fall through by default
- Creating infinite loops accidentally
- Using break in switch inside a loop (breaks switch, not loop)
- Complex nested conditions that are hard to understand
- Overusing goto when better structures exist

## Chapter Summary

By the end of Chapter 3, Marcus has mastered:
- ✓ Making decisions with if-else statements
- ✓ Handling multiple cases with switch
- ✓ Creating loops for repetitive tasks
- ✓ Controlling loop flow with break and continue
- ✓ Understanding why goto should be avoided

## What's Next?

With control structures mastered, Marcus is ready to learn about functions - the building blocks of modular, reusable code. In Chapter 4, he'll discover how to write clean, testable functions that make the coffee shop's operations more efficient.

---

[Continue to Chapter 4: Functions →](../04-functions/Chapter4_Functions_And_Methods.md)