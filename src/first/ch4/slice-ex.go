package ch4

import (
	"errors"
	"unicode"
	"unicode/utf8"
)

// Note that copying a slice (which happens implicitly when you pass a slice, as an arg, to a function) creates
// an alias for that slice and the slice elements can be changed through this alias.

// Reverse reverses the given slice "in place"
func Reverse(a []int) {
	n := len(a)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

// LeftRotate rotates the given slice by r places to the left
func LeftRotate(a []int, r int) {
	r %= len(a)
	Reverse(a[:r])
	Reverse(a[r:])
	Reverse(a)
}

// Given a list of strings, the nonempty function returns the non-empty ones
func nonempty(strings []string) []string {
	var ne []string
	for _, s := range strings {
		if len(s) > 0 {
			ne = append(ne, s)
		}
	}
	return ne
}

// Given a list of strings, the nonempty function returns the non-empty ones (in place; underlying array is modified
func nonemptyInPlace(strings []string) []string {
	i := 0
	for _, s := range strings {
		if len(s) > 0 {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

// A slice can be used to implement a stack
var stack []string
var EmptyStackError = errors.New("empty stack; can't pop")

func push(s string) {
	stack = append(stack, s)
}
func pop() (string, error) {
	n := len(stack)
	if n == 0 {
		return "", EmptyStackError
	}
	e := stack[n-1]
	stack = stack[:n-1]
	return e, nil
}

// Rewrite reverse to use an array pointer instead of a slice.
func reverseArray(p *[10]int) {
	for i, j := 0, len(p)-1; i < j; i, j = i+1, j-1 {
		p[i], p[j] = p[j], p[i]
	}
}

//  Write a version of rotate (that left rotates a given slice by r places) that operates in a single pass
func LeftRotateOnePass(a []int, r int) {
	// Is it possible to do this in place?
	// of course, an O(n)-space solution is readily available :)
	// one O(r)-space solution is possible
	n := len(a)
	r %= n
	buf := make([]int, r)
	copy(buf, a[:r])
	copy(a[:n-r], a[r:])
	//copy(a[n-r:], buf)
	a = append(a[:n-r], buf...)
	// The above is O(r)-space complexity, and it is reasonable if r << n.
	// Can we do better?
}

// Write an in-place function to eliminate adjacent duplicates in a []string slice
func UniqAdj(a []string) []string {
	j := 0 // last index (incl) of all unique strings: a[0] through a[j] are all adjacently unique
	for i := 1; i < len(a); i++ {
		if a[i] != a[j] {
			a[j+1] = a[i]
			j++
		}
	}
	return a[:j+1]
}

// Write an in-place function that squashes each run of adjacent Unicode spaces (see unicode.IsSpace)
// in a UTF-8-encoded []byte slice into a single ASCII space.
func SquashSpace(buf []byte) []byte {
	return SquashRuns(buf, func(curr rune, prev rune) bool {
		return unicode.IsSpace(curr) && unicode.IsSpace(prev)
	})
}

func SquashAnyRuneRun(buf []byte) []byte {
	return SquashRuns(buf, func(curr rune, prev rune) bool {
		return curr == prev
	})
}

func SquashRuneRun(buf []byte, r rune) []byte {
	return SquashRuns(buf, func(curr rune, prev rune) bool {
		return curr == r && curr == prev
	})
}

// SquashRuns removes "runs" of the given rune r from the given byte slice; returns the byte slice that has no runs
// of the given rune.
func SquashRuns(buf []byte, isInsideRun func(curr rune, prev rune) bool) []byte {
	// So, this function gets a UTF-8-encoded []byte slice and removes the repeated runs of a rune given a function that determines the run
	// from that slice; some tricky UTF-8 manipulation is required
	// We may assume that the encoded slice is valid UTF-8, but should protect against it
	totalBytes := len(buf)
	bytesRead := 0
	wStart := 0
	wEnd := 0
	prevCh := '\u0000' // previous character
	for ; bytesRead < totalBytes; {
		ch, chSize := utf8.DecodeRune(buf[bytesRead:])
		// assert that chSize is not zero; it may not be, otherwise it is a bug
		// even if ch == RuneError, we don't care, we just carry the encoding errors forward
		if isInsideRun(ch, prevCh) {
			// do nothing
		} else {
			copy(buf[wStart:], buf[bytesRead:bytesRead+chSize])
			wEnd = wStart + chSize
			wStart = wEnd
		}
		bytesRead += chSize
		prevCh = ch
	}
	return buf[:wEnd]
}

// Modify reverse to reverse the characters of a []byte slice that represents a UTF-8-encoded string, in place.
// Can you do it without allocating new memory?
func reverseRunesInPlace(a []byte) {
	// I think we can reverse the bytes of each rune and then reverse the entire slice of bytes
	totalBytes := len(a)
	bytesRead := 0
	for ; bytesRead < totalBytes; {
		_, chSize := utf8.DecodeRune(a[bytesRead:])
		// assert chSize != 0, if chSize == 1, then we take it as-is
		reverseBytesInPlace(a[bytesRead:bytesRead+chSize])
		bytesRead += chSize
	}
	reverseBytesInPlace(a)
}

func reverseBytesInPlace(a []byte) {
	n := len(a)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}
