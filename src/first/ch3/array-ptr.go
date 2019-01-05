package ch3

import "fmt"

// Arrays are "values" or rather "immediate values" in Go. When they are passed to functions,
// they are "copied" i.e. their contents are copied. The pointers to arrays, being references however, although
// copied, can be used to avoid the copying of the contents.

func studyPointersToArrays() {
	a := [...]int{1, 20, 300} // an array
	ap := &a // pointer to that array; type of ap is *[3]int
	addressOf0th := &a[0] // the type of addressOf0th is *int
	fmt.Printf("value of pointer to array: %p\n", ap)
	fmt.Printf("address of the 0th element of the array: %v\n", addressOf0th)
	arrayZerothElementAddressPrinter1(a)
	arrayZerothElementAddressPrinter2(ap)
}

// Just prints the address of the 0th element of the given array
func arrayZerothElementAddressPrinter1(a [3]int) {
	fmt.Printf("address of the 0th element of the array: %v\n", &a[0])
}
// Just prints the address of the 0th element of the given array ptr; should be the same as the one in caller
func arrayZerothElementAddressPrinter2(ap *[3]int) {
	fmt.Printf("address of the 0th element of the array: %v\n", &ap[0])
}

