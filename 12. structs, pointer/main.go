package main

import "fmt"

type ContactInfo struct {
	email   string
	phone   string
	zipCode int
}

type Person struct {
	firstName string
	lastName  string
	age       int
	contact   ContactInfo
}

func main() {
	p := Person{
		firstName: "A",
		lastName:  "B",
		age:       30,
	}
	fmt.Println(p)

	var emptyPerson Person
	fmt.Println(emptyPerson)
	fmt.Printf("%+v", emptyPerson)

	fmt.Println("\n--- Struct Literal with Field Names ---")
	fmt.Println("Updating Fields of an Empty Struct")

	emptyPerson.firstName = "John"
	emptyPerson.lastName = "Doe"
	emptyPerson.age = 25

	fmt.Println(emptyPerson)

	fmt.Println("\n--- Nested Structs ---")

	contactPerson := Person{
		firstName: "Jane",
		lastName:  "Smith",
		age:       28,
		contact: ContactInfo{
			email:   "a@gmail.com",
			phone:   "123-456-7890",
			zipCode: 12345,
		},
	}

	// contactPPointerAddress := &contactPerson // optional & to get pointer address, go automatically handles it for *Person receiver
	// contactPPointerAddress.updateName("Sabbir")

	contactPerson.updateName("Sabbir")
	contactPerson.print()

}

func (contactPointer *Person) updateName(firstName string) {
	(*contactPointer).firstName = firstName
}

func (p Person) print() {
	fmt.Printf("%+v\n", p)
}

/*

Pointers in Go:

	& --> Pointer hold memory address of that variable
	* --> Dereference pointer to get value

	Reference type: Slices, Maps, Channels, Pointers, Functions
	Value type: int, float, string, bool, arrays, structs

*/
