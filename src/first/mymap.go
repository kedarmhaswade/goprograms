// Demonstrates my own version of strings.map
package main

import "fmt"

func myMap(mapping func(rune) rune, in string) string {
	m := make([]rune, len(in))
	for _, c := range in {
		r := mapping(c)
		m = append(m, r)
		fmt.Printf("%c->%c ", c, r)
	}
	out := string(m[:])
	fmt.Printf("length of given string: %d, length of mapped string: %d\n", len(in), len(out))
	return out
}
func main() {
	p1 := func(r rune) rune { return r + 1 }
	m1 := func(r rune) rune { return r - 1 }
	o := "hello, world"
	s := myMap(p1, o)
	fmt.Printf("s: %s, len(s): %d\n", s, len(s))
	t := myMap(m1, s)
	fmt.Printf("t: %s, len(t): %d\n", t, len(t))
	fmt.Printf("got it back? %v\n", o == t)
}
