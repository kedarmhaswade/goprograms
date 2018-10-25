// demonstrates the spirit of Go with a simple map.get like operation
package main

import "fmt"

func main() {
	monthNames := make(map[int]string) // a map from int to string
	monthNames[0] = "Jan"
	monthNames[1] = "Feb"
	i := 2
	z := "nil"
	ok := false
	if z, ok = monthNames[i]; ok { // note: z is overwritten here!
		fmt.Printf("Name of 0th month: %v\n", z)
	} else {
		fmt.Printf("No name for the month %d\n", i)
		fmt.Printf("type of z: %T, value of z: %v\n", z, z)
		fmt.Printf("type of nil is %T\n", nil)
	}
}
