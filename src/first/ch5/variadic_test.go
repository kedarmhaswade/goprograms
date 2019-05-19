package ch5

import (
	"fmt"
	"strings"
	"testing"
)

func TestMax1(t *testing.T) {
	a := 22
	m := max(a)
	if m != a {
		t.Errorf("expected: %d, got: %d\n", a, m)
	}
	b := 33
	c := 0
	m = max(a, b, c)
	if m != b {
		t.Errorf("expected: %d, got: %d\n", b, m)
	}
}

func TestMin1(t *testing.T) {
	a := 42
	m := min(a)
	if m != a {
		t.Errorf("expected: %d, got: %d\n", a, m)
	}
	b := 33
	c := 0
	m = min(a, b, c)
	if m != c {
		t.Errorf("expected: %d, got: %d\n", c, m)
	}
}

func TestEmptyJoin(t *testing.T) {
	sep := ","
	var a []string
	expected := strings.Join(a, sep)
	actual := Join(sep, a...)
	if expected != actual {
		t.Errorf("expected: %v, actual: %v\n", expected, actual)
	}
}
func TestOneJoin(t *testing.T) {
	sep := ","
	a := []string{"Larry"}
	expected := strings.Join(a, sep)
	actual := Join(sep, a...)
	if expected != actual {
		t.Errorf("expected: %v, actual: %v\n", expected, actual)
	}
}
func TestMoreJoin(t *testing.T) {
	sep := ","
	a := []string{"Larry", "Moe", "Curly"}
	expected := strings.Join(a, sep)
	actual := Join(sep, a...)
	if expected != actual {
		t.Errorf("expected: %v, actual: %v\n", expected, actual)
	}
}

func TestMaxMax1MinMin1(t *testing.T) {
	a := []int{-2, -34, 15, 2, 0, 10}
	m := callIntIntEllipsisIntFunction(max, a[0], a[1:]...)
	m1 := callIntIntEllipsisIntFunction(max1, a[0], a[1:]...)
	fmt.Printf("m: %d, m1: %d\n", m, m1)
	if (m != m1) {
		t.Errorf("m: %d should have matched m1: %d", m, m1)
	}
	m = callIntIntEllipsisIntFunction(min, a[0], a[1:]...)
	m1 = callIntIntEllipsisIntFunction(min1, a[0], a[1:]...)
	fmt.Printf("m: %d, m1: %d\n", m, m1)
	if (m != m1) {
		t.Errorf("m: %d should have matched m1: %d", m, m1)
	}
}

func callIntIntEllipsisIntFunction(f func(int, ...int) int, first int, rest ...int) int {
	return f(first, rest...)
}
