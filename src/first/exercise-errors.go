package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot take a real square root of a negative number: %v", float64(e))
}

func Sqrt1(x float64) (float64, error) {
	if x >= 0 {
		return sqrtInner(x), nil
	}
	return 0, ErrNegativeSqrt(x)
}

func sqrtInner(x float64) float64 {
	z := 1.0
	zsq := z * z
	diff := absInner(zsq - x)
	for diff > 0.000001 {
		if zsq > x {
			z -= (z*z - x) / (2 * z)
		} else {
			z += (x - z*z) / (2 * z)
		}
		zsq = z * z
		diff = absInner(zsq - x)
	}
	return z
}
func absInner(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
func main() {
	fmt.Println(Sqrt1(2))
	fmt.Println(Sqrt1(-2))
}
