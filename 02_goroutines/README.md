### Goroutines

References
    - [Go Doc](https://golang.org/doc/effective_go#goroutines)
    - [GoByExample](https://gobyexample.com/goroutines)

Go eases the way to execute a function concurrently, with a simple syntax:
```golang
func main() {

    f := func(sec int) {
       for i := 1; i <= sec; i++ {
            time.Sleep(time.Second * i)
            fmt.Printf("%d second(s) have passed!\n", i)
       }
    }

    // Now can run concurrently with:
    go f(10)

}
```

If you run the code above, it won't print a thing, because the "f() function" continued running in the background concurrently, so the "main() scope" ended before the start of the goroutine.

So, to fix it, we'll need to import a package called [sync](https://golang.org/pkg/sync/), it provides us basic tools to manage our routines; implementing it the code will look like this:
```golang
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Wait group variable
	var wg sync.WaitGroup

	// Adding one "task" to wait group
	wg.Add(1)

	// Function stored in a variable
	f := func(sec int) {
		for i := 1; i <= sec; i++ {
			// Sleeping for 1 second
			time.Sleep(time.Second * 1)

			fmt.Printf("%d second(s) have passed!\n", i)
		}

		// Decreasing by 1 the WaitGroup total
		wg.Done()
	}

	// Running goroutine
	go f(10)

	// The program won't pass this line until the WaitGroup reaches 0
	wg.Wait()
}
```

Now a pratical example, the following example the program will try to ping various websites, concurrently:
```golang
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
```
