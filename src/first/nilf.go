// default value of a function is nil
package main

import (
	"fmt"
	"log"
)

func main() {
	var f func(int) int
	fmt.Printf("type of f: %T\n", f)
	if f == nil {
		log.Fatalf("oh nos! nil function f, not calling it")
	}
	f(1)
}
