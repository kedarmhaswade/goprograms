package ch4

import (
	"fmt"
	"testing"
)

func TestStructBasic(t *testing.T) {
	fmt.Printf("Before: %v ", EmployeeByID(4).Salary)
	GiveRaise(4, 100)
	fmt.Printf("After : %v\n", EmployeeByID(4).Salary)
	fmt.Printf("Look ma, no raise!\n")
}
