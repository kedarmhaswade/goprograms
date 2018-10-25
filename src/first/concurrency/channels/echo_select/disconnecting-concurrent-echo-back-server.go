// Demonstrates a simple, concurrent echo server that echoes back to a client, as in concurrent-echo-back-server.go
// In addition, this server disconnects a client if it does not send anything to echo back in some timeout (10s)
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	fmt.Printf("Welcome to a disconnecting echo-back server, say something, or else ...\n")
	server := NewEchoServer(2021) // TODO: take port on cmd line
	err := startListening(server) // no need to wait as this uses blocking IO
	if err != nil {
		fmt.Printf("error listening: %v\n", err)
	}
}

type Server struct {
	listenPort int
}

func NewEchoServer(port int) Server {
	var s Server
	s.listenPort = port
	return s
}
func startListening(server Server) error {
	listener, e := net.Listen("tcp", fmt.Sprint("localhost:", server.listenPort))
	if e != nil {
		return e
	}
	// first start a goroutine to handle the communication, possibly indefinitely
	go doCommunication(listener)
	// in the main goroutine, i.e. this one, just wait for the user to enter stop!
	waitForStopCommand()
	return nil // success
}
func doCommunication(listener net.Listener) {
	for {
		conn, e := listener.Accept()
		if e != nil {
			fmt.Fprintf(os.Stderr, "error accepting connections: %v\n", e)
			d := 10 * time.Second
			fmt.Fprintf(os.Stderr, "continuing still after a delay of %v\n...", d)
			time.Sleep(d)
			continue
		}
		// to be a concurrent server we need to hand this connection off to another goroutine
		fmt.Printf("new remote client connected: %v\n", conn.RemoteAddr())
		go echo(conn)
	}
}
func echo(conn net.Conn) {
	defer conn.Close() // we need this and another conn.Close() call in the actual goroutine below as well
	// simply echoes back to the connection -- this is a blocking call
	// if the client says nothing, we wait for some timeout before disconnecting
	ticker := time.NewTicker(5 * time.Second)
	reader := make(chan struct{})
	defer ticker.Stop()
	go func() {
		defer conn.Close() // we are doing this to clean up after, but should we, since we are handed over this conn?
		tickerStopped := false
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			if !tickerStopped {
				tickerStopped = true
				ticker.Stop()
			}
			fmt.Fprintf(conn, "%s\n", scanner.Text())
		}
		if scanner.Err() == nil {
			fmt.Printf("disconnected client: %v\n", conn.RemoteAddr())
		}
	}()
	select {
	case <-ticker.C: // we received no echo and time's up
		_, err := fmt.Fprintf(conn, "%s\n", "Time's up!")
		if err != nil {
			fmt.Printf("Remote client:%v disconnected, but couldn't write to it: %v\n", conn.RemoteAddr(), err)
		}
		return
	case <-reader: // we received something from client, don't disconnect
		// do nothing
	}
}
func waitForStopCommand() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		cmd := scanner.Text()
		if cmd == "stop!" {
			fmt.Println("good bye!")
			break
		}
	}
}
