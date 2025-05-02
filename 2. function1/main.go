package main

import "fmt"

func welcomeMessage() {
	fmt.Println("Welcome to the application")
}

func getUserName() string {
	var name string
	fmt.Println("Enter your name : ")
	fmt.Scanln(&name)
	return name
}

func getTwoNumber() (int, int) {
	var num1, num2 int
	fmt.Println("Enter two numbers with space <> : ")
	fmt.Scanln(&num1, &num2)
	return num1, num2
}

func add(num1 int, num2 int) int {
	return num1 + num2
}

func display(name string, sum int) {
	fmt.Println("Hello, ", name)
	fmt.Println("The result is : ", sum)
}

func printGoodByMsg() {
	fmt.Println("Thanks for using our application")
	fmt.Println("Good Bye")
}

func main() {
	welcomeMessage()
	name := getUserName()
	num1, num2 := getTwoNumber()
	sum := add(num1, num2)
	display(name, sum)
	printGoodByMsg()
}
