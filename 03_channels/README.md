### Channels

#### Summary
1. [References](#references)
2. [Introduction](#introduction)
    1. [Example 1](#example-1)
3. [Buffering](#buffering)
    1. [Example 2](#example-2)
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

#### Buffering
_"Like maps, channels are allocated with make, and the resulting value acts as a reference to an underlying data structure. If an optional integer parameter is provided, it sets the buffer size for the channel. The default is zero, for an unbuffered or synchronous channel."_

```golang
// As seen on "Effective Go - Channels"

ci := make(chan int)            // unbuffered channel of integers
cj := make(chan int, 0)         // unbuffered channel of integers
cs := make(chan *os.File, 100)  // buffered channel of pointers to Files
```

##### Example 2

```golang
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
```
