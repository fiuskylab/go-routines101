package main

import (
	"fmt"
	"time"
)

// Message sets a basic data structure about messaging.
type Message struct {

	// Text is the content being sent
	Text string

	// SentAt is the string value shown at front-end
	SentAt string

	// timeSent is set when the Message is created
	// using the method time.Now()
	timeSent time.Time
}

func main() {
	// msgChan receives created Messages
	msgChan := make(chan *Message)
	// quit is flag, signalizes the end of sending messages
	quit := make(chan int)

	n := 10

	// GoRoutine that will trigger a N amount of messages
	go func() {
		// Loop iterating a N amount of times
		for i := 1; i <= n; i++ {
			// Sending data to send
			send(&Message{
				Text:     fmt.Sprintf("Message %d", i),
				timeSent: time.Now(),
			}, msgChan)

			// Just a silly way to control the time between sent messages
			if i%2 == 0 {
				time.Sleep(1 * time.Second)
				continue
			}
			time.Sleep(2 * time.Second)
		}

		// Send a value to quit
		// This can be any value and/or type of value
		quit <- 0
	}()

	// Stop the program here, until quit flag is triggered
	awaitMsg(msgChan, quit)
}

// send receives:
// msg a *Message
// msgChan a receive-only channel of *Message
func send(msg *Message, msgChan chan<- *Message) {
	// Transforming time.Time to string and
	// allocating the value to Message's SentAt field
	msg.SentAt = msg.timeSent.String()
	// msgChan receiving msg value
	msgChan <- msg
}

// awaitMsg show the sent messages and end the execution when the flag is triggered, receives:
// msgChan a send-only channel, that receives *Message
// quit a send-only channel, that receives a int, it's the flag that will stop the function
func awaitMsg(msgChan <-chan *Message, quitChan <-chan int) {
	// looping through infinite
	for {
		// select blocks the program until one channel receives a value
		select {
		// Receives a msg and trigger a message
		case msg := <-msgChan:
			fmt.Printf("Message Received \"%s\" at: %s\n", msg.Text, msg.SentAt)
		// Receives the quit flag, and exits the for loop
		case <-quitChan:
			return
		}
	}
}
