// Demonstrates use of bytes.Buffer
package main

import (
	"bytes"
	"fmt"
	"strings"
)

// recComma inserts commas in a decimal string which could optionally have a floating point and a sign.
func recComma(s string) string {
	if len(s) == 0 {
		return s
	}
	n := len(s)
	iDot := strings.IndexByte(s, '.')
	if (n <= 3 && iDot == -1) || (n <= 5 && iDot == n-2) || (n <= 6 && iDot == n-3) || (n <= 7 && iDot == n-4) {
		return s
	}
	if s[0] == '+' || s[0] == '-' {
		return string(s[0]) + recComma(s[1:])
	}
	if iDot == -1 {
		return recComma(s[:n-3]) + "," + s[n-3:]
	} else if iDot == n-2 {
		return recComma(s[:n-5]) + "," + s[n-5:]
	} else if iDot == n-3 {
		return recComma(s[:n-6]) + "," + s[n-6:]
	} else if iDot == n-4 {
		return recComma(s[:n-7]) + "," + s[n-7:]
	} else {
		//invalid string?
		panic(fmt.Sprintf("invalid string as index of decimal point is %d and length of string is %d", iDot, n))
	}
}

// This version (recComma) uses the string concatenation and recursion.

// Write a non-recursive version of comma, using bytes.Buffer instead of string concatenation.

func comma(s string) string {
	if len(s) == 0 {
		return s
	}
	var buf bytes.Buffer
	if s[0] == '+' || s[0] == '-' {
		buf.WriteByte(s[0])
		s = s[1:]
	}
	iDot := strings.IndexByte(s, '.')
	var n int
	if iDot == -1 {
		n = len(s)
	} else {
		n = len(s[:iDot])
	}
	for i := n; i > 0; i-- {
		if i%3 == 0 && i < n {
			buf.WriteByte(',')
		}
		buf.WriteByte(s[n-i])
	}
	if iDot != -1 {
		buf.WriteString(s[iDot:])
	}
	return buf.String()
}
func main() {
	s := "3181235663"
	fmt.Printf("%v\n", comma(s) == recComma(s))
	s = "318123"
	fmt.Printf("%v\n", comma(s) == recComma(s))
	s = "112"
	fmt.Printf("%v\n", comma(s) == recComma(s))
	s = "12"
	fmt.Printf("%v\n", comma(s) == recComma(s))
	s = "2"
	fmt.Printf("%v\n", comma(s) == recComma(s))
	s = ""
	fmt.Printf("%v\n", comma(s) == recComma(s))
	s = "54305463523612365463"
	fmt.Printf("%v\n", comma(s) == recComma(s))
	s = "-123450.59"
	println(comma(s))
	println(recComma(s))
	fmt.Printf("%v\n", comma(s) == recComma(s))
	s = "-123450.59"
	fmt.Printf("%v\n", comma(s) == recComma(s))
	s = "-123450.535"
	fmt.Printf("%v\n", comma(s) == recComma(s))
	s = "-12353450.5"
	fmt.Printf("%v\n", comma(s) == recComma(s))
}
