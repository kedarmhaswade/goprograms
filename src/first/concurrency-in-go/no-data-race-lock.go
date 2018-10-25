package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var memoryAccess sync.Mutex // (1)
	var value int
	go func() {
		memoryAccess.Lock() // (2)
		value++
		memoryAccess.Unlock() // (3)
	}()
	// comment/uncomment out up to HERE to see the diff
	//p := new (big.Int)
	//p.SetString(randomDecimalString(2), 10)
	//if p.ProbablyPrime(10) {
	//	//fmt.Fprintf(ioutil.Discard, "%v is probably prime!\n", p)
	//	fmt.Fprintf(os.Stdout, "%v is probably prime!\n", p)
	//} else {
	//	fmt.Fprintf(os.Stdout, "%v is probably _not_ a prime!\n", p)
	//}
	// HERE
	memoryAccess.Lock() // (4)
	if value == 0 {
		fmt.Printf("the value is %v.\n", value)
	} else {
		fmt.Printf("the value is %v.\n", value)
	}
	memoryAccess.Unlock() // (5)
}
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
