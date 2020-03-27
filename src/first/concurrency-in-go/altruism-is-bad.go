package main

import "sync"

// Porting https://stackoverflow.com/a/8863671/437506 to Go

// Spoon is a shared resource
type Spoon struct {
	lock sync.WaitGroup
}
type Diner struct {
	name string
	isHungry bool
}
