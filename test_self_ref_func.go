package main

import (
	"fan1"
	"fmt"
)

func main() {
	// client of fan1
	f := fan1.New()
	f.Option(fan1.Speed(3), fan1.Color("white"), fan1.Kind(fan1.Ceiling))
	fmt.Printf("%v\n", f)

}
