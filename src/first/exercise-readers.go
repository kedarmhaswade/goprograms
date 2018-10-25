package main

// Implement a Reader type that emits an infinite stream of the ASCII character 'A'.

import (
	"fmt"
	"golang.org/x/tour/reader"
)

type MyReader struct{}

type NilArray string

func (_ NilArray) Error() string {
	return "passed a nil buffer"
}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (reader MyReader) Read(b []byte) (int, error) {
	if b == nil {
		return 0, NilArray("")
	}
	for i := range b {
		b[i] = byte('A')
	}
	return len(b), nil
}

func main() {
	r := MyReader{}
	if _, error := r.Read(nil); error != nil {
		fmt.Printf("%v\n", error)
	}
	reader.Validate(r)
}
