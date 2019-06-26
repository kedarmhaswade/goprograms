// Sums up all the int32 numbers and produces a big result in parallel :-)
package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func sumAll(loInc, hiInc int32, nc int) {
	c := make(chan int64)
	var r = int(hiInc-loInc) / nc
	//fmt.Printf("range: %v\n", r)
	l := loInc
	h := int32(0)
	for i := 0; i < nc - 1; i++ {
		h = l + int32(r-1)
		//fmt.Printf("call %v: lo: %v, hi: %v\n", i, l, h)
		go sum(l, h, c)
		l = h + 1
	}
	go sum(l, hiInc, c)
	var all int64
	for i := 0; i < nc; i++ {
		all += <-c
	}
	fmt.Printf("loInc: %v, hiInc: %v, sum: %v\n", loInc, hiInc, all)
}

func sum(first int32, last int32, c chan int64) {
	var res int64
	var i uint32
	for i = uint32(first); i <= uint32(last); i++ {
		res += int64(i);
	}
	//fmt.Printf("%v to %v = %v\n", first, last, res)
	c <- res
}
func main() {
	var nc int64
	var err error
	if len(os.Args) < 2 {
		nc = 10
	} else {
		nc, err = strconv.ParseInt(os.Args[1], 10, 32)
	}
	if err != nil {
		_ = fmt.Errorf("Usage: %s [number-of-goroutines = 10]\n", nc)
	}
	for i := 0; i < 10000; i++ {
		sumAll(0, math.MaxInt32, int(nc))
	}
}
