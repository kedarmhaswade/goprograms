package main

import "math"

type Point struct {
	X, Y float64
}

type Path []Point // a slice of Points


//Distance returns the distance between points p and q
func (p Point) Distance(q Point) float64 {
	//dx := p.X - q.X
	//dy := p.Y - q.Y
	//dsq := dx*dx + dy*dy
	//return math.Sqrt(dsq)
	// Knowing libraries is important!
	return math.Hypot(p.X - q.X, p.Y - q.Y)
}

//Distance returns the distance traveled along the path
func (p Path) Distance() float64 {
	sum := 0.0
	for i := 1; i < len(p); i++ {
		sum += p[i].Distance(p[i-1])
	}
	return sum
}
func main() {
	path := Path{{1, 1}, {5, 1}, {5, 4}, {1, 1}, {}}
	// no need to declare the structs like below
	//path := Path{Point{1, 1}, Point{5, 1}, Point{5, 4}, Point{1,1}}
}
