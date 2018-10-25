package main

import "fmt"

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
)

func main() {
	var arg2 interface{} = Big
	fmt.Printf("Big is of type %T\n", arg2)

	fmt.Printf("Big is of type %T\n", Big)
}
