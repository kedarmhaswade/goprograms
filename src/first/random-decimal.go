package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomDecimalString(nDigits int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, nDigits)
	for {
		d := rand.Intn(10)
		if d > 0 {
			b[0] = byte(d + '0')
			break
		}
	}
	for i := 1; i < nDigits; i++ {
		b[i] = byte(rand.Intn(10) + '0')
	}
	return string(b[:])
}
func main() {
	fmt.Printf("%s\n", randomDecimalString(20))
}
