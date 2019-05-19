// There can be many different binary trees with the same sequence of values.
// A function to check whether two binary trees store the same sequence is quite complex in most languages.
// We'll use Go's concurrency and channels to write a simple solution.
// https://tour.golang.org/concurrency/7
package main

import (
	"fmt"
)

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func inorder(root *Tree, ch chan int) {
	if root == nil {
		return
	}
	if root.Left != nil {
		inorder(root.Left, ch)
	}
	ch <- root.Value
	if root.Right != nil {
		inorder(root.Right, ch)
	}
}

func callInorder(root *Tree, ch chan int) {
	inorder(root, ch)
	close(ch)
}

func main() {
	isSame(tree1(), tree2())
	isSame(tree1(), tree3())
	isSame(tree2(), tree3())
	isSame(tree3(), tree3()) // hee hee
	isSame(tree1(), tree4())
	isSame(tree3(), tree5())
}

func isSame(tree1, tree2 *Tree) bool {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	go callInorder(tree1, ch1)
	go callInorder(tree2, ch2)
	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		if !ok1 && !ok2 { // both closed
			fmt.Printf("Trees represent the same sequence\n")
			return true
		}
		if !(ok1 && ok2) { // only one of them closed
			fmt.Printf("Trees do *not* represent the same sequence [One of the trees is bigger]\n")
			return false
		}
		if v1 != v2 {
			fmt.Printf("Trees do *not* represent the same sequence [Values (%d, %d) differ]\n", v1, v2)
			return false
		}
	}
}
func tree1() *Tree {
	root := &Tree{Value: 3}
	root.Left = &Tree{Value: 1}
	root.Left.Left = &Tree{Value: 1}
	root.Left.Right = &Tree{Value: 2}

	root.Right = &Tree{Value: 8}
	root.Right.Left = &Tree{Value: 5}
	root.Right.Right = &Tree{Value: 13}
	return root
}
func tree2() *Tree {
	root := &Tree{Value: 13}
	root.Left = &Tree{Value: 8}
	root.Left.Left = &Tree{Value: 5}
	root.Left.Left.Left = &Tree{Value: 3}
	root.Left.Left.Left.Left = &Tree{Value: 2}
	root.Left.Left.Left.Left.Left = &Tree{Value: 1}
	root.Left.Left.Left.Left.Left.Left = &Tree{Value: 1}
	return root
}

func tree3() *Tree {
	root := &Tree{Value: 10}
	root.Left = &Tree{Value: 8}
	root.Right = &Tree{Value: 14}
	return root
}
func tree4() *Tree {
	return nil
}

func tree5() *Tree {
	root := &Tree{Value: 8}
	root.Right = &Tree{Value: 10}
	root.Right.Right = &Tree{Value: 14}
	return root
}
