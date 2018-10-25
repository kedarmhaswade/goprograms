// sorts strings, there are no generics in Go yet ;)
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//s := []string{"c", "b", "a"}// getStrings(3)
	s := getStrings(10000)
	//fmt.Println("before.........")
	//print(s)
	fmt.Printf("sorted? %v\n", isSortedStrings(s))
	s = Sort(s)
	//fmt.Println("after..........")
	//print(s)
	fmt.Printf("sorted? %v\n", isSortedStrings(s))
}

//func main() {
//	s := make([]int, 2)
//	s[0] = 13
//	s[1] = 2
//	fmt.Println("before.........")
//	print(s)
//	iSortInt(s)
//	fmt.Println("after..........")
//	print(s)
//}

func Sort(s []string) []string {
	return iSortStringNoSwap(s)
}
func iSortStringSwap(s []string) []string {
	for i := 1; i < len(s); i++ {
		for j := i - 1; j >= 0 && s[j] > s[j+1]; j-- {
			tmp := s[j]
			s[j] = s[j+1]
			s[j+1] = tmp
		}
	}
	return s
}

func iSortStringNoSwap(s []string) []string {
	for i := 1; i < len(s); i++ {
		t := s[i]
		j := i - 1
		for ; j >= 0 && t < s[j]; j-- {
			s[j+1] = s[j]
		}
		s[j+1] = t
	}
	return s
}

func iSortInt(s []int) []int {
	for i := 0; i < len(s)-1; i++ {
		j := i + 1
		t := s[j]
		for ; j >= 1 && s[j] < s[j-1]; j-- {
			s[j-1] = s[j]
		}
		s[i+1] = t
	}
	return s
}

func getStrings(n int) []string {
	s := make([]string, n)
	for i := 0; i < n; i++ {
		s[i] = randomEnglishAlphabetString(5)
	}
	return s
}
func randomEnglishAlphabetString(len int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, len)
	for i := 0; i < len; i++ {
		b[i] = byte(rand.Intn(26) + 'a')
		//fmt.Printf("ch: %c\n", b[i])
	}
	return string(b[:])
}
func print(s []string) {
	for _, x := range s {
		fmt.Printf("%v\n", x)
	}
}
func isSortedStrings(a []string) bool {
	for i := 0; i < len(a)-1; i++ {
		if a[i] > a[i+1] {
			return false
		}
	}
	return true
}
