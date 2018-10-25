package problem24

import (
	"fmt"
	"testing"
)

func TestSimple(t *testing.T) {
	actual, err := EvaluatePostfix([]string{"1", "2.5", "+", "3", "-"})
	if err != nil {
		t.Errorf("nil error expected, got %v", err)
	}
	expected := 0.5
	fmt.Printf("actual: %v, expected: %v\n", actual, expected)
	if actual != expected {
		t.Errorf("expected: [%v], actual: [%v]", expected, actual)
	}
}

func TestDivByZero(t *testing.T) {
	_, err := EvaluatePostfix([]string{"1", "0", "/"})
	if err == nil {
		t.Errorf("non-nil error expected, got nil")
	}
}
func TestInvalidPostfix(t *testing.T) {
	_, err := EvaluatePostfix([]string{"0", "1"})
	if err == nil {
		t.Errorf("non-nil error expected, got nil")
	}
}

func Test7c3(t *testing.T) {
	c := ncr(7, 3)
	if c != 35 {
		t.Errorf("7 choose 3 must be 35, not %d", c)
	}
}

func TestPlace2In5(t *testing.T) {
	a := []string{"1", "3", "4", "6"}
	n := 7
	r := len(a)
	ways := place(a, n) // place the numbers a in n places
	fmt.Printf("num ways: %+v\n", ways)
	act := len(ways)
	exp := ncr(n, r)
	if act != exp {
		t.Errorf(fmt.Sprintf("number of ways of placing %d in %d places, expected: %d, actual: %d", r, n, exp, act))
	}
}
func TestPostfixCandidates(t *testing.T) {
	numbers := []string{"6", "4", "3", "1"}
	operators := []string{"+", "-", "+"}
	candidates := postfixCandidates(operators, numbers)
	fmt.Printf("%+v\n", candidates)
	r := len(numbers)
	n := r + len(operators)
	act := len(candidates)
	exp := ncr(n, r)
	if act != exp {
		t.Errorf(fmt.Sprintf("expected: %d, actual: %d\n", exp, act))
	}
	for _, candidate := range candidates {
		v, e := EvaluatePostfix(candidate)
		if e != nil {
			fmt.Printf("invalid postfix: %s\n", candidate)
		} else {
			fmt.Printf("%s = %v\n", candidate, v)
			if v == 24 {
				fmt.Printf("BINGO")
				break
			}
		}
	}
}

func TestOperatorTuples(t *testing.T) {
	tuples := operatorTuples(3)
	fmt.Printf("%+v\n", tuples)
}

func TestSolveIt(t *testing.T) {
	SolveIt()
}