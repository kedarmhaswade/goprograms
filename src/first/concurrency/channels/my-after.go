// The time.After function immediately returns a channel, and
// starts a new goroutine that sends a single Value on that channel after the specified time.
// Here's how I implemented it; is that good?
package main

import (
	"fmt"
	"time"
)

func myAfter(duration time.Duration) <-chan time.Time {
	out := make(chan time.Time)
	go func() {
		c := time.Tick(duration) // returns a channel on which a single Value would be sent after time elapses
		out <- <-c
	}()
	return out
}
func main() {
	fmt.Printf("Now: %v\n", time.Now().Format("03:04:05"))
	ch := myAfter(1 * time.Second)
	fmt.Printf("and: %v\n", (<-ch).Format("03:04:05"))
}
