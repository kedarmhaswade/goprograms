package ch4

import (
	"fmt"
)

// struct types tend to be verbose because they often involve a line for each field.
// Although we could write out the whole type each time it is needed, the repetition would get tiresome.
// Instead, struct types usually appear within the declaration of a named type like Employee.

// Here's the ultimate unnamed type specification -- look at the amount of repetition
var larry = struct {
	ID        int
	LastName  string
	FirstName string
	Salary    int
}{1, "Stooge", "Larry", 100}
var moe = struct {
	ID        int
	LastName  string
	FirstName string
	Salary    int
}{1, "Stooge", "Moe", 110}

// Here's how we make that type into a named type
type E struct {
	ID        int
	LastName  string
	FirstName string
	Salary    int
}
var curly = E{2, "Stooge", "Curly", 120}

func list() {
	fmt.Printf("%T\n", larry)
	fmt.Printf("%T\n", moe)
	fmt.Printf("%T\n", curly)
}
