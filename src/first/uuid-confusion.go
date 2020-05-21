package main

import (
	"fmt"
	"github.com/pborman/uuid"
)

// UUIDHelper is a util function to wrap UUID logic
type UUIDHelper interface {
	New() uuid.UUID
}

// NewUUIDHelper is a constructor function for UUIDHelper
func NewUUIDHelper() UUIDHelper {
	return &uuidHelper{}
}

type uuidHelper struct{}

// New returns a new UUID
func (o *uuidHelper) New() uuid.UUID {
	return uuid.NewRandom()
}
func main() {
	uh := uuidHelper{}
	n := uh.New()
	fmt.Printf("uuid: %v\n", n)
}
