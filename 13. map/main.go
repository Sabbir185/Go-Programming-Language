package main

import "fmt"

// map is a collection of a same type of data
// map[keyType]valueType
// maps are reference types

func printMap(c map[string]string) {
	c["red"] = "#ff00" // has changed the original map, due to reference type
	for key, value := range c {
		fmt.Println(key, value)
	}
}

func main() {
	// Map
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}

	for color, hex := range colors {
		fmt.Printf("The hex code for %s is %s\n", color, hex)
	}

	printMap(colors)

	// array
	numbers := [3]int{1, 2, 3}
	for ind, number := range numbers {
		fmt.Println(ind, number)
	}

	// built in function
	// var numbersMap map[int]string  --> nil map, will cause runtime error
	numbersMap := make(map[int]string)
	numbersMap[1] = "one"
	numbersMap[2] = "two"
	fmt.Println(numbersMap)

	delete(numbersMap, 1)
	fmt.Println(numbersMap)
}

// by iterating over a map, you can access both keys and values
// by iterating over an array, you can access both indices and values
// struct can not be iterated over directly
// built-in function make() can be used to create maps
// built-in function delete() can be used to remove key-value pairs from a map
