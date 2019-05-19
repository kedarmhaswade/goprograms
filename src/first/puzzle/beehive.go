package puzzle

import (
	"fmt"
)

// Chooses any r items from an n-set. IOW, returns in acc, the r-sets of the given n-set.
// c is the current, incomplete set which this function will add to acc. len(c) is always r.
// The caller should provide a nil or empty slice as c.
// Returns the number of such r-sets, in other words, the value of "len(n) choose r"
func choose(n []int, r int, c []int, acc [][]int) [][]int {
	if r > len(n) {
		return acc // the trivial base case, do nothing
	}
	if r == 0 {
		// we have found an r-set, accumulate it and return
		acc = append(acc, c)
		//fmt.Printf("%v\n", acc)
		return acc
	}
	c1 := make([]int, len(c))
	c2 := make([]int, len(c))
	copy(c1, c)
	c1 = append(c1, n[0])
	copy(c2, c)
	c1 = append(c2, n[0])
	// the Pascal triangle formula!
	return append(choose(n[1:], r-1, c1, acc), choose(n[1:], r, c2, acc)...)
}

func insane38(r int) [][]int {
	n := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	c := make([]int, 0)
	acc := make([][]int, 0)
	acc = choose(n, r, c, acc)
	var acc38 [][]int
	for _, row := range acc {
		if sum(row) == 38 {
			acc38 = append(acc38, row)
		}
	}
	fmt.Printf("len(acc38): %v\n", len(acc38))
	fmt.Printf("acc38:\n%v\n", acc38)
	fmt.Printf("Frequencies: \n")
	frequencies := make(map[int]int)
	for _, row := range acc38 {
		for _, n := range(row) {
			frequencies[n] += 1
		}
	}
	for i := 1; i <= 19; i++ {
		fmt.Printf("%d: %d\n", i, frequencies[i])
	}
	return acc38
}

func sum(ints []int) int {
	sum := 0
	for _, i := range ints {
		sum += i
	}
	return sum
}
