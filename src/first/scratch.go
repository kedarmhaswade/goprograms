package main

import (
	"fmt"
	"strconv"
)

var status = true

func main() {
	a, _ := strconv.ParseUint("1011", 2, 4)
	b, _ := strconv.ParseUint("0101", 2, 4)
	r, _ := strconv.ParseUint("1010", 2, 4)
	fmt.Printf("%v\n", a&^b == r)
	a, _ = strconv.ParseUint("11010110", 2, 8)
	b, _ = strconv.ParseUint("10110110", 2, 8)
	r, _ = strconv.ParseUint("01000000", 2, 8)
	fmt.Printf("%v\n", a&^b == r)
	fmt.Printf("%v\n", 2.4/0.2)
	// a = 1 << 100 // compiler error
	fmt.Printf("a=%d\n", a)
	var unsigned uint
	fmt.Printf("%T %[1]v %T %[2]v %T %[3]v \n", unsigned, unsigned-1, 1-unsigned)
	i, j := 0, 1 // "short variable declaration" of multiple variables
	fmt.Printf("before: %d %d\n", i, j)
	i, j = j, i // tuple "assignment"
	fmt.Printf("after: %d %d\n", i, j)
}
