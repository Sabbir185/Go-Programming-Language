package main

import (
	"fmt"
	"sync"
)

func resolveRaceConditionUsinglockingMutex() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	var sum int = 0

	for i := 1; i <= 1000; i++ {

		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			mu.Lock()
			defer mu.Unlock()
			sum = sum + i

		}(i)

	}

	wg.Wait()

	fmt.Println("Total Sum: ", sum)
}
