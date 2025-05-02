package main

import "fmt"

func add(a float32, b float32) {
	c := a + b
	fmt.Println("Sum : ", c)
}

func sub(b float32, a float32) float32 {
	c := b - a
	return c
}

func getResults(a int, b int) (int, int) {
	sum := a + b
	multi := a * b
	return sum, multi
}

func main() {
	num := 7
	if num >= 18 {
		fmt.Println("Eligiable to married")
	} else if num < 18 && num >= 15 {
		fmt.Println("Not eligible")
	} else {
		fmt.Println("Sorry, baby!!")
	}

	// switch
	a := 3
	switch a {
	case 1:
		{
			fmt.Println("One")
		}
	case 2, 3:
		{
			a = a + a
			fmt.Println("2 or 3 : ", a)
		}
	default:
		{
			fmt.Println("Zero")
		}
	}

	// function calling with passing args
	add(3.5, 2.4)

	// function returning the result
	sub_result := sub(20.5, 10.0)
	fmt.Println("Sub : ", sub_result)

	// return multiple result
	sum, multi := getResults(2, 3)
	fmt.Printf("Sum is = %d, multi is = %d\n", sum, multi)
}
