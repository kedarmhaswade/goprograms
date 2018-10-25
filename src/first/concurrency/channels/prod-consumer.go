// The buffer in a buffered channel separates the sending and receiving operations to quite an extent.
// This mitigates the blocking and synchronization related to two or more goroutines operating on the same channel concurrently.
// This is similar to a buffered IO implemented in the standard C library so that the blocking can be minimized.
// This program demonstrates how the ubiquitous producer-consumer pattern can be implemented
// using a buffered channel shared between a producing goroutine and a consuming goroutine which produce
// and consume items concurrently at will.
// An additional benefit is that by employing a buffered channel, one can reduce the likelihood of
// deadlocks with concurrently executing goroutines.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// produces a stream of first n natural numbers to the given channel
func produce(n int, out chan<- int) {
	for i := 1; i <= n; i++ {
		out <- i
		//time.Sleep( rand.Intn(500) * time.Millisecond)
		//time.Sleep(500 * time.Millisecond)
		//fmt.Printf("%T\n", 500)
		time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
	}
	close(out)
	fmt.Printf("Produced: %v items, producing finished at: %v\n", n, time.Now().Format(time.StampMilli))
}

// consumes the given stream of integers and returns their average. The average should be same as (n+1)/2,
// the function always keeps a running average.
func consume(in <-chan int) float64 {
	sum := 0
	count := 0
	avg := 0.0
	for i := range in {
		count += 1
		sum += i
		avg = float64(sum) / float64(count)
		//fmt.Printf("sum: %d, running avg: %f\n", sum, avg)
		time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
	}
	fmt.Printf("Consumed: %v items, consuming finished at: %v\n", count, time.Now().Format(time.StampMilli))
	return avg
}

func main() {
	q := make(chan int, 1000) // play with the buffer size
	//r := make(chan float64, 1) // it is just the transfer object, just one is enough
	go produce(10, q)
	a := consume(q) //synchronous call!
	fmt.Printf("the average of numbers: %f\n", a)

}
