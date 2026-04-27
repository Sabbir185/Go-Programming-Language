package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type User struct {
	Name       string `json:"name"`
	Department string `json:"department"`
}

type NewUser struct {
	Name  string `json:"name"`
	Class string `json:"class"`
}

func main() {
	user := User{
		Name:       "Sabbir",
		Department: "CSE",
	}
	fmt.Println(user)

	// making object to json data using Marshal
	jsonData, err := json.Marshal(user)
	if err != nil {
		log.Fatalln("Error during json conversion")
	}
	fmt.Println("Byte slice: ", jsonData)
	fmt.Println("Json data: ", string(jsonData))

	// json to object or struct
	var user2 User
	err = json.Unmarshal(jsonData, &user2)
	if err != nil {
		log.Fatalln("Error during user conversion")
	}
	fmt.Println("User created from json data: ", user2)

	// ----------------------------------------------------
	// json encoding or decoding without json marshal / unmarshal
	var row_json string = `{"Name": "ABC", "Class": "7"}`
	reader := strings.NewReader(row_json)
	decoder := json.NewDecoder(reader)

	var user3 NewUser
	err = decoder.Decode(&user3)
	if err != nil {
		log.Fatalln("Error to decode json")
	}
	fmt.Println("User3 object: ", user3)

	// making json using NewEncoder
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	err = encoder.Encode(user)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("User json data: ", buf.String())

}
