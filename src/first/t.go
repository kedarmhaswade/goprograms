package main

import (
	"fmt"
	"strconv"
)

type WorkflowType struct {
	WorkflowName string
	MajorVersion uint64
	MinorVersion uint64
	PatchVersion uint64
}

func main() {
	t := WorkflowType{WorkflowName: "record", MajorVersion: 1, MinorVersion: 0, PatchVersion: 0}
	s := t.WorkflowName + "V" + strconv.Itoa(int(t.MajorVersion))

	fmt.Println(s)
	var x map[int]string
	if val, ok := x[1]; ok {
		fmt.Printf("%s\n", val)
	} else {
		fmt.Printf("%v\n", "no mapping")
	}
	type Foo struct {

	}
	foo := &Foo{}
	fmt.Printf("%v\n", foo)
}
