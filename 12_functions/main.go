package main

import "fmt"

// 🔹 FUNCTIONS IN GO

// Functions = reusable blocks of code
// Syntax:
// func functionName(params) returnType {}

// =========================================================
// 🔹 1. BASIC FUNCTION
// =========================================================

// func add(a int, b int) int {
// 	return a + b
// }

// --- Short syntax (same type params) ---
// func add(a, b int) int {
// 	return a + b
// }

// =========================================================
// 🔹 2. MULTIPLE RETURN VALUES
// =========================================================

// func getLanguages() (string, string, string) {
// 	return "golang", "js", "ts"
// }

// --- Usage ---
// l1, l2, l3 := getLanguages()
// l1, l2, _ := getLanguages() // ignore value using _

// =========================================================
// 🔹 3. FUNCTIONS AS PARAMETERS (FIRST-CLASS FUNCTIONS)
// =========================================================

// func processIt(fn func(a int) int) {
// 	fn(1)
// }

// --- Passing function ---
// fn := func(a int) int {
// 	return 2
// }
// processIt(fn)

// =========================================================
// 🔹 4. FUNCTION RETURNING FUNCTION (HIGHER-ORDER FUNCTION)
// =========================================================

func processIt() func(a int) int {
	return func(a int) int {
		return 4
	}
}

// =========================================================
// 🔹 MAIN FUNCTION
// =========================================================

func main() {

	// --- Calling function ---
	// result := add(3, 6)
	// fmt.Println(result)

	// --- Multiple return ---
	// l1, l2, l3 := getLanguages()
	// fmt.Println(l1, l2, l3)

	// --- Function as variable ---
	// fn := func(a int) int {
	// 	return 2
	// }

	// --- Function returning function ---
	a := processIt()
	fmt.Println(a(6))
	// processIt() returns a function
	// a holds that function
	// a(6) executes returned function

	// Output: 4 (argument ignored)
}

// =========================================================
// 🔥 IMPORTANT CONCEPTS
// =========================================================

// --- 1. Functions are FIRST-CLASS CITIZENS ---
// - Can assign to variables
// - Can pass as arguments
// - Can return from functions

// --- 2. Anonymous Functions ---
// func(a int) int { return a }

// --- 3. Type Signature Matters ---
// func(a int) int → input int, output int

// --- 4. Return Function Execution Flow ---
// processIt() → returns function → then call with ()

// =========================================================
// 🔹 QUICK SUMMARY
// =========================================================

// - func keyword defines function
// - Can return multiple values
// - Functions can be passed/returned
// - Anonymous functions allowed
// - Type safety strict in Go

// =========================================================
// 🧠 HINGLISH EASY NOTES
// =========================================================

// function = reusable code block

// Go me function powerful hai:
// - variable me store kar sakte ho
// - function ko function me pass kar sakte ho
// - function return bhi kar sakte ho

// processIt case:
// processIt() → function return karta hai
// a := processIt()
// a(6) → returned function run hota hai

// IMPORTANT:
// argument diya (6) use nahi ho raha
// isliye output always 4

// Ye concept backend + functional programming me bahut important hai 🚀