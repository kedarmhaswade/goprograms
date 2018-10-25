package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	m := make(map[string]int, 10)
	for i := 0; i < 10; i++ {
		name := r1(5)
		m[name] = i
		if strings.HasPrefix(name, "a") {
			fmt.Printf("deleting %s, starts with an 'a'\n", name)
			delete(m, name)
		}
	}
	for name, id := range m {
		fmt.Printf("%s:%d\n", name, id)
	}
}

func r1(len int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, len)
	for i := 0; i < len; i++ {
		b[i] = byte(rand.Intn(26) + 'a')
		//fmt.Printf("ch: %c\n", b[i])
	}
	return string(b[:])
}
