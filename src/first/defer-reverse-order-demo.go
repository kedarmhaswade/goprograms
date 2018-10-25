// When a panic occurs, all deferred functions are run in reverse order,
// starting with those of the topmost function  on the stack and
// proceeding up to main, as the program below demonstrates:
package main

import "fmt"

func main() {
	f(3)
}

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x) // panics if x == 0
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}