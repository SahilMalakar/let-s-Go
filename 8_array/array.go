package main 
import "fmt"

// Array = fixed-size sequence of elements of SAME type
func main() {

	var nums [4]int 
	// Declares an array of size 4
	// Default value = [0 0 0 0] (zero value for int)

	fmt.Println(len(nums)) 
	// len() gives array length (fixed at compile time)

	nums[0] = 1 // Arrays are index-based (0 starts)

	fmt.Println(nums[0]) 
	fmt.Println(nums)    
	// Prints entire array

	var boolean [4]bool 
	// Default value = [false false false false]

	boolean[2] = true
	fmt.Println(boolean)

	var str [4]string
	// Default value = ["" "" "" ""] (empty string is zero value)

	str[1] = "hey buddy"
	fmt.Println(str)

	num1 := [3]int{1, 4, 7}
	// Array literal (size must match number of elements)
	fmt.Println(num1)

	// 2D array (matrix)
	num2 := [2][2]int{{3, 4}, {8, 5}}
	// 2 rows, 2 columns
	fmt.Println(num2)

	// Arrays in Go:
	// - Fixed size (cannot grow/shrink)
	// - Stored contiguously in memory
	// - Constant time access O(1)
}