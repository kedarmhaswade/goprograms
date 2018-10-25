// Methods may be declared on any "named type" defined in the same package,
// so long as its underlying type is neither a pointer nor an interface.
package ch6

import "strconv"

type Integer int

//ToString defines a method on type Integer. It represents the string value of the integer represented by n.
func (n Integer) ToString() string {
	return strconv.Itoa(int(n))
}

//Multiply accepts a pointer to Integer as a receiver and multiplies its value by the given parameter
func (ptr *Integer) Multiply(times int) {
	*ptr *= Integer(times)
}

// Values maps a string key to a list of values (copied from net/url)
type Values map[string][]string  // value type is []string => multimap

// Get returns the first value associated with the given key,
// or "" if there are none.
func (v Values) Get(key string) string {
	if vs := v[key]; len(vs) > 0 {
		return vs[0]
	}
	return ""
}

// Add adds the value to key.
// It appends to any existing values associated with key.
func (v Values) Add(key, value string) {
	v[key] = append(v[key], value)
}
