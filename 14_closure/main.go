package main

import "fmt"

/*
========================================
🔹 1. BASIC CLOSURE (COUNTER)
========================================

✔ Remembers state (count)
*/

func counter() func() int {
	count := 0

	return func() int {
		count++
		return count
	}
}

/*
========================================
🔹 2. MULTIPLE INSTANCES
========================================

✔ Each closure has its own memory
*/

func multiCounter() func() int {
	count := 0

	return func() int {
		count++
		return count
	}
}

/*
========================================
🔹 3. ADDER (FUNCTION FACTORY)
========================================

✔ Outer variable "x" remembered
*/

func adder(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

/*
========================================
🔹 4. MULTIPLIER (CUSTOM FUNCTION)
========================================

✔ Creates customized functions
*/

func multiplier(factor int) func(int) int {
	return func(num int) int {
		return num * factor
	}
}

/*
========================================
🔹 5. LOGGER (STATEFUL FUNCTION)
========================================

✔ Keeps track of how many times called
*/

func logger() func(string) {
	count := 0

	return func(msg string) {
		count++
		fmt.Println("Log", count, ":", msg)
	}
}

/*
========================================
🔹 6. ACCUMULATOR (SUM TRACKER)
========================================

✔ Maintains running total
*/

func accumulator() func(int) int {
	total := 0

	return func(val int) int {
		total += val
		return total
	}
}

/*
========================================
🔹 MAIN FUNCTION (TEST ALL)
========================================
*/

func main() {

	fmt.Println("===== 1. Counter =====")
	inc := counter()
	fmt.Println(inc()) // 1
	fmt.Println(inc()) // 2
	fmt.Println(inc()) // 3

	fmt.Println("\n===== 2. Multiple Instances =====")
	inc1 := multiCounter()
	inc2 := multiCounter()

	fmt.Println(inc1()) // 1
	fmt.Println(inc1()) // 2

	fmt.Println(inc2()) // 1 (separate state)
	fmt.Println(inc2()) // 2

	fmt.Println("\n===== 3. Adder =====")
	add5 := adder(5)
	fmt.Println(add5(3))  // 8
	fmt.Println(add5(10)) // 15

	fmt.Println("\n===== 4. Multiplier =====")
	double := multiplier(2)
	triple := multiplier(3)

	fmt.Println(double(5)) // 10
	fmt.Println(triple(5)) // 15

	fmt.Println("\n===== 5. Logger =====")
	log := logger()
	log("Start")
	log("Processing")
	log("Done")

	fmt.Println("\n===== 6. Accumulator =====")
	acc := accumulator()
	fmt.Println(acc(10)) // 10
	fmt.Println(acc(5))  // 15
	fmt.Println(acc(20)) // 35
}

/*
========================================
🔹 FINAL SUMMARY
========================================

✔ Closure = Function + captured variable
✔ Variables persist after function ends
✔ Each function call creates new closure instance
✔ Used for:
   - State management
   - Function factories
   - Logging
   - Caching
   - Middleware

========================================
🔹 COMMON INTERVIEW POINTS
========================================

✔ Stored in heap (not stack)
✔ Not thread-safe by default
✔ Avoid overuse (can increase memory usage)

========================================
*/