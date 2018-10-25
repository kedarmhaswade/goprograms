package ch5

import (
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

