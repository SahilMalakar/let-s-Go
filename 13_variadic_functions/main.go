package main

import "fmt"

/*
========================================
🔹 1. Variadic Function Definition
========================================

func sum(nums ...int) int

✔ "..." means variable number of arguments
✔ Accepts 0, 1, or many values
✔ Internally treated as a slice: []int

Examples:
sum(1,2,3)
sum(10,20)
sum()
*/

func sum(nums ...int) int {

    /*
    ========================================
    🔹 2. Internal Working
    ========================================

    ✔ nums behaves like: []int
    ✔ range loop iterates over slice
    ✔ "_" ignores index
    ✔ "num" is the actual value
    */

    total := 0

    for _, num := range nums {
        total = total + num
    }

    /*
    ========================================
    🔹 3. Return Value
    ========================================

    ✔ Returns sum of all elements
    ✔ Return type: int
    */

    return total
}

/*
========================================
🔹 4. Passing Slice to Variadic Function
========================================

nums := []int{3,5,6,8,1,9,10,15}

✔ To pass slice → use "..."
✔ nums... converts slice into arguments
*/

func main() {

    nums := []int{3, 5, 6, 8, 1, 9, 10, 15}

    result := sum(nums...) // slice unpacking

    fmt.Println(result)
}

/*
========================================
🔹 5. Important Concepts
========================================

✔ Variadic parameter must be LAST parameter
✔ Only ONE variadic parameter allowed
✔ Inside function → treated as slice
✔ Can pass:
    - individual values → sum(1,2,3)
    - slice → sum(nums...)

========================================
🔹 6. Why interface{} version is wrong?
========================================

func sum(nums ...interface{}) int

❌ Problem:
- interface{} can hold ANY type
- Cannot directly do: total + num
- Requires type assertion:

    num.(int)

✔ Correct usage:
for _, num := range nums {
    total += num.(int)
}

⚠ Risk: panic if type is not int

========================================
🔹 7. Common Mistakes
========================================

❌ Forgetting "..." when passing slice
    sum(nums)     // ERROR
    sum(nums...)  // CORRECT

❌ Mixing types without interface handling

❌ Using variadic parameter not at last position

========================================
🔹 8. Real-world Use Cases
========================================

✔ Logging functions
✔ Sum / aggregation utilities
✔ Flexible APIs (unknown number of inputs)

========================================
🔹 9. Quick Summary (Revision)
========================================

✔ ... → variadic (multiple inputs)
✔ Internally → slice ([]type)
✔ Use nums... to unpack slice
✔ Only last parameter can be variadic
✔ interface{} needs type assertion

========================================
*/