// Use panic and recover to write a function that contains no return statement yet returns a non-zero value.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func functionOfNoReturnWithFuncLiteral() (x int) {
	defer func() {
		p := recover()
		x = p.(int)
	}()
	rand.Seed(time.Now().UnixNano())
	panic(rand.Intn(10))
}
func main() {
	fmt.Printf("%d\n", functionOfNoReturnWithFuncLiteral())
	fmt.Printf("%d\n", functionOfNoReturn())
}

func functionOfNoReturn() (x int) {
	rand.Seed(time.Now().UnixNano())
	y := rand.Intn(10)
	defer ff(&x)
	panic(y)
}
func ff(ptr *int) {
	p := recover()
	*ptr = p.(int)
}
