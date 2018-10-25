package nick

import (
	"fmt"
	"testing"
)

func TestIsSubsequencePositive(t *testing.T) {
	sub := "hac"
	seq := "cathartic"
	ssrec := IsSubsequence(sub, seq)
	if !ssrec {
		t.Errorf(fmt.Sprintf("%s is a subsequence of %s, but the function returned false", sub, seq))
	}
	ssiter := IsSubsequenceIter(sub, seq)
	if ssiter != ssrec {
		t.Errorf(fmt.Sprintf("iterative: %v, recursive: %v", ssiter, ssrec))
	}
	sub = "abc"
	seq = "aabc"
	if !ssrec {
		t.Errorf(fmt.Sprintf("%s is a subsequence of %s, but the function returned false", sub, seq))
	}
	ssiter = IsSubsequenceIter(sub, seq)
	if ssiter != ssrec {
		t.Errorf(fmt.Sprintf("iterative: %v, recursive: %v", ssiter, ssrec))
	}
}
func TestIsSubsequenceNegative(t *testing.T) {
	sub := "bat"
	seq := "table"
	ssrec := IsSubsequence(sub, seq)
	if ssrec {
		t.Errorf(fmt.Sprintf("%s is not a subsequence of %s, but the function returned true", sub, seq))
	}
	ssiter := IsSubsequenceIter(sub, seq)
	if ssiter != ssrec {
		t.Errorf(fmt.Sprintf("iterative: %v, recursive: %v", ssiter, ssrec))
	}
	sub = "abcd"
	seq = "aabcefa"
	if ssrec {
		t.Errorf(fmt.Sprintf("%s is not a subsequence of %s, but the function returned true", sub, seq))
	}
	ssiter = IsSubsequenceIter(sub, seq)
	if ssiter != ssrec {
		t.Errorf(fmt.Sprintf("iterative: %v, recursive: %v", ssiter, ssrec))
	}
}
