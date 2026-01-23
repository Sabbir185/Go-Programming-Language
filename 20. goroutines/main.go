package main

import (
	"fmt"
	"time"
)

// A goroutine is a lightweight thread managed by the Go runtime.
// Goroutines in Go let functions run concurrently, using less memory than traditional threads.
// Every Go program starts with a main Goroutine, and if it exits, all others stop.

func task(ind int) {
	fmt.Printf("Task %d is running\n", ind)
}

func main() {
	for i := 0; i <= 10; i++ {
		go task(i)

		go func(i int) {
			fmt.Printf("Immediate function for %d\n", i)
		}(i)
	}

	time.Sleep(time.Second * 2)
}
