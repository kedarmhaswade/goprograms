package main

import "fmt"

type Vertex struct {
	x, y int
}

func main() {
	var p *int
	fmt.Printf("%T=%d\n", p, p)
	// struct literals
	fmt.Printf("type: %T, value: %v\n", Vertex{y: 1}, Vertex{y: 2})
	i := 2
	p = &i
	*p += 1
	fmt.Println(i)
}
