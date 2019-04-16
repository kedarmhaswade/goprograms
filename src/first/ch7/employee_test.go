package ch7

import (
	"fmt"
	"sort"
	"testing"
)

func TestSortStooges(t *testing.T) {
	expFirstByID := &Employee{2, "Curly", 19}
	employees := []*Employee{{3, "Larry", 20}, {10, "Moe", 21}, expFirstByID, {4, "Sally", 29}}
	sort.Sort(ByID(employees))
	if expFirstByID != employees[0] {
		t.Errorf(fmt.Sprintf("error sorting, expected first: %v, actual: %v", expFirstByID, employees[0]))
	}
	if expFirstByID.ID != employees[0].ID {
		t.Errorf(fmt.Sprintf("error sorting, expected id of the first: %v, actual: %v", expFirstByID.ID, employees[0].ID))
	}
	sort.Sort(Reverse(ByID(employees)))
	if expFirstByID != employees[3] {
		t.Errorf(fmt.Sprintf("error sorting, expected last: %v, actual: %v", expFirstByID, employees[3]))
	}
	if expFirstByID.ID != employees[3].ID {
		t.Errorf(fmt.Sprintf("error sorting, expected id of the last: %v, actual: %v", expFirstByID.ID, employees[3].ID))
	}
	expFirstByName := expFirstByID // Curly is first by ID (2) and his name appears first in lexicographic order
	sort.Sort(ByName(employees))
	if expFirstByName != employees[0] {
		t.Errorf(fmt.Sprintf("error sorting, expected first: %v, actual: %v", expFirstByName, employees[0]))
	}
	// reuse the reverse function
	sort.Sort(Reverse(ByName(employees)))
	if expFirstByName != employees[3] {
		t.Errorf(fmt.Sprintf("error sorting, expected last: %v, actual: %v", expFirstByName, employees[3]))
	}
}

