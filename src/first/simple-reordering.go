package main

import (
	"fmt"
	"time"
)

func main() {
	x := 1
	go func() {
		x += 2
	}()
	go func() {
		x *= 2
	}()
	time.Sleep(10 * time.Millisecond)
	fmt.Printf("x: %d\n", x)
}
