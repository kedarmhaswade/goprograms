// Demonstrates what happens with bit patterns and uint integers
package main

import "fmt"

func main() {
	var us uint8 = 0xFF
	var s int8 = 0x7F            // 0x8F through 0xFF are not allowed as the "constants" in that range [128-255] overflow int8
	fmt.Printf("%v %v\n", us, s) //255 127
	ss := 0xFF
	fmt.Printf("%T %v\n", ss, ss) //int 255
}
