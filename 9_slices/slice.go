package main 

import (
	"fmt"    
	"slices" // Utility package (Go 1.21+) for slice helpers like Equal
)

// Slice = dynamic, flexible view over an underlying array
func main() {

	// --- Nil slice ---
	var nums []int
	// Declared but not initialized → nil slice
	var a []int       // nil
    b := []int{}      // empty
	fmt.Println(a)
	fmt.Println(b)
	// Zero value of slice is nil

	fmt.Println(nums == nil) // true
	fmt.Println(len(nums))   // 0
	// ⚠️ nil slice has length 0 but is different from empty slice []int{}

	// --- make() function ---
	var num1 = make([]int, 2, 5)
	// make(type, length, capacity)
	// length = 2 (accessible elements)
	// capacity = 5 (max space before reallocation)

	fmt.Println(len(num1))   // 2
	fmt.Println(num1)        // [0 0]
	fmt.Println(num1 == nil) // false (it's initialized)
	fmt.Println(cap(num1))   // 5

	// --- append() ---
	num1 = append(num1, 2)
	num1 = append(num1, 7)
	num1 = append(num1, 9)
	num1 = append(num1, 3)
	num1 = append(num1, 134)
	num1 = append(num1, 8)
	num1 = append(num1, 2)
	// ⚠️ When capacity exceeds, Go allocates new bigger array automatically

	fmt.Println(num1)
	fmt.Println(cap(num1)) // Capacity grows (usually doubles)

	// --- Empty slice (not nil) ---
	num2 := []int{}
	// This is NOT nil, just empty

	num2 = append(num2, 2)
	num2 = append(num2, 9)

	fmt.Println(num2)
	fmt.Println(cap(num2))
	fmt.Println(len(num2))

	// --- copy() ---
	var num4 = make([]int, 0, 5)
	num4 = append(num4, 2)

	var num3 = make([]int, len(num4))
	copy(num3, num4)
	// copy(dst, src)

	fmt.Println(num4, num3)
	// ⚠️ copy creates independent slice (no shared backing array)

	// --- Slice operator ---
	var num5 = []int{4, 5, 3, 1, 7, 8}

	fmt.Println(num5[0:2]) // index 0 to 1
	fmt.Println(num5[:4])  // start to index 3
	fmt.Println(num5[2:5]) // index 2 to 4
	fmt.Println(num5[3:])  // index 3 to end
	fmt.Println(num5[:])   // full slice

	// ⚠️ Syntax: slice[start:end]
	// ⚠️ end is exclusive (not included)

	// --- Comparing slices ---
	sl1 := []int{1, 2}
	sl2 := []int{1, 2}

	fmt.Println(slices.Equal(sl1, sl2))
	// ⚠️ You CANNOT use sl1 == sl2 (only allowed for nil comparison)

	// --- 2D slice ---
	sl3 := [][]int{{8, 9, 4}, {2, 3, 4}}
	// Slice of slices (not fixed like 2D array)

	fmt.Println(sl3)

	// --- Slices Are Reference Types ---
	s1 := []int{1,2}
    s2 := s1
    s2[0] = 100
	fmt.Println(s1)
	fmt.Println(s2)
	//Now BOTH change because they share same backing array.
	//Slices share memory unless copied.
}