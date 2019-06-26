// Should we ignore timezone in time.Date?
package main

import (
	"fmt"
	"time"
)


func main() {
	var s int64 = 0
	for {
		t := time.Unix(s, 0)
		if t.Weekday().String() == "Friday" {
			fmt.Println("s: ", s)
			break
		}
		s++
	}
}
