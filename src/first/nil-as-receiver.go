// It's important to realize that in Go nil can be a valid receiver in a method call.
// This is different from languages like Java where the JVM throws a NullPointerException at runtime.
package main

// Consider a pair of strings
type Pair [2]string

// Members returns the members of the pair as an ordered sequence of strings
// For nil pairs, it returns a sequence of two empty strings
func (p *Pair) Members() (string, string) {
	if p == nil {
		return "", ""
	}
	return p[0], p[1]
}
//func main() {
//	var a = [2]string{"Omelet", "Hash Browns"}
//	p := Pair{a}
//}