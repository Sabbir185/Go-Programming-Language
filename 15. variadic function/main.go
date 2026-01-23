package main

// Variadic functions in Go allow you to pass a variable number of arguments to a function.
// This feature is useful when you don’t know beforehand how many arguments you will pass.
// A variadic function accepts multiple arguments of the same type and can be called with any number of arguments, including none.

import "fmt"

// ...int accept same type arguments
// ...interface{} accept different type arguments
// Here, nums is a slice of integers. var nums []int

func sum(nums ...int) {
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println("Sum:", total)
}

func main() {
	// sum(1, 2, 3, 4, 5)
	numbers := []int{1, 2, 3, 4, 5}
	sum(numbers...) // Unpacking the slice using '...'
}
