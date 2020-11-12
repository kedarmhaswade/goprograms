// The problem 24 has remained tantalizing and as of 9/22/2018 I am not able to solve it on paper.
// Another attempt to cheat (write a program) has failed and now I am attempting it using a postfix evaluation scheme.
// Let's see if it succeeds.

package problem24

import (
	"fmt"
	"github.com/golang-collections/collections/stack"
	"github.com/pkg/errors"
	"math"
	"sort"
	"strconv"
)

const BLANK = "␢"

var operators = map[rune]bool{
	'+': true,
	'-': true,
	'*': true,
	'/': true,
}

// IsBinaryOperator determines if the given string represents a supported binary operator
func IsBinaryOperator(x string) bool {
	if len(x) < 0 {
		return false
	}
	_, ok := operators[rune(x[0])] // possible problem, but not a real concern
	return ok
}
type Multimap map[float64][][]string
func (mm Multimap) add(result float64, expr []string) {
	mm[result] = append(mm[result], expr)
}
func (mm Multimap) get(result float64) [][]string {
	return mm[result]
}
// Solves the problem for an expResult and slice of numbers (e.g. SolveIt(24, "1", "3", "4", "6")
func SolveIt(expResult float64, ints... string) {
	permutations := [][]string{{}}
	permutations = permute(ints, 0, permutations)
	operators := operatorTuples(3)
	mm := Multimap{}
	//fmt.Printf("%+v\n", permutations)
	for _, perm := range permutations {
		for _, operatorTuple := range operators {
			candidates := postfixCandidates(operatorTuple, perm)
			for _, candidate := range candidates {
				result, e := EvaluatePostfix(candidate)
				if e == nil && result == expResult {
					fmt.Printf("BINGO %s = %v\n", candidate, result)
					mm.add(result, candidate)
					//return
				} else if e == nil {
					mm.add(result, candidate)
				}
			}
		}
	}
	sortedKeys := make([]float64, 0)
	for key := range mm {
		sortedKeys = append(sortedKeys, key)
	}
	sort.Slice(sortedKeys, func(i, j int) bool {
		return sortedKeys[i] < sortedKeys[j]
	})
	for _, result := range sortedKeys {
		if math.Round(result) == result {
			//fmt.Printf("[%v]: {%v}\n", result, mm[result])
			fmt.Printf("[%v]: {%v}\n", result, len(mm[result]))
		}
		//if result == -24 {
		//	fmt.Printf("[%v]: {%v}\n", result, mm[result])
		//}
	}
}
func operatorTuples(n int) [][]string {
	if n == 1 {
		var tuples [][]string
		for key := range operators {
			s := string(key)
			tuple := make([]string, 1)
			tuple[0] = s
			tuples = append(tuples, tuple)
		}
		return tuples // [["+"], ["-"], ["*"], ["/"]]
	}
	prevTuples := operatorTuples(n-1)
	var currTuples [][]string
	for _, prevTuple := range prevTuples {
		for key := range operators {
			s := string(key)
			currTuple := make([]string, len(prevTuple))
			copy(currTuple, prevTuple)
			currTuple = append(currTuple, s)
			currTuples = append(currTuples, currTuple)
		}
	}
	return currTuples
}

//EvaluatePostfix evaluates the given expression assuming it is a postfix expression. It returns an error if the expression
// is invalid postfix in which case the result is not be relied upon. It returns the expected result and nil error if
// it is a valid postfix expression (e.g. ['+', '1', 5', '3', '-'] returns 3, nil) that can be evaluated. In case of
// evaluation errors (e.g. division by zero) it returns an error.
func EvaluatePostfix(expr []string) (float64, error) {
	stack := stack.New()
	for _, str := range expr {
		if IsBinaryOperator(str) {
			slen := stack.Len()
			if slen < 2 {
				s := fmt.Sprintf("expected: [2] operands for operator: %s, found:[%d]", str, slen)
				return 0, errors.New(s)
			}
			val := stack.Pop()
			b, conversionOk := val.(float64)
			if !conversionOk {
				return 0, errors.New(fmt.Sprintf("%v can not be converted into float64", val))
			}
			val = stack.Pop()
			a, conversionOk := val.(float64)
			if !conversionOk {
				return 0, errors.New(fmt.Sprintf("%v can not be converted into float64", val))
			}
			r, err := eval(a, b, str)
			if err != nil {
				return 0, err
			}
			stack.Push(r)
		} else {
			operand, err := strconv.ParseFloat(str, 64)
			if err != nil {
				return 0, err
			}
			stack.Push(operand)
		}
	}
	slen := stack.Len()
	if slen != 1 {
		return 0, errors.New(fmt.Sprintf("programming error, the length of stack expected [1], found: %d\n", slen))
	}
	val := stack.Pop()
	result, conversionOk := val.(float64)
	if !conversionOk {
		return 0, errors.New(fmt.Sprintf("%v can not be converted into float64", val))
	}
	return result, nil
}

// postfixCandidates places the given r binary operators and (n-r) operands in a slice in such a way that operators are first placed
// in r places of an n-slice and the remaining (n-r) operands are then placed in available places from left (zero-index)
// to right. There are n-choose-r such permutations which the function returns.
// An example return slice is ["1", "3", "+", "4", "*", "6", "-"] from operators = ["+, "*", "-"] and
// numbers = ["1", "3", "4", "6"] (in that order). This returned slice, as a postfix expression, would evaluate to "29".
// Note that this function will return invalid postfix expressions as well so that the calling function may keep
// a count of the valid/invalid postfix expressions arising from these operators and operands.
func postfixCandidates(operators []string, numbers []string) [][]string {
	r := len(operators)
	n := r + len(numbers)
	candidates := place(numbers, n)
	fillOperators(candidates, operators)
	return candidates
}
func fillOperators(candidates [][]string, operators []string) {
	lops := len(operators)
	for _, expr := range candidates {
		n := len(expr)
		for i, j := 0, 0; i < n && j < lops; {
			if expr[i] == BLANK {
				expr[i] = operators[j]
				i += 1
				j += 1
			} else {
				i += 1
			}
		}
	}
}

// place places r numbers in n places
func place(values []string, n int) [][]string {
	r := len(values)
	if r == 0 || n == 0 || r > n {
		return [][]string{}
	}
	if r == 1 { // the "main" base case
		theOnly := values[0] // the only number to place
		var acc [][]string
		for i := 0; i < n; i++ {
			a := makeNew(n) // with BLANKs for identification of a blank space
			a[i] = theOnly
			acc = append(acc, a)
		}
		return acc
	}
	// recursive case
	first := values[0]
	rest := values[1:]
	var currs [][]string
	for r := n - 1; r >= len(rest); r-- {
		prevs := place(rest, r)
		for _, prev := range prevs {
			var curr []string
			curr = append(curr, makeNew(n-1-r)...) //prefix
			curr = append(curr, first)
			curr = append(curr, prev...)
			currs = append(currs, curr)
		}
	}
	return currs
}

func makeNew(n int) []string {
	a := make([]string, n)
	for i := 0; i < n; i++ {
		a[i] = BLANK
	}
	return a
}
func ncr(n int, r int) int {
	// other conditions need to be considered TODO
	if r > n {
		return 0
	}
	if 2*r > n {
		r = n - r
	}
	c := 1
	for i, j := 1, n; i <= r; i, j = i+1, j-1 {
		c *= j
		c /= i
	}
	return c
}

// ====================== UnExported =============
func eval(a float64, b float64, r string) (float64, error) {
	switch r[0] {
	case '+':
		return a + b, nil
	case '-':
		return a - b, nil
	case '*':
		return a * b, nil
	case '/':
		if b == 0 {
			return 0, errors.New(fmt.Sprintf("division by 0, dividend: %v", a))
		}
		return a / b, nil
	}
	return 0, errors.New(fmt.Sprintf("invalid rune %c", r[0]))
}
func permute(numbers []string, index int, acc [][]string) [][]string {
	if index < len(numbers) {
		prevLen := len(acc)
		for _, prevP := range acc {
			for i := 0; i < len(prevP); i++ {
				currP := make([]string, 0)
				for j := 0; j < i; j++ {
					currP = append(currP, prevP[j])
				}
				currP = append(currP, numbers[index])
				for j := i; j < len(prevP); j++ {
					currP = append(currP, prevP[j])
				}
				acc = append(acc, currP)
			}
			acc = append(acc, append(prevP, numbers[index]))
		}
		return permute(numbers, index+1, acc[prevLen:])
	}
	return acc
}
