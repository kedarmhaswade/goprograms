package main

import "fmt"

func main() {
	var target []string
	//ss := []string{"joe", "molly"}
	// names := nil
	// name := "joe"
	src := source()
	// copy source to target
	target = make([]string, len(src))
	fmt.Printf("len = %v\n", len(target))
}

func source() []string  {
	return []string{"larry", "moe", "curly"}
}