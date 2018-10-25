// Demo: Implicitly, the caller allocates an array, copies the arguments into it, and passes a slice of
// the entire array to the function.
package main

import "fmt"

func sumSlice(vals []int) int {
	total := 0
	for i, val := range vals {
		total += val
		fmt.Printf("address of %d: %p (sumSlice)\n", val, &vals[i])
	}
	return total
}
func sum(vals ...int) int { // declaration: ellipsis before the type
	total := 0
	for i, val := range vals {
		total += val
		fmt.Printf("address of %d: %p (sum)\n", val, &vals[i])
	}
	return total
}
func main() {
	a := 1
	b := 2
	c := 3
	fmt.Printf("address of %d: %p (main)\n", a, &a)
	fmt.Printf("address of %d: %p (main)\n", b, &b)
	fmt.Printf("address of %d: %p (main)\n", c, &c)
	_ = sum(a, b, c)
	slice := make([]int, 3)
	slice[0] = a
	slice[1] = b
	slice[2] = c
	_ = sumSlice(slice)
	fmt.Printf("address of %d: %p (main)\n", slice[0], &slice[0])
	fmt.Printf("address of %d: %p (main)\n", slice[1], &slice[1])
	fmt.Printf("address of %d: %p (main)\n", slice[2], &slice[2])
}
