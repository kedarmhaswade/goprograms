// Removes the headstart constraint to measure the difference, see no-headstart.go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ng := 40 // comparable to number of processors on this computer, but benchmarking here
	startTimes := make(chan int64, ng)
	var endGate sync.WaitGroup
	for i := 1; i <= ng; i++ {
		endGate.Add(1)
		go func(id int) {
			defer endGate.Done()
			startTimes <- time.Now().UnixNano()
		}(i)
	}
	// now open the startgate!
	endGate.Wait()
	// now clear the channel, since we know that goroutines have ended
	min := <-startTimes // the first goroutine to start
	for i := 1; i <= ng-2; i++ {
		<-startTimes // consume others
	}
	max := <-startTimes // the last goroutine to start
	fmt.Printf("the max headstart: %vns\n", max-min)
}
