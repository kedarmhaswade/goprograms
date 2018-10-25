package main

import "fmt"

func main() {
	stooges := []string{"larry", "moe", "curly"}
	fmt.Println(stooges)
	m := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(m)
	ints5 := [5]int{}
	fmt.Println(ints5)
	s := make([]byte, 5, 5)
	fmt.Println(len(s), cap(s))
	s = s[2:cap(s)]
	fmt.Println(len(s), cap(s))
	s = s[0:2]
	fmt.Println(len(s), cap(s))
}
