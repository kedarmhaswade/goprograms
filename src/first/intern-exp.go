package main

import (
	"fmt"
	"time"
)

// Are strings interned in Go?

func main() {
	minDate := time.Now().AddDate(-8, 0, 0)
	fmt.Printf("%v", minDate)
}
