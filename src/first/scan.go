// Demonstrates the use of Scanner interface
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Enter text ...\n")
	input := bufio.NewScanner(os.Stdin) // a new scanner with a default split function that splits input on lines
	//input.Split(bufio.ScanWords) // scans the input as tokens of "words"
	for input.Scan() {
		text := input.Text()
		fmt.Println(text)
	}
	if input.Err() != nil {
		fmt.Printf("input error: %v\n", input.Err())
	} else { // // ^D, or EOF is a typical, not exceptional behavior
		fmt.Printf("Good bye!\n")
	}
}
