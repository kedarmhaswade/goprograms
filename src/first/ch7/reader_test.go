package ch7

import (
	"fmt"
	"io"
	"testing"
)

func TestStringReader1(t *testing.T) {
	size := 10
	src := "strlen: 10"
	reader := NewReader(src)
	halfSize := size / 2
	p := make([]byte, halfSize)
	exp := halfSize
	act, err := reader.Read(p)
	if err != nil {
		t.Errorf(fmt.Sprintf("error -- expected: nil, found: %v", err))
	}
	if exp != act {
		t.Errorf(fmt.Sprintf("bytes -- expected: %v, found: %v", exp, act))
	}
	// do it again, this time we should receive no error this time as well
	act, err = reader.Read(p)
	if err != nil {
		t.Errorf(fmt.Sprintf("error -- expected: nil, found: %v", err))
	}
	if exp != act {
		t.Errorf(fmt.Sprintf("bytes -- expected: %v, found: %v", exp, act))
	}
	// now we should receive an EOF!
	expErr := io.EOF
	act, err = reader.Read(p)
	if err != expErr {
		t.Errorf(fmt.Sprintf("error -- expected: %v, found: %v", expErr, err))
	}
	if exp != act {
		t.Errorf(fmt.Sprintf("bytes -- expected: %v, found: %v", exp, act))
	}
}
