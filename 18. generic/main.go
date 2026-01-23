package main

import "fmt"

// this func accepts any types of slice and prints its details
func printDetails[T any](items []T) {
	for i, item := range items {
		fmt.Println(i, " --> ", item)
	}
}

// this func accepts only slices of int or string types
func printStrInt[T int | string](items []T) {
	for i, item := range items {
		fmt.Println(i, " --> ", item)
	}
}

// interface{} can be used to accept any type, similar to 'any'
func printAny[T interface{}](items []T) {
	for i, item := range items {
		fmt.Println(i, " --> ", item)
	}
}

// this func accepts only slices of comparable types
// comparable হচ্ছে Go-এর একটি predeclared constraint
// যেসব type কে == বা != দিয়ে compare করা যায়,
// যেমন int, string, bool, struct, array, pointer
func printAnyWithComparable[T comparable](items []T) {
	for i, item := range items {
		fmt.Println(i, " --> ", item)
	}
}

// generic with struct
type Stacks[T any] struct {
	elements []T
}

func main() {
	ints := []int{1, 2, 3, 4, 5}
	strs := []string{"apple", "banana", "cherry"}
	bools := []bool{true, false, true}

	printDetails(strs)
	printDetails(ints)
	printDetails(bools)

	// ...................

	printStrInt(strs)
	printStrInt(ints)
	// printStrInt(bools) // this will give error as bool is not in the constraint

	// ...................

	printAny(strs)
	printAny(ints)
	printAny(bools)

	// ...................

	printAnyWithComparable(strs)
	printAnyWithComparable(ints)
	printAnyWithComparable(bools)

	// ...................
	// generic with struct

	my_stack := Stacks[int]{
		elements: []int{10, 20, 30, 40, 50},
	}

	my_stack2 := Stacks[string]{
		elements: []string{"ten", "twenty", "thirty", "forty", "fifty"},
	}

	fmt.Println("My Stack: ", my_stack.elements)
	fmt.Println("My Stack2: ", my_stack2.elements)
}
