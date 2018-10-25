// Provides some io utils
package util

import (
	"bufio"
	"fmt"
	"os"
)

func Prompt(format string, a ...interface{}) (s string, err error) {
	_, e := fmt.Printf(format, a)
	if e != nil {
		return "", e
	}
	p := "Press Enter to Continue ..."
	fmt.Printf(p)
	reader := bufio.NewReader(os.Stdin)
	return reader.ReadString('\n')
}

