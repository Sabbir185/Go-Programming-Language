package main

import "fmt"

type PaymentGateway interface {
	Pay(amount float32, currence string) (string, error)
	Refund(paymentId string) error
}

// No need "implement" or "extends" keyword for any base class
// Just implement the interface methods, following the same method signatures

// -- Stripe ---
type Stripe struct{}

func (s Stripe) Pay(amount float32, currence string) (string, error) {
	return fmt.Sprintf("Payment Success: %.2f %s using Stripe\n", amount, currence), nil
}

func (s Stripe) Refund(paymentId string) error {
	fmt.Printf("Refund Success for payment ID: %s using Stripe\n", paymentId)
	return nil
}

// -- Razorpay --
type Razorpay struct{}

func (r Razorpay) Pay(amount float32, currence string) (string, error) {
	return fmt.Sprintf("Razorpay success: %.2f %s\n", amount, currence), nil
}

func (r Razorpay) Refund(paymentId string) error {
	fmt.Printf("Razorpay refund success for payment ID: %s\n", paymentId)
	return nil
}

// Note: Stripe আর Razorpay -> interface ব্যবহার করছে, কিন্তু নিজেদের মতো করে logic লিখছে

// -- Common function to process payment --
func ProcessPayment(pg PaymentGateway, amount float32, currence string) {
	result, _ := pg.Pay(amount, currence)
	fmt.Println(result)

	pg.Refund("PAY123")
}

func main() {
	stripe := Stripe{}
	razorpay := Razorpay{}

	ProcessPayment(stripe, 100.0, "USD")
	ProcessPayment(razorpay, 200.0, "INR")
}
