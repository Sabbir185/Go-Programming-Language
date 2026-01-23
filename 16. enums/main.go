package main

import "fmt"

// integer enums
type OrderStatusInt int

const (
	Pending OrderStatusInt = iota
	Shipped
	Delivered
	Canceled
)

// string enums
type OrderStatusStr string

const (
	PendingStr   OrderStatusStr = "PENDING"
	ShippedStr   OrderStatusStr = "SHIPPED"
	DeliveredStr OrderStatusStr = "DELIVERED"
	CanceledStr  OrderStatusStr = "CANCELED"
)

func orderStatusInt(status OrderStatusInt) {
	fmt.Println("Order Status:", status)
}

func orderStatusStr(status OrderStatusStr) {
	fmt.Println("Order Status:", status)
}

func main() {
	orderStatusInt(Pending)
	orderStatusStr(DeliveredStr)
}
