package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {
	var w io.Writer = os.Stdout
	staticDynamic(w)
}

func staticDynamic(v interface{}) {

	d := reflect.TypeOf(v)
	vv := reflect.ValueOf(v)
	fmt.Printf("dynamic type of v: %v, value: %v\n", d, vv)
}
