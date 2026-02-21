package main 

import (
	"fmt" 
	"time" // For working with time/date
)

func main() {

	// --- Simple switch ---
	i := 5

	switch i { // Compares value of i with each case
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default: // Runs if no case matches
		fmt.Println("other")
	}
	// ⚠️ No 'break' needed in Go switch (it stops automatically)
	// ⚠️ Go does NOT fall through by default (unlike C/C++)


	// --- Multiple condition switch ---
	switch time.Now().Weekday() { 
	// time.Now() → current time
	// Weekday() → returns type time.Weekday (enum-like)

	case time.Saturday, time.Sunday:
		// Multiple values in one case
		fmt.Println("it is weekend")

	default:
		fmt.Println("it is workday")
	}
	// ⚠️ time.Saturday is a constant from time package
	// ⚠️ Switch can work with expressions, not just variables


	// --- Type switch ---
	whoAmI := func(i interface{}) {
		// interface{} means: can accept ANY type (empty interface)

		switch t := i.(type) {
		// Type assertion switch
		// t becomes the actual value with its concrete type

		case int:
			fmt.Println("its an integer")

		case string:
			fmt.Println("its a string")

		case bool:
			fmt.Println("its a boolean")

		default:
			fmt.Println("other", t)
		}
	}

	// Go uses double quotes for strings
	whoAmI(true)

	
    // 	Go uses:
    // "double quotes" → string
    // 'single quotes' → rune (single character, Unicode)
}