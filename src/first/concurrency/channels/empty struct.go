// Demonstrates the use of "empty struct"
package main

import "fmt"

type empty struct{}

func main() {
	s := empty{}
	// this can be short-circuited to: e := struct{}{} which can be a bit confusing, but
	// think that empty is the "type" which can be declared in place as struct{} whose
	// literal Value is {}, and therefore, one can declare an empty struct "Value" as struct{}{}
	t := struct{}{}
	fmt.Printf("is s == t? %v\n", s == t)
}
