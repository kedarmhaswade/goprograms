package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func main() {
	hwhex := "\xe4\xb8\x96\xe7\x95\x8c"
	s := hwhex
	for i := 0; i < len(s); i++ {
		fmt.Printf("byte %d = %b\n", i, s[i])
	}
	kedar := "\u0900"
	fmt.Printf("string: %s nbytes = %d, nchars = %d\n", kedar, len(kedar), utf8.RuneCountInString(kedar))
	testValid([]byte(kedar))
	testValid([]byte(hwhex))
	hwu := "\u4e16\u754c"
	testValid([]byte(hwu))
	if hwhex == hwu {
		fmt.Printf("%s and %s are the same string\n", hwhex, hwu)
	}
}

func testValid(b []byte) {
	e, _ := isValidUtf8(b)
	if e == nil {
		fmt.Printf("valid utf-8 encoded byte sequence: %x\n", b)
	} else {
		fmt.Printf("%v\n", e)
	}
}

// Checks whether the given byte array represents valid utf-8-encoded characters.
// Returns an error and the index of the byte (non negative) that in incorrectly encoded.
// If the error is nil, then the second return value is -1 (negative) which means a valid utf-8 encoding.
func isValidUtf8(b []byte) (error, int) {
	n := len(b)
	v := true
	for i := 0; i < n; {
		for j := 1; j <= 4; j++ {
			if shouldGet(b, i, j) {
				if e := valid(b, i, j); e == nil {
					i += j
					break // inner loop exits here if every character is correctly encoded
				} else {
					return e, i
				}
			} else if j == 4 {
				v = false
			}
		}
		if !v {
			// the inner loop failed to find either a 1-, 2-, 3-, or 4-byte valid UTF-encoding at index i, so we error out
			msg := fmt.Sprintf("The byte %b at index %d is none of [0xxxxxx, 110xxxxx, 1110xxxx, 11110xxx]", b[i], i)
			return errors.New(msg), i
		}
	}
	return nil, -1
}

// Determines if the j bytes at index i (inclusive) of b represent valid UTF-8 encoding
func valid(b []byte, i int, j int) error {
	switch j {
	case 1:
		if b[i]>>7 == 0 { // 7-bit ascii is okay!
			return nil
		}
		msg := fmt.Sprintf("Invalid; the byte %b at index %d should have represented 7-bit ASCII value", b[i], i)
		return errors.New(msg)
	case 2:
		if !has(b, j, i) { // common for 2, 3, 4
			msg := fmt.Sprintf("Invalid; not enough bytes at index %d following byte %x", j, b[i])
			return errors.New(msg)
		}
		// ensured that index is within the bounds
		if b[i]>>5 != 6 { // b[i] does not start with 110
			msg := fmt.Sprintf("Invalid; the byte %b at index %d must start with 110", b[i], i)
			return errors.New(msg)
		}
		if b[i+1]>>6 != 2 { // b[i+1] does not start with 10
			msg := fmt.Sprintf("Invalid; the byte %b at index %d must start with 10", b[i+1], i+1)
			return errors.New(msg)
		}
		return nil
	case 3:
		if !has(b, j, i) { // common for 2, 3, 4
			msg := fmt.Sprintf("Invalid; not enough bytes at index %d following byte %x", j, b[i])
			return errors.New(msg)
		}
		// ensured that index is within the bounds
		if b[i]>>4 != 14 { // b[i] does not start with 1110
			msg := fmt.Sprintf("Invalid; the byte %b at index %d must start with 1110", b[i], i)
			return errors.New(msg)
		}
		if b[i+1]>>6 != 2 { // b[i+1] does not start with 10
			msg := fmt.Sprintf("Invalid; the byte %b at index %d must start with 10", b[i+1], i+1)
			return errors.New(msg)
		}
		if b[i+2]>>6 != 2 { // b[i+2] does not start with 10
			msg := fmt.Sprintf("Invalid; the byte %b at index %d must start with 10", b[i+2], i+2)
			return errors.New(msg)
		}
		return nil
	case 4:
		if !has(b, j, i) { // common for 2, 3, 4
			msg := fmt.Sprintf("Invalid; not enough bytes at index %d following byte %x", j, b[i])
			return errors.New(msg)
		}
		// ensured that index is within the bounds
		if b[i]>>3 != 30 { // b[i] does not start with 11110
			msg := fmt.Sprintf("Invalid; the byte %b at index %d must start with 11110", b[i], i)
			return errors.New(msg)
		}
		if b[i+1]>>6 != 2 { // b[i+1] does not start with 10
			msg := fmt.Sprintf("Invalid; the byte %b at index %d must start with 10", b[i+1], i+1)
			return errors.New(msg)
		}
		if b[i+2]>>6 != 2 { // b[i+2] does not start with 10
			msg := fmt.Sprintf("Invalid; the byte %b at index %d must start with 10", b[i+2], i+2)
			return errors.New(msg)
		}
		if b[i+3]>>6 != 2 { // b[i+3] does not start with 10
			msg := fmt.Sprintf("Invalid; the byte %b at index %d must start with 10", b[i+3], i+3)
			return errors.New(msg)
		}
		return nil
	}
	// this is impossible!
	msg := fmt.Sprintf("This is impossible: index: %d, #bytes: %d", i, j)
	return errors.New(msg)
}

// Decides if the given byte slice b has j bytes starting at index i (inclusive)
func has(b []byte, j int, i int) bool {
	return (i + j) <= len(b)
}

func shouldGet(b []byte, i int, j int) bool {
	switch j {
	case 1:
		return b[i]>>7 == 0 // b[i] is 7-bit ascii
	case 2:
		return b[i]>>5 == 6 // b[i] starts with 110
	case 3:
		return b[i]>>4 == 14 // b[i] starts with 1110
	case 4:
		return b[i]>>3 == 30 // b[i] starts with 11110
	default:
		return false
	}
}
