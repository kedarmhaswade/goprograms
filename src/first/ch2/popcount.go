package ch2

// pc is an array that holds the population count of each of bytes whose values are from 1 through 256.
// pc[i] = popcount(i) for all 1 <= i < 256
var pc [256]byte

func init() {
	// populates the popcount of every number from its half and its last bit!
	for i:= 0; i < 256; i++ {
		pc[i] = pc[i>>1] + byte(i&1)
	}
}

// PopCount finds the number of set bits in a given int64 number.
func PopCount(x int64) int {
	c := 0
	for i := 0; i < 8; i++ {
		c += int(pc[byte(x & 255)])
		x >>= 8
	}
	return c
}