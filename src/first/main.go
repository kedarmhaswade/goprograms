// temporary main for all programs
package main

import (
	"fmt"

	"first/ch4"
)
func main() {
	data, _ := ch4.Marshal()
	fmt.Printf("%s\n", data)
}
