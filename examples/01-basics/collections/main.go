package main

import "fmt"

/*
=============================================================================
TOPIC: Arrays, Slices, and Maps in Go
=============================================================================

Go's built-in data structures for storing collections.

Key Concepts:
- Arrays (fixed size)
- Slices (dynamic arrays)
- Slice operations (append, copy, slicing)
- Maps (key-value pairs)
- Map operations
- Iterating over collections
=============================================================================
*/

func main() {
	fmt.Println("=== Arrays, Slices, and Maps ===\n")
	
	// ========================================
	// 1. ARRAYS
	// ========================================
	// Fixed size, size is part of the type
	
	// Declare array
	var numbers [5]int
	fmt.Printf("Empty array: %v (length: %d)\n", numbers, len(numbers))
	
	// Initialize array
	primes := [5]int{2, 3, 5, 7, 11}
	fmt.Printf("Primes: %v\n", primes)
	
	// Let compiler count the size
	days := [...]string{"Mon", "Tue", "Wed", "Thu", "Fri"}
	fmt.Printf("Days: %v (length: %d)\n", days, len(days))
	
	// Access elements
	fmt.Printf("First prime: %d, Last prime: %d\n", primes[0], primes[4])
	
	// Modify elements
	primes[0] = 1
	fmt.Printf("Modified primes: %v\n", primes)
	primes[0] = 2 // fix it back
	
	// Arrays are value types (copied when assigned)
	primes2 := primes
	primes2[0] = 99
	fmt.Printf("primes: %v, primes2: %v (separate copies)\n", primes, primes2)
	
	// ========================================
	// 2. SLICES
	// ========================================
	// Dynamic arrays, most commonly used in Go
	
	// Create slice using make
	scores := make([]int, 3)  // length 3, capacity 3
	fmt.Printf("\nSlice with make: %v (len: %d, cap: %d)\n", 
		scores, len(scores), cap(scores))
	
	// Create slice with literal
	colors := []string{"red", "green", "blue"}
	fmt.Printf("Colors: %v\n", colors)
	
	// Slice of a slice (shares underlying array)
	subColors := colors[1:3]  // from index 1 to 2
	fmt.Printf("Sub-slice [1:3]: %v\n", subColors)
	
	// Slicing syntax
	fmt.Printf("colors[:2]: %v\n", colors[:2])   // first 2
	fmt.Printf("colors[1:]: %v\n", colors[1:])   // from index 1 to end
	fmt.Printf("colors[:]: %v\n", colors[:])     // all elements
	
	// ========================================
	// 3. SLICE OPERATIONS
	// ========================================
	
	// Append (most important slice operation)
	colors = append(colors, "yellow")
	fmt.Printf("\nAfter append: %v\n", colors)
	
	// Append multiple
	colors = append(colors, "purple", "orange")
	fmt.Printf("After multiple append: %v\n", colors)
	
	// Append slice to slice
	moreColors := []string{"black", "white"}
	colors = append(colors, moreColors...)  // ... unpacks the slice
	fmt.Printf("After appending slice: %v\n", colors)
	
	// Copy slices
	colorsCopy := make([]string, len(colors))
	copy(colorsCopy, colors)
	fmt.Printf("Copied slice: %v\n", colorsCopy)
	
	// Remove element (no built-in remove, use slicing)
	indexToRemove := 2
	colors = append(colors[:indexToRemove], colors[indexToRemove+1:]...)
	fmt.Printf("After removing index %d: %v\n", indexToRemove, colors)
	
	// ========================================
	// 4. SLICE INTERNALS
	// ========================================
	
	// Slice has length and capacity
	numbers2 := make([]int, 3, 5)  // length 3, capacity 5
	fmt.Printf("\nSlice: %v (len: %d, cap: %d)\n", 
		numbers2, len(numbers2), cap(numbers2))
	
	numbers2 = append(numbers2, 10, 20)
	fmt.Printf("After append: %v (len: %d, cap: %d)\n", 
		numbers2, len(numbers2), cap(numbers2))
	
	// When capacity is exceeded, a new underlying array is allocated
	numbers2 = append(numbers2, 30)  // exceeds capacity
	fmt.Printf("After exceeding cap: %v (len: %d, cap: %d)\n", 
		numbers2, len(numbers2), cap(numbers2))
	
	// ========================================
	// 5. MULTIDIMENSIONAL SLICES
	// ========================================
	
	// 2D slice
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("\n2D Slice (Matrix):")
	for i, row := range matrix {
		fmt.Printf("  Row %d: %v\n", i, row)
	}
	
	// ========================================
	// 6. MAPS
	// ========================================
	// Key-value data structure (like dictionary/hash table)
	
	// Create map using make
	ages := make(map[string]int)
	ages["Alice"] = 25
	ages["Bob"] = 30
	ages["Charlie"] = 35
	fmt.Printf("\nAges map: %v\n", ages)
	
	// Create map with literal
	capitals := map[string]string{
		"USA":     "Washington DC",
		"UK":      "London",
		"France":  "Paris",
		"Germany": "Berlin",
	}
	fmt.Printf("Capitals: %v\n", capitals)
	
	// Access map values
	fmt.Printf("Capital of France: %s\n", capitals["France"])
	
	// Check if key exists (important pattern!)
	capital, exists := capitals["Japan"]
	if exists {
		fmt.Printf("Capital of Japan: %s\n", capital)
	} else {
		fmt.Println("Japan not found in map")
	}
	
	// Add/Update
	capitals["Japan"] = "Tokyo"
	fmt.Printf("After adding Japan: %v\n", capitals)
	
	// Delete
	delete(capitals, "UK")
	fmt.Printf("After deleting UK: %v\n", capitals)
	
	// Map length
	fmt.Printf("Map size: %d\n", len(capitals))
	
	// ========================================
	// 7. ITERATING OVER COLLECTIONS
	// ========================================
	
	fmt.Println("\nIterating over slice:")
	fruits := []string{"apple", "banana", "cherry"}
	for index, fruit := range fruits {
		fmt.Printf("  %d: %s\n", index, fruit)
	}
	
	fmt.Println("Iterating over map:")
	for country, capital := range capitals {
		fmt.Printf("  %s -> %s\n", country, capital)
	}
	
	// ========================================
	// 8. PRACTICAL EXAMPLES
	// ========================================
	
	// Example: Count word frequency
	words := []string{"apple", "banana", "apple", "cherry", "banana", "apple"}
	frequency := make(map[string]int)
	
	for _, word := range words {
		frequency[word]++
	}
	
	fmt.Println("\nWord frequency:")
	for word, count := range frequency {
		fmt.Printf("  %s: %d\n", word, count)
	}
	
	// Example: Filter slice
	numbers3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evens := []int{}
	for _, num := range numbers3 {
		if num%2 == 0 {
			evens = append(evens, num)
		}
	}
	fmt.Printf("\nEven numbers from %v: %v\n", numbers3, evens)
	
	// Example: Map transformation
	transformed := make(map[string]int)
	for key, value := range ages {
		transformed[key] = value * 2
	}
	fmt.Printf("Ages doubled: %v\n", transformed)
	
	fmt.Println("\n✅ Arrays, Slices, and Maps completed!")
}
