package main

import "fmt"

type Name string

func main()  {
	myName := Name("foo")
	yrName := "bar"
	//fmt.Printf("does not compile! %v", myName == yrName)
	fmt.Printf("explicit type conversion is needed: %v\n", myName == Name(yrName))
	fmt.Printf("explicit type conversion is needed: %v\n", yrName == string(myName))
}
