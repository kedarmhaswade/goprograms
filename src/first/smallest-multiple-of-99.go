package main

import "fmt"

// A sloppy computer solution to this beautiful problem:
// https://mathschallenge.net/view/divisible_by_99
// Find the smallest number that is made up of each of the digits 1 through 9 exactly once
// and is divisible by 99.

// Of course, a "math solution" should first be devised and a stupid program could verify the answer :-)

func main() {
	var a int64 = 123456789
	for {
		if a%99 == 0 && uniqueNonzeroDigits(a) {
			fmt.Printf("%v\n", a)
			return
		}
		a++
	}
}

func uniqueNonzeroDigits(a int64) bool {
	var digits [9]int
	for i := 8; i >= 0; i-- {
		d := int(a % 10)
		if d == 0 {
			return false
		}
		digits[i] = d
		a /= 10
	}
	// stupid n^2 algo :-(
	for i := 0; i < 8; i++ {
		for j := i + 1; j < 9; j++ {
			if digits[i] == digits[j] {
				return false
			}
		}
	}
	return true
}
