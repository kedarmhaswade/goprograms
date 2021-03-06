// Just demonstrates the use of DecodeRuneInString
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "ॐ"
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
		fmt.Printf("size: %d\n", size)
		i += size
	}
	// Go's range loop UTF-decodes a string implicitly
	v := "Hello, केदार"
	for i, r := range v {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}
	fmt.Printf("num bytes in v = %v\n", len(v))

}
