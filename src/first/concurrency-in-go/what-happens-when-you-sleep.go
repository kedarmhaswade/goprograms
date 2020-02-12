package main

import (
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

func main() {
	for i := 0; i < 100; i++ {
		b := work()
		sleep(i, b)
	}
}

func sleep(i int, b bool) {
	fmt.Printf("zzz... i = %v, b = %v\n", i, b)
	time.Sleep(1*time.Second)
}
func work() bool {
	b := false
	for i := 0; i < 100000; i++ {
		var n big.Int
		n.SetUint64(423432442233424342)
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		var z big.Int
		z.Rand(r, &n)
		b = z.ProbablyPrime(40)
	}
	return b
}
