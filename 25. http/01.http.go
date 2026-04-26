package main

import (
	"fmt"
	"io"
	"net/http"
)

func HelloHttp01() {
	client := &http.Client{}
	res, err := client.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println("Network error. Failed to fetch data: ", err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Failed to read data")
		return
	}

	fmt.Println(string(body))
}
