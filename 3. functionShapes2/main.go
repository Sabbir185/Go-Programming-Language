package main

import "fmt"

var num = 10

func main() {
	fmt.Println("Hello World: ", num)
}

// called before main function
func init() {
	fmt.Println("init function: ", num)
	num = 100

	// anonymous function and IIFE
	func(a int, b int) {
		c := a + b
		fmt.Println("Anonymous function: ", c)
	}(2, 3)

	// Function expression or Assign function in variable
	// doSum(2, 3) --> this is not possible to call before declaration
	doSum := func(a int, b int) {
		c := a + b
		fmt.Println("Anonymous function: ", c)
	}
	doSum(2, 3)
}

// standard or named function
// init function --> you can not call it, computer calls this automatically, no params, no return
// Anonymous function --> no name
// IIFE --> Immediately Invoked Function Expression
// Function expression or Assign function in variable
