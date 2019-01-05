// A few tests for chapter 2
package ch2

import (
	"math"
	"testing"
)

func TestPopCount1(t *testing.T) {
	var m1 int64 = -1
	act := PopCount(m1)
	exp := 64
	if act != exp {
		t.Errorf("expected: %d, actual: %d", exp, act)
	}
	var m2 uint64 = math.MaxUint64 // all 64 1's
	signed := int64(m2) // same as -1
	act = PopCount(signed)
	exp = 64
	if act != exp {
		t.Errorf("expected: %d, actual: %d", exp, act)
	}
}

func TestF(t *testing.T) {
	f()
}

func TestLoopVar(t *testing.T) {
	loopVars()
}

