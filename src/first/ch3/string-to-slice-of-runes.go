// how does the conversion by rune[] of string look?
package ch3

import (
	"fmt"
	"unicode/utf8"
)

func convert() {
	s := "プログラム"      // "program" in Japanese Katakana
	fmt.Printf("% x\n", s) // space between % and x -- fmt tricks
	// the above simply prints the bytes in the string!
	// the bytes an be returned by a []byte conversion
	for _, b := range []byte(s) {
		fmt.Printf("% x", b)
	}
	fmt.Println()
	a := []rune(s) // convert to a
	fmt.Printf("\n%x\n", a)
	// simply prints the runes in the string!
	//DecodeRune implements a UTF-decoder. For a three-byte sequence, for example, the unicode code point is determined as:
	//rune(b[0]&0x0F)<<12 | rune(b[1]&0x3F)<<6 | rune(b[2]&0x3F)
	//Thus, it clears the
	//  4 higher order bits of the most significant byte, i.e. byte 0,
	//	2 higher order bits of the byte 1,
	//	2 higher order bits of the least significant byte, i.e. byte 2
	//
	//And simply concatenates the resulting bytes in the same order.

	var i, j int
	for _, c := range s {
		i = j
		j += utf8.RuneLen(c)
		fb, _ := utf8.DecodeRune([]byte(s)[i:j])	
		fmt.Printf("rune: %v = %v\n", c, fb)
	}
}
