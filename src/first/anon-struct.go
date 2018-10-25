// structs are a bit confusing
package main

import "fmt"

type Inner struct {
	Name string
}

type Outer struct {
	Inner
	ID int
}

func main() {
	var outer  = &Outer{Inner{Name:"Larry"}, 1}
	fmt.Printf("%v\n", outer.Name)
}
