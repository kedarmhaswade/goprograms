// Demonstrates the := assignment statement and its effects
package main

import "fmt"

func main() {
	medals := []string{"gold", "silver", "bronze"}
	for i := len(medals) - 1; i >= 0; i-- {
		fmt.Println(medals[i])
	}
	// the above works because len returns a signed integer, if it were returning unsigned integer, then ...
	i := myLen(medals) // note: i is now an unsigned int, not just an int
	fmt.Printf("type of i is: %T and its value = %v\n", i, i)
	fmt.Printf("i - 1 = %v\n", i-1)
	fmt.Printf("i - 1 = %v\n", i-2)
	fmt.Printf("i - 1 = %v\n", i-3)
	fmt.Printf("i - 1 = %v\n", i-4)
}
func myLen(a []string) uint {
	return uint(len(a))
}
