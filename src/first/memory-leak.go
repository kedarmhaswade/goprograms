package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"regexp"
	"time"
)

var digitRegexp = regexp.MustCompile(".*1.*2.*")

func FindDigits(filename string) []byte {
	b, _ := ioutil.ReadFile(filename)
	return digitRegexp.Find(b) // this leaks memory

}

func CopyDigits(filename string) []byte {
	b, _ := ioutil.ReadFile(filename)
	b = digitRegexp.Find(b)
	c := make([]byte, len(b))
	copy(c, b)
	return c
}

func AppendDigits(filename string) []byte {
	b, _ := ioutil.ReadFile(filename)
	b = append([]byte(nil), digitRegexp.Find(b)...)
	return b
}
func main() {
	tempFile, err := ioutil.TempFile("/tmp", "go-mem-leak")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tempFile.Name())
	size := 5
	fill(tempFile, size)
	tempFile.Sync()
	tempFile.Close()
	readFile(tempFile.Name())
	fmt.Printf("matching: %c\n", CopyDigits(tempFile.Name()))
	fmt.Printf("matching: %c\n", AppendDigits(tempFile.Name()))
}
func readFile(file string) {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Printf("error reading: %v\n", err)
	} else {
		fmt.Printf("contents of file: %c\n", bytes)
	}
}
func fill(file *os.File, total int) {
	current := 0
	var z int = '0'
	rand.Seed(time.Now().UnixNano())
	for current < total {
		var buffer []byte
		for i := 0; i < 10; i++ {
			b := byte(z + rand.Intn(10))
			buffer = append(buffer, b)
		}
		size := len(buffer)
		if current+size > total {
			buffer = buffer[:total-current]
		}
		//fmt.Printf("buffer: %c\n", buffer)
		_, err := file.Write(buffer)
		if err != nil {
			log.Println(nil)
		}
		current += size
	}
}
