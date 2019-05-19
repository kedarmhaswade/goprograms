// The defer statement is novel. Some experiments are in order. There are other Go features that help make this happen of course.
package ch5

import (
	"fmt"
	"math"

	"first/ch6"
)

// Print the parameters and results of a function
func distance(p1 ch6.Point, p2 ch6.Point) (d float64) {
	defer func() {
		fmt.Printf("Parameters: %v and %v, return value: %f\n", p1, p2, d)
	}()
	dx := p2.X - p1.X
	dy := p2.Y - p1.Y
	d = math.Sqrt((dx * dx) + (dy * dy))
	return d
}
