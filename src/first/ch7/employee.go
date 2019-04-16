// Study the sort interface
package ch7

import "sort"

type Employee struct {
	ID   int
	Name string
	Age  int
}
type ByID []*Employee

func (employees ByID) Len() int {
	return len(employees)
}
func (employees ByID) Less(i, j int) bool {
	return employees[i].ID < employees[j].ID
}
func (employees ByID) Swap(i, j int) {
	employees[i], employees[j] = employees[j], employees[i]
}

// unexported interface type
type rev struct {
	s sort.Interface
}

func (r rev) Len() int {
	return r.s.Len()
}
func (r rev) Less(i, j int) bool {
	return r.s.Less(j, i)
}
func (r rev) Swap(i, j int) {
	r.s.Swap(i, j)
}

// Similar to sort.Reverse, provides a Reverse function that transforms any sort order s to its reverse.
func Reverse(s sort.Interface) sort.Interface {
	return rev{s}
}

// Another type that satisfies sort.Interface; sorts the given slice by their names
type ByName []*Employee

func (employees ByName) Len() int {
	return len(employees)
}
func (employees ByName) Less(i, j int) bool {
	return employees[i].Name < employees[j].Name
}
func (employees ByName) Swap(i, j int) {
	employees[i], employees[j] = employees[j], employees[i]
}

// Here's my implementation of sort.IsSorted function that determines if a given sequence is sorted
func isSorted(p sort.Interface) bool {
	n := p.Len()
	for i := 1; i < n; i++ {
		if !p.Less(i-1, i) {
			return false
		}
	}
	return true
}
