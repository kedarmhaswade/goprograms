// Demonstrates the declaration of
package main

import "fmt"

func main() {
	s := "ाबहिसहेक अभिषेक ब भ क ख uber"
	// this was a new syntax for me:
	tests := []struct {
		ID   int
		Name string
	}{
		{1, "Foo"},
		{2, "Bar"},
		{3, "Baz"},
	}
	for _, test := range tests {
		fmt.Printf("%+v\n", test)
	}
}
