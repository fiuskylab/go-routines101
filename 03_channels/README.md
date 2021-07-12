### Channels

#### Summary
1. [References](#references)
2. [Introduction](#introduction)
    1. [Example 1.1](#example-1.1)
3. [Buffering](#buffering)
    1. [Example 2.1](#example-2.1)
4. [Channels of channels](#channels-of-channels)
    1. [Example 3.1](#example-3.1)
5. [Channel Directions](#channel-directions)
    1. [Example 4.1](#example-4.1)
6. [Select](#select)
    1. [Example 5.1](#example-5.1)
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


##### Example 1.1
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

##### Example 2.1

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

#### Channels of Channels
- Props to:
    - [Gergő Huszty at gitconnected](https://levelup.gitconnected.com/channels-inside-channels-pattern-in-golang-3d0e73a545cc)
    - [YourBasic](https://yourbasic.org/golang/data-races-explained/)
Because _"[...] a channel is a first-class value that can be allocated and passed around like any other. [...]"_, that means we can create a channel of type __T__ that inside of it have a channel.

This can brings some negative effects like "data-race"(_"A data race happens when two goroutines access the same variable concur­rently, and at least one of the accesses is a write"_)

##### Example 3.1
- Run with: `go test . -race`
```golang
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
```

#### Channel directions
- Channels as _function parameters_ can have directions:
    - `f(myChannel <-chan T)` - _myChannel_ will only be able to __receive__ values
    - `f(myChannel chan<- T)` - _myChannel_ will only be able to __send__ values
- More about it at:
    - [GoByExample - Channel Directions](https://gobyexample.com/channel-directions)
    - [GolangR](https://golangr.com/channel-directions/)
    - [GolangByExample](https://golangbyexample.com/channel-direction-go/)


##### Example 4.1
```golang
// minMiddleware receives:
// arrChan - a channel that only RECEIVES values
// arr - a slice of  ints
func minMiddleware(arrChan chan<- []int, arr []int) {
	// Writing arr into arrChan
	arrChan <- arr
}

// min receives:
// minChan - a channel that only RECEIVES value
// arrChan - a channel that only WRITES value
func min(minChan chan<- int, arrChan <-chan []int) {
	min := 0
	for key, val := range <-arrChan {
		if key == 0 {
			min = val
		}

		if min > val {
			min = val
		}
	}

	// Writing the minimum value into minChan
	minChan <- min
}

func main() {
	arrChan := make(chan []int, 1)
	minChan := make(chan int, 1)

	minMiddleware(arrChan, []int{9, 2, 10, 15, -7, 8, 5})
	min(minChan, arrChan)

	fmt.Printf("The minimum value: %d", <-minChan)
}
```

#### Select 
- More about it at:
    - [Tour Golang](https://tour.golang.org/concurrency/5)
    - [YourBasic](https://yourbasic.org/golang/select-explained/)
    - [GoByExample](https://gobyexample.com/select)
    - [GolangDocs](https://golangdocs.com/select-statement-in-golang)

The `select` statement is used to choose from multiple send/receive channel operations, it blocks until one of the send/receive operations is ready. If multiple operations are ready, one of them is chosen at random.

##### Example 5.1
```golang
```
