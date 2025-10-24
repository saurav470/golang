package main

import "fmt"

/*
=============================================================================
TOPIC: Control Flow in Go
=============================================================================

Control flow determines the order in which code executes.
Go has simple and clean control flow structures.

Key Concepts:
- if/else statements
- switch statements
- for loops (the only loop in Go!)
- break, continue, goto
- defer statements
=============================================================================
*/

func main() {
	fmt.Println("=== Control Flow in Go ===\n")
	
	// ========================================
	// 1. IF-ELSE STATEMENTS
	// ========================================
	
	age := 25
	
	// Simple if
	if age >= 18 {
		fmt.Println("You are an adult")
	}
	
	// if-else
	if age < 13 {
		fmt.Println("You are a child")
	} else if age < 18 {
		fmt.Println("You are a teenager")
	} else {
		fmt.Println("You are an adult")
	}
	
	// if with short statement (common pattern in Go)
	if score := 85; score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else {
		fmt.Println("Grade: C")
	}
	// Note: 'score' is only available within the if-else block
	
	// ========================================
	// 2. SWITCH STATEMENTS
	// ========================================
	
	day := "Monday"
	
	// Basic switch
	switch day {
	case "Monday":
		fmt.Println("Start of the work week")
	case "Tuesday", "Wednesday", "Thursday":
		fmt.Println("Middle of the week")
	case "Friday":
		fmt.Println("TGIF!")
	case "Saturday", "Sunday":
		fmt.Println("Weekend!")
	default:
		fmt.Println("Invalid day")
	}
	
	// Switch with condition (no expression)
	temperature := 25
	switch {
	case temperature < 0:
		fmt.Println("Freezing")
	case temperature < 15:
		fmt.Println("Cold")
	case temperature < 25:
		fmt.Println("Mild")
	default:
		fmt.Println("Warm")
	}
	
	// Switch with short statement
	switch hour := 14; {
	case hour < 12:
		fmt.Println("Good morning!")
	case hour < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}
	
	// Type switch (for interfaces - we'll cover this more later)
	var value interface{} = "Hello"
	switch v := value.(type) {
	case string:
		fmt.Printf("String: %s\n", v)
	case int:
		fmt.Printf("Integer: %d\n", v)
	case bool:
		fmt.Printf("Boolean: %t\n", v)
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}
	
	// ========================================
	// 3. FOR LOOPS
	// ========================================
	// Go has only one loop keyword: 'for'
	// But it can be used in multiple ways!
	
	fmt.Println("\n--- For Loops ---")
	
	// Traditional for loop
	fmt.Print("Traditional: ")
	for i := 0; i < 5; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()
	
	// While-style loop (no while keyword in Go)
	fmt.Print("While-style: ")
	counter := 0
	for counter < 5 {
		fmt.Print(counter, " ")
		counter++
	}
	fmt.Println()
	
	// Infinite loop (use with break)
	fmt.Print("Infinite with break: ")
	n := 0
	for {
		if n >= 5 {
			break
		}
		fmt.Print(n, " ")
		n++
	}
	fmt.Println()
	
	// For-range loop (iterate over arrays, slices, maps, strings)
	fmt.Print("Range over slice: ")
	numbers := []int{10, 20, 30, 40, 50}
	for index, value := range numbers {
		fmt.Printf("[%d:%d] ", index, value)
	}
	fmt.Println()
	
	// Range with index only
	fmt.Print("Range index only: ")
	for i := range numbers {
		fmt.Print(i, " ")
	}
	fmt.Println()
	
	// Range with value only (use _ to ignore index)
	fmt.Print("Range value only: ")
	for _, num := range numbers {
		fmt.Print(num, " ")
	}
	fmt.Println()
	
	// Range over string (iterates over runes)
	fmt.Print("Range over string: ")
	for index, char := range "Go!" {
		fmt.Printf("[%d:%c] ", index, char)
	}
	fmt.Println()
	
	// Range over map
	fmt.Println("Range over map:")
	person := map[string]string{
		"name":  "Alice",
		"city":  "NYC",
		"role":  "Engineer",
	}
	for key, value := range person {
		fmt.Printf("  %s: %s\n", key, value)
	}
	
	// ========================================
	// 4. BREAK, CONTINUE, GOTO
	// ========================================
	
	fmt.Println("\n--- Break and Continue ---")
	
	// Break - exit the loop
	fmt.Print("Break example: ")
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Print(i, " ")
	}
	fmt.Println()
	
	// Continue - skip to next iteration
	fmt.Print("Continue example (skip evens): ")
	for i := 0; i < 10; i++ {
		if i % 2 == 0 {
			continue
		}
		fmt.Print(i, " ")
	}
	fmt.Println()
	
	// Labels with break (for nested loops)
	fmt.Println("Labeled break:")
OuterLoop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				break OuterLoop
			}
			fmt.Printf("  i=%d, j=%d\n", i, j)
		}
	}
	
	// Goto (use sparingly - generally discouraged)
	fmt.Print("Goto example: ")
	i := 0
loop:
	fmt.Print(i, " ")
	i++
	if i < 5 {
		goto loop
	}
	fmt.Println()
	
	// ========================================
	// 5. DEFER STATEMENT
	// ========================================
	// defer delays execution until surrounding function returns
	// Very useful for cleanup (closing files, connections, etc.)
	
	fmt.Println("\n--- Defer Statement ---")
	
	defer fmt.Println("This prints last (deferred)")
	fmt.Println("This prints first")
	fmt.Println("This prints second")
	
	// Multiple defers execute in LIFO (Last In, First Out) order
	fmt.Println("\nMultiple defers (LIFO):")
	for i := 1; i <= 3; i++ {
		defer fmt.Printf("  Deferred %d\n", i)
	}
	fmt.Println("  Regular statement")
	
	// Practical defer example
	demonstrateDefer()
	
	fmt.Println("\n✅ Control Flow completed!")
}

func demonstrateDefer() {
	fmt.Println("\nPractical defer example:")
	defer fmt.Println("  4. Cleanup completed")
	defer fmt.Println("  3. Closing connections")
	
	fmt.Println("  1. Opening resources")
	fmt.Println("  2. Processing data")
}
