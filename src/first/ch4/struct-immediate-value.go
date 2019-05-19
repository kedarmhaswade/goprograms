package ch4

// Demonstrates the difference between struct as an immediate value and pointer to a struct

type Employee struct {
	ID        int
	LastName  string
	FirstName string
	Salary    int
}

var Larry = Employee{1, "Stooge", "Larry", 100}
var Moe = Employee{2, "Stooge", "Moe", 110}
var Curly = Employee{3, "Stooge", "Curly", 120}
var Me = Employee{4, "Stooge", "Radek", 0}
var None = Employee{}

var employees map[int]*Employee = make(map[int]*Employee)

func init() {
	employees[Larry.ID] = &Larry
	employees[Moe.ID] = &Moe
	employees[Curly.ID] = &Curly
	employees[Me.ID] = &Me
}

func EmployeeByID(id int) Employee {
	e, _ := employees[id] // assume that id exists
	return *e
}

func GiveRaise(id int, raise int) {
	a := EmployeeByID(id)
	a.Salary += raise

	// There is so much wrong with this func.
	// if you change it to:
	// EmployeeID(id).Salary += raise
	// you get -- yes, compilation error with this cryptic message: cannot assign to EmployeeByID(id).Salary
	// the reason is in the assignment statement: EmployeeID(id).Salary += raise, the left hand side
	// does NOT identify a variable!
	// A short snippet that demonstrates the same "issue" is in cannot-assign.go
	// Secondly, in its current form, this function does not do what you want. It creates a copy of given
	// employee struct and mutates that copy, leaving the original employee in the map alone.
	// NEVER pass/return structs, always pass/return pointers to structs

}
