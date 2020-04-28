package main

import (
	"fmt"
	"strconv"
)

// I was intrigued by the type switch like the following:
// func sprint(x interface{}) string {
//	type stringer interface {
//		String() string
//	}
//	switch x := x.(type) { // type switch
// ...
// how is it that the variable declared x has a different type? -- that is because a new scope is created by switch
// wanted to investigate

func foo1(x int) {
	{ // without this block, x := strconv.Itoa(x) won't work
		x := strconv.Itoa(x) // the x defined here is different from the function's argument
		fmt.Printf("%v\n", x)
	} // without this block, x := strconv.Itoa(x) won't work

}

func main() {
	foo1(0x64)
}
