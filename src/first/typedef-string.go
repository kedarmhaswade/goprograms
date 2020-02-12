package main

import "fmt"
type foo int

func (f foo) String() string {
	x := (string)(f)
	return fmt.Sprint(x)
}
func main() {
	var f foo = foo(2315)
	fmt.Printf("%v\n", string(f))
	fmt.Printf("%v\n", f.String())
}