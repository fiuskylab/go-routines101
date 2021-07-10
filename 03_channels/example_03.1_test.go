package main

import (
	"math/rand"
	"testing"
	"time"
)

// Data is a dummy/moq struct
// If any goroutine wants the data, it sends anything into the ReadRequest
// channels and starts listening on the ReadResponse channel.
type Data struct {
	secret      int           // A randomly generated integer
	readRequest chan chan int // readRequest responsible for reading what is received (a chan int)
}

// Run writes the data periodically(1 second) to the struct. This meant
// to run in it's own goroutine.
// It waits for anything pushed into the ReadRequest channels and pushes
// the current data to the ReadResponse.
func (d *Data) Run() {
	// Generating a seed rand.Source
	seed := rand.NewSource(time.Now().UnixNano())

	// *rand.Rand
	gen := rand.New(seed)

	// Defining a ticker of 1 second
	ticker := time.NewTicker(1 * time.Second)

	// Looping through infinity
	for {
		// Select wait on channel operation
		select {
		// When the ticker passes 1 second
		case <-ticker.C:
			// Define the Data secret field
			d.secret = gen.Int()

		// When a chan int is sent to d.readRequest(chan chan int)
		// respChan(chan int) is defined
		case respChan := <-d.readRequest:
			// Secret(int) is set into respChan(chan int)
			respChan <- d.secret
		}
	}
}

// Get the data secret
func (d *Data) Get() int {
	respChan := make(chan int)

	// respChan(chan int) is set into readRequest(chan chan int)
	d.readRequest <- respChan
	// response(int) receives int from respChan(chan int)
	response := <-respChan
	return response
}

// TestConcurrent sees if there's a data race
func TestConcurrent(t *testing.T) {
	data := Data{readRequest: make(chan chan int)}
	go data.Run()
	time.Sleep(1 * time.Second)
	data.Get()
}
