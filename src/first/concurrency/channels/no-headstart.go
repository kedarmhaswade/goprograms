// When one starts goroutines in a loop, the first goroutine gets an (unfair?) advantage of a head-start. Is
// it possible to make all the goroutines make progress at the same time?
// Perhaps it is not worth the trouble, but it demonstrates the use of a "latch" (similar to: https://docs.oracle.com/javase/8/docs/api/java/util/concurrent/CountDownLatch.html)

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ng := 40 // comparable to number of processors on this computer, but benchmarking here
	startTimes := make(chan int64, ng)
	var startGate, endGate sync.WaitGroup
	for i := 1; i <= ng; i++ {
		startGate.Add(1)
		endGate.Add(1)
		go func(id int) {
			defer endGate.Done()
			startGate.Wait()
			startTimes <- time.Now().UnixNano()
		}(i)
	}
	// now open the startgate!
	startGate.Add(-ng)
	endGate.Wait()
	// now clear the channel, since we know that goroutines have ended
	min := <-startTimes // the first goroutine to start
	for i := 1; i <= ng-2; i++ {
		<-startTimes // consume others
	}
	max := <-startTimes // the last goroutine to start
	fmt.Printf("the max headstart: %vns\n", max-min)
}
