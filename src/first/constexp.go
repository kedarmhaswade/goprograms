package main

import "fmt"

type Flags uint

const (
	FlagUp           Flags = 1 << iota // is up
	FlagBroadcast                      // supports broadcast access capability
	FlagLoopback                       // is a loopback interface
	FlagPointToPoint                   // belongs to a point-to-point link
	FlagMulticast                      // supports multicast access capability
)

func main() {
	const s = "some name"
	// v := "some name"
	// fmt.Printf("%c %c\n", s[8], v[10])  // fails at run time
	// fmt.Printf("%c %c\n", s[10], v[10])  // fails at compile time
	const (
		a = 2
		b
		c
		d = 4
	)
	fmt.Printf("%v %v %v %v\n", a, b, c, d)
	type Weekday uint
	const (
		Sunday Weekday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	day := Tuesday - 1
	fmt.Printf("The day is a %T with value %v\n", day, day)

	flag := FlagMulticast
	fmt.Printf("flag: %v\n", flag)
	flag = Flags(17) // note: Flags is just another name for int
	fmt.Printf("new flag: %v isUp? %v\n", flag, isUp(flag))
	turnDown(&flag)
	fmt.Printf("new flag: %v isUp? %v\n", flag, isUp(flag))
	v := FlagMulticast | FlagUp
	fmt.Printf("%b %t\n", v, isUp(v)) // "10001 true"
	turnDown(&v)
	fmt.Printf("%b %t\n", v, isUp(v)) // "10000 false"
	setBroadcast(&v)
	fmt.Printf("%b %t\n", v, isUp(v))   // "10010 false"
	fmt.Printf("%b %t\n", v, isCast(v)) // "10010 true"

	var t, u = FlagUp, FlagUp
	toggleUpDown(&t)
	toggleUpDown(&t)
	fmt.Printf("restored? %v\n", t == u)
}
func isUp(v Flags) bool {
	return v&FlagUp == FlagUp
}
func turnDown(v *Flags) {
	*v &^= FlagUp // &^ is the bit clear operator, ^= is toggle!
}
func toggleUpDown(v *Flags) {
	*v ^= FlagUp
}
func setBroadcast(v *Flags) {
	*v |= FlagBroadcast
}
func isCast(v Flags) bool {
	return v&(FlagBroadcast|FlagMulticast) != 0
}
