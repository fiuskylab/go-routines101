package main

import (
	"fmt"
)

// Factorial function
func fac(n int) int {
	if n > 0 {
		return n * fac(n-1)
	}
	return 1
}

func main() {
	// Creating a int channel
	nums := make(chan int, 10)

	// Goroutine
	go func(n int) {
		for i := 0; i < n; i++ {
			nums <- fac(i)
			// Sending a value to "nums" channel
		}
	}(10)

	// counter
	i := 0

	// Closing channel before iterating
	// Or the application will show a message: fatal error: all goroutines are asleep - deadlock!
	close(nums)

	// Iterating channel values
	for n := range nums {
		// Printing what was received
		fmt.Printf("Fac(%d) = %d\n", i, n)
		i++
	}
}
