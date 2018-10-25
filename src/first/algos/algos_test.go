package algos

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimpleArrays(t *testing.T) {
	a := []int{-9, -8, -4, -3, -1, 2, 5, 6, 10}
	k := 2
	kss := KthSmallestSquareInSortedArray(a, k) // 2nd smallest square
	if kss != 4 {
		t.Errorf("2nd smallest square -- expected: %d, actual: %d", 9, kss)
	}
	a = []int{-20, -10, -9, 8, 11, 13, 15, 18}
	k = 3
	kss = KthSmallestSquareInSortedArray(a, k) // 3rd smallest square
	if kss != 100 {
		t.Errorf("2nd smallest square -- expected: %d, actual: %d", 100, kss)
	}
}
func TestA1(t *testing.T) {
	a := []int{-5, -4, -3, -1}
	kss := KthSmallestSquareInSortedArray(a, 2) // 2nd smallest square
	fmt.Printf("%d\n", kss)
}
func TestZeroes(t *testing.T) {
	a := []int{-4, 0, 0, 2, 4, 8}
	kss := KthSmallestSquareInSortedArray(a, 4) // 4th smallest square
	fmt.Printf("%d\n", kss)
}
func TestEquals(t *testing.T) {
	a := []int{-4, -1, 0, 2, 2, 4, 8}
	kss := KthSmallestSquareInSortedArray(a, 4) // 2nd smallest square
	fmt.Printf("%d\n", kss)
}

func TestAllPositive(t *testing.T) {
	a := []int{1, 3, 4, 8, 12, 14, 23}
	kss := KthSmallestSquareInSortedArray(a, 4) // 2nd smallest square
	fmt.Printf("%d\n", kss)
}

func TestThreeDigits(t *testing.T) {
	digits := []string{"1", "2", "3"}
	exprs := expressions(digits)
	assert.Contains(t, exprs, "1+2+3")
	assert.Contains(t, exprs, "1+2-3")
	assert.Contains(t, exprs, "1+23")
	assert.Contains(t, exprs, "1-2+3")
	assert.Contains(t, exprs, "1-2-3")
	assert.Contains(t, exprs, "1-23")
	assert.Contains(t, exprs, "12+3")
	assert.Contains(t, exprs, "12-3")
	assert.Contains(t, exprs, "123")
	for _, expr := range exprs {
		fmt.Printf("%v\n", evaluate(expr))
	}
}

func TestCenturies(t *testing.T) {
	printCenturies()
}
