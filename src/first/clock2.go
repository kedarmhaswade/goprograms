// Clock1 is a TCP server that periodically writes the time.
package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s port_num\n", os.Args[0])
		return
	}
	args := os.Args[1:]
	port, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalf("Invalid port number: %d\n", port)
		return
	}
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Listening at: %v, Time Zone: %v\n", listener.Addr(), os.Getenv("TZ"))
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		//handleConnConc(conn) // handle one connection at a time
		go handleConnConc(conn) // handles multiple connections at a time, uses one goroutine per connection
	}
}

func handleConnConc(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}
