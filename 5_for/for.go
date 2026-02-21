package main

import "fmt"

// In Go, 'for' is the ONLY looping construct (no separate while/do-while)

func main() {

	// --- While-style loop ---
	i := 1 // Short declaration (int by default)

	for i <= 10 { // Works like a while loop
		fmt.Println(i)
		i++ // Important: no i++ inside Println; Go doesn't allow ++ in expressions
	}
	// ⚠️ If you forget i++, this becomes an infinite loop

	// --- Infinite loop ---
	for { // No condition → runs forever
		fmt.Println("infinite loop")
		break // Break exits the nearest loop immediately
	}
	// ⚠️ Used in servers, listeners, background workers

	// --- Classic for loop ---
	for i := 0; i < 5; i++ { // Initialization ; Condition ; Post statement

		if i == 2 {
			continue // Skips current iteration (does NOT stop loop)
		}

		fmt.Println(i)
	}
	// ⚠️ 'continue' jumps to next iteration
	// ⚠️ New 'i' here is scoped only inside this loop

	// --- Range loop (Go 1.22+ feature) ---
	for i := range 6 { // Loops from 0 to 5
		fmt.Print(i)
	}
	// ⚠️ range 6 means 0..(6-1)
	// ⚠️ Before Go 1.22, this syntax was not supported



	
	i++ // valid
    // ++i // ❌ invalid
	// fmt.Println(i++) // ❌ not allowed
}