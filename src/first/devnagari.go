// These are devnagari runes!
package main

import "fmt"

func main() {
	for c := '\u0900'; c <= '\u097F'; c++ {
		fmt.Printf("devnagari unicode code point (decimal/hex): %d/%X = %c\n", c, c, c)
	}
	fmt.Printf("total devnagari: %d\n", '\u097f'-'\u0900')
	fmt.Println()
	for c := '\u0A80'; c <= '\u0AF9'; c++ {
		fmt.Printf("gujarati unicode code point (decimal/hex): %d/%X = %c\n", c, c, c)
	}
	fmt.Printf("total gujarati: %d\n", '\u0AF9'-'\u0A80')
	"इ°℃"
}
