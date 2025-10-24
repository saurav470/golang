package main

import (
	"errors"
	"fmt"
)

/*
=============================================================================
TOPIC: Functions in Go
=============================================================================

Functions are the building blocks of Go programs.
Go functions are first-class citizens and can be passed around.

Key Concepts:
- Function declaration
- Parameters and return values
- Multiple return values
- Named return values
- Variadic functions
- Anonymous functions and closures
- Higher-order functions
- Recursion
=============================================================================
*/

func main() {
	fmt.Println("=== Functions in Go ===\n")
	
	// ========================================
	// 1. BASIC FUNCTIONS
	// ========================================
	
	greet("Alice")
	greetWithAge("Bob", 30)
	
	// ========================================
	// 2. RETURN VALUES
	// ========================================
	
	sum := add(5, 3)
	fmt.Printf("5 + 3 = %d\n", sum)
	
	// Multiple return values (very common in Go!)
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("10 / 2 = %.2f\n", result)
	}
	
	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}
	
	// Named return values
	q, r := divideWithRemainder(17, 5)
	fmt.Printf("17 / 5 = %d remainder %d\n", q, r)
	
	// ========================================
	// 3. VARIADIC FUNCTIONS
	// ========================================
	// Functions that accept variable number of arguments
	
	total := sumAll(1, 2, 3, 4, 5)
	fmt.Printf("Sum of 1,2,3,4,5 = %d\n", total)
	
	displayInfo("Alice", "Engineer", "NYC", "alice@example.com")
	
	// ========================================
	// 4. ANONYMOUS FUNCTIONS
	// ========================================
	// Functions without a name (lambda functions)
	
	// Define and call immediately
	func() {
		fmt.Println("Anonymous function called!")
	}()
	
	// Assign to variable
	multiply := func(a, b int) int {
		return a * b
	}
	fmt.Printf("5 * 3 = %d\n", multiply(5, 3))
	
	// ========================================
	// 5. CLOSURES
	// ========================================
	// Functions that reference variables from outside their body
	
	counter := makeCounter()
	fmt.Println("Counter:", counter()) // 1
	fmt.Println("Counter:", counter()) // 2
	fmt.Println("Counter:", counter()) // 3
	
	counter2 := makeCounter()
	fmt.Println("Counter2:", counter2()) // 1 (new closure, new state)
	
	// ========================================
	// 6. HIGHER-ORDER FUNCTIONS
	// ========================================
	// Functions that take functions as parameters or return functions
	
	nums := []int{1, 2, 3, 4, 5}
	
	doubled := mapInts(nums, func(n int) int {
		return n * 2
	})
	fmt.Printf("Original: %v, Doubled: %v\n", nums, doubled)
	
	evens := filterInts(nums, func(n int) bool {
		return n%2 == 0
	})
	fmt.Printf("Even numbers: %v\n", evens)
	
	// ========================================
	// 7. RECURSION
	// ========================================
	
	fmt.Printf("Factorial of 5: %d\n", factorial(5))
	fmt.Printf("Fibonacci of 8: %d\n", fibonacci(8))
	
	// ========================================
	// 8. DEFER IN FUNCTIONS
	// ========================================
	
	processFile("data.txt")
	
	fmt.Println("\n✅ Functions completed!")
}

// ========================================
// FUNCTION DEFINITIONS
// ========================================

// Simple function with parameter
func greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

// Multiple parameters of same type
func greetWithAge(name string, age int) {
	fmt.Printf("Hello, %s! You are %d years old.\n", name, age)
}

// Function with return value
func add(a int, b int) int {
	return a + b
}

// Can shorten parameter list if types are same
func subtract(a, b int) int {
	return a - b
}

// Multiple return values (idiomatic Go for error handling)
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// Named return values
func divideWithRemainder(a, b int) (quotient int, remainder int) {
	quotient = a / b
	remainder = a % b
	return // "naked return" - returns named values
}

// Variadic function (variable number of arguments)
func sumAll(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// Variadic with other parameters
func displayInfo(name string, details ...string) {
	fmt.Printf("Name: %s\n", name)
	for i, detail := range details {
		fmt.Printf("  Detail %d: %s\n", i+1, detail)
	}
}

// Closure - function that returns a function
func makeCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// Higher-order function: takes a function as parameter
func mapInts(numbers []int, fn func(int) int) []int {
	result := make([]int, len(numbers))
	for i, n := range numbers {
		result[i] = fn(n)
	}
	return result
}

// Filter function
func filterInts(numbers []int, predicate func(int) bool) []int {
	result := []int{}
	for _, n := range numbers {
		if predicate(n) {
			result = append(result, n)
		}
	}
	return result
}

// Recursive function
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

// Another recursive example
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// Function demonstrating defer
func processFile(filename string) {
	fmt.Printf("\nProcessing file: %s\n", filename)
	defer fmt.Println("  Closed file")
	defer fmt.Println("  Cleanup completed")
	
	fmt.Println("  Opened file")
	fmt.Println("  Reading data")
	fmt.Println("  Processing data")
}
