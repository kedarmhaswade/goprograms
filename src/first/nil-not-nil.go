// https://yourbasic.org/golang/gotcha-why-nil-error-not-equal-nil/
package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	//var err error = nil
	// …
	return err
}

func main() {
	err := Foo()
	fmt.Printf("type of return value: %T\n", err)
	fmt.Println(err)        // <nil>
	fmt.Println(err == nil) // false
}
