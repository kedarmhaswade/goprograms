// Demonstrates a pipeline of communicating goroutines with a possible "sentinel"
// This sentinel is my solution (before I was told that a close function can close a channel)
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Packet struct {
	data int
	done bool
}

func main() {
	naturals := make(chan Packet)
	squares := make(chan Packet)
	// Counter
	go func() {
		x := 1
		for {
			rand.Seed(time.Now().UnixNano())
			if rand.Float64() < 0.1 {
				naturals <- Packet{0, true}
				break
			}
			naturals <- Packet{x, false}
			x += 1
		}
	}()
	// Squarer
	go func() {
		for {
			p := <-naturals
			if p.done {
				squares <- p
				break
			}
			x := p.data
			s := x * x
			squares <- Packet{s, false}
		}
	}()
	// Printer
	for {
		p := <-squares
		if p.done {
			break
		}
		fmt.Printf("received: %v from squares channel\n", p.data)
	}
}
