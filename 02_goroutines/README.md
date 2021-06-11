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

