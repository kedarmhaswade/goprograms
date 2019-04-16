package ch4

import (
	"fmt"
	"testing"
)

func TestMarshal(t *testing.T) {
	data, e := Marshal()
	if e != nil {
		t.Errorf("error: expected: nil, actual: %v\n", e)
	}
	fmt.Printf("%s\n", data)
}
