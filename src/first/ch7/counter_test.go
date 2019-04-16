package ch7

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
)

func TestSimpleWordCounter(t *testing.T) {
	s := " hey! this has eight  \n words and two lines "
	var wc WordCounter
	act, _ := wc.Write([]byte(s)) // a pointer to wc is actually passed
	exp := 8
	if exp != act {
		t.Errorf(fmt.Sprintf("expected: %d, actual: %d words", exp, act))
	}
	if exp != int(wc) {
		t.Errorf(fmt.Sprintf("expected: %d, actual: %d words", exp, int(wc)))
	}
}

func TestSimpleLineCounter(t *testing.T) {
	text := `
one
two
three`
	var lc LineCounter
	exp := 3
	act, err := lc.Write([]byte(text)) // note: actually a pointer to a LineCounter value is passed
	if exp != act {
		t.Errorf(fmt.Sprintf("number of lines -- expected: %d, actual: %d", exp, act))
	}
	if err != nil {
		t.Errorf(fmt.Sprintf("error -- expected: %v, actual: %v ", nil, err))
	}
	if exp != int(lc) { // type conversion either way is necessary
		if LineCounter(exp) != lc { // why not make it a part of the test itself?
			t.Errorf(fmt.Sprintf("number of lines -- expected: %v, actual: %v ", exp, lc))
		}
	}
}

func TestCountingWriter1(t *testing.T) {
	// wrap os.Stdout, the writer to stdout
	ww, cp := CountingWriter(os.Stdout) // a writer and a *int64
	s := "four"                         // the only cardinal number whose name is as long as the number itself :-)
	ww.Write([]byte(s))                 // return value ignored! (writes to stdout)
	exp := len(s)
	act := *cp
	if int64(exp) != act {
		t.Errorf(fmt.Sprintf("expected: %v, actual: %v", exp, act))
	}
	// once more
	ww.Write([]byte(s))
	exp = 2 * len(s)
	act = *cp
	if int64(exp) != act {
		t.Errorf(fmt.Sprintf("expected: %v, actual: %v", exp, act))
	}
}

func TestCountingWriterEnglishLetters(t *testing.T) {
	// wrap os.Stdout, the writer to stdout
	ww, cp := CountingWriter(os.Stdout) // a writer and a *int64
	exp := 0
	for i := 0; i < 20; i++ {
		exp += i
		ww.Write(randomBytes(i))
	}
	act := *cp
	if int64(exp) != act {
		t.Errorf(fmt.Sprintf("expected: %v, actual: %v", exp, act))
	}
}
func randomBytes(n int) []byte {
	s := 'a'
	var buf []byte
	for i := 0; i < n; i++ {

		b := byte(int(s) + rand.Intn(26))
		buf = append(buf, b)
	}
	return buf
}
