// Just demonstrates the use of DecodeRuneInString, especially for the "replacement character"
package main

import "fmt"

func main() {
	fmt.Printf("%c = %d\n", '\uFFFD', '\uFFFD')
}
