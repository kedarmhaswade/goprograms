// build a program that reports the disk usage of one or more directories specified on the command line,
// like the Unix du command.
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	name := "/tmp"
	s, e := os.Open(name)
	if e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		return
	}
	size := int64(0)
	num := 0
	du(s, &size, &num)
	fmt.Printf("%d files %d bytes\n", num, size)
}

func du(f *os.File, size *int64, num *int) {
	infos, e := f.Readdir(-1)
	if e != nil {
		fmt.Printf("error:%v traversing %s, giving up\n", e, f.Name())
		return
	}
	for _, finfo := range infos {
		path := filepath.Join(f.Name(), finfo.Name())
		if finfo.IsDir() {
			dir, e := os.Open(path)
			if e != nil {
				//log
				fmt.Printf("%v\n", e)
				continue
			}
			du(dir, size, num)
		} else {
			*size += finfo.Size()
			*num = *num + 1
		}
	}
}
