// Imagine that a program starts a number of gorotines in a loop.
// Each goroutine runs for a random amount of time up to a second.
// How do we print (every so often) a count of the number of concurrently executing goroutines?
package main

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"math/rand"
	"os"
	"os/signal"
	"runtime"
	"sync/atomic"
	"syscall"
	"time"
)

func main() {
	c := int32(0)
	rand.Seed(time.Now().UnixNano())
	go captureSigQuit()
	for j := 1; j <= 1000; j++ {
		go func() {
			change(&c, 1)
			defer change(&c, -1)
			work()
		}()
	}
	val := atomic.LoadInt32(&c)
	for val > 0 {
		fmt.Printf("current: %d\n", val)
		time.Sleep(100 * time.Millisecond)
		val = atomic.LoadInt32(&c)
	}
	fmt.Printf("current: %d\n", val)
}

// atomically changes the Value of the integer with given pointer by the given delta
func change(c *int32, delta int32) {
	for {
		old := atomic.LoadInt32(c)
		cur := old + delta
		s := atomic.CompareAndSwapInt32(c, old, cur)
		if s {
			return
		}
	}
}

// returns a random string of decimal digits of given length
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

// simulates work by testing if a number of random decimal numbers are probably prime
func work() {
	p := new(big.Int)
	times := rand.Intn(100)
	for j := 0; j < times; j++ {
		p.SetString(randomDecimalString(20), 10)
		//fmt.Printf("number: %v\n", p)
		if p.ProbablyPrime(10) {
			fmt.Fprintf(ioutil.Discard, "%v is probably prime!\n", p)
		}
	}
}
func sleepOnTheJob() {
	d := time.Duration(rand.Intn(1000))
	time.Sleep(d * time.Millisecond)
}
func captureSigQuit() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGQUIT)
	buf := make([]byte, 1<<20)
	for {
		<-sigs
		len := runtime.Stack(buf, true)
		fmt.Printf("=== received ctrl+d ===\n*** goroutine dump...\n%s\n*** end\n", buf[:len])
	}
}
