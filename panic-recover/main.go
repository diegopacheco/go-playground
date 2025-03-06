package main

import "fmt"

func safeOperation() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	fmt.Println("Starting operation...")
	performRiskyOperation(0)

	// This line won't execute if there's a panic
	fmt.Println("Operation completed successfully")
}

func performRiskyOperation(divisor int) {
	fmt.Println("Performing risky operation...")

	// This will cause a panic if divisor is 0
	result := 10 / divisor

	// Won't execute if panic occurs
	fmt.Println("Result:", result)
}

func main() {
	safeOperation()
	fmt.Println("Program completed normally")
}
