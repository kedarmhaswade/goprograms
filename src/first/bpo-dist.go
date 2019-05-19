package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	names := []string{"A", "V", "K", "N", "P", "KK", "S", "T", "SS", "AA", "KKK"}
	groups := make(map[int][]string, 3)
	distOk := false
	attempt := 0
	for ; !distOk; {
		for i := 0; i < len(names); i++ {
			rand.Seed(time.Now().UnixNano())
			idx := rand.Intn(3)
			groups[idx] = append(groups[idx], names[i])
			//fmt.Printf("idx: %d, len(groups[idx]): %d\n", idx, len(groups[idx]))
		}
		distOk = isDistOk(groups)
		if !distOk {
			attempt += 1
			fmt.Printf("failed attempt: %d\n", attempt)
			groups = make(map[int][]string, 3)
		}
	}
	fmt.Printf("%v\n", groups)
}

func isDistOk(groups map[int][]string) bool {
	for _, v := range groups {
		if len(v) <=2 || len(v) >=5 {
			//fmt.Printf("group %d too small or too big -- size: %d\n", k, len(v))
			return false
		}
	}
	return true
}
