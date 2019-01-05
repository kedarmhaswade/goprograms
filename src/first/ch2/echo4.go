// Echo4 prints its command-line arguments.
package ch2

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline") // type of n is *bool
var sep = flag.String("s", " ", "separator") // type of sep is *string

func f() {
	fmt.Printf("type of n: %T, type of sep: %T\n", n, sep)
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}