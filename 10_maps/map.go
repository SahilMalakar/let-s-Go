package main

import (
	"fmt"  
	"maps" // Go 1.21+ utility package for comparing maps
)

// Map = built-in reference type (hash table / dictionary)
func main() {

	// --- Creating a map ---
	m := make(map[string]string)
	// make(map[keyType]valueType)
	// MUST initialize before use
	// Default value of map is nil (cannot assign to nil map)

	// --- Setting elements ---
	m["name"] = "golang"
	m["area"] = "backend"
	// Key must be unique
	// Assigning same key again overwrites value

	// --- Getting elements ---
	fmt.Println(m["name"], m["area"])
	// If key does not exist → returns zero value of value type

	// fmt.Println(m["phoneNo"])
	// Would print "" (empty string) — NOT an error

	// --- Map with int values ---
	n := make(map[string]int)
	n["age"] = 25

	fmt.Println(n["phoneNo"])
	// "phoneNo" doesn't exist → returns 0 (zero value of int)

	fmt.Println(len(m)) 
	// len() gives number of key-value pairs

	// --- delete() ---
	delete(m, "name")
	// Safe even if key doesn't exist
	fmt.Println(m)

	// --- clear() ---
	clear(m)
	// Removes all key-value pairs (Go 1.21+)
	fmt.Println(m)

	// --- Map literal ---
	s := map[string]int{"price": 39, "phone": 23}
	fmt.Println(s)

	// --- Checking key existence ---
	value, ok := s["price"]
	// ok = true if key exists
	// This is the SAFE way to check key presence

	fmt.Println(value)

	if ok {
		fmt.Println("all ok")
	} else {
		fmt.Println("not ok")
	}

	// --- Comparing maps ---
	t := map[string]int{"price": 39, "phone": 23}

	fmt.Println(maps.Equal(s, t))
	// Cannot use s == t (only allowed for nil comparison)


	a := map[string]int{"x":1}
    b := a
    b["x"] = 100
	fmt.Println(a)
	fmt.Println(b)
	// Now BOTH a and b change.
    // Because maps do NOT copy data on assignment.

    // When printing map:
	// Order may change every run.
    // Maps are unordered.
}