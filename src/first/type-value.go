package main

import "fmt"

func half(x float64) float64 {
	return x / 2
}
func main() {
	fmt.Printf("type of pure value: %T, returned value: %precise\n", 5, half(5))
}
