### Channels

#### Summary
1. [References](#references)
2. [Introduction](#introduction)
    1. [Example 1](#example-1)
3. [](#)
4. [](#)
5. [](#)
6. [](#)
7. [](#)
8. [](#)

#### References
- [GoByExample - Channels](https://gobyexample.com/channels)
- [GoByExample - Channels Buffering](https://gobyexample.com/channel-buffering)
- [GoByExample - Channels Synchronization](https://gobyexample.com/channel-synchronization)
- [GoByExample - Channels Directions](https://gobyexample.com/channel-directions)
- [Effective Go - Channels](https://golang.org/doc/effective_go#channels)

#### Introduction
What is a _channel_?
_"Channels are the pipes that connect concurrent goroutines. You can send values into channels from one goroutine and receive those values into another goroutine."_

Channels follows the following template for declaration: 
    - _make(chan type)_ - You should use _make_, and the _type_ is any type, can even be another channels, we'll see this example at [Example 2](#example-2)


##### Example 1
Simple example of a channel receiving a string
```golang
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
```
