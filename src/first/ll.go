package main

import "fmt"

//1  class Cell {
//2    int key;
//3    Cell next;
//4
//5    static Cell recReverse(Cell c) {
//6      if (c == null || c.next == null)
//7        return c; // the base case, returns the head!
//8      //send next to this method, get the head of reversed list
//9      Cell headOfRemainingReversed = recReverse(c.next);
//10     //c.next is current tail of reversed linked list,
//11     //so the current element should goto its tail
//12     c.next.next = c;
//13     // and the current element’s next is the end!
//14     c.next = null;
//15     //return the head that we got from recursive call
//16     return headOfRemainingReversed;
//17   }
//18 }
type Cell struct  {
	key int
	next *Cell
}

func recReverse(c *Cell) *Cell {
	if c == nil || c.next == nil {
		return c
	}
	headOfRem  := recReverse(c.next)
	c.next.next = c  // Anish thinks that this results in a nil pointer deref :-)
	c.next = nil
	return headOfRem
}
func (receiver *Cell) path() string {
	if receiver == nil {
		return ""
	}
	return fmt.Sprintf("%d->%s", receiver.key, receiver.next.path())
}
func (receiver *Cell) print() {
	fmt.Printf("%s\n", receiver.path())
}
// 5    static Cell tailReverse(Cell c, Cell acc) {
//6      while (c != null) {
//7        Cell n = c.next; // save the next
//8        c.next = acc;  // current.next = accumulator
//9        c = n;
//10       acc = c;
//11     }
//12     return acc;
func tailReverse(c *Cell, acc *Cell) *Cell {
	for c != nil {
		n := c.next
		c.next = acc
		acc = c
		c = n
	}
	return acc
}
func main() {
	// create the ll
	two := Cell {2, nil}
	one := Cell{1, &two}
	//one := Cell{1, nil}
	(&one).print()
	//rev := recReverse(&one)
	//rev.print()
	rev2 := tailReverse(&one, nil)
	rev2.print()
}
