package main

import "fmt"

func evenOdd() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, num := range arr {
		if num%2 == 0 {
			fmt.Printf("The number is event %d", num)
		} else {
			fmt.Printf("The number is odd %d", num)
		}
		fmt.Println()
	}
}