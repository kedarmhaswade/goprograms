package main

import "fmt"

type Vertex struct {
	x, y float64
}
func (vp *Vertex) ScaleBy(f float64) *Vertex {
	vp.x *= f
	vp.y *= f
	return vp
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
	v := (&Vertex{1, 1}).ScaleBy(2)
	fmt.Printf("scaled: %v\n", v)
}
