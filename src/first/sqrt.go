package main

import (
	"fmt"
	"math"
)

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
func Sqrt(x float64) float64 {
	z := 1.0
	zsq := z * z
	diff := abs(zsq - x)
	for diff > 0.000001 {
		if zsq > x {
			z -= (z*z - x) / (2 * z)
		} else {
			z += (x - z*z) / (2 * z)
		}
		zsq = z * z
		diff = abs(zsq - x)
	}
	return z
}

func main() {
	x := 0.5
	fmt.Printf("mine: %g, library: %g\n", Sqrt(x), math.Sqrt(x))
}
