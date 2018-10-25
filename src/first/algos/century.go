// You see the following decimal digits on a whiteboard:
// 1 2 3 4 5 6 7 8 9 0
// and are asked to place either a '+', a '-', or nothing to get an expression which evaluates to 100. When you
// place none of '+' and '-', you simply concatenate the digits to form a decimal number, e.g. 45 is forty five.
// One possible solution is:
// 123+45-67+8-9+0
// Can you find all such expressions?
package algos

import (
	"fmt"
	"strconv"
)

func printCenturies() {
	exprs := expressions([]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"})
	fmt.Printf("total number of expressions: %d\n", len(exprs))
	ways := 0
	for _, expr := range exprs {
		result := evaluate(expr)
		if result == 100 {
			ways += 1
			fmt.Printf("(%+v) = %d\n", expr, result)
		}
	}
	fmt.Printf("total number of expressions to sum to a 100: %d\n", ways)
}
// expressions returns all the expressions created in the given way from given slice of digits
func expressions(expr []string) []string {
	n := len(expr)
	if n < 2 {
		return []string{}
	} else if n == 2 {
		plus := expr[0] + "+" + expr[1]
		minus := expr[0] + "-" + expr[1]
		join := expr[0] + expr[1]
		return []string{plus, minus, join}
	} else {
		partials := expressions(expr[:n-1]) // recursive call
		last := expr[n-1]
		fulls := make([]string, 0)
		for _, p := range partials {
			plus := p + "+" + last
			fulls = append(fulls, plus)
			minus := p + "-" + last
			fulls = append(fulls, minus)
			join := p + last
			fulls = append(fulls, join)
		}
		return fulls
	}
}

func evaluate(expression string) int {
	tokens := getTokens(expression)
	n := len(tokens)
	if n == 0 {
		return 0
	}
	result, _ := strconv.Atoi(tokens[0])
	for i := 1; i < n; i++ {
		token := tokens[i]
		if token != "+" && token != "-" {
			operand, _ := strconv.Atoi(token)
			prev := tokens[i-1]
			if prev == "+" {
				result += operand
			} else if prev == "-" {
				result -= operand
			} else {
				panic("invalid token: " + token)
			}
		}
	}
	return result
}
func getTokens(s string) []string {
	tokens := make([]string, 0)
	n := len(s)
	for i := 0; i < n;  {
		b := s[i]
		if b == '+' || b == '-' {
			tokens = append(tokens, string(b))
			i += 1
		} else {
			j := i
			for ;j < n && s[j] != '+' && s[j] != '-'; {
				j += 1
			}
			tokens = append(tokens, s[i:j])
			i = j
		}
	}
	return tokens
}
