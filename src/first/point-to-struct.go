//// Demonstrates returning a struct or a pointer to a struct

// Demonstrates returning a struct or a pointer to a struct
package main

import (
	"fmt"
)

type Emp struct {
	Id     int
	Name   string
	Salary int
}

//var directory = make(map[int]Emp)
var directory []*Emp // it's critical to have this as a slice of pointers to Emp

func main() {
	initialize()
	var e = getEmployeeById(2) // type of e is *Emp
	fmt.Printf("type of return value: %T\n", e)
	fmt.Printf("outside: employee's address: %p\n", e)
	if e != nil {
		e.Salary = 0
		fmt.Printf("%s demoted, Salary: %d\n", e.Name, e.Salary)
	}
	printSalaries()
}
func printSalaries() {
	for i := 0; i < len(directory); i++ {
		emp := directory[i]
		fmt.Printf("%d\t%s\t%d\n", emp.Id, emp.Name, emp.Salary)
		//fmt.Printf("%+v\n", *emp)
	}
}
func getEmployeeById(id int) *Emp {
	for i := 0; i < len(directory); i++ {
		e := directory[i]
		if e.Id == id {
			fmt.Printf("inside: employee's address: %p\n", e)
			return e
		}
	}
	return nil
}

func initialize() {
	employees := []Emp{{1, "Larry", 1000}, {2, "Moe", 1000}, {3, "Curly", 1000}}
	for i := 0; i < len(employees); i++ {
		directory = append(directory, &employees[i])
	}
	printSalaries()
}

//package main
//
//import (
//	"fmt"
//)
//
//type Emp struct {
//	Id     int
//	Name   string
//	Salary int
//}
//
////var directory = make(map[int]Emp)
//var directory []Emp
//
//func main() {
//	initialize()
//	var e = getEmployeeById(2) // type of e is *Emp
//	fmt.Printf("type of return value: %T\n", e)
//	fmt.Printf("outside: employee's address: %p\n", e)
//	if e != nil {
//		e.Salary = 0
//		fmt.Printf("%s demoted, Salary: %d\n", e.Name, e.Salary)
//	}
//	printSalaries()
//}
//func printSalaries() {
//	for i := 0; i < len(directory); i++ {
//		emp := directory[i]
//		fmt.Printf("%d\t%s\t%d\n", emp.Id, emp.Name, emp.Salary)
//	}
//}
//func getEmployeeById(id int) *Emp {
//	for i := 0; i < len(directory); i++ {
//		e := directory[i]
//		if e.Id == id {
//			fmt.Printf("inside: employee's address: %p\n", &e)
//			return &e
//		}
//	}
//	return nil
//}
//
//func initialize() {
//	employees := []Emp{{1, "Larry", 1000}, {2, "Moe", 1000}, {3, "Curly", 1000}}
//	for _, e := range employees {
//		directory = append(directory, e)
//	}
//	printSalaries()
//}
