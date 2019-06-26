// It's not _pedantic_ to ensure stability of a sorting algorithm
package main

import (
	"fmt"
	"sort"
)

type ByValue []*int

func (a ByValue) Len() int {
	return len(a)
}

func (a ByValue) Less(i, j int) bool {
	// stable?
	// return *a[i] < *a[j]
	return !(*a[i] > *a[j])
}

func (a ByValue) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {
	numbers := makeIt(3)
	fmt.Printf("before: 0: %v %v\n", numbers[0], *numbers[0])
	sort.Stable(ByValue(numbers))
	fmt.Printf(" after: 0: %v %v\n", numbers[0], *numbers[0])
}

func makeIt(n int) []*int {
	numbers := make([]*int, 3)
	numbers[0] = new(int)
	*numbers[0] = 1
	numbers[1] = new(int)
	*numbers[1] = 1
	numbers[2] = new(int)
	*numbers[2] = 2
	return numbers
}
