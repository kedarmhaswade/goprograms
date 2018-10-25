package main

import (
	"fmt"
	"time"
)

func fibSlow(n int) int {
	if n <= 1 {
		return n
	}
	return fibSlow(n-1) + fibSlow(n-2)
}

func showSpinner(delay time.Duration) {
	fmt.Printf("%s\n", "working ...")
	for {
		for _, c := range `-\|/` {
			fmt.Printf("\r%c", c) // \r is carriage return: resets the cursor to the beginning of the line
			time.Sleep(delay)
		}
	}
}
func main() {
	go showSpinner(100 * time.Millisecond) // show spinner in a separate goroutine, args evaluated in main goroutine
	const n = 45
	fibN := fibSlow(n) // slow
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}
