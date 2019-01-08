// Use alternate way to get the pop-count of an int64
package ch2

// PopCount1 finds the number of set bits in a given int64 number.
func PopCount1(x int64) int {
	pc := 0
	for x != 0 {
		x &= x-1 // clears the rightmost set bit (i.e. a 1) in x
		pc++
	}
	return pc
}
