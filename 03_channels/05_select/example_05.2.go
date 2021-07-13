package main

import (
	"time"
	"fmt"
)


// RegisterTimestamp just send a string to ch in a period
// of 500 * N milliseconds between each.
// 0: 0ms
// 1: 500ms
// 2: 1000ms
// 3: 1500ms
// n: (n * 500)ms
func RegisterTimestamp(ch chan<- string) {
	i := -1
	for {
		i++

		ch <- fmt.Sprintf("Passed %d miliseconds!", i * 500)

		milli := i * 500

		time.Sleep(time.Duration(milli) * time.Millisecond)
	}
}

func main() {

	// stopwatch receives messages that registers
	// timestamps
	stopwatch := make(chan string)

	// GoRoutine that receives the stopwatch
	go RegisterTimestamp(stopwatch)

	// Loop through the infinity
	for {
		// Select block
		select {
		// Receiving stopwatch value
		case msg := <-stopwatch:
			fmt.Println(msg)

		// After 1 second it breaks the loop
		case <-time.After(3 * time.Second):
			fmt.Println("Timeout program")
			return
		}
	}
}
