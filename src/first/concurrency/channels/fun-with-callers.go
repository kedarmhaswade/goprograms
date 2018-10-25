// fun with runtime.Callers
package main

import (
	"fmt"
	"runtime"
)

func main() {
	f()
}
func f() {
	g()
}
func g() {
	fmt.Println("this is the callee, i.e. the called function g")
	pc, file, line, ok := runtime.Caller(1)
	fmt.Printf("pc: %v, file: %v, line: %v, ok: %v\n", pc, file, line, ok)
	var pc1 [3]uintptr
	_ = runtime.Callers(3, pc1[:])
	for _, y := range pc1 {
		fmt.Printf("%v\n", y)
	}
}
