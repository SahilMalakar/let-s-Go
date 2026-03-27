package main

import (
	"fmt"
	"time"
)

/*
========================================
📌 GO NOTES – STRUCTS, METHODS, EMBEDDING
========================================
*/

// -------------------------------------
// 🔹 1. Struct (Basic)
// -------------------------------------

type customer struct {
	name  string
	phone string
}

// -------------------------------------
// 🔹 2. Struct Embedding (IMPORTANT 🔥)
// -------------------------------------

// customer is embedded inside order
type order struct {
	id        string
	amt       float32
	status    string
	createdAt time.Time
	customer  // 👈 embedded struct (composition)
}

// -------------------------------------
// 🔹 3. Constructor Pattern
// -------------------------------------

func newOrder(id string, amt float32, status string) *order {
	return &order{
		id:     id,
		amt:    amt,
		status: status,
	}
}

// -------------------------------------
// 🔹 4. Methods
// -------------------------------------

// Pointer receiver → modifies original
func (o *order) changeStatus(status string) {
	o.status = status
}

// Value receiver → works on copy
func (o order) getAmount() float32 {
	return o.amt
}

// -------------------------------------
// 🔹 5. Main Function
// -------------------------------------

func main() {

	// ---------------------------------
	// ✅ Creating Customer
	// ---------------------------------

	newCustomer := customer{
		name:  "sahil",
		phone: "1234567890",
	}

	// ---------------------------------
	// ✅ Struct with Embedded Struct
	// ---------------------------------

	newOrder2 := order{
		id:       "1",
		amt:      30.78,
		status:   "received",
		customer: newCustomer,
	}

	// ---------------------------------
	// 🔥 Accessing Embedded Fields
	// ---------------------------------

	// Full access
	newOrder2.customer.name = "sweety"

	// Promoted field access (Go feature)
	newOrder2.name = "rahul"

	fmt.Println("Full Order:", newOrder2)
	fmt.Println("Customer:", newOrder2.customer)

	// ---------------------------------
	// ✅ Zero Values Example
	// ---------------------------------

	partialOrder := order{
		id: "2",
	}

	fmt.Println("Partial Order (Zero Values):", partialOrder)

	// ---------------------------------
	// ✅ Constructor Usage
	// ---------------------------------

	newOrderObj := newOrder("3", 50.25, "pending")
	fmt.Println("New Order (Constructor):", newOrderObj)

	// ---------------------------------
	// ✅ Anonymous Struct
	// ---------------------------------

	language := struct {
		name   string
		isGood bool
	}{
		name:   "Golang",
		isGood: true,
	}

	fmt.Println("Anonymous Struct:", language)
}

/*
========================================
🧠 INTERVIEW NOTES (VERY IMPORTANT)
========================================

🔥 1. Struct Embedding (Composition)
   - Go uses composition instead of inheritance
   - Embedded struct fields are "promoted"

   Example:
   order.customer.name
   OR directly:
   order.name  ✅ (promoted field)

----------------------------------------

🔥 2. Pointer vs Value Receiver
   - Pointer → modifies original
   - Value → copy (read-only behavior)

----------------------------------------

🔥 3. Constructor Pattern
   - No constructors in Go
   - Use factory functions

----------------------------------------

🔥 4. Zero Values
   - Automatic initialization
   - No undefined/null bugs

----------------------------------------

🔥 5. Anonymous Struct
   - Used for quick temporary objects

========================================
💡 Hinglish Explanation (Easy Way)
========================================

- Go me inheritance nahi hota
  → instead "composition" use hota hai

- Tumne jo kiya:
  order ke andar customer daala
  → this is embedding

- Special feature:
  → order.customer.name ❌ nahi likhna padega
  → direct order.name bhi kaam karega

- Pointer receiver:
  → original object change karega

- Value receiver:
  → copy pe kaam karega

========================================
*/