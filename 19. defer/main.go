package main

import "fmt"

// defer statements delay the execution of the function or method or an anonymous method until the nearby functions returns.
// In other words, defer function or method call arguments evaluate instantly,
// but they don't execute until the nearby functions returns.

func add(a1, a2 int) int {
	res := a1 + a2
	fmt.Println("Result: ", res)
	return 0
}

func main() {
	fmt.Println("Start...")

	// Multiple defer statements
	// Executes in LIFO order
	defer fmt.Println("End")
	defer add(1, 3)
	defer add(2, 3)

}
