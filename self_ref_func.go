package main

import (
	"fan1"
	"fan2"
	"fmt"
)

func main() {
	// client of fan1
	f1 := fan1.New()
	f1.Option(fan1.Speed(3), fan1.Color("white"), fan1.Kind(fan1.Ceiling))
	// This call above, according to Rob Pike, is "nice to use for clients".
	// Well, beauty lies in the eyes of the beholder.
	// Would a builder-like call: f1.Speed(3).Color("white") be beautiful? I would think so.
	fmt.Printf("%v\n", f1)

	// The client can use this the same as before, but if the client also wants to restore a previous value,
	// all that's needed is to save the return value from the first call, and then restore it.
	// Here Rob just uses one option
	f2 := fan2.New()
	fmt.Printf("before running: %v\n", f2)
	prevSpeed := f2.Option(fan2.Speed(3))
	f2.Run()
	f2.Option(fan2.Speed(prevSpeed.(int))) // this is definitely clumsy, Rob says ...
	fmt.Printf("after running: %v\n", f2)
	// What if I use multiple options?
	fmt.Printf("before running: %v\n", f2)
	prevColor := f2.Option(fan2.Speed(3), fan2.Color("white"))
	f2.Run()
	f2.Option(fan2.Color(prevColor.(string)))
	fmt.Printf("after running: %v\n", f2)
	// I don't think it is as intuitive as Rob makes it sound
	// Here I am losing my interest in reading that article further.
	// https://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html
}
