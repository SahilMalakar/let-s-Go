package main

import "fmt"

/*
========================================
📌 GO NOTES – INTERFACES & SOLID (OCP)
========================================
*/

// -------------------------------------
// 🔹 1. Interface Definition
// -------------------------------------

type paymenter interface {
	pay(amt float32)
	refund(amt float32, account string)
}

// -------------------------------------
// 🔹 2. Concrete Implementation (Razorpay)
// -------------------------------------

type razorPay struct{}

func (r razorPay) pay(amt float32) {
	fmt.Println("Making payment using Razorpay:", amt)
}

func (r razorPay) refund(amt float32, account string) {
	fmt.Println("Refunding", amt, "to account:", account)
}

// -------------------------------------
// 🔹 3. Another Implementation (Stripe)
// -------------------------------------

type stripePay struct{}

func (s stripePay) pay(amt float32) {
	fmt.Println("Making payment using Stripe:", amt)
}

func (s stripePay) refund(amt float32, account string) {
	fmt.Println("Refunding", amt, "to account:", account)
}

// -------------------------------------
// 🔹 4. Payment Struct (Dependency Injection)
// -------------------------------------

type payment struct {
	gateway paymenter // interface dependency
}

// Method using abstraction
func (p payment) makePayment(amt float32) {
	p.gateway.pay(amt)
}

// -------------------------------------
// 🔹 5. Main Function
// -------------------------------------

func main() {

	// Inject Razorpay
	razor := razorPay{}
	p1 := payment{
		gateway: razor,
	}
	p1.makePayment(100)

	// Inject Stripe
	stripe := stripePay{}
	p2 := payment{
		gateway: stripe,
	}
	p2.makePayment(200)
}

/*
========================================
🧠 INTERVIEW NOTES (VERY IMPORTANT)
========================================

🔥 1. Interface in Go
   - Defines behavior (NOT data)
   - Any struct that implements methods → automatically satisfies interface

----------------------------------------

🔥 2. Dependency Injection
   - payment struct does NOT depend on concrete implementation
   - It depends on abstraction (paymenter)

----------------------------------------

🔥 3. Open/Closed Principle (OCP)

❌ BAD (violating OCP):
   if gateway == "razorpay" → call razorPay
   if gateway == "stripe" → call stripePay

👉 Requires modifying existing code → NOT scalable

✅ GOOD (your approach):
   - Add new gateway → just implement interface
   - No change in existing logic

----------------------------------------

🔥 4. Polymorphism in Go

Same method call:
   p.gateway.pay()

But behavior changes based on:
   razorPay / stripePay

----------------------------------------

🔥 5. Real Backend Relevance

Used in:
- Payment systems
- Logging systems
- Database adapters
- Notification services

========================================
💡 Hinglish Explanation (Easy Way)
========================================

- Interface = contract
  → "jo bhi ye methods implement karega wo valid hai"

- Tumne smart kaam kiya:
  → payment ko direct razorPay se link nahi kiya
  → interface use kiya

- Result:
  → future me Paytm, PhonePe add kar sakte ho
  → bina existing code change kiye

- Ye hi Open/Closed Principle hai:
  → code open for extension
  → closed for modification

========================================
*/