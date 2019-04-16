// From gopl.io, Chapter 6:
// If all the methods of a named type T have a receiver type of T itself (not *T), it is safe to copy
// instances of that type; calling any of its methods necessarily makes a copy. For example, time.Duration
// values are liberally copied, including as arguments to functions. But if any method has a pointer receiver,
// you should avoid copying instances of T because doing so may violate internal invariants.
// For example, copying an instance of bytes.Buffer would cause the original and the copy to alias (§2.3.2)
// the same underlying array of bytes. Subsequent method calls would have unpredictable effects.
// Can we demonstrate that?
package ch6

import "errors"

// Let's try creating a type T and two methods, one that works on T and other on *T and see if
// an inconsistency can be produced rather easily.

// A naive circular buffer
type CircularBuffer struct {
	buf        []int
	cap, head, tail int
}

func (q *CircularBuffer) add(d int) error { // note: this method works on a pointer to a value, not the value of the type
	if q.tail-q.head > q.cap {
		return errors.New("buffer full")
	}
	q.buf[q.tail%q.cap] = d
	q.tail += 1
	return nil
}

func (q CircularBuffer) remove() (int, error) { // note: this method works on a value of the type, not a pointer
	if q.tail == q.head {
		return -1, errors.New("buffer empty")
	}
	x := q.buf[q.head%q.cap]
	q.head += 1
	return x, nil
}
