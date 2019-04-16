package ch6

import (
	"fmt"
	"testing"
)

func TestIntegerToString(t *testing.T) {
	var n  = Integer(123)
	exp := "123"
	act := n.ToString()
	if act != exp {
		t.Errorf("expected: %s, actual: %s", exp, act)
	}
	n = Integer(0x123)
	exp = "291"
	act = n.ToString()
	if act != exp {
		t.Errorf("expected: %s, actual: %s", exp, act)
	}
	fmt.Printf("method type: %T\n", Integer.ToString)
}

func TestMultiplySelf(t *testing.T) {
	var act = Integer(12)
	act.Multiply(3) // same as (&act).Multiply(3)
	var exp Integer = 36
	if act != exp {
		t.Errorf("expected: %d, actual: %d", exp, act)
	}
	fmt.Printf("method type: %T\n", (*Integer).Multiply)
}

func TestString_Length(t *testing.T) {
	var empty String
	var exp = 0
	// Note: method Length requires a pointer to String (*String) receiver. The type of empty is only a String.
	// By making a String a receiver, we are asking the compiler to do the implicit conversion from String to *String
	var act = empty.Length()
	if exp != act {
		t.Errorf("expected: %d, actual: %d", exp, act)
	}
	var nonEmpty String = "Hello"
	exp = len(string(nonEmpty))
	act = nonEmpty.Length() // String and not *String is the receiver
	if exp != act {
		t.Errorf("expected: %d, actual: %d", exp, act)
	}
}

func TestValues(t *testing.T) {
	m := Values{"lang": {"en"}}
	m = nil
	fmt.Println(m.Get("item")) // ""
}
