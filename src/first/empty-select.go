package main

import (
	"bufio"
	"os"
)

func main() {
	//go func() {
	//	t := time.NewTicker(time.Millisecond * 10)
	//	select {
	//	case <-t.C:
	//		// Initialize two big ints with the first two numbers in the sequence.
	//		a := big.NewInt(0)
	//		b := big.NewInt(1)
	//
	//		// Initialize limit as 10^99, the smallest integer with 100 digits.
	//		var limit big.Int
	//		limit.Exp(big.NewInt(10), big.NewInt(1000), nil)
	//
	//		// Loop while a is smaller than 1e100.
	//		for a.Cmp(&limit) < 0 {
	//			// Compute the next Fibonacci number, storing it in a.
	//			a.Add(a, b)
	//			// Swap a and b so that b is the next number in the sequence.
	//			a, b = b, a
	//		}
	//		fmt.Println(a) // 100-digit Fibonacci number
	//
	//		// Test a for primality.
	//		// (ProbablyPrimes' argument sets the number of Miller-Rabin
	//		// rounds to be performed. 20 is a good value.)
	//		fmt.Println(a.ProbablyPrime(200))
	//	}
	//}()
	go func() {
		//fmt.Printf("inside another goroutine\n")
		//time.Sleep(1 * time.Minute)
		select {}
	}()
	for bufio.NewScanner(os.Stdin).Scan() {}
}
