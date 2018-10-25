// Implement a concurrent File Transfer Protocol (FTP) server. The server should interpret commands from each
// client such as cd to change directory, ls to list a directory, get to send the contents of a file, and
// close to close the connection. You can use the standard ftp command as the client, or write your own.
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

func main() {

	args := os.Args
	port := 2021 // default: 21 + 2000
	if len(args) >= 2 {
		p, err := strconv.Atoi(args[1])
		if err != nil {
			log.Printf("invalid port: %s, using default: %d", args[1], port)
		}
		port = p
	}
	listener, e := net.Listen("tcp", fmt.Sprint("localhost:", port))
	if e != nil {
		log.Fatalf("listen error: %v\n", e)
		return
	}
	log.Printf("listening at: %s\n", listener.Addr())
	for {
		conn, e := listener.Accept()
		if e != nil {
			log.Fatalf("error while accepting a connection request: %v, continuing ...\n", e)
			continue
		}
		go doFtp(conn)
	}
}

type Context struct {
	cwd string
}

func NewContext() *Context {
	var c Context
	p, e := os.Getwd()
	if e != nil {
		// no need to panic, maybe?
		log.Fatalf("fatal error getting server's cwd: %v\n", e)
	}
	c.cwd = p
	return &c
}

func doFtp(conn net.Conn) {
	defer conn.Close()
	for {
		var b bytes.Buffer
		_, err := io.Copy(bufio.NewWriter(&b), conn)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error reading command: %v, ignoring for now\n", err)
			continue
		}
		context := NewContext()
		s := string(b.Bytes())
		fmt.Printf("cmd: %s\n", s)
		return
		switch cmd := strings.Fields(s); cmd[0] {
		case "ls":
			fallthrough
		case "cd":
			fallthrough
		case "get":
			io.WriteString(conn, "not yet implemented")
			break
		case "pwd":
			io.WriteString(conn, context.cwd)
		default:
			io.WriteString(conn, "good bye!")
			return // returns from the function
		}
	}
}
