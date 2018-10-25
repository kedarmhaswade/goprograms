package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

func raiseSalary(e *Employee, amt int) {
	e.Salary += amt // preferred -- dot notation applied to pointers verbatim, no need to deref the pointer
	// same as
	//(*e).Salary += amt
}

func main() {
	var dilbert, asok, boss Employee
	boss.ID = 1
	boss.Name = "Mr. pointy-haired-boss"
	dilbert.ID = 2
	dilbert.Name = "Mr. Dilbert"
	asok.ID = 3
	asok.Name = "Mr. Asok"
	m := make(map[int]Employee)
	m[0] = boss
	m[1] = dilbert
	m[2] = asok
	dilbert.Position = "Programmer"
	ptr := &dilbert.Position
	*ptr += ", Sr."
	fmt.Printf("position: %s\n", *ptr)
	raise := 1
	origSal := dilbert.Salary
	raiseSalary(&dilbert, raise)
	fmt.Printf("Salary raised by $%d ? %v\n", raise, dilbert.Salary == origSal+raise)

	e := lookupByID(m, 1)
	fmt.Printf("name: %s\n", e.Name)
}
func lookupByID(employees map[int]Employee, id int) Employee {
	return employees[id]
}
