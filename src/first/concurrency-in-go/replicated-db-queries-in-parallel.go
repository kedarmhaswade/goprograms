package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Conn struct {
	URL string
}

func randomResult(c *Conn) string {
	return fmt.Sprintf("%v from db: %v", rand.Intn(500), c.URL)
}
func (c *Conn) DoQuery(query string) string {
	// simulate blocking
	os.Stdin.Read(make([]byte, 1)) // read a single byte
	return randomResult(c)
}

func Query(conns []Conn, query string) string {
	ch := make(chan string)
	for _, conn := range conns {
		go func(c Conn) {
			select {
			case ch <- c.DoQuery(query):
			default:
				ch <- "non"
			}
			//ch <- c.DoQuery(query)
		}(conn)
	}
	time.Sleep(100 * time.Millisecond) // beware, not prod code
	//time.After()
	return <-ch
}
func main() {
	rand.Seed(time.Now().UnixNano())
	conns := []Conn{{URL: "one"}, {URL: "two"}, {URL: "three"}}
	result := Query(conns, "foo")
	fmt.Printf(" result: %v\n", result)
}
