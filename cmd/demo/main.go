package main

import (
        "fmt"
        "os"
)

/*
=============================================================================
GO LEARNING PROJECT - MAIN DEMO
=============================================================================

This is the main entry point that demonstrates all learning modules.

Each example can be run individually from its directory:
  go run examples/01-basics/variables-types/main.go
  go run examples/01-basics/control-flow/main.go
  go run examples/02-structs-interfaces/structs/main.go
  ... etc

Or run this demo to see an overview.
=============================================================================
*/

func main() {
        fmt.Println("╔════════════════════════════════════════════════════════╗")
        fmt.Println("║     Welcome to Go Learning Project!                   ║")
        fmt.Println("║     From Beginner to Advanced                         ║")
        fmt.Println("╚════════════════════════════════════════════════════════╝")
        fmt.Println()
        
        showMenu()
}

func showMenu() {
        fmt.Println("📚 Learning Modules Available:\n")
        
        fmt.Println("🔹 Beginner Level (01-basics):")
        fmt.Println("   1. Variables and Types")
        fmt.Println("   2. Control Flow (if, switch, for, defer)")
        fmt.Println("   3. Functions (closures, recursion, higher-order)")
        fmt.Println("   4. Collections (arrays, slices, maps)")
        
        fmt.Println("\n🔹 Intermediate Level (02-structs-interfaces):")
        fmt.Println("   5. Structs and Methods")
        fmt.Println("   6. Interfaces and Polymorphism")
        fmt.Println("   7. Error Handling")
        fmt.Println("   8. Pointers")
        
        fmt.Println("\n🔹 Web Development (03-04):")
        fmt.Println("   9. HTTP Server Basics")
        fmt.Println("   10. RESTful API with JSON")
        
        fmt.Println("\n🔹 Advanced Concurrency (05-06):")
        fmt.Println("   11. Goroutines and Channels")
        fmt.Println("   12. Context, WaitGroups, Mutex")
        
        fmt.Println("\n🔹 Database (07):")
        fmt.Println("   13. SQL Operations with SQLite")
        
        fmt.Println("\n🔹 Testing (08):")
        fmt.Println("   14. Unit Tests and Benchmarks")
        
        fmt.Println("\n🔹 Design Patterns (09):")
        fmt.Println("   15. Factory, Singleton, Builder, Strategy, Observer")
        
        fmt.Println("\n🔹 Packages & Modules (10):")
        fmt.Println("   16. Creating and Using Custom Packages")
        
        fmt.Println("\n🔹 Performance Optimization (11):")
        fmt.Println("   17. sync.Pool, sync.Map, sync.Once, String Builder, Worker Pools")
        
        fmt.Println("\n🔹 Large Data Processing (12):")
        fmt.Println("   18. 1BRC Techniques: Chunking, Sharding, Buffer Pooling")
        
        fmt.Println("\n" + string(make([]byte, 56)))
        for i := 0; i < 56; i++ {
                fmt.Print("─")
        }
        fmt.Println()
        
        fmt.Println("\n💡 How to run examples:")
        fmt.Println("   cd examples/01-basics/variables-types && go run main.go")
        fmt.Println("   cd examples/03-http-server/basic-server && go run main.go")
        fmt.Println("   cd examples/04-rest-api && go run main.go")
        fmt.Println("   cd examples/05-concurrency/goroutines && go run main.go")
        fmt.Println("   cd examples/07-database/sql-basics && go run main.go")
        fmt.Println("   cd examples/09-patterns && go run main.go")
        fmt.Println("   cd examples/10-packages-modules/app && go run main.go")
        fmt.Println("   cd examples/11-performance && go run main.go")
        fmt.Println("   cd examples/12-large-data-processing && go run main.go")
        
        fmt.Println("\n💡 Run tests:")
        fmt.Println("   cd examples/08-testing && go test -v")
        fmt.Println("   cd examples/08-testing && go test -bench=.")
        
        fmt.Println("\n📖 Project Structure:")
        fmt.Println("   examples/  - All learning examples organized by topic")
        fmt.Println("   cmd/       - Command-line applications")
        fmt.Println("   pkg/       - Reusable packages")
        fmt.Println("   internal/  - Private application code")
        
        fmt.Println("\n🚀 Quick Start Recommendations:")
        fmt.Println("   1. Start with basics: variables, control flow, functions")
        fmt.Println("   2. Learn structs and interfaces")
        fmt.Println("   3. Build HTTP servers and APIs")
        fmt.Println("   4. Master concurrency (goroutines, channels)")
        fmt.Println("   5. Practice with databases and testing")
        fmt.Println("   6. Study design patterns and packages")
        fmt.Println("   7. Optimize performance for production")
        fmt.Println("   8. Process large datasets efficiently")
        
        fmt.Println("\n✨ Happy Learning! ✨\n")
        
        os.Exit(0)
}
