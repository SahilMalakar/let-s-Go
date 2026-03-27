package main

import "fmt"

/*
========================================
📌 GO NOTES – ENUMS (TYPED CONSTANTS)
========================================
*/

// -------------------------------------
// 🔹 1. Enum Type Definition
// -------------------------------------

// Custom type (strong typing)
type OrderStatus string

// -------------------------------------
// 🔹 2. Enum Values (Constants)
// -------------------------------------

const (
	Received  OrderStatus = "received"
	Confirmed OrderStatus = "confirmed"
	Prepared  OrderStatus = "prepared"
	Delivered OrderStatus = "delivered"
)

// -------------------------------------
// 🔹 3. Function Using Enum
// -------------------------------------

func changeOrderStatus(status OrderStatus) {
	fmt.Println("Changing order status to:", status)
}

// -------------------------------------
// 🔹 4. Main Function
// -------------------------------------

func main() {

	// Valid usage
	changeOrderStatus(Delivered)

	// ❌ Invalid (uncomment → compile-time safety)
	// changeOrderStatus("random_status")

}

/*
========================================
🧠 INTERVIEW NOTES (VERY IMPORTANT)
========================================

🔥 1. Go does NOT have real enums
   - Instead uses "typed constants"

----------------------------------------

🔥 2. Why use custom type?

type OrderStatus string

👉 Gives type safety:
   changeOrderStatus("hello") ❌ (not allowed)

----------------------------------------

🔥 3. Why string enum (BEST PRACTICE)?

✅ Human readable
✅ Easy for JSON / APIs / DB
✅ Debugging friendly

Example:
   "delivered" instead of 3

----------------------------------------

🔥 4. Alternative → iota (int-based enum)

type OrderStatus int

const (
	Received OrderStatus = iota
	Confirmed
	Prepared
	Delivered
)

👉 Values:
Received = 0
Confirmed = 1
...

----------------------------------------

🔥 5. When to use what?

| Type   | Use Case |
|--------|--------|
| string | APIs, DB, logs (recommended) |
| int    | performance-critical systems |

----------------------------------------

🔥 6. Real Backend Usage

Used in:
- Order status
- Payment status
- User roles
- Job states

========================================
💡 Hinglish Explanation (Easy Way)
========================================

- Go me direct enum nahi hota
  → hum custom type + const use karte hain

- Tumne jo kiya:
  type OrderStatus string
  → best approach hai backend ke liye

- Fayda:
  → galat value pass nahi kar sakte
  → readable hai (debug easy)

- Agar iota use karte:
  → fast hota but readable nahi

========================================
*/