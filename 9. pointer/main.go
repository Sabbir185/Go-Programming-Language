package main

import "fmt"

// pass by value -> copy -> send value, receive value
// pass by reference -> send adress, get value through pointer

func print(nums *[3]int) { // pointer -> reference value of passing address
	fmt.Println(nums)
}

func main() {
	// pointer or address of memory (RAM)
	num := 10
	println(num)

	// & (ampersand) -> address of something
	// direact variable diye amra, value access korte pari
	// & -> diye oi variable er address access korte pari
	println(&num)

	// * (asterisk) -> value of something
	fmt.Println(*&num)

	//----------------

	y := 20
	fmt.Println("Y = ", y)
	p := &y // address of y
	*p = 30 // put value in address of y
	fmt.Println("Y = ", y)

	//----------------

	s := [3]int{1, 2, 3}
	print(&s) // passing address of s
}
