// Implements an io.Reader that reads from a backing string
package ch7

import (
	"io"
)

type Text struct {
	src string // denotes the source of input data
	cur int    // denotes where in src to start reading the data from
}

// Returns a new io.Reader that reads from given text
func NewReader(src string) io.Reader {
	return &Text{src: src}
}

// Give Test some methods

// One of these methods is what makes Text pass as an io.Reader:
//type Reader interface {
//	Read(p []byte) (n int, err error)
//}

// Implements the Reader contract described in io/io.go
func (t *Text) Read(p []byte) (n int, err error) {
	i := 0
	for ; i < len(p) && (t.cur+i) < len(t.src); i++ {
		p[i] = t.src[t.cur+i]
	}
	if i == len(p) { // reached the end of slice given to us to write to
		t.cur += len(p)
		return len(p), nil
	}
	// we are at the end of input source, so return an EOF, along with number of bytes read
	t.cur += len(p)
	return len(p), io.EOF
}
