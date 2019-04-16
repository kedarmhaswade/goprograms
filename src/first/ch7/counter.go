package ch7

import (
	"io"
	"unicode"
)

type ByteCounter int

// Writes the length of the given slice of bytes to a counter and also returns it
// Note that the method accepts a pointer to a ByteCounter whose underlying type is just an int.
func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert len(p) which is an int to ByteCounter
	//*c += len(p) // won't work as the left hand side points to a "ByteCounter" whereas the right hand side is an int
	return len(p), nil
}

// Using the ideas from ByteCounter, implement counters for words and for lines.
// You will find bufio.ScanWords useful.

type WordCounter int

type LineCounter int

// Returns the number of words in the given byte slice. By implementing this method the contract of
// io.Writer is satisfied by a value of type *WordCounter.
func (wc *WordCounter) Write(p []byte) (int, error) {
	// we assume that words are separated by whitespaces
	words := 0
	for i := 0; i < len(p); {
		if ! unicode.IsSpace(rune(p[i])) { // a new word began
			words++
			i++
			for ; i < len(p) && !unicode.IsSpace(rune(p[i])); {
				i++
			}
		} else { // it is a space, so continue reading next byte
			i++
		}
	}
	*wc = WordCounter(words) // it is required to change the contents of the receiver!
	return words, nil
}

// Writes the number of lines in the given slice of bytes to the given receiver.
// Note: it modifies the receiver!
func (lc *LineCounter) Write(p []byte) (int, error) {
	for _, b := range p {
		if rune(b) == '\n' {
			*lc += 1
		}
	}
	return int(*lc), nil
}

// Write a function CountingWriter with the signature below that, given an io.Writer,
// returns a new Writer that wraps the original, and a pointer to an int64 variable
// that at any moment contains the number of bytes written to the new Writer.
// (note: this sounds like the decorator pattern)
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var count int64
	nw := NewWriter{w, &count}
	return nw, &count
}

type NewWriter struct {
	w io.Writer
	count *int64
}

func (nw NewWriter) Write(b []byte) (int, error) {
	n, err := nw.w.Write(b)
	if err == nil {
		*(nw.count) += int64(n)
	}
	return n, err
}