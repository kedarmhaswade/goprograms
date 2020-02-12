package main

import (
	"fmt"
	"math"
)

func main() {
	//var f float32 = 0.15625
	//f, _ := strconv.ParseFloat("0.1", 32)
	//var g float64 = 1.0/2
	printBits32(0.5)
	//var g float32 = 0.199999999
	//println(g == float32(f * 2.0))
	//printBits64(g)
}

func printBits32(f float32) {
	bits := math.Float32bits(f)
	fmt.Printf("%b\n", bits)
	bs := make([]int, 32)
	i := 0
	for bits != 0 {
		b := bits & 1
		bs[i] = int(b)
		bits >>= 1
		i += 1
	}
	sign := "+"
	for i := 31; i >= 0; i-- {
		switch  {
		case i == 31:
			fmt.Printf("s")
			if bs[31] & 1 == 1 {
				sign = "-"
			}
		case i <= 30 && i >= 23:
			fmt.Printf("e")
		default:
			fmt.Printf("f")
		}
	}
	println()
	for i := 31; i >= 0; i-- {
		fmt.Printf("%d", bs[i])
	}
	println()
	println(sign)
}