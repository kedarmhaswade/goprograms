// Demonstrates a simple, concurrent echo server that echoes back to a client
package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func main() {
	fmt.Printf("Welcome to the echo-back server\n")
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
	defer conn.Close() // we are doing this to clean up after, but should we, since we are handed over this conn?
	// simply echoes back to the connection, of course this is a blocking call
	// if the client says nothing, we wait forever till the client disconnects (sends ^D, e.g.)
	_, err := io.Copy(conn, conn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error writing to connection %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "remote client disconnected successfully: %v\n", conn.RemoteAddr())
}
func waitForStopCommand() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		cmd := scanner.Text()
		if cmd == "stop!" {
			fmt.Printf("good bye!\n")
			break
		}
	}
}
