package main

import (
	"math/rand"
	"time"
)

type Person struct {
	address StreetAddress
}
type StreetAddress struct {
	HouseNumber int
}

func NewAddress() StreetAddress {
	rand.Seed(time.Now().UnixNano())
	return StreetAddress{rand.Intn(100)}
}
func NewPerson(address StreetAddress) Person {
	return Person{address}
}
