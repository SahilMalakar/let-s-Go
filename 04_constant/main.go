package main 

import "fmt"

const age = 34 
// Global constant → available everywhere in this file
// Constants MUST be known at compile time

var name2 = "sahil" 
// Global variable → accessible inside main()
// Type inferred as string

func main() {

	const name = "golang" 
	// Local constant → scope limited to main()
	// Cannot be reassigned after declaration

	fmt.Println(age)   // Accessing global constant
	fmt.Println(name)  // Accessing local constant
	fmt.Println(name2) // Accessing global variable

	const (
		port = 5000        // Untyped constant (acts like int)
		host = "localhost" // Untyped string constant
	)
	// Grouped constant declaration (cleaner for related values)

	fmt.Println(port, host)
	// Multiple values printed in one line

	const x = 10 + 5 
	// Constant expressions are evaluated at compile time

	// const x = someFunction()
	// ❌ Function runs at runtime
}