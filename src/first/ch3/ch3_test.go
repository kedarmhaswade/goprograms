package ch3

import "testing"

func Test_convert(t *testing.T) {
	convert()
}

func TestBitDiff1(t *testing.T) {
	a := []byte("a")
	b := []byte("A")
	actual, e := BitDiff(a, b)
	if e != nil {
		t.Errorf("expected: nil, found: %v", e)
	}
	if actual != 1 {
		t.Errorf("expected: 1, found: %d", actual) // 'a' and 'A' differ in exactly one bit
	}
	actual, e = BitDiff(a, a)
	if e != nil {
		t.Errorf("expected: nil, found: %v", e)
	}
	if actual != 0 {
		t.Errorf("expected: 0, found: %d", actual) // 'a' and 'a' do not differ
	}
}

func TestBitDiff2(t *testing.T) {
	a := []byte("aa")
	b := []byte("AA")
	actual, e := BitDiff(a, b)
	if e != nil {
		t.Errorf("expected: nil, found: %v", e)
	}
	if actual != 2 {
		t.Errorf("expected: 2, found: %d", actual) // 'aa' and 'AA' differ in two bits
	}

}

func TestShaFun(t *testing.T) {
	shaFun()
}

func TestStudyArrayPointers(t *testing.T) {
	studyPointersToArrays()
}

