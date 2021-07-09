package main

import (
	"fmt"
	"net/http"
	"sync"
)

var websites map[string]string

func main() {
	// Wait group variable
	var wg sync.WaitGroup

	// Map of sites and their URL
	websites = map[string]string{
		"Google":      "https://google.com",
		"GitHub":      "https://github.com",
		"Golange":     "https://golang.org",
		"Drone":       "https://drone.io",
		"GoByExample": "https://gobyexample.com",
		"Wikipedia":   "https://wikipedia.org",
		"Twitter":     "https://twitter.com",
	}

	// Iterating all websites
	for website, url := range websites {
		// Adding one "task" to wait group
		wg.Add(1)

		// Running go routine
		go func(url, website string, wg *sync.WaitGroup) {
			// HTTP Request into websites's URL
			res, err := http.Get(url)

			// Conditional if had any error
			if err != nil {
				fmt.Printf("Not able to ping %s, status code: %d\n", website, res.StatusCode)
			} else {
				fmt.Printf("%s is working fine\n", website)
			}

			// Decreasing by 1 the WaitGroup total
			wg.Done()

		}(website, url, &wg)
	}

	// The program won't pass this line until the WaitGroup reaches 0
	wg.Wait()
}
