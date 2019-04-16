// Demonstrates methods using simple constructs like a point and a path which is a sequence of points
package ch6

import "math"

// Point is a named type whose underlying type is a struct of
// two floating point numbers, namely X- and Y- coordinates of a 2-D point
type Point struct {
	X, Y float64
}

// Path is a named type whose underlying type is a slice (sequence) of Points
// We can define new methods of this slice. In this respect, Go is quite unique that it allows
// the programmer to define arbitrary methods on values of any type
type Path []Point

// As a convention, we'll only define methods whose receivers are pointers to named types.
// Go helps do the necessary type conversions when a value (and not a pointer) is provided as a receiver.

//Distance returns the cartesian distance of the given Point 'q' from the receiver Point
func (p *Point) Distance(q *Point) float64 {
	//return math.Sqrt((p.X - q.X)*(p.X - q.X) - (p.Y - q.Y)*(p.Y - q.Y))
	return math.Hypot(p.X - q.X, p.Y - q.Y) // takes care of several edge cases
}

//Distance calculates the "length" of the given Path. Remember, since Path is a slice of Points, it is already a reference
func (path Path) Distance() float64 {
	s := len([]Point(path))
	if s <= 1 {
		return 0
	}
	pathLen := 0.0
	for i := 1; i < s; i++ {
		pathLen += path[i].Distance(&path[i-1])
	}
	return pathLen
}