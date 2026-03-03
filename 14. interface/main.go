package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct{}
type Cat struct{}

func (d Dog) Speak() string {
	return "Woof"
}

func (c Cat) Speak() string {
	return "Meow"
}

// common, reusable function
func MakeSound(a Animal) {
	fmt.Println(a.Speak())
}

func main() {
	dog := Dog{}
	cat := Cat{}

	MakeSound(dog)
	MakeSound(cat)
}



/*

	1. এখানে explicitly লিখতে হয় না যে Dog implements Animal — Go automatically বুঝে নেয়।
	2. এখানে polymorphism হচ্ছে — একই function, different behavior (MakeSound).
		Polymorphism = same interface, different behavior
		MakeSound() function টা সবসময় Animal নেয়।
		কিন্তু runtime এ:
			যদি Dog পাঠাও → Dog.Speak() call হবে
			যদি Cat পাঠাও → Cat.Speak() call হবে
		👉 একই function, different output
		👉 behavior depend করছে actual object এর উপর

*/
