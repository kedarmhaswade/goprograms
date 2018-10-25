// Using log.FatalF
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f := "/tmp/xxx"
	log.SetPrefix("xxx: ")
	log.SetFlags(log.Ltime | log.Ldate)
	fmt.Printf("Ldate: %d, Llongfile: %d, Lmicroseconds: %d, Lshortfile: %d, Ltime: %d, LUTC: %d\n", log.Ldate, log.Llongfile, log.Lmicroseconds, log.Lshortfile, log.Ltime, log.LUTC)
	if _, err := os.Lstat(f); err != nil {
		log.Fatalf("File %s does not exist yet!", f) // should exit here
	}
	fmt.Printf("You should not see this!!")
}
