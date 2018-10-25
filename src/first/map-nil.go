// using nil key to get the mapping?
package main

import (
	"fmt"
)

func main() {
	var m map[string]struct{}
	m = make(map[string]struct{})
	var s string
	_, ok := m[s]
	fmt.Printf("%v\n", ok)
	for k, v := range m {
		fmt.Printf("k: %v, v: %v\n", k, v)
	}
	var n map[string][]string // e.g. a map of "student name" -> slice of her/his grades
	n = nil
	key := "foo"
	fmt.Printf("value of a non-nil key: %s in a 'nil' map: %v (the default value of the value type)\n", key, n[key])
	// can we do a "get" on a nil map?
}
