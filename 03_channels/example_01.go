package main

import (
	"fmt"
)

func main() {
	// Creating a string channel
	msgs := make(chan string)

	// Goroutine
	go func() {
		// To prove that the application waits for the value to be sent to channel
		// time.Sleep(2 * time.Second)

		// Sending a value to "msgs" channel
		msgs <- "Hello, World!"
	}()

	// Receiving value from "msgs" channel
	// The code will only pass this point when the value is received
	// it's possible to notice it adding a "time.Sleep()" inside the goroutine
	msg := <-msgs

	// Printing what was received
	fmt.Printf("MSG received: %s\n", msg)
}
