package ch5

import (
	"testing"

	"first/ch6"
)

func TestDistance(t *testing.T) {
	origin := ch6.Point{0, 0}
	p := ch6.Point{3, 4}
	distance(p, origin)
}

