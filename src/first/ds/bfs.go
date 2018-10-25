// Implements a breadth-first traversal based topological sort
package ds

import (
	"github.com/golang-collections/collections/stack"
)

//TopologicallySortedBFS returns a topologically sorted order of the given graph.
//Implementation detail: Uses BFS and detects cycles
func TopologicallySortedBFS(graph map[string][]string) ([]string, bool) {
	// the given graph is of the form: "a":{"b", "c"}, "c": {"d", "e"}, "d": {"e"}
	var pile = stack.New()
	v := len(graph)
	if v == 0 {
		return []string{}, false
	}
	seen := make(map[string]bool, v) // set of seen vertices
	var q []string
	for vertex := range graph {
		q = append(q, vertex)
	}
	for len(q) > 0 {
		v1 := q[0]
		q = q[1:]
		if ok := seen[v1]; ok {
			//fmt.Printf("already seen: %v\n", v1)
			continue
		} else {
			seen[v1] = true
			pile.Push(v1)
			for _, v2 := range graph[v1] {
				if ok := seen[v2]; !ok {
					q = append(q, v2)
				}
			}
		}
	}
	//fmt.Printf("%v\n", order)
	var order []string
	for pile.Len() > 0 {
		e := (pile.Pop()).(string)
		order = append(order, e)
	}
	return order, false
}
