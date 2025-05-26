package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("=== Typed vs Untyped Constants ===\n")

	// Untyped constants (flexible)
	const (
		untypedInt    = 100
		untypedFloat  = 3.14
		untypedString = "GoCoffee"
	)

	// Typed constants (strict)
	const (
		typedInt    int     = 100
		typedFloat  float64 = 3.14
		typedString string  = "GoCoffee"
	)

	// Untyped constants are flexible
	var myInt32 int32 = untypedInt
	var myInt64 int64 = untypedInt
	var myFloat float64 = untypedInt

	fmt.Println("UNTYPED CONSTANT FLEXIBILITY:")
	fmt.Printf("untypedInt as int32: %d\n", myInt32)
	fmt.Printf("untypedInt as int64: %d\n", myInt64)
	fmt.Printf("untypedInt as float64: %.1f\n", myFloat)

	// Typed constants are strict
	// var wrongType int32 = typedInt // Error: cannot use typedInt (type int)
	var rightType int = typedInt
	fmt.Printf("\nTyped constant must match: %d\n", rightType)

	// Constants in expressions
	const coffeePrice = 4.50
	const taxRate = 0.085

	// These work with any numeric type
	var totalFloat32 float32 = coffeePrice * (1 + taxRate)
	var totalFloat64 float64 = coffeePrice * (1 + taxRate)

	fmt.Printf("\nPrice calculations:")
	fmt.Printf("\nAs float32: $%.2f", totalFloat32)
	fmt.Printf("\nAs float64: $%.2f\n", totalFloat64)

	// Constant expressions
	const (
		secondsPerMinute = 60
		minutesPerHour   = 60
		hoursPerDay      = 24

		secondsPerHour = secondsPerMinute * minutesPerHour
		secondsPerDay  = secondsPerHour * hoursPerDay
	)

	fmt.Printf("\nTime constants:")
	fmt.Printf("\nSeconds per hour: %d", secondsPerHour)
	fmt.Printf("\nSeconds per day: %d\n", secondsPerDay)

	// Using reflection to check types
	fmt.Println("\nTYPE CHECKING:")
	checkType("untypedInt used as int32", myInt32)
	checkType("untypedInt used as float64", myFloat)
	checkType("typedFloat", typedFloat)
}

func checkType(name string, v interface{}) {
	fmt.Printf("%s: %v (type: %v)\n", name, v, reflect.TypeOf(v))
}