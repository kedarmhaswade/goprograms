package main

import (
	"fmt"
	"strconv"
)

func sprint(x interface{}) string {
	type stringer interface {
		String() string
	}
	switch x := x.(type) { // type switch
	case stringer:
		return x.String()
	case fmt.Stringer:
		return x.String()
	case string:
		return x
	case int:
		return strconv.Itoa(x)
	// ...similar cases for int16, uint32, and so on...
	case bool:
		if x {
			return "true"
		}
		return "false"
	default:
		// array, chan, func, map, pointer, slice, struct
		return "???"
	}
}

type myName struct {
	first string
	last  string
}

func (receiver *myName) String() string {
	fmt.Printf("String was called on this object of type %T\n", receiver)
	return fmt.Sprintf("%s %s", receiver.last, receiver.first)
}

func main() {
	s := myName{
		first: "Joe",
		last:  "Blo",
	}
	fmt.Printf(sprint(&s))
}
