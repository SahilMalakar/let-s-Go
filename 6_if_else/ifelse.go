package main 

import "fmt" 

func main() {

	age := 17 // Short declaration (type inferred as int)

	// --- Basic if-else ladder ---
	if age >= 18 { // Condition must be boolean (no parentheses required)

		fmt.Println("u r an adult")

	} else if age >= 12 { 
		// Checked only if first condition is false

		fmt.Println("u r a teenager")

	} else {

		fmt.Println("u r not an adult")

	}
	// ⚠️ Order matters! Conditions are checked top to bottom.


	// --- Logical operators ---
	var role = "admin"        // string
	var hasPermission = false // bool

	if role == "admin" && hasPermission {
		// && means BOTH must be true
		fmt.Println("yes")
	}
	// ⚠️ Since hasPermission is false, this block will NOT execute


	// --- If with short statement ---
	if age2 := 15; age >= 18 {
		// age2 exists only inside this if-else block
		fmt.Println("person is an adult", age2)

	} else if age2 >= 12 {
		// age2 is still accessible here
		fmt.Println("person is teenager", age2)
	}
	// ⚠️ After this block, age2 does NOT exist


	// Go does NOT support ternary operator (? :)
}