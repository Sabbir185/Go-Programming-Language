package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	ch := make(chan int)

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Sending data to channel...")
		ch <- 10
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Receiving data from channel...")
		data := <-ch
		fmt.Println("Received:", data)
	}()

	wg.Wait()
	fmt.Println("Main go routine")
}
