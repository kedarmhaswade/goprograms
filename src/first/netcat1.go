// read-only TCP client
package main

// my version
//import (
//	"net"
//	"fmt"
//)
//
//func main() {
//	conn, err := net.Dial("tcp4", "localhost:8000")
//	defer conn.Close()
//	if (err != nil) {
//		fmt.Println(err)
//		return
//	}
//	for {
//		b := make([]byte, 10)
//		_, err := conn.Read(b)
//		if err != nil {
//			fmt.Println(err)
//			break
//		}
//		fmt.Printf("%s", b[0:len(b) - 1])
//	}
//}

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

// the version in the book
