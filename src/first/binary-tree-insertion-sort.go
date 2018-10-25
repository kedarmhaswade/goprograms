// Demonstrates the use of a binary tree to implement insertion sort
package main

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

// sort sorts values in place.
func sortInPlace(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// add adds n to the bst pointed to by root and returns the root of the resulting bst
func add(root *tree, n int) *tree {
	if root != nil {
		if n <= root.value {
			root.left = add(root.left, n)
		} else {
			root.right = add(root.right, n)
		}
	} else {
		root = &tree{value: n}
	}
	return root
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}
func main() {
	a := []int{32, 2, 3, 1, 5, 12, 7, 1, 24, 0}
	//a := []int{32, 2, 3}
	sortInPlace(a)
	fmt.Printf("%+v\n", a)
}
