// In Go, calling a function makes a copy of each argument value, if a function needs to update a variable,
// or if an argument is so large that we wish to avoid copying it, we must pass the address of the variable
// using a pointer.
// Slices are already references. How can you prove that?
package ch3

import "fmt"

func callerOfFunctionWithSliceArg() {
	a := []int {1, 2, 3}
	fmt.Printf("in caller -- type: %T, value: %[1]p\n", a)
	argValuePrinter(a)
}
func argValuePrinter(a []int) {
	fmt.Printf("in called -- type: %T, value: %[1]p\n", a)
}

func callerOfFunctionWithArrayArg() {
	a := [3]int {1, 2, 3}
	fmt.Printf("in caller -- type: %T, value: %p\n", a, &a)
	arrayArgValuePrinter(a)
}
func arrayArgValuePrinter(a [3]int) {
	fmt.Printf("in called -- type: %T, value: %p\n", a, &a)
	// entire array is copied since the addresses are different, Go passes arguments by value
}

type Employee struct {
	Name string
	Id int
}
func callerOfFunctionWithStructArg() {
	e := Employee{"Kedar", 7}
	fmt.Printf("in caller -- type: %T, value: %p\n", e, &e)
	structArgValuePrinter(e)
}
func structArgValuePrinter(e Employee) {
	fmt.Printf("in called -- type: %T, value: %p\n", e, &e)
	// entire struct is copied since the addresses are different, Go passes arguments by value
}
