package main

import (
	"fmt"
)

func squarer(x int) int {
	return x*x
}
func cuber(x int) int {
	return x*x*x
}
func main() {
	m := map[string]func(x int) int {}
	m["squarer"] = squarer
	m["cuber"] = cuber
	f := m["nil"]
	if f != nil {
		f(2)
	}
	if _, ok := m["nil"]; !ok {
		fmt.Println("the function does not exist in the map!")
	}
}
