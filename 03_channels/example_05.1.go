package main

import (
	"fmt"
	"time"
)

type Message struct {
	Text     string
	SentAt   string
	timeSent time.Time
}

func main() {
	msgChan := make(chan *Message)
	quit := make(chan int)

	go func() {
		for i := 1; i <= 2; i++ {
			send(&Message{
				Text:     fmt.Sprintf("Message %d", i),
				timeSent: time.Now(),
			}, msgChan)
			if i%2 == 0 {
				time.Sleep(1 * time.Second)
				continue
			}
			time.Sleep(2 * time.Second)
		}

		close(msgChan)

		quit <- 0
	}()

	awaitMsg(msgChan, quit)

}

func send(msg *Message, msgChan chan<- *Message) {
	msg.SentAt = msg.timeSent.String()
	msgChan <- msg
}

func awaitMsg(msgChan <-chan *Message, quitChan <-chan int) {
	for {
		select {
		case msg := <-msgChan:
			fmt.Printf("Message Received \"%s\" at: %s\n", msg.Text, msg.SentAt)
		case <-quitChan:
			break
		}
	}
}
