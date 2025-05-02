package main

import "fmt"

type User struct {
	name string
	dept string // member variable or property
	age  int
}

func main() {
	sabbir := User{ // instantiate
		name: "Sabbir",
		dept: "CSE",
		age:  23,
	}

	fmt.Println(sabbir)

	motiur := User{ // Object or Instance
		name: "Motiur",
		dept: "EEE",
		age:  22,
	}

	fmt.Println(motiur.dept)
}
