package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Printf("%d\n", 1<<500)
	const n = 1 << 500 / (1 << 450)
	b := n == (1 << 50)
	fmt.Printf("%d\n", n)
	fmt.Printf("%v\n", b)

	x := 0.1
	y := 0.3
	runtimeAddition := x+x+x == y             // false!
	compileTimeAddition := 0.1+0.1+0.1 == 0.3 // true!
	fmt.Printf("compile %v run %v\n", compileTimeAddition, runtimeAddition)
	precise()
	imprecise()
	conv()
	untypedConstConv()
}
func precise() {
	var x float32 = math.Pi
	var y = math.Pi
	var z complex128 = math.Pi
	fmt.Printf("%v %v %v\n", x, y, z)
}
func imprecise() {
	const Pi64 float64 = math.Pi

	var x = float32(Pi64)
	var y = Pi64
	var z = complex128(Pi64)
	fmt.Printf("%v %v %v\n", x, y, z)
}
func conv() {
	var f float64 = 212
	fmt.Println(5.0 / 9.0 * (f - 32))
	n := 5.0
	d := 9.0
	fmt.Println(n / d * (f - 32))
}
func untypedConstConv() { // rather weird, but you know constants in Go are untyped, and hence unusual
	var f float64 = 3 + 0i // untyped complex -> float64
	fmt.Printf("f (with complex), value: %v, type: %T\n", f, f)
	f = 2 // untyped integer -> float64
	fmt.Printf("f (with complex), value: %v, type: %T\n", f, f)
	f = 1e123 // untyped floating-point -> float64
	fmt.Printf("f (with complex), value: %v, type: %T\n", f, f)
	f = 'a' // untyped rune -> float64 this would be 97
	fmt.Printf("f (with complex), value: %v, type: %T\n", f, f)
	// Above are equivalent to:
	//	var f float64 = float64(3 + 0i)
	//	f = float64(2)
	//	f = float64(1e123)
	//	f = float64('a')
}
func convertConstTypes() {
	const (
		deadbeef = 0xdeadbeef        // untyped int with value 3735928559
		a        = uint32(deadbeef)  // uint32 with value 3735928559
		b        = float32(deadbeef) // float32 with value 3735928576 (rounded up)
		c        = float64(deadbeef) // float64 with value 3735928559 (exact)
		//d = int32(deadbeef)   // compile error: constant 3735928559 overflows int32
		//e = float64(1e309)    // compile error: constant 1e+309 overflows float64
		//f = uint(-1)          // compile error: constant -1 overflows uint
	)
}
