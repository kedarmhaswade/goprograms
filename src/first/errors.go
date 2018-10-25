package main

import (
	"fmt"
	"strconv"
)

func main() {
	i, err := strconv.Atoi("110")
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Println("Converted integer:", i)
}
