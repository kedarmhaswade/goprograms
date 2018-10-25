// Demonstrates the syntactic issue of unidirectional channels
package main

import (
	"fmt"
	"io/ioutil"
)

// our main goroutine creates a receive-only goroutine (input) where it sends a set of numbers to another goroutine
// which in turn returns the sum of those numbers on another channel that is send-only.

func summer(in <-chan int, out chan<- int) {
	sum := 0
	for n := range in {
		fmt.Fprintf(ioutil.Discard, "received: %d\n", n)
		sum += n
	}
	// in <- 2 // error: in is a receive-only channel
	out <- sum
	// x := <-out // error: out is a send-only channel
}
func main() {
	in := make(chan int)
	out := make(chan int)
	n := 0
	go summer(in, out)
	for n <= 100 {
		in <- n
		n += 1
	}
	//close(in)
	fmt.Printf("sum: %v\n", <-out)
}
