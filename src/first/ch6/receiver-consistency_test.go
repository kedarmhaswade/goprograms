package ch6

import (
	"fmt"
	"testing"
)

func TestAddRemove(t *testing.T) {
	q := CircularBuffer{buf: make([]int, 5), cap: 5} //[], 10, 0, 0
	exp := 1
	_ = q.add(exp)       // [10, 0, 0, 0, 0], 10, 0, 1
	act, _ := q.remove() // note: a copy of q is made and remove is called _on that copy_.
	if act != exp {
		// we got lucky if they are the same
		t.Errorf(fmt.Sprintf("removed element, expected: %d, actual: %d", exp, act))
	}
	exp = 2
	_ = q.add(exp)
	act, _ = q.remove() // note: a copy of q is made and remove is called _on that copy_.
	if act != exp {
		// and yes, this test fails because add works on a pointer, and remove on the value directly!
		// uncomment the following to see it in action
		 t.Errorf(fmt.Sprintf("removed element, expected: %d, actual: %d", exp, act))
	}
}
