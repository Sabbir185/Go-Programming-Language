package main

import "fmt"

type Person struct {
	Name     string
	Age      int
	Address  string
	favFoods []string
}

func main() {

	sabbir := Person{
		Name:     "Sabbir",
		Age:      24,
		Address:  "Dhaka",
		favFoods: []string{"Burger", "Pizza", "Ice Cream"},
	}

	fmt.Println(&sabbir)

}
