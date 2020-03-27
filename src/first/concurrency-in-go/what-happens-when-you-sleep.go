package main

import (
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

func main() {
	for i := 0; i < 100; i++ {
		go workLifeImbalance()
	}
}

func workLifeImbalance() {
	b := false
	for {
		for i := 0; i < 1000000; i++ {
			var n big.Int
			n.SetUint64(423432442233424342)
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			var z big.Int
			z.Rand(r, &n)
			b = b && z.ProbablyPrime(40)
		}
		fmt.Printf("final: %v\n", b)
		time.Sleep(1 * time.Second)
	}
}
