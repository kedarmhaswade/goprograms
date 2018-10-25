// Sunil Singh asks this problem:
// Take the four numbers: 1, 3, 4, 6
// and four operations (binary operands): +, -, ×, ÷
// and using the given binary operands, normal evaluation rules, and all the digits exactly once,
// find an expression that evaluates to every number from 0 to 24.
// By hand, I created expressions for everything from 0 to 23. But 24 is illusive as Sunil said.
// This program creates all the permutations of these four numbers and applies the binary operands
// to every two of them recursively. Let's see what we get.
package main

import (
	"fmt"
	"strings"
)

// permute returns a list of permutations of the given numbers
//func permute(numbers []int) [][]int {
//	if len(numbers) == 0 {
//		return [][]int{}
//	}
//	if len(numbers) == 1 {
//		return [][]int{numbers}
//	}
//	curr := make([][]int, 0)     // current permutations each of which is one greater than each of prev
//	for i := 1; i < len(numbers); i++ {
//		prev := permute(numbers[:i]) // previous permutations
//		for j := 0; j < len(prev); j++ { // do for each of the permutations so far
//			prevP := prev[j]
//			for k := 0; k <= len(prevP); k++ { // increment the length of permutation
//				currP := make([]int, len(prevP)+1)
//				//subslice [0:k)
//				currP = append(currP, prevP[0:k]...)
//				// the new number
//				currP = append(currP, numbers[i])
//				// subslice[k:len(prevP)]
//				currP = append(currP, prevP[k:]...)
//				curr = append(curr, currP)
//			}
//		}
//	}
//	return curr
//}

func permute(numbers []float64, index int, acc [][]float64) [][]float64 {
	if index < len(numbers) {
		prevLen := len(acc)
		for _, prevP := range acc {
			for i := 0; i < len(prevP); i++ {
				currP := make([]float64, 0)
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
func main() {
	bOps := []rune{'+', '-', '×', '÷'}
	acc := make([][]float64, 1)
	//var acc [][]int
	// assumption: none of the operands is zero
	operands := []float64{1, 3, 4, 6}
	acc = permute(operands, 0, acc)
	for i, operands := range acc {
		fmt.Printf("%d: %v\n", i, operands)
		var results []float64
		var expr strings.Builder
		eval(operands, bOps, 0, results, expr)
	}
}
func eval(operands []float64, bOps []rune, time int, result []float64, expr strings.Builder) {
	if len(operands) <= 0 {
		for i := 0; i < len(result); i++ {
			fmt.Printf("result %v: %v = %v\n", i, expr.String(), result[i])
			if result[i] == 24 {
				fmt.Printf("hurrah! %v\n", result[i])
				break
			}
		}
		return
	}
	if time == 0 { // first time around
		r2 := eval2(operands[0], operands[1], bOps, expr)
		eval(operands[2:], bOps, time+1, r2, expr)
	} else {
		n := len(result)
		for i := 0; i < n; i++ {
			r2 := eval2(result[i], operands[0], bOps, expr)
			result = append(result, r2...)
		}
		eval(operands[1:], bOps, time+1, result[n:], expr)
	}
}

func eval2(a float64, b float64, bOps []rune, expr strings.Builder) []float64 {
	var results []float64
	for _, op := range bOps {
		switch op {
		case '+':
			results = append(results, float64(a)+float64(b))
			expr.WriteString(fmt.Sprintf("(%v + %v)", a, b))
		case '-':
			results = append(results, float64(a)-float64(b))
			expr.WriteString(fmt.Sprintf("(%v - %v)", a, b))
		case '×':
			results = append(results, float64(a)*float64(b))
			expr.WriteString(fmt.Sprintf("(%v × %v)", a, b))
		case '÷':
			if b == 0 {
				fmt.Printf("refusing to divide %d by %d! \n", a, b)
				break
			}
			results = append(results, float64(a)/float64(b))
			expr.WriteString(fmt.Sprintf("(%v ÷ %v)", a, b))
		default:
			fmt.Printf("Got something I don't know of %v, results may be wrong\n", op)
		}
	}
	return results
}
