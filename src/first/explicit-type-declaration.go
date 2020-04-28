// The named type provides a way to separate different and perhaps incompatible uses of the
// underlying type so that they can’t be mixed unintentionally.
package main

import "fmt"

func main() {
	type Celcius float64
	var t1 Celcius = 1.0
	var t2 = t1 + 1.0
	//var incr = 1.0
	//var t2 = t1 + incr // this won't compile: invalid operation: t1 + incr (mismatched types Celsius and float64)
	fmt.Printf("temp t2: %f, type of addition: %T\n", t2, t1 + 1.0)
}
