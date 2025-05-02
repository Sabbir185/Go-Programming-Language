package main

import "fmt"

type User struct {
	name string
	dept string // member variable or property
	age  int
}

// only for custom type
func (usr User) getDept() {
	fmt.Println(usr.name)
	fmt.Println(usr.dept)
	fmt.Println(usr.age)
}

func (usr User) getSub(sub string) {
	fmt.Println(usr.name)
	fmt.Println(sub)
}

func main() {
	sabbir := User{ // instantiate
		name: "Sabbir",
		dept: "CSE",
		age:  23,
	}
	sabbir.getDept()
	sabbir.getSub("CSE 201")

	motiur := User{
		name: "Motiur",
		dept: "EEE",
		age:  22,
	}
	motiur.getDept()
	motiur.getSub("EEE 201")
}
