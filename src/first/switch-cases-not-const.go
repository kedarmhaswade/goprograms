package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	s("darwin")
	today := time.Now().Weekday()
	fmt.Printf("%d\n", today)
	fmt.Printf("%d\n", time.Saturday)
	fmt.Println(time.Now())
	t1, _ := time.Parse("UTC", "2018-07-16 14:51:15")
	fmt.Println(t1)
}

func s(osRequired string) {
	switch os := runtime.GOOS; os {
	case osRequired:
		fmt.Println("Yes, the OS is the same as required OS: " + osRequired)
	default:
		fmt.Println("No, the OS and required OS differ")
	}
}
