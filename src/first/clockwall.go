// acts as a client of several clock servers at once, reading the times from each one and
// displaying the results in a table, akin to the wall of clocks seen in some business offices.
package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	for i, cs := range os.Args[1:] {
		tokens := strings.Split(cs, "=")
		if len(tokens) == 2 {
			go displayTime(tokens[1], i)
		} else {
			fmt.Printf("invalid format %s, ignoring\n", tokens)
		}
	}
	time.Sleep(2000 * time.Second)
}

func displayTime(server string, pos int) {
	conn, e := net.Dial("tcp", server)
	if e != nil {
		log.Fatalf("%v\n", e)
	}
	defer conn.Close()
	for {
		b := make([]byte, 9) // read exactly nine bytes => time that server writes to connection every second
		// alternatively, we should try to read the exact name string format that server writes periodically to stdout
		_, err := conn.Read(b)
		if err != nil {
			log.Fatal(err)
			break
		}
		for i := 0; i < pos; i++ {
			fmt.Printf("\t\t\t")
		}
		fmt.Printf("%s", string(b))
	}
}
