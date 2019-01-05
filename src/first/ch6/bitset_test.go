package ch6

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestSimpleHasAddHas(t *testing.T) {
	testOne(t, &IntSet{}, 65)
}

func testOne(t *testing.T, bv *IntSet, n int) {
	actual := bv.Has(n)
	expected := false
	if actual {
		t.Errorf(fmt.Sprintf("membership for %d, expected: %v, actual: %v", n, expected, actual))
	}
	added := bv.Add(n)
	if !added {
		t.Errorf(fmt.Sprintf("Expected to add %d to the set, but adding the element failed", n))
	}
	actual = bv.Has(n)
	expected = true
	if !actual {
		t.Errorf(fmt.Sprintf("membership for %d, expected: %v, actual: %v", n, expected, actual))
	}
}

func TestManyHasAddHas(t *testing.T) {
	rand.Seed(1234)
	for i := 1; i <= 100; i++ {
		testOne(t, &IntSet{}, rand.Intn(10000))
	}
}

func TestAddAll(t *testing.T) {
	a := []int{1, 4, 6, 253, 533, 8899, 1343590}
	var s = &IntSet{}
	s.AddAll(a...)
	for _, n := range a {
		if !s.Has(n) {
			t.Errorf(fmt.Sprintf("Expected %d to be in the set, but not found", n))
		}
	}
}

func contains(a []int, x int) bool {
	for _, n := range a {
		if n == x {
			return true
		}
	}
	return false
}
func TestUnion(t *testing.T) {
	a := []int{1, 10, 64, 129, 200, 1000}
	b := []int{2, 10, 65, 129, 230, 999}
	var aSet = &IntSet{}
	var bSet = &IntSet{}
	aSet.AddAll(a...)
	bSet.AddAll(b...)
	union := aSet.UnionWith(bSet)
	for n := 0; n <= 2000; n++ {
		if contains(a, n) || contains(b, n) { // inefficient
			if !union.Has(n) {
				t.Errorf(fmt.Sprintf("Expected %d in the union, but did not find it", n))
			}
		} else {
			if union.Has(n) {
				t.Errorf(fmt.Sprintf("Did not expect %d in the union, but found it", n))
			}
		}
	}
}

func TestLen(t *testing.T) {
	var s = &IntSet{}
	actual := s.Size()
	expected := 0
	if actual != expected {
		t.Errorf(fmt.Sprintf("Empty set, expected size: %d, actual size: %d", expected, actual))
	}
	// add 10 elements
	expected = 1000
	for i := 0; i < expected; i++ {
		s.Add(i)
	}
	actual = s.Size()
	if actual != expected {
		t.Errorf(fmt.Sprintf("Empty set, expected size: %d, actual size: %d", expected, actual))
	}

}

func TestRemove(t *testing.T) {
	var set = &IntSet{}
	lim := 100
	for i := 0; i < lim; i++ {
		set.Add(i)
	}
	for i := 0; i < lim; i++ {
		if !set.Has(i) {
			t.Errorf(fmt.Sprintf("Expected %d to be present in the set, but found it absent!", i))
		}
		removed := set.Remove(i)
		if !removed {
			t.Errorf(fmt.Sprintf("Expected Remove method to return true, but it returned false!"))
		}
		if set.Has(i) {
			t.Errorf(fmt.Sprintf("Expected %d to be absent in the set, but found it there!", i))
		}
	}
}

func TestClearRelease(t *testing.T) {
	var set = &IntSet{}
	lim := 10000
	for i := 0; i < lim; i++ {
		set.Add(i)
	}
	set.Clear(true) // release the memory
	n := set.Size()
	if n != 0 {
		t.Errorf(fmt.Sprintf("Expected %d, but got %d in a set that was cleared", 0, n))
	}
}

func TestClearRetain(t *testing.T) {
	var set = &IntSet{}
	lim := 10000
	for i := 0; i < lim; i++ {
		set.Add(i)
	}
	set.Clear(false) // retain the memory
	n := set.Size()
	if n != 0 {
		t.Errorf(fmt.Sprintf("Expected %d, but got %d in a set that was cleared", 0, n))
	}
}

func TestEquals(tt *testing.T) {
	s := &IntSet{}
	t := &IntSet{}
	if !s.Equals(t) {
		tt.Errorf(fmt.Sprintf("Expected the empty sets to be equal, but Equals returns false!"))
	}
	s.Add(10)
	t.Add(10)
	if !s.Equals(t) {
		tt.Errorf(fmt.Sprintf("Expected the sets to be equal, but Equals returns false!"))
	}
	s.Remove(10)
	if s.Equals(t) {
		tt.Errorf(fmt.Sprintf("Expected the sets to be unequal, but Equals returns true!"))
	}
	t.Remove(10)
	if !s.Equals(t) {
		tt.Errorf(fmt.Sprintf("Expected the sets to be equal, but Equals returns false!"))
	}
	s = &IntSet{}
	t = &IntSet{}
	for _, i := range []int{1, 3, 53, 232} {
		s.Add(i)
		t.Add(i)
	}
	if !s.Equals(t) {
		tt.Errorf(fmt.Sprintf("Expected the sets to be equal, but Equals returns false!"))
	}
	t.Add(0)
	if s.Equals(t) {
		tt.Errorf(fmt.Sprintf("Expected the sets to be unequal, but Equals returns true!"))
	}
	s = &IntSet{}
	for i := 0; i < 10000; i++ {
		s.Add(i)
	}
	t = s
	if !s.Equals(t) {
		tt.Errorf(fmt.Sprintf("Wow! a set does not equal itself!"))
	}
}
func TestCopy(tt *testing.T) {
	s := &IntSet{}
	for _, i := range []int{1, 3, 53, 232} {
		s.Add(i)
	}
	t := s.Copy()
	if !t.Equals(s) {
		tt.Errorf(fmt.Sprintf("Expected the sets to be equal, but Equals returns false!"))
	}
	s = &IntSet{}
	t = s.Copy()
	if !t.Equals(s) {
		tt.Errorf(fmt.Sprintf("Expected the sets to be equal, but Equals returns false!"))
	}
	ss := s.UnionWith(t)
	if !t.Equals(ss) {
		tt.Errorf(fmt.Sprintf("Expected the sets to be equal, but Equals returns false!"))
	}
	if !s.Equals(ss) {
		tt.Errorf(fmt.Sprintf("Expected the sets to be equal, but Equals returns false!"))
	}
}

func TestIntersectEmpty(tt *testing.T) {
	s := &IntSet{}
	s.Add(1)
	t := &IntSet{}
	i1 := s.IntersectWith(t)
	i2 := t.IntersectWith(s)
	if !i1.Equals(i2) {
		tt.Errorf("Expected intersection of two sets to work both ways, but that was not the case")
	}
}

func TestIntersectOddEven(tt *testing.T) {
	w1 := []int{1, 3, 5, 7, 9, 11, 101}
	w2 := []int{0, 2, 4, 6, 8, 10, 100}
	s := &IntSet{}
	s.AddAll(w1...)
	t := &IntSet{}
	t.AddAll(w2...)
	i1 := s.IntersectWith(t)
	i2 := t.IntersectWith(s)
	if !i1.Equals(i2) {
		tt.Errorf("Expected intersection of two sets to work both ways, but that was not the case")
	}
	coi := i1.Size()
	if coi != 0 {
		tt.Errorf(fmt.Sprintf("Cardinality of intersection, expected: %d, actual: %d", 0, coi))
	}
}

func TestIntersectComplete(tt *testing.T)  {
	w1 := []int{1, 3, 5, 7, 9, 11, 101, 1001, 100001}
	w2 := []int{1, 100001}
	s := &IntSet{}
	s.AddAll(w1...)
	t := &IntSet{}
	t.AddAll(w2...)
	i1 := s.IntersectWith(t)
	i2 := t.IntersectWith(s)
	if !i1.Equals(i2) {
		tt.Errorf("Expected intersection of two sets to work both ways, but that was not the case")
	}
	coi := i1.Size()
	if coi != t.Size() {
		tt.Errorf(fmt.Sprintf("Cardinality of intersection, expected: %d, actual: %d", t.Size(), coi))
	}
}

func TestDifferenceWith(tt *testing.T) {
	w1 := []int{1, 3, 4, 6, 10, 101, 1001, 10001, 100001, 1000001}
	w2 := []int{1, 4, 10, 1001, 100001}
	s := &IntSet{}
	s.AddAll(w1...)
	t := &IntSet{}
	t.AddAll(w2...)
	expected1 := (&IntSet{}).AddAll(3, 6, 101, 10001, 1000001)
	actual1 := s.DifferenceWith(t)
	if !expected1.Equals(actual1) {
		tt.Errorf(fmt.Sprintf("expected1 %v, actual %v", expected1, actual1))
	}
	expected2 := &IntSet{}
	actual2 := t.DifferenceWith(s)
	if !expected2.Equals(actual2) {
		tt.Errorf(fmt.Sprintf("expected1 %v, actual %v", expected2, actual2))
	}
}


