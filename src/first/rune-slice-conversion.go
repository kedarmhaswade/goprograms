// A []rune conversion applied to a UTF-8-encoded string returns
// the sequence of Unicode code points that the string encodes
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "プログラム"
	fmt.Printf("% x\n", s) // "e3 83 97 e3 83 ad e3 82 b0 e3 83 a9 e3 83 a0" (% x inserts space!)
	r := []rune(s)
	fmt.Printf("%x\n", r) // "[30d7 30ed 30b0 30e9 30e0]"
	// let's decode a single character that is encoded as UTF-8
	s = "क"
	bytes := []byte(s)
	fmt.Printf("bytes of %s: (binary) %b (hex) %x (decimal) %d\n", s, bytes, bytes, bytes)
	// here we see that the bytes of क are: [11100000 10100100 10010101] or [e0 a4 95] or [224 164 149],
	// however its unicode code point is 2325 decimal:
	c := 'क'
	fmt.Printf("unicode code point of %c is %d\n", c, c)
	// so, how is the sequence of 3 bytes: e0 a4 95 or 11100000 10100100 10010101
	// decoded to the code point of 2325 decimal using UTF-8 decoding?
	// the answer is in DecodeRune call. Debug that by stepping into this call:
	d, size := utf8.DecodeRune([]byte(s))
	fmt.Printf("code point (decimal) of %c is %d and it occupies %d bytes", c, d, size)
}
