package main

import "fmt"

// Higher order functions
// function o receive kore alongside other data types
// function as a parameter
// function as a return type
// call back function (op)
func processOperation(a int, b int, op func(p int, q int)) {
	op(a, b)
}

// function as a return -> HOF
func call() func(a int, b int) {
	return add
}

func add(a int, b int) {
	fmt.Println(a + b)
}

func main() {
	fmt.Println("Hello, world!")
	processOperation(1, 2, add)

	sum := call()
	sum(2, 3)
}

/*

	1. Parameters vs arguments
	2. First order functions
		i. standard or named function
		ii. anonymous function
		iii. IIFE
		iv. function expression
	3. Higher order functions or first class functions(treat ed as first class citizen)
	4. Call back functions
	5. First class citizen => random variable assigned value -> a:=10/true/"hello"/funcs

	Function paradigm --> haskel, racket
	math -> logic -> discrete math
	1. first order logic -> object, property, relation
	2. higher order logic -> different kinds of rules, logics, functions

*/
