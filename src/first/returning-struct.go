// Demonstrates what happens when a struct or a pointer thereof is returned by a function
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Family struct {
	Name    string
	Members []string
}

var clowns = Family{"Clown", []string{"John", "Mary", "Jim", "Cathy"}}
var darks = Family{"Washmade", []string{"Peeda", "Radek", "Artuju", "Proova"}}
var stooges = Family{"Stooge", []string{"Larry", "Moe", "Curly"}}
var families = []Family{clowns, darks, stooges}

func getFamilyPtr() *Family {
	rand.Seed(time.Now().UnixNano())
	idx := rand.Intn(len(families))
	return &families[idx]
	// This is a subtle point! When we assign a struct value to a variable, the entire struct is copied
	// because struct is always an "immediate value" and not a reference!
	// if, instead, the above line: return &families[idx] were uncommented, it would return the
	// address of the struct that was already present in the families slice!
	// r := families[idx]
	//return &r
}
func getFamily() Family {
	rand.Seed(time.Now().UnixNano())
	r := families[rand.Intn(len(families))]
	return r
}
func main() {
	getFamilyPtr().Name = "Random"
	//getFamily().Name = "Random"
	// Another subtle point. Since getFamily() returns a struct, the left hand side of this assignment
	// does not identify a variable and as such, a compile error is generated:
	// ./returning-struct.go:38:19: cannot assign to getFamily().Name
	for i := 0; i < len(families); i++ {
		fmt.Printf("%+v\n", families[i])
	}
}
