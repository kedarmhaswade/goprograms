// Plays with Octal values
package main

import (
	"fmt"
)

func main() {
	b := "\377"
	fmt.Printf("len(b) = %d, value of first byte: %d\n", len(b), b[0])
	b = "\101"
	fmt.Printf("len(b) = %d, value of first byte as a char: %c\n", len(b), b[0])
	//buf := make([]byte, 10)
	//c, e := os.Stdin.Read(buf)
	//for e != io.EOF {
	//	for i := 0; i < c; i++ {
	//		fmt.Printf("byte: %d\n", buf[i])
	//	}
	//	c, e = os.Stdin.Read(buf)
	//}
	b = `\`
	fmt.Printf("len(b) = %d, ASCII value of %c = %d\n", len(b), b[0], b[0])
}
