// Write a function expand(s string, f func(string) string) string
// that replaces each substring “$foo” within s by the text returned by f("foo").

package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	expand("$foooobarfoo$foo", func(s string) string {
		return strings.ToUpper(s)
	})
}

func expand(s string, f func(string) string) {
	var buf bytes.Buffer
	if len(s) == 0 {
		return
	}
	i := 0
	token := "foo"
	for i < len(s) {
		if s[i] == '$' && len(s) > i+3 && isFoundFree(s, i+1, token) {
			buf.WriteString(f(token))
			i += 1 + len(token)
		} else {
			buf.WriteByte(s[i])
			i += 1
		}
	}
	fmt.Printf("expanded: %s\n", buf.String())
}
func isFoundFree(s string, i int, token string) bool {
	for j := 0; j < len(token); j++ {
		if s[i] != token[j] {
			return false
		}
		i += 1
	}
	return true
}
