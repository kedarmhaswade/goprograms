// The channel ch, whose buffer size is 1, is alternately empty then full,
// so only one of the cases can proceed, either the send when i is even,
// or the receive when i is odd. It always prints 0 2 4 6 8.
package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch: // blocks, or waits when i is even
			fmt.Println("chan rece: ", x, ", i: ", i) // "0" "2" "4" "6" "8"
		case ch <- i: // blocks, or waits when i is odd
			fmt.Println("chan send: ", i, ", i: ", i)
		}
	}
	//ch := make(chan int, 1)
	//i := 0
	//for i <= 8 {
	//	select {
	//	case i = <-ch: // receive
	//		fmt.Printf("%d ", i)
	//		i += 1
	//	case ch <- i: // send
	//		i += 1
	//	default:
	//		fmt.Printf("nothing to do\n")
	//	}
	//}
}
