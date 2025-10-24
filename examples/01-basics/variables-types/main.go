package main

import "fmt"

/*
=============================================================================
TOPIC: Variables and Data Types in Go
=============================================================================

Go is a statically typed language, meaning variable types are known at compile time.
This file covers all fundamental data types and variable declaration methods.

Key Concepts:
- Variable declarations (var, :=)
- Basic types (int, float, string, bool)
- Zero values
- Type conversion
- Constants
=============================================================================
*/

func main() {
	fmt.Println("=== Variables and Types in Go ===\n")
	
	// ========================================
	// 1. VARIABLE DECLARATION METHODS
	// ========================================
	
	// Method 1: var keyword with explicit type
	var name string = "John Doe"
	fmt.Println("Method 1 - Explicit type:", name)
	
	// Method 2: var with type inference
	var age = 25
	fmt.Println("Method 2 - Type inference:", age)
	
	// Method 3: Short declaration (most common in Go) - only inside functions
	email := "john@example.com"
	fmt.Println("Method 3 - Short declaration:", email)
	
	// Method 4: Multiple variable declaration
	var (
		firstName string = "Alice"
		lastName  string = "Smith"
		salary    int    = 75000
	)
	fmt.Printf("Method 4 - Multiple: %s %s earns $%d\n", firstName, lastName, salary)
	
	// ========================================
	// 2. BASIC DATA TYPES
	// ========================================
	
	// Integer types (signed)
	var smallInt int8 = 127           // -128 to 127
	var regularInt int16 = 32767      // -32768 to 32767
	var bigInt int32 = 2147483647     // -2^31 to 2^31-1
	var hugeInt int64 = 9223372036854775807 // -2^63 to 2^63-1
	var platformInt int = 123456      // Platform dependent (32 or 64 bit)
	
	fmt.Printf("\nInteger Types:\nint8: %d, int16: %d, int32: %d, int64: %d, int: %d\n",
		smallInt, regularInt, bigInt, hugeInt, platformInt)
	
	// Unsigned integers
	var unsignedByte uint8 = 255      // 0 to 255 (also called 'byte')
	var unsignedSmall uint16 = 65535  // 0 to 65535
	var unsignedBig uint32 = 4294967295 // 0 to 2^32-1
	var unsignedHuge uint64 = 18446744073709551615 // 0 to 2^64-1
	
	fmt.Printf("Unsigned: uint8: %d, uint16: %d, uint32: %d, uint64: %d\n",
		unsignedByte, unsignedSmall, unsignedBig, unsignedHuge)
	
	// Floating point numbers
	var price float32 = 19.99          // 32-bit floating point
	var precise float64 = 3.14159265359 // 64-bit floating point (preferred)
	
	fmt.Printf("\nFloating Point:\nfloat32: %.2f, float64: %.11f\n", price, precise)
	
	// Boolean
	var isActive bool = true
	var isDeleted bool = false
	
	fmt.Printf("\nBoolean:\nisActive: %t, isDeleted: %t\n", isActive, isDeleted)
	
	// String
	var greeting string = "Hello, Go!"
	var multiLine string = `This is a
multi-line string
using backticks`
	
	fmt.Printf("\nStrings:\ngreeting: %s\nmultiLine: %s\n", greeting, multiLine)
	
	// Rune (represents a Unicode code point - alias for int32)
	var letter rune = 'A'              // Single character
	var emoji rune = '😀'
	
	fmt.Printf("\nRunes:\nletter: %c (value: %d), emoji: %c (value: %d)\n",
		letter, letter, emoji, emoji)
	
	// Byte (alias for uint8)
	var b byte = 'Z'
	fmt.Printf("Byte: %c (value: %d)\n", b, b)
	
	// ========================================
	// 3. ZERO VALUES
	// ========================================
	// Variables declared without initialization get "zero values"
	
	var defaultInt int           // 0
	var defaultFloat float64     // 0.0
	var defaultBool bool         // false
	var defaultString string     // "" (empty string)
	
	fmt.Printf("\nZero Values:\nint: %d, float64: %.1f, bool: %t, string: '%s'\n",
		defaultInt, defaultFloat, defaultBool, defaultString)
	
	// ========================================
	// 4. TYPE CONVERSION
	// ========================================
	// Go requires explicit type conversion (no implicit conversion)
	
	var integer int = 42
	var floating float64 = float64(integer)  // int to float64
	var unsigned uint = uint(integer)        // int to uint
	
	fmt.Printf("\nType Conversion:\nint: %d -> float64: %.1f, uint: %d\n",
		integer, floating, unsigned)
	
	// String to bytes and vice versa
	text := "Hello"
	bytes := []byte(text)           // string to []byte
	backToString := string(bytes)   // []byte to string
	
	fmt.Printf("String '%s' -> Bytes: %v -> String: '%s'\n",
		text, bytes, backToString)
	
	// ========================================
	// 5. CONSTANTS
	// ========================================
	// Constants cannot be changed after declaration
	
	const Pi = 3.14159
	const MaxConnections = 100
	const AppName = "Go Learning App"
	
	// Multiple constants
	const (
		StatusOK       = 200
		StatusNotFound = 404
		StatusError    = 500
	)
	
	fmt.Printf("\nConstants:\nPi: %.5f, MaxConnections: %d, AppName: %s\n",
		Pi, MaxConnections, AppName)
	fmt.Printf("HTTP Status Codes: OK=%d, NotFound=%d, Error=%d\n",
		StatusOK, StatusNotFound, StatusError)
	
	// iota - enumeration constant generator
	const (
		Sunday = iota    // 0
		Monday           // 1
		Tuesday          // 2
		Wednesday        // 3
		Thursday         // 4
		Friday           // 5
		Saturday         // 6
	)
	
	fmt.Printf("\nDays using iota:\nSunday=%d, Monday=%d, Friday=%d\n",
		Sunday, Monday, Friday)
	
	// ========================================
	// 6. TYPE INFERENCE
	// ========================================
	// Go can infer the type from the value
	
	inferredInt := 42              // int
	inferredFloat := 3.14          // float64
	inferredString := "Go rocks!"  // string
	inferredBool := true           // bool
	
	fmt.Printf("\nType Inference:\n%T, %T, %T, %T\n",
		inferredInt, inferredFloat, inferredString, inferredBool)
	
	fmt.Println("\n✅ Basics: Variables and Types completed!")
}
