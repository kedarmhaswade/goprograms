package main

import (
	"fmt"

	uuid2 "github.com/gofrs/uuid"
)

func main() {
	// numeric types
	// is byte unsigned?
	var b byte = 255 // 256 // overflow
	fmt.Printf("%v\n", b)
	// conversion
	var x = 104
	fmt.Println(string(x)) // the underlying type of x and

	uuid, err := uuid2.NewV4()
	if err != nil {
		_ = fmt.Errorf("got %v\n", err)
	}
	fmt.Printf("uuid.String() = %v\n", uuid.String())
	bytes := [16]byte(uuid)
	fmt.Printf("%v\n", string(bytes[:]))
}
