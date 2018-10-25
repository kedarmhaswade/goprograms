package main

import (
	"fmt"
	"time"
)

var data int

func main() {
	go func() { // (1)
		data++
	}()
	time.Sleep(2 * time.Second)
	if data == 0 {
		fmt.Printf("the value is %v.\n", data)
	}
}
