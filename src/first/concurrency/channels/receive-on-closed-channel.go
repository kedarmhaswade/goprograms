// What happens when a goroutine tries to receive on a closed channel?
package main

import (
	"fmt"
)

func closer(ch chan int) {
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)
}

func main() {
	ch := make(chan int)
	go closer(ch)
	for {
		v, ok := <-ch
		fmt.Printf("channel is closed, ok: %v, value: %d\n", ok, v)
		if !ok {
			return
		}
	}
}
