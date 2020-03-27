package main

import "fmt"

// this iota "constant generator" is really confusing and perhaps whimsical
func main() {
	const (
		mutexLocked = 1 << iota // mutex is locked
		mutexWoken
		mutexStarving
		unused = 100
		mutexWaiterShift = iota
	)
	fmt.Printf("mutexLocked = %v\n", mutexLocked)
	fmt.Printf("mutexWoken = %v\n", mutexWoken)
	fmt.Printf("mutexStarving = %v\n", mutexStarving)
	fmt.Printf("unused = %v\n", unused)
	fmt.Printf("mutexWaiterShift = %v\n", mutexWaiterShift)
}
