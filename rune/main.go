package main

import (
	"fmt"
)

func main() {
	name := "সাব্বির"

	rname := []rune(name)

	fmt.Println("Updated name:", rname)

	for i, value := range rname {
		fmt.Printf("Index %d, value %c, Unicode value %U\n", i, value, value)
	}

	rname[0] = 'ছ'
	uname := string(rname)

	fmt.Println("Updated name:", uname)
	
	fmt.Printf("String length: %d, \nrune length: %d\n", len(name), len(rname))
}
