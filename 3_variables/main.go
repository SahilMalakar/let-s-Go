package main

import "fmt" 

func main() { 

	// var name string = "golang"
	// Explicit type declaration (string)

	// var name = "golang"
	// Type inference → Go automatically detects type (string)

	var isAdult bool = true // Explicit bool type (true/false only, lowercase)

	var age int = 30 // int type (size depends on system: usually 64-bit on modern PCs)

	name := "golang"
	// name := "go" ❌ Error (already declared)

	// Short declaration syntax (:=)
	// Can ONLY be used inside functions
	// Automatically infers type (string here)

	name = "go" // reassignment

	// var price float32 = 50.5
	// float32 → 32-bit precision (less memory, less precision)

	// var price2 float32 = 50
	// 50 is int, but Go allows constant assignment if compatible

	// var price3 float64 = 46.34
	// float64 → default float type in Go

	price := 30.2
	// 30.2 becomes float64 by default (IMPORTANT: decimals → float64)

	fmt.Println(name)    // Prints string
	fmt.Println(isAdult) // Prints bool
	fmt.Println(age)     // Prints int
	fmt.Println(price)   // Prints float64



	var x int = 10
    var y float64 = 2.5
    // fmt.Println(x + y) // ❌ Type mismatch
	fmt.Println(float64(x) + y)
	// Go does NOT do automatic type conversion.
}