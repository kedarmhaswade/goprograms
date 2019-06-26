// When one starts goroutines in a loop, the first goroutine gets an (unfair?) advantage of a head-start. Is
// it possible to make all the goroutines make progress at the same time?
// Perhaps it is not worth the trouble, but it demonstrates the use of a "latch" (similar to: https://docs.oracle.com/javase/8/docs/api/java/util/concurrent/CountDownLatch.html)

package main

import (
	"fmt"
	"sync"
	"time"
)

func getMaxHeadStart(useStartGate bool) int64 {
	ng := 10 // comparable to number of processors on this computer, but benchmarking here
	startTimes := make(chan int64, ng)
	var startGate, endGate sync.WaitGroup
	for i := 1; i <= ng; i++ {
		if useStartGate {
			startGate.Add(1)
		}
		endGate.Add(1)
		go func(id int) {
			defer endGate.Done()
			startGate.Wait()
			startTimes <- time.Now().UnixNano()
			// do your work
		}(i)
	}
	// now open the start-gate!
	if useStartGate {
		startGate.Add(-ng)
	}
	endGate.Wait()
	// now clear the channel, since we know that goroutines have started and ended
	min := <-startTimes // the first goroutine to start
	for i := 1; i <= ng-2; i++ {
		t := <-startTimes // consume others
		if t < min {
			fmt.Printf("Strange! %v (t) should have been greater than %v (min)\n", t, min)
		}
	}
	max := <-startTimes // the last goroutine to start
	close(startTimes)
	m := max - min
	fmt.Printf("use startGate: %5v, the max headstart: %10v ns\n", useStartGate, m)
	return m
}

func main() {
	m1 := getMaxHeadStart(false)
	m2 := getMaxHeadStart(true)
	reducedTo := m1 - m2
	fmt.Printf("Headstart without startGate: %v ns, with startGate: %v ns\n", m1, m2)
	fmt.Printf("Headstart is reduced by: %v ns (bigger the better)\n", reducedTo)
}