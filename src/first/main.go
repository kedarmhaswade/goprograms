// temporary main for all programs
package main

import (
	"first/problem24"
	"fmt"
)

func main() {
	perms := make([][]int, 1)
	operands := []int{1, 3, 4, 6}
	perms = problem24.permute(operands, 0, perms)
	fmt.Printf("%+v\n", perms)
	result, err := problem24.EvaluatePostfix([]string{"1", "2.5", "+", "3", "-"})
	if err != nil {
		fmt.Printf("error: %v\n", err)
	} else {
		fmt.Printf("result: %v\n", result)
	}
}
