package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// Go is certainly unusual when it comes to nil method receivers.
// You just have to read the code to know if a nil receiver won't result in a panic.
// From gopl.io:
// For some types, such as *os.File, nil is a valid receiver (§6.2.1),
// but *bytes.Buffer is not among them. The method is called, but it panics
// as it tries to access the buffer.

func main() {
	var buf *bytes.Buffer
	var out *os.File
	var writer io.Writer
	writer = out
	_, err := writer.Write([]byte("hw"))
	if err != nil {
		fmt.Printf("error writing, no panic!: %v\n", err)
	}
	writer = buf
	_, err = buf.WriteString("hw")
	if err != nil {
		fmt.Printf("error writing, no panic!: %v\n", err)
	}
}
