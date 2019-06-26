// Let's say that at compile time you want to make sure that a type implements a specific interface.
// The trick used here is based on the type conversion specified in Go Language Spec: https://golang.org/ref/spec#Conversions
// An explicit conversion is an expression of the form T(x) where T is a type and x is an expression that can be converted to type T.

package main

import (
	"fmt"
)

type Greeter interface {
	Greet()
}
 type Man struct {
	Name string
}

// elsewhere ...
func (m *Man) Greet()  {
	fmt.Printf("Hello, My name is %s\n", m.Name)
}

var _ Greeter = (*Man)(nil) // this line ensures that *Man implements the Greeter interface.

func main() {
	var g Greeter = &Man{"Kedar"}
	g.Greet()
}
