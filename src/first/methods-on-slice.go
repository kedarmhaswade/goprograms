package main

import "fmt"

type Set []string

func (s Set) println()  {
	for _, v := range s {
		fmt.Printf("%s\n", v)
	}
}
func main() {
	s := Set([]string{"a", "b", "c"})
	s.println()
}
