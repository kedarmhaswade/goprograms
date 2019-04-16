package ch4

import (
	"fmt"
	"testing"
)

func TestBasicRev(t *testing.T) {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	a := arr[:]
	Reverse(a) // changes a's elements through the alias
	fmt.Printf("%v\n", a)
	a = a[0:1]
	Reverse(a)
	fmt.Printf("%v\n", a)
}

func TestNonemptyInPlace(t *testing.T) {
	strings := []string{"non", "empty", "", "foo"}
	fmt.Printf("before: %v\n", strings)
	ne := nonemptyInPlace(strings)
	fmt.Printf("%v\n", ne)
	fmt.Printf("%v\n", strings)
	fmt.Printf("len: %d, cap: %d\n", len(ne), cap(ne)) // should be 3, 4
}

func TestStack(t *testing.T) {
	push("a")
	push("b")
	push("c")
	exp := "c"
	act, err := pop()
	if err != nil {
		t.Errorf("error: expected: %v, actul: %v", nil, err)
	}
	if exp != act {
		t.Errorf("expected: %s, actul: %s", exp, act)
	}
	exp = "b"
	act, err = pop()
	if err != nil {
		t.Errorf("error: expected: %v, actul: %v", nil, err)
	}
	if exp != act {
		t.Errorf("expected: %s, actul: %s", exp, act)
	}
	exp = "a"
	act, err = pop()
	if err != nil {
		t.Errorf("error: expected: %v, actul: %v", nil, err)
	}
	if exp != act {
		t.Errorf("expected: %s, actul: %s", exp, act)
	}
	act, err = pop()
	if err != EmptyStackError {
		t.Errorf("expected error, got %v", err)
	}
}

func TestReverseArray(t *testing.T) {
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	p := &a
	reverseArray(p)
	fmt.Printf("%v\n", *p)
}

func TestBasicLeftRotate(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6}
	LeftRotateOnePass(a, 2)
	fmt.Printf("%v\n", a)
	b := []int{1, 2, 3, 4, 5, 6}
	LeftRotate(b, 2)
	fmt.Printf("%v\n", b)
}

func TestUniq(t *testing.T) {
	a := []string{"a", "b", "b", "b", "a", "a", "c"}
	b := UniqAdj(a)
	fmt.Printf("%v\n", b)
}

func TestSquashSpace(t *testing.T) {
	s := "a		b"
	fmt.Printf("%s\n", SquashSpace([]byte(s)))
	s = " 			 A leading space, three tabs, and a space squashed to a space"
	fmt.Printf("%s\n", SquashSpace([]byte(s)))
	s = "अंकित त्यागी अौर केदार म्हसवडे कोडिंग् कर      	रहे 	हैं।"
	fmt.Printf("%s\n", SquashSpace([]byte(s)))
}

func TestAnyRuneRun(t *testing.T) {
	s := "aaaaaaab"
	fmt.Printf("%s\n", SquashAnyRuneRun([]byte(s)))
	s = "ककककदददददमममम बढाये जा a a abbb      "
	fmt.Printf("%s\n", SquashAnyRuneRun([]byte(s)))
}

func TestSquashSpecificRuneRuns(t *testing.T) {
	s := "squash aaaaa's but not bbbbbb's"
	fmt.Printf("%s\n", SquashRuneRun([]byte(s), 'a'))
}

func TestBasicReverseRunes(t *testing.T) {
	s := []byte("x = नयन") // reverse should be: "नयन = x"
	reverseRunesInPlace(s)
	fmt.Printf("%s\n", s)
	s = []byte("1:१,2:२,3:३,4:४")
	reverseRunesInPlace(s) // reverse should be: ४:4,३:3,२:2,१:1
	fmt.Printf("%s\n", s)
}



