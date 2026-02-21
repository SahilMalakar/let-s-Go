package main

import "fmt"

func main(){
	//simple value
	
	fmt.Println(1+1) //int
	
	fmt.Println("hello buddy") //string
	
	fmt.Println(true) //bool
	
	fmt.Println(10.4) //float
	
	fmt.Println((7.0/3.0)) // 2.3
	// Extra brackets not required
	
	fmt.Println(7/3) // 2 
	// Both numbers are int
	// Go removes decimal part (no rounding)

	fmt.Println(1 + 2.5) // ❌ Error (int + float64 not allowed)
	
	//Are float64 by default, not float32.
	//This is important later when working with math functions.
}



// Go automatically infers basic types.

// Integer division behaves differently from float division.

// Types must match — Go is strict.

// main() is mandatory for executable programs