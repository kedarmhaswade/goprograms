package main

import (
    "fmt"
    "os"
    "strconv"
    "first/problem24"
)

func main() {
    if len(os.Args) < 3 {
        fmt.Fprintf(os.Stderr, "Usage: %v result operands...\n", os.Args[0])
        return
    }
    var result float64
    var err error
    if result, err = strconv.ParseFloat(os.Args[1], 64); err != nil {
        fmt.Fprintf(os.Stderr, "parse error, provide a number instead of %v\n", os.Args[1])
        return
    }
    problem24.SolveIt(result, os.Args[2:]...)
}
