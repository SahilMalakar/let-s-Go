package main

import "fmt"

/*
========================================
📌 GO NOTES – GENERICS (TYPE PARAMETERS)
========================================
*/

// -------------------------------------
// 🔹 1. Basic Generic Function
// -------------------------------------

// T = any type
func printSlice[T any](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

// -------------------------------------
// 🔹 2. Generic with Multiple Types
// -------------------------------------

func printSliceWithLabel[T any, V any](items []T, label V) {
	for _, item := range items {
		fmt.Println(item, label)
	}
}

// -------------------------------------
// 🔹 3. Type Constraints
// -------------------------------------

// Only allow specific types
func printLimited[T int | string | bool](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

// comparable → supports ==, !=
func checkEqual[T comparable](a, b T) bool {
	return a == b
}

// -------------------------------------
// 🔹 4. Generic Struct (Stack Example)
// -------------------------------------

// LIFO (Last In First Out)
type stack[T any] struct {
	elements []T
}

func (s *stack[T]) push(item T) {
	s.elements = append(s.elements, item)
}

func (s *stack[T]) pop() T {
	if len(s.elements) == 0 {
		var zero T // zero value of type T
		return zero
	}
	last := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return last
}

// -------------------------------------
// 🔹 5. Main Function
// -------------------------------------

func main() {

	// ---------------------------------
	// ✅ Using Generic Function
	// ---------------------------------

	nums := []int{1, 2, 3, 4}
	names := []string{"sahil", "sweety"}

	printSlice(nums)
	printSlice(names)

	// ---------------------------------
	// ✅ Generic with Label
	// ---------------------------------

	printSliceWithLabel(names, "miss you")

	// ---------------------------------
	// ✅ Type Constraint Function
	// ---------------------------------

	vals := []bool{true, false, true}
	printLimited(vals)

	// ---------------------------------
	// ✅ Comparable Example
	// ---------------------------------

	fmt.Println("Are equal:", checkEqual(10, 10))

	// ---------------------------------
	// ✅ Generic Stack
	// ---------------------------------

	myStack := stack[string]{}

	myStack.push("sahil")
	myStack.push("sweety")

	fmt.Println("Pop:", myStack.pop())
	fmt.Println("Pop:", myStack.pop())
}

/*
========================================
🧠 INTERVIEW NOTES (VERY IMPORTANT)
========================================

🔥 1. Generics introduced in Go 1.18+

----------------------------------------

🔥 2. Syntax

func functionName[T any](param T)

T → type parameter

----------------------------------------

🔥 3. any vs interface{}

any = alias of interface{}
(preferred modern syntax)

----------------------------------------

🔥 4. Type Constraints

T comparable → supports ==, !=

T int | string → union types

----------------------------------------

🔥 5. Why Generics?

❌ Without generics:
   printIntSlice
   printStringSlice

👉 Code duplication

✅ With generics:
   One reusable function

----------------------------------------

🔥 6. Generic Structs

Used in:
- Stack
- Queue
- Cache systems
- Repositories

----------------------------------------

🔥 7. Zero Value Trick

var zero T

👉 important for generic return

========================================
💡 Hinglish Explanation (Easy Way)
========================================

- Generics ka matlab:
  → ek hi function multiple type ke liye kaam kare

- Pehle:
  int ke liye alag function
  string ke liye alag function

- Ab:
  ek hi function sab handle karega

- T any:
  → koi bhi type allowed

- comparable:
  → == use kar sakte ho

- Stack example:
  → real backend me bahut use hota hai

========================================
*/