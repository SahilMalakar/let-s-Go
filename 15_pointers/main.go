package main

import "fmt"

/*
========================================
🔹 1. PASS BY VALUE (DEFAULT IN GO)
========================================

✔ Go always passes arguments by value
✔ Means: copy of variable is passed

❌ Original variable does NOT change
*/

// func changeNum(num int) {
// 	num = 5
// 	fmt.Println("Inside changeNum:", num)
// }

/*
========================================
🔹 2. PASS BY REFERENCE (USING POINTER)
========================================

✔ Pointer = variable that stores memory address
✔ "*" → dereference (access actual value)
✔ "&" → get address of variable

✔ This allows modifying original variable
*/

func changeNum(num *int) {
	*num = 5 // dereferencing pointer
	fmt.Println("Inside changeNum:", *num)
}

/*
========================================
🔹 3. MAIN FUNCTION
========================================
*/

func main() {

	num := 1

	/*
	========================================
	🔹 MEMORY ADDRESS
	========================================

	✔ &num → gives memory address
	*/

	// fmt.Println("Address of num:", &num)

	/*
	========================================
	🔹 PASSING POINTER
	========================================

	✔ &num → passing address
	✔ changeNum modifies original value
	*/

	changeNum(&num)

	fmt.Println("After changeNum in main:", num)
}

/*
========================================
🔹 4. OUTPUT
========================================

Inside changeNum: 5
After changeNum in main: 5

✔ Original variable changed
*/

/*
========================================
🔹 5. IMPORTANT SYMBOLS
========================================

| Symbol | Meaning |
|--------|--------|
| &      | address of variable |
| *      | pointer / dereference |

Example:
num := 10
ptr := &num   // pointer stores address
*ptr = 20     // changes num to 20
*/

/*
========================================
🔹 6. VALUE vs POINTER COMPARISON
========================================

| Feature          | Pass by Value | Pointer (Reference-like) |
|------------------|--------------|--------------------------|
| Memory           | Copy         | Address                  |
| Original Change  | ❌ No        | ✅ Yes                   |
| Performance      | Slower (copy)| Faster (no copy)         |
| Safety           | Safer        | Risky if misused         |

*/

/*
========================================
🔹 7. WHEN TO USE POINTERS
========================================

✔ When you want to modify original value
✔ When struct is large (avoid copying)
✔ For performance optimization
✔ In APIs / DB operations

*/

/*
========================================
🔹 8. COMMON MISTAKES
========================================

❌ Forgetting "*" while modifying
	num = 5        // WRONG
	*num = 5       // CORRECT

❌ Forgetting "&" while passing
	changeNum(num)   // WRONG
	changeNum(&num)  // CORRECT

❌ Confusing pointer with reference
	Go has NO direct reference like C++
*/

/*
========================================
🔹 9. INTERVIEW QUESTIONS
========================================

Q1: Does Go support pass by reference?
✔ No, Go is always pass by value
✔ But pointers simulate reference behavior

Q2: Why use pointers?
✔ Avoid copying large data
✔ Modify original variable

Q3: What is zero value of pointer?
✔ nil

*/

/*
========================================
🔹 10. QUICK SUMMARY
========================================

✔ Go = pass by value
✔ Pointer = stores address
✔ "*" → access value
✔ "&" → get address
✔ Use pointer to modify original variable

========================================
*/