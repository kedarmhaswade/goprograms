package ch7

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestLimit1(t *testing.T) {
	var expLim int64 = 4
	var expPrefix = "four"
	reader := LimitReader(bytes.NewBufferString(expPrefix+ " and more, but clip at 4"), expLim)
	buf := make([]byte, expLim) // exactly four bytes long
	act, err := reader.Read(buf)
	// we should be good here!
	if err != nil {
		t.Errorf(fmt.Sprintf("error -- expected: %v, actual: %v", nil, err))
	}
	if expLim != int64(act) {
		t.Errorf(fmt.Sprintf("bytes read -- expected: %v, actual: %v", expLim, act))
	}
	actPrefix := string(buf)
	if expPrefix != actPrefix {
		t.Errorf(fmt.Sprintf("prefix read -- expected: %v, actual: %v", expPrefix, actPrefix))
	}
	// read again, and we should have an EOF
	act, err = reader.Read(buf) // overwrite buf
	if err != io.EOF {
		t.Errorf(fmt.Sprintf("expected EOF, but got: %v", err))
	}
	fmt.Printf("%v\n", string(act))
}

