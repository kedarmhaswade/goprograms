// In Go, slices are references. Thus, generally, there is no need to have pointers to slices, at least not
// to save the "so-called" copying of contents. This is so because what is copied (when passing slices to
// functions that accept them) on the stack are copies of address to the slice on which the function operates.
package main

import (
	"fmt"
)

func main() {
	a := make([]int, 10)
	fmt.Printf("I have a slice of 10 ints. Let's print the info of its zeroth element.\n")
	//n := len(a)
	fmt.Printf("The 'address' of element[%d] is %p and its 'value' is %d.\n", 0, &a[0], a[0])
	fmt.Printf("The value contained in the variable a that denotes the above slice (by e.g.: 'a := make([]int, 10)') is: %p.\n", a)
	fmt.Printf("But note that the address of the variable a itself is %p\n", &a)
	fmt.Printf("Thus, the contents of location %p are %p\n", &a, a)
	fmt.Printf("Let's pass this slice to a function that accepts an int- slice:\n")
	initSlice(a)
}
func initSlice(b []int) {
	fmt.Printf("The 'address' of the element[%d] of slice inside the function is %p and its 'value' is %d.\n", 0, &b[0], b[0])
	fmt.Printf("The value contained in the variable b that denotes the slice is: %p.\n", b)
	fmt.Printf("But note that the address of the variable b itself is %p\n", &b)
	fmt.Printf("Thus, the contents of location %p are %p\n", &b, b)
}

