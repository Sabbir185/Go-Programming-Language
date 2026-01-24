package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://go.dev",
		"http://amazon.com",
	}

	ch := make(chan string)

	for _, link := range links {
		go checkLink(link, ch)
	}

	for i := 0; i < len(links); i++ {
		fmt.Println(<-ch)
	}
}

func checkLink(link string, ch chan string) {
	_, err := http.Get(link)
	if err != nil {
		println(link, "might be down!")
		ch <- "Might be down I think!"
		return
	}
	println(link, "is up!")
	ch <- "Yep, it's up!"
}
