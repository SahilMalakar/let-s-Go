package main

import (
	"fmt"
	"sync"
)

/*
========================================
📌 GO NOTES – CONCURRENCY (GOROUTINES)
========================================
*/

// -------------------------------------
// 🔹 1. Goroutine Task Function
// -------------------------------------

func task(id int, wg *sync.WaitGroup) {
	defer wg.Done() // IMPORTANT: signals completion

	fmt.Println("Doing task:", id)
}

// -------------------------------------
// 🔹 2. Main Function
// -------------------------------------

func main() {

	fmt.Println("Before goroutines execution")

	// ---------------------------------
	// 🔹 3. WaitGroup (Synchronization)
// ---------------------------------

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)           // register task
		go task(i, &wg)     // run concurrently
	}

	wg.Wait() // block until all goroutines finish

	fmt.Println("After goroutines execution")

	// ---------------------------------
	// 🔹 4. Closure Gotcha (IMPORTANT)
// ---------------------------------

	// ❌ Wrong (captures same i)
	/*
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println(i) // unpredictable output
		}()
	}
	*/

	// ✅ Correct (pass i as parameter)
	/*
	for i := 0; i < 5; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	*/
}

/*
========================================
🧠 INTERVIEW NOTES (VERY IMPORTANT)
========================================

🔥 1. Goroutine
   - Lightweight thread managed by Go runtime
   - Starts with keyword: go

   Example:
   go task()

----------------------------------------

🔥 2. WaitGroup (sync package)

Used to wait for multiple goroutines

Methods:
- Add(n)   → number of tasks
- Done()   → task finished
- Wait()   → block until all done

----------------------------------------

🔥 3. Why WaitGroup?

Without it:
main() exits before goroutines finish ❌

----------------------------------------

🔥 4. Closure Bug (VERY COMMON 🔥)

Wrong:
   go func() { fmt.Println(i) }()

👉 All goroutines may print same value

Fix:
   pass i as argument

----------------------------------------

🔥 5. Execution Order

❌ NOT guaranteed:
goroutines run asynchronously

----------------------------------------

🔥 6. Real Backend Usage

Used in:
- API parallel calls
- Worker pools
- Background jobs
- File processing
- Microservices communication

========================================
💡 Hinglish Explanation (Easy Way)
========================================

- Goroutine = lightweight thread
  → "go" keyword lagao aur kaam parallel ho jayega

- Problem:
  → main function jaldi exit ho jata hai

- Solution:
  → WaitGroup use karo

- wg.Add(1)
  → ek task add

- defer wg.Done()
  → task khatam hone pe signal

- wg.Wait()
  → sab tasks khatam hone ka wait

----------------------------------------

🔥 Closure bug:
- loop ka i directly use mat karo
- warna sab same value print karega

========================================
*/