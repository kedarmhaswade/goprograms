// Consider the problem of computing a sequence of computer science courses that satisfies the
// prerequisite requirements of each one. The prerequisites are given in the prereqs table below,
// which is a mapping from each course to the list of courses that must be completed before it.

// This kind of problem is known as topological sorting.
// Conceptually, the prerequisite information forms a directed graph with a node for
// each course and edges from each course to the courses that it depends on.
// The graph is acyclic: there is no path from a course that leads back to itself.
// We can compute a valid sequence using depth-first search through the graph with the code below.

// In the first iteration, this program did not detect cycles. In the second, it was enhanced to do so.
package main

import (
	"fmt"
	"sort"
)

// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}
var prereqsAa = map[string][]string{
	"algorithms": {"algorithms"}, // hee hee
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

var prereqsAba = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"linear algebra":        {"calculus"}, // a cycle of the form: a->b->a
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	traverseCourses(prereqs)
	//traverseCourses(prereqsAa)
	//traverseCourses(prereqsAba)
}

var seen = make(map[string]bool, 0) // set of visited vertices
var seq = make([]string, 0)         // sequence of visited vertices

func traverseCourses(courses map[string][]string) {
	fmt.Printf("\t\t\tCourse Sequence\n")
	for i, course := range topoSort(courses) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}
func topoSort(graph map[string][]string) []string {
	// code between start and end is not necessary for correctness, but ensures a deterministic order
	// start
	sorted := make([]string, 0)
	for src := range graph {
		sorted = append(sorted, src)
	}
	sort.Strings(sorted)
	// end
	//for src := range graph {
	for _, src := range sorted {
		if !visited(src) {
			if visitDfsWithCycle(graph, src, make(map[string]bool, 0)) {
				break
			}
			seen[src] = true
			seq = append(seq, src)
		}
	}
	return seq
}
func visitDfsWithCycle(graph map[string][]string, src string, path map[string]bool) bool {
	if _, srcInSrcPath := path[src]; srcInSrcPath {
		fmt.Printf("Cycle: %s depends on itself!\n", src)
		return true
	}
	path[src] = true
	for _, tgt := range graph[src] {
		if visited(tgt) {
			continue
		}
		if _, ok := graph[tgt]; ok { // tgt has prereqs, visit them first
			cycle := visitDfsWithCycle(graph, tgt, path)
			if cycle {
				return true
			}
		}
		seq = append(seq, tgt)
		seen[tgt] = true
	}
	return false
}
func visited(vertex string) bool {
	_, ok := seen[vertex]
	return ok
}
