// Demonstrate the lifetime of a loop variable
package ch2

import "fmt"

func loopVars() {
	for i := 0; i < 3; i++ {
		var j  = byte(i + 1)
		fmt.Printf("&i: %p, &j: %p\n", &i, &j)
	}
}
