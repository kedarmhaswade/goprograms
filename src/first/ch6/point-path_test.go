package ch6

import (
	"fmt"
	"testing"
)

func TestSimplePoint(t *testing.T) {
	origin := Point{0, 0}
	p1 := Point{3, 0}
	p2 := Point{3, 4}
	expP1 := 3.0
	act := p1.Distance(&origin)
	if expP1 != act {
		t.Errorf(fmt.Sprintf("expected: %f, actual: %f", expP1, act))
	}
	expP1P2 := 4.0
	act = p1.Distance(&p2)
	if expP1P2 != act {
		t.Errorf(fmt.Sprintf("expected: %f, actual: %f", expP1P2, act))
	}
	expP2Origin := 5.0
	act = p2.Distance(&origin)
	if expP2Origin != act {
		t.Errorf(fmt.Sprintf("expected: %f, actual: %f", expP2Origin, act))
	}
}

func TestSimplePath(t *testing.T) {
	origin := Point{0, 0}
	p1 := Point{3, 0}
	p2 := Point{3, 4}
	path := Path{origin, p1, p2}
	exp := 7.0
	act := path.Distance()
	if exp != act {
		t.Errorf(fmt.Sprintf("expected: %f, actual: %f", exp, act))
	}
}

func TestTwoPointPath(t *testing.T)  {
	origin := Point{0, 0}
	p := Point{3, 4}
	path := Path{origin, p}
	exp := 5.0
	act := path.Distance()
	if exp != act {
		t.Errorf(fmt.Sprintf("expected: %f, actual: %f", exp, act))
	}
}

func TestCycle(t *testing.T) {
	origin := Point{0, 0}
	p1 := Point{3, 0}
	p2 := Point{3, 4}
	path := Path{origin, p1, p2, origin}
	exp := 12.0
	act := path.Distance()
	if exp != act {
		t.Errorf(fmt.Sprintf("expected: %f, actual: %f", exp, act))
	}
}
