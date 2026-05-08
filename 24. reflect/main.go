package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := 10
	y := reflect.ValueOf(x)
	a := y.Type()

	fmt.Println("value of y : ", y)
	fmt.Println("is zero of y : ", y.IsZero())
	fmt.Println("Type of y : ", a)
	fmt.Println("kind of y : ", a.Kind())
	fmt.Println("Is int : ", a.Kind() == reflect.Int)
	fmt.Println("Is string : ", a.Kind() == reflect.String)

	// Run time value change using reflect package
	id := 101
	valueOfIdAddr := reflect.ValueOf(&id)    // returns memory address
	valueOfId := reflect.ValueOf(&id).Elem() // returns original value

	typeOfId := valueOfIdAddr.Type()

	fmt.Println("Value of ID address is: ", valueOfIdAddr)
	fmt.Println("Value of ID address is: ", valueOfId)
	fmt.Println("Type of id: ", typeOfId)
	fmt.Println("Original value: ", valueOfId.Int())

	valueOfId.SetInt(100) // change value in runtime
	fmt.Println("modified value: ", valueOfId.Int())

	// interface{}
	var unKnownValue interface{} = "This is unKnown type"
	v3 := reflect.ValueOf(unKnownValue)

	if v3.Kind() == reflect.String {
		fmt.Println("This is string type: ", unKnownValue)
	}
}
