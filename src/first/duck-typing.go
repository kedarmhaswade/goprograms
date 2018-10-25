package main

import "fmt"

// Go seems to use duck typing: A type implements an interface by implementing all its methods

// define an interface and its methods
type Duck interface {
	// returns the number of legs of this Duck
	Legs() int
	// walks like a Duck
	Walk()
	// talks like a Duck
	Talk()
}
type Swan struct {
	legs int
}

func (swan Swan) Legs() int {
	return swan.legs
}
func (swan Swan) Walk() {
	fmt.Printf("Yay! %T Walking with %d legs!\n", swan, swan.legs)
}
func (swan Swan) Talk() {
	fmt.Println("...")
}
func main() {
	var s Duck = Swan{4}
	s.Walk()
	s.Talk()
}
