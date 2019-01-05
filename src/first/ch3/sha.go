package ch3

import (
	"crypto/sha256"
	"errors"
	"first/ch6"
	"fmt"
	"strconv"
)

func formatBinary() {
	x1 := strconv.FormatInt(int64('a'), 2)
	x2 := strconv.FormatInt(int64('A'), 2)
	fmt.Printf("%v\n", x1)
	fmt.Printf("%v\n", x2)
}

//BitDiff returns the number of bits in which the given byte slices differ. To calculate the bit diff of
// two slices, the slices must have the same length, otherwise an error is returned. The bytes at the same
// index are compared bit-by-bit and the cumulative difference is calculated. If the total number of differences
// exceeds int64 max value, then an error is returned. Returned error is nil in case of success.
func BitDiff(first, second []byte) (int64, error) {
	l1 := len(first)
	l2 := len(second)
	if l1 != l2 {
		return -1, errors.New(fmt.Sprintf("Unequal array length, first: %d, second: %d", l1, l2))
	}
	var difbits int64
	for i := 0; i < l1; i++ {
		d := int64(ch6.PopCount(uint64(first[i] ^ second[i])))
		difbits += d
	}
	//todo check overflow
	return difbits, nil
}
func shaFun() {
	x := "a"
	y := "A"
	i, _ := BitDiff([]byte(x), []byte(y))
	fmt.Printf("input slices differ in %d bits\n", i)
	shax := sha256.Sum256([]byte(x))
	shay := sha256.Sum256([]byte(y))
	i, _ = BitDiff(shax[:], shay[:])
	fmt.Printf("their sha's differ in %d bits\n", i)
}