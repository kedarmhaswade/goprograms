package main

import (
	"fmt"
	"sync"
)

func main() {
	x := 1
	var start, done sync.WaitGroup
	start.Add(1)
	done.Add(2)
	go func() {
		defer done.Done()
		start.Wait()
		x += 2
	}()
	go func() {
		defer done.Done()
		start.Wait()
		x *= 2
	}()
	// no goroutines have made progress! Let the games begin NOW.
	start.Done()
	done.Wait()
	fmt.Printf("x: %d\n", x)
}