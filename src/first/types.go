package main

import "fmt"

func main() {
	// numeric types
	// is byte unsigned?
	var b byte = 255 // 256 // overflow
	fmt.Printf("%v\n", b)
	// conversion
	var x = 104
	fmt.Println(string(x)) // the underlying type of x and
}
