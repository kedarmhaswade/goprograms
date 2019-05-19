package puzzle

import (
	"fmt"
	"testing"
)

func TestChoose(t *testing.T) {
	n := []int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	r := 3
	c := make([]int, 0)
	acc := make([][]int, 0)
	exp := 19*3*17
	act := choose(n, r, c, acc)
	fmt.Printf("exp: %d, act: %d\n", exp, len(act))
	fmt.Printf("%v\n", act)
}

func TestInsane38(t *testing.T) {
	insane38(3)
	insane38(4)
	insane38(5)
}


