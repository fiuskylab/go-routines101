package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Wait group variable
	var wg sync.WaitGroup

	// Adding one "task" to wait group
	wg.Add(1)

	// Function stored in a variable
	f := func(sec int) {
		for i := 1; i <= sec; i++ {
			// Sleeping for 1 second
			time.Sleep(time.Second * 1)

			fmt.Printf("%d second(s) have passed!\n", i)
		}

		// Decreasing by 1 the WaitGroup total
		wg.Done()
	}

	// Running goroutine
	go f(10)

	// The program won't pass this line until the WaitGroup reaches 0
	wg.Wait()
}
