// deferred call to a stack printing routine
// Go's panic mechanism runs the deferred functions BEFORE it unwinds the stack!
package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	defer printStack()
	fn(3)
}

func fn(x int) {
	fmt.Printf("fn(%d)\n", x+0/x) // panics if x == 0
	defer fmt.Printf("defer %d\n", x) // similar to calling defer in a loop!
	fn(x - 1)
}
func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}
