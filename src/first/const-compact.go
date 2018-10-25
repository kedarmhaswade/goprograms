// Gopl.io
// Exercise 3.13: Write const declarations for KB, MB, up through YB as compactly as you can.
// KB, MB, (the powers of 10 can't written directly using the iota mechanism as there is no exponentiation operator)

package main

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

//const (
//	KB = 1000
//	MB = 1000 * 1000
//	GB = 1000 * 1000 * 1000
//	TB = 1000 * 1000 * 1000 * 1000
//	PB = 1000 * 1000 * 1000 * 1000 * 1000
//	EB = 1000 * 1000 * 1000 * 1000 * 1000 * 1000
//	ZB = 1000 * 1000 * 1000 * 1000 * 1000 * 1000 * 1000 // zetta: overflows int
//	YB = 1000 * 1000 * 1000 * 1000 * 1000 * 1000 * 1000 * 1000 // yotta: overflows int
//	XB = 1000 * 1000 * 1000 * 1000 * 1000 * 1000 * 1000 * 1000 * 1000 // xona: overflows int
//)
// writing it as compactly as I can
const (
	KB = 1000    // 3
	MB = KB * KB // 6  =  3 + 3
	GB = KB * MB //  9 =  3 + 6
	TB = KB * GB // 12 =  3 + 9
	PB = KB * TB // 15 =  3 + 12
	EB = KB * PB // 18 =  3 + 15
	ZB = KB * EB // zetta: overflows int 21 = 3 + 18
	YB = KB * ZB // yotta: overflows int 24 = 3 + 21
	XB = KB * YB // xona: overflows int 27 = 3 + 24
)

func main() {
	p := message.NewPrinter(language.English)
	//p.Printf(" kilo: %v\n mega: %v\n giga: %v\n tera: %v\n peta: %v\n  exa: %v\n", KB, MB, GB, TB, PB, EB, ZB, YB, XB)
	p.Printf(" kilo: %v\n mega: %v\n giga: %v\n tera: %v\n peta: %v\n  exa: %v\n", KB, MB, GB, TB, PB, EB)
}
