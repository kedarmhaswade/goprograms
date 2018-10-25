// What happens when you pass a struct as a value to a function that modifies its fields
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Product struct {
	ID    int
	Name  string
	Likes int
}

func main() {
	p1 := Product{1, "Computer", 0}
	fmt.Printf("likes then: %d\n", p1.Likes)
	likeVal(p1)
	fmt.Printf("likes now: %d\n", p1.Likes)
	likePointer(&p1)
	fmt.Printf("likes now: %d\n", p1.Likes)
}
func likeVal(product Product) {
	rand.Seed(time.Now().UnixNano())
	product.Likes += rand.Intn(10)
}
func likePointer(product *Product) {
	rand.Seed(time.Now().UnixNano())
	product.Likes += rand.Intn(10) // note: the . notation works on pointers too!
}
