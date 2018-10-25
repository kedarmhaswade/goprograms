package main

import "fmt"

func main() {
	ch := make(chan int)
	fmt.Printf("type of channel: %T\n", ch)
	//n := 1 << 70 // constant 1180591620717411303424 overflows int
	//fmt.Printf("type of n: %T\n", n)
}
