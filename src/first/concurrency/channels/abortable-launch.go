// Let’s make our launch program print the countdown.
// The select statement below causes each iteration of the loop to wait up
// to 1 second for an abort, but no longer.
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	ticker := time.Tick(1 * time.Second)
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		abort <- struct{}{}
	}()
	for i := 10; i >= 1; i-- {
		select {
		case <-ticker:
			fmt.Printf("%d\n", i)
		case <-abort:
			fmt.Printf("launch aborted!\n")
			return // break will not do because then we'll print "blast off" when aborted!
		}
	}
	fmt.Printf("blast off!")
}
