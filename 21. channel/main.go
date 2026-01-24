package main

import (
	"net/http"
	"time"
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

	/*

		for i := 0; i < len(links); i++ {
			fmt.Println(<-ch)
		}

		for {
			go checkLink(<-ch, ch)
		}

		for l := range ch {
			go func(link string) {
				time.Sleep(5 * time.Second)
				checkLink(link, ch)
			}(l)
		}

	*/

	for l := range ch {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, ch)
		}(l)
	}

}

func checkLink(link string, ch chan string) {
	_, err := http.Get(link)
	if err != nil {
		println(link, "might be down!")
		ch <- link
		return
	}
	println(link, "is up!")
	ch <- link
}
