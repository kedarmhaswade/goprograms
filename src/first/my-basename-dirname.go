// Demonstrates implementation of basename and dirname
// The basename function below was inspired by the Unix shell utility of the same name.
// In our version, basename(s) removes
// - any prefix of s that looks like a file system path with components separated by slashes, and
// - any suffix that looks like a file type
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(basename("a/b/c.go")) // "c"
	fmt.Println(basename("c.d.go"))   // "c.d"
	fmt.Println(basename("abc"))      // "abc"
	fmt.Println(basename("/x/abc/"))  // "abc"
}

func basename(s string) string {
	n := len(s)
	if n == 0 {
		return s
	}
	if strings.HasSuffix(s, "/") {
		s = s[:n-1]
	}
	i := strings.LastIndexByte(s, '/')
	j := strings.LastIndexByte(s, '.')
	if i == -1 && j == -1 {
		return s
	} else if i != -1 && j == -1 {
		return s[i+1:]
	} else if i == -1 && j != -1 {
		return s[:j]
	} else {
		// ensure j > i?
		return s[i+1 : j]
	}
}
