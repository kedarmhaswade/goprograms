package main

import "fmt"

type Severity int

const (
	// SeverityWarn is warning level
	SeverityWarn Severity = iota + 1
	// SeverityError is error level
	SeverityError
)

func main() {
	fmt.Printf("%d, %d\n", SeverityWarn, SeverityError)
}
