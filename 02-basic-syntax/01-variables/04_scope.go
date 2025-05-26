package main

import "fmt"

// Package-level variable
var storeName = "GoCoffee Downtown"

func main() {
	// Function-level variable
	dailySales := 0.0

	fmt.Println("Store:", storeName)

	// Block scope
	{
		morningShift := 8
		dailySales = 1250.50
		fmt.Printf("Morning shift starts at: %d AM\n", morningShift)
	}

	// morningShift is not accessible here
	// fmt.Println(morningShift) // This would cause an error

	fmt.Printf("Daily sales: $%.2f\n", dailySales)

	// Call function that uses package variable
	printStoreInfo()
}

func printStoreInfo() {
	// Can access package-level variable
	fmt.Println("Info for:", storeName)

	// Cannot access main's variables
	// fmt.Println(dailySales) // This would cause an error
}