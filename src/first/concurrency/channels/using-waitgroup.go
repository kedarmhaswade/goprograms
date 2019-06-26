// Demonstrates the use of a WaitGroup which waits for a collection of goroutines to finish.
// This could be utilized in starting a group of workers and then waiting for all of them to finish.
// Is this similar to semaphores, or perhaps latches in Java?
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	rand.Seed(time.Now().UnixNano())
	n := 5
	ch := make(chan int, n)
	for i := 1; i <= n; i++ {
		wg.Add(1) // must happen in the main goroutine!
		go func(id int) {
			//wg.Add(1)  // wrong placement of Add!
			defer wg.Done()
			duration := time.Duration(rand.Intn(1000)) * time.Millisecond
			time.Sleep(duration) // sleep for a random number of subsecond milliseconds to getMaxHeadstart work
			fmt.Printf("returning from work goroutine: %d after working for %v\n", id, duration)
			ch <- int(duration / time.Millisecond)
		}(i)
	}
	// closer
	wg.Wait()
	close(ch)
	max := 0
	for s := range ch {
		if s > max {
			max = s
		}
	}
	fmt.Printf("max working time: %dms\n", max)
}
