package main

import "fmt"

// Types are fundamental in any programming language. From gopl.io:
// Conversions are also allowed between numeric types, and between string and some slice types.
// These conversions may change the representation of the value.
// Converting a string to a []byte slice allocates a copy of the string data.

// This program demonstrates the above.

func main() {
	type aname string
	s := "foo" // some test string
	b := []byte(s) // just by "type conversion", this allocates a copy of the string data
	b[0] = 'b'
	fmt.Printf("%v\n", s)
	for _, c := range b {
		fmt.Printf("%c", c)
	}
	println()
	var n  = aname(s) // in this type conversion, however, there is no copy of the data, both variables are names to the same location
	fmt.Printf("address of s: %p\n", &s)
	fmt.Printf("address of n: %p\n", &n)
	// another demonstration that the value of a type may not change
	type Celsius float64
	type Fahrenheit float64
	var c Celsius = 100 // 100 degrees celsius
	var f Fahrenheit = 100 // 100 degrees fahrenheit
	//fmt.Println(f == c) // does not compile, since f and c are values of different types
	fmt.Println(f == Fahrenheit(c))  // prints true!, there is no change in the value -- this also does not mean that 100°c == 100°f
}