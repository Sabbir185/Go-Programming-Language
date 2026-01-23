package main

import (
	"fmt"
	"sync"
)

func doTask(ind int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Do Task %d is running\n", ind)
}

func execute() {
	var wg sync.WaitGroup

	for i := 0; i <= 3; i++ {
		wg.Add(1)
		go doTask(i, &wg)
	}

	wg.Wait()
}
