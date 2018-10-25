// Studies the relationship between the utf-8 encoded bytes and unicode code point
package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// Consider the devnagari character om: ॐ which has a unicode code point of 950 (hex) or 2384 (decimal).
	// Now, let's write this character to a file and read it back as bytes
	s := "\u0950" // om as a rune with unicode escape followed by the hex value of its code point, s contains only that rune
	fname := "/tmp/uuu"
	write(s, fname)
	b := read(fname)
	fmt.Printf("Characters %s was written as UTF-8 encoded bytes %v\n", s, b)
}
func read(fname string) []byte {
	b, e := ioutil.ReadFile(fname)
	if e != nil {
		fmt.Fprintf(os.Stderr, "Read error %v\n", e)
		return nil
	}
	return b
}
func write(s string, fname string) {
	f, err := os.Create(fname)
	if err == nil {
		f.WriteString(s)
		err = f.Sync()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Sync error %v\n", err)
		}
	} else {
		fmt.Fprintf(os.Stderr, "Write error %v\n", err)
	}
}
