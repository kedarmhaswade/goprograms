// Solves a recurrence using a computer program: How many 00-free n-bit strings are there?
package main

import (
	"fmt"
	"strings"
)

func main() {
	n := 8
	strs := make([]string, 0)
	strs = append(strs, "0", "1")
	c, strs := countZZFree(n, strs, 2) // there are 2 00-free strings when n = 1
	fmt.Printf("number of %v-bit strings that are 00-free is: %v\n", n, c)
	for _, s := range strs {
		fmt.Printf("00-free: %v\n", s)
	}
}

func countZZFree(n int, strs []string, acc int) (int, []string) {
	//fmt.Printf("bitstring %v, its len: %v, n: %v\n", strs[0], len(strs[0]), n)
	if n < 1 || len(strs[0]) >= n {
		return acc, strs
	}
	newStrs := make([]string, 0)
	newAcc := 0
	bits := []string{"0", "1"}
	for _, s := range strs {
		for _, b := range bits {
			if strings.HasSuffix(s, "1") || b == "1" {
				newAcc += 1 // one more 00-free string
				newStrs = append(newStrs, s + b)
			}
		}
	}
	return countZZFree(n, newStrs, newAcc)
}
