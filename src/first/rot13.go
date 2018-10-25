// Implement a rot13Reader that implements io.Reader and reads from an io.Reader,
// modifying the stream by applying the rot13 substitution cipher to all alphabetical characters.
package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

type rot13Writer struct {
	r io.Writer
}

func (r13 rot13Reader) Read(b []byte) (int, error) {
	n, error := r13.r.Read(b)
	if error == nil {
		for i, a := range b {
			if (a >= 'a' && a <= 'm') || (a >= 'A' && a <= 'M') {
				b[i] = a + 13
			} else if (a >= 'n' && a <= 'z') || (a >= 'N' && a <= 'Z') {
				b[i] = a - 13
			}
		}
	}
	return n, error
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
