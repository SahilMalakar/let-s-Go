package main

import "fmt"

// 🔁 RANGE IN GO (Iteration)

// range = used to iterate over slices, arrays, maps, strings

func main() {

	// =========================================================
	// 🔹 1. RANGE WITH SLICE
	// =========================================================

	nums := []int{6, 7, 8, 9}

	// --- Using index only ---
	for i := range nums {
		fmt.Println(nums[i])
		// i → index
		// need nums[i] to get value
	}

	// --- Using index + value ---
	for i, num := range nums {
		fmt.Println(i, num)
		// i → index
		// num → value (direct)
	}

	// --- Ignoring index ---
	for _, num := range nums {
		fmt.Println(num)
		// _ → ignore index
	}

	// --- Sum example ---
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// =========================================================
	// 🔹 2. RANGE WITH MAP
	// =========================================================

	m := map[string]string{
		"firstName": "Sahil",
		"lastName":  "Malakar",
	}

	// --- Key + Value ---
	for k, v := range m {
		fmt.Println(k, v)
		// k → key
		// v → value
	}

	// --- Only values ---
	for _, v := range m {
		fmt.Println(v)
	}

	// ⚠️ Maps are unordered → output order changes every run

	// =========================================================
	// 🔹 3. RANGE WITH STRING
	// =========================================================

	for i, c := range "Golang" {
		fmt.Println(i, c, string(c))
		// i → byte index
		// c → rune (unicode code point)
		// string(c) → actual character
	}

	// =========================================================
	// 🔥 IMPORTANT CONCEPTS
	// =========================================================

	// --- 1. Values are COPIED ---
	for _, v := range nums {
		v = 100
		// Does NOT modify original slice
		fmt.Println(v)
	}
	fmt.Println("nums after copy loop:", nums)

	// --- 2. To modify slice → use index ---
	for i := range nums {
		nums[i] = 100
	}
	fmt.Println("nums after modification:", nums)

	// --- 3. Range on nil slice/map ---
	var a []int
	for range a {
		// Will NOT run, but NO panic
	}

	var b map[string]int
	for range b {
		// Safe, no panic
	}

	// =========================================================
	// 🔹 RANGE VS FOR LOOP
	// =========================================================

	// --- Traditional for loop ---
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

	// --- Range loop ---
	for _, v := range nums {
		fmt.Println(v)
	}

	// =========================================================
	// 🔹 STRING UTF-8 BEHAVIOR
	// =========================================================

	for i, c := range "Go❤️" {
		fmt.Println(i, c, string(c))
		// ❤️ takes multiple bytes → index jumps
	}

	// =========================================================
	// 🔥 SUMMARY
	// =========================================================

	// range:
	// - Iterates over slices, maps, strings
	// - Returns index/key and value
	// - Use _ to ignore variables
	// - Values are copied (not directly mutable)
	// - Maps are unordered
	// - Strings return runes (Unicode)

	// =========================================================
	// 🧠 HINGLISH QUICK NOTES
	// =========================================================

	// range = shortcut loop likhne ka
	// slice → index + value milta hai
	// map → key + value milta hai
	// string → unicode characters milte hai
	// value copy hota hai → original change nahi hota
	// change karna hai → index use karo
	// map ka order kabhi fix nahi hota
}