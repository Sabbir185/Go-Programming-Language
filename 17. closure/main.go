package main

import "fmt"

// Go closure is a nested function that allows us to access variables of the outer function even after the outer function is closed.
// Nested Functions, and Anonymous Functions
// Returning a function

// --> Nested Functions
func greet() {
	name := "Sabbir"
	displayName := func() {
		fmt.Printf("Hi, %s\n", name)
	}
	displayName()
}

// --> Anonymous Functions and returning a function
func greetWithReturn(name string) func() {
	return func() {
		fmt.Printf("Hi, I am %s\n", name)
	}
}

// --> Closures
// A closure is a function where inner function has access to the outer function's variables.
// Even after the outer function has finished executing.
// The inner function is returned from the outer function and called later.
func addSum() func() {
	sum := 0
	return func() {
		sum += 1
		fmt.Println("Sum:", sum)
	}
}

func main() {
	greet()

	var mySelf func() = greetWithReturn("Sabbir")
	mySelf()

	result := addSum()
	result()
	result()
	result()

}
