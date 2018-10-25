package main

import "fmt"

/** Complete the following function. Double the capacity if the slice needs to "grow"*/
func AppendByte(slice []byte, data ...byte) []byte {
	m := len(slice) // existing
	n := len(data)  // new
	if (m + n) > cap(slice) {
		newSlice := make([]byte, 2*(m+n+1))
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:(m + n)]
	copy(slice[m:m+n], data)
	return slice
}
func main() {
	slice := []byte{'a', 'b', 'c'}
	slice = AppendByte(slice, 'd', 'e')
	fmt.Printf("slice: %c len(slice): %d cap(slice): %d\n", slice, len(slice), cap(slice))
	a := []string{"John", "Paul"}
	b := []string{"George", "Ringo", "Pete"}
	a = append(a, b...) // equivalent to "append(a, b[0], b[1], b[2])"
	fmt.Printf("%v\n", a)
}
