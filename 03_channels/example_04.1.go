package main

import "fmt"

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
